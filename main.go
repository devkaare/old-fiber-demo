package main

import "github.com/gofiber/fiber"

// Dummy handler
func handler() func(*fiber.Ctx) {
        return func(c *fiber.Ctx) {
                c.Send("This is a dummy route")
        }
}

func main() {
        app := fiber.New()

        // Create new sample GET routes
        app.Get("/demo", handler())
        app.Get("/list", handler())

        // Middleware to match anything
        app.Use(func(c *fiber.Ctx) {
                c.SendStatus(404) // => 404 "Not Found"
        })

        app.Listen(3000)
}
