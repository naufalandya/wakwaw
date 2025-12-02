package main

import (
	"log"

	"belajar/app/config"
	"belajar/app/db"
	"belajar/app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()

	if err := db.InitOracle(cfg); err != nil {
		log.Fatalln("Oracle init error:", err)
	}
	defer db.Oracle.Close()

	app := fiber.New()
	routes.SetupRoutes(app)

	log.Println("Server running on :8080")
	log.Fatal(app.Listen(":8080"))
}
