package main

import (
  "fmt"
  "log"
  "github.com/gofiber/fiber/v2"
)

func main() {
  fmt.Print("hollup... let him cook")

  app := fiber.New()

  app.Get("/healthcheck", func(c *fiber.Ctx) error {
    return c.SendString("OK bruh")
  })

  log.Fatal(app.Listen(":4000"))
}
