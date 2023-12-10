package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thtrangphu/go-oauth2/config"
	"github.com/thtrangphu/go-oauth2/controllers" //imoprting the controllers package
)

func main() {
    app := fiber.New()
		config.GoogleConfig()
    app.Get("/google_login", controllers.GoogleLogin)
		app.Get("/google_callback", controllers.GoogleCallback)
    app.Listen(":3000")

}