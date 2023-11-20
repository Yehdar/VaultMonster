package main

import (
  "encoding/json"
  "log"

  "github.com/gofiber/fiber/v2"
 	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct{
  ID    int     `json:"id"`
  Title string  `json:"title"`
  Done  bool    `json:"done"`
  Body  string  `json:"body"`
}

func main() {
  app := fiber.New()

  app.Use(cors.New(cors.Config{
    AllowOrigins: "http://localhost:5173",
    AllowHeaders: "Origin, Content-Type, Accept",
  }))

  app.Options("/*", func(c *fiber.Ctx) error {
    return c.SendStatus(fiber.StatusNoContent)
  })


  todos := []Todo{}

  app.Get("/healthcheck", func(c *fiber.Ctx) error {
    return c.SendString("OK")
  })

  app.Post("/api/todos", func(c *fiber.Ctx) error {
    body := c.Body()

    todo := &Todo{}

    if err := json.Unmarshal(body, todo); err != nil {
      log.Println("Error parsing request body:", err)
      return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
    }

    todo.ID = len(todos) + 1
    todos = append(todos, *todo)

    return c.JSON(todos)
  })

  app.Patch("/api/todos/:id/done",  func(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")

    if err != nil {
      return c.Status(401).SendString("Invalid ID")
    }

    for i, t := range todos {
      if t.ID == id {
        todos[i].Done = true
        break
      }
    }

    return c.JSON(todos)
  })

  app.Get("/api/todos", func(c *fiber.Ctx) error {
    return c.JSON(todos)
  })

  log.Fatal(app.Listen(":4000"))
}
