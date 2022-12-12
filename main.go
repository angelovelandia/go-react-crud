package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./public")
	app.Listen(":3000")
	fmt.Println("Server listening on 3000")
}
