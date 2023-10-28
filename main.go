package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	
)


func main() {
	var lis[]string
	
	app:=fiber.New()

	app.Use(cors.New(cors.Config{AllowOrigins: "*",}))
	
	app.Static("/","a.html",)

	app.Get("/add",func(c *fiber.Ctx) error {
		lis = append(lis, c.Query("adatok"))
		println(c.OriginalURL())
		return c.SendString(c.Query("adatok"))
	})
	

	app.Get("/print",func(c *fiber.Ctx) error {
		return c.JSON(map[string][]string{"adatok":lis})
	})

	app.Listen(":8080")

}
