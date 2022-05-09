package main

import (
	"go-fiber-starter/api/database"
	"go-fiber-starter/api/router"
	"log"

	_ "go-fiber-starter/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title Fiber Starter API
// @version 1.0
// @description This is a sample swagger for Fiber Starter API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
