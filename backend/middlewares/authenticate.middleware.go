package middlewares

import (
	"fmt"
	"log"

	"github.com/abhibamnote/go-vue/backend/initializers"
	"github.com/abhibamnote/go-vue/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticate(c *fiber.Ctx) error {
	var authToken = c.Get("Authorization")
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		config, err := initializers.LoadConfig(".")
		if err != nil {
			log.Fatalln("Failed to load environment variables! \n", err.Error())
		}

		JWT := []byte(config.JWT)
		return JWT, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Token not found"})

	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["user_id"])
		var user models.User
		if erro := initializers.DB.Where("id = ?", claims["user_id"]).First(&user).Error; erro != nil {
			return c.Status(401).JSON(fiber.Map{"error": "User not found"})
		}

		c.Locals("user", user)
	} else {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}


	return c.Next() 
} 