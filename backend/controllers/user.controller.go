package controllers

import (
	// "strconv"
	"strings"
	"time"
	"log"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/abhibamnote/go-vue/backend/initializers"
	"github.com/abhibamnote/go-vue/backend/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	// "gorm.io/gorm"
)

func GenerateJWT(userID uint) (string, error) {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // 1 day expiry
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token.SignedString(config.JWT))
	return token.SignedString([]byte(config.JWT))
}

func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		config, err := initializers.LoadConfig(".")
		if err != nil {
			log.Fatalln("Failed to load environment variables! \n", err.Error())
		}

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return config.JWT, nil
	})
}


func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func CreateUser(c *fiber.Ctx) error {
	payload := new(models.CreateUserSchema)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	hashed, err := HashPassword(payload.Password)
	if err != nil {
		log.Fatal(err)
	}

	newUser := models.User{
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Email: payload.Email,
		Phone: &payload.Phone,
		CountryCode: &payload.CountryCode,
		Password: hashed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := initializers.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate entry") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Title already exist, please use another title"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	token, err := GenerateJWT(newUser.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": newUser, "token": token}})
}

func LoginUser(c *fiber.Ctx) error {
	payload := new(models.LoginUserSchema)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	var user models.User
	if err := initializers.DB.Where("email = ?", payload.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}
	
	fmt.Println(user.ID)
	doesMatch := CheckPasswordHash(payload.Password, user.Password)
	fmt.Println(doesMatch)
	
	token, err := GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": "user", "token": token}})

}