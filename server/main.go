package main

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
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

	app.Post("/api/upload", func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		files := form.File["file"]

		for _, file := range files {
			filename := filepath.Base(file.Filename)
			if err := c.SaveFile(file, "./uploads/"+filename); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
			return c.JSON(fiber.Map{"fileName": filename})
		}

		return c.SendStatus(fiber.StatusBadRequest)
	})

	app.Get("/api/download/:filename", func(c *fiber.Ctx) error {
		filename := c.Params("filename")
		filepath := "./uploads/" + filename

		return c.SendFile(filepath)
	})

	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
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

	// Listen on port 4000
	if err := app.Listen(":4000"); err != nil {
		panic(err)
	}
}
