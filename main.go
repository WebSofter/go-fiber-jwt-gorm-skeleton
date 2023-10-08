package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"wsofter.com/database"
	"wsofter.com/routes"
)

func main() {
	if err := godotenv.Load("local.env"); err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	app.Listen(":8000")
}
