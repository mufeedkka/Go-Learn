package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	fmt.Println("Hello world")

	app := fiber.New()

	todos := []Todo{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	app.Patch("/api/updt/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(401).SendString("Invalid Id")
		}
		for i, t := range todos {
			if t.ID == id {
				todos[i].Done = true
				break
			}
		}
		return c.JSON(todos)

	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		err := c.BodyParser(todo)

		if err != nil {
			return err
		}
		todo.ID = len(todos) + 1

		todos = append(todos, *todo)

		return c.JSON(todos)
	})

	app.Get("/api/get", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})
	app.Get("/api/get/true", func(c *fiber.Ctx) error {
		todotrue := []Todo{}
		for _, data := range todos {
			if data.Done == true {
				todotrue = append(todotrue, data)
			}
		}
		return c.JSON(todotrue)
	})

	log.Fatal(app.Listen(":8001"))

}
