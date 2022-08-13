package main

import (
	"fmt"
	"log"
	"os"


	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
}