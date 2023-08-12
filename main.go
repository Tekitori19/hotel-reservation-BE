package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {

	listenAddr := flag.String("listen",":5000","listen address")
	flag.Parse()
	app := fiber.New()
	api :=  app.Group("api/v1")

	app.Get("/", handleHello)
	api.Get("/user",handleUser)
	app.Listen(*listenAddr)
}

func handleHello(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg":"hello world"})
}
func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg":"hello user"})
}