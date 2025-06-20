package main

import (
	"log"

	"github.com/abhibamnote/go-vue/backend/controllers"
	"github.com/abhibamnote/go-vue/backend/initializers"
	"github.com/abhibamnote/go-vue/backend/middlewares"
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

	api := app.Group("/api")
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

    auth := api.Group("/auth")
	auth.Post("/signup", controllers.CreateUser)
	auth.Post("/login", controllers.LoginUser)

	chartData := api.Group("/chart-data", middlewares.Authenticate)
	chartData.Post("/", controllers.CreateChartDataBatch)
	chartData.Get("/", controllers.GetChartData)


	optionChain := api.Group("/option-chain", middlewares.Authenticate)
	optionChain.Post("/", controllers.CreateOptionChain)
	optionChain.Get("/", controllers.GetOptionChain)

	masterData := api.Group("/master-data", middlewares.Authenticate)
	masterData.Post("/", controllers.TriggerData)
	masterData.Get("/", controllers.GetMasterData)

	watchlist := api.Group("/watchlist", middlewares.Authenticate)
	watchlist.Post("/", controllers.PostWatchList)
	watchlist.Get("/", controllers.GetWatchList)
	watchlist.Delete("/:id", controllers.DeleteWatchList)

	err := controllers.CreateScheduler()
	if err != nil {
		log.Fatal(err)
	}



    log.Fatal(app.Listen(":8000"))
};
