package main

import (
	"log"

	"github.com/abhibamnote/go-vue/backend/controllers"
	"github.com/abhibamnote/go-vue/backend/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
    config, err := initializers.LoadConfig(".")
    if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
}

func main() {
    // fmt.Println("Hello, World! two")
    app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

    micro.Route("/auth", func(router fiber.Router) {
        router.Post("/signup", controllers.CreateUser)
        router.Post("/login", controllers.LoginUser)
    })

    log.Fatal(app.Listen(":8000"))
};
