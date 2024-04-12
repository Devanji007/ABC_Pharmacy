package main

import (
	"ABC_PHARMACY/controllers"
	"ABC_PHARMACY/routes"
	"log"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &controllers.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	err = controllers.NewConnection(config)

	if err != nil {
		log.Fatal("couldnot load the datsbase")
	}

	app := fiber.New()
	routes.SetupItemRoutes(app)
	routes.SetupInvoiceRoutes(app)
	app.Listen(":8080")

}
