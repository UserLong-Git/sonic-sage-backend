package main

import (
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/daniellong/sonic-sage-backend/routes"
)

func main() {
    app := fiber.New()

    app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"status": "ok"})
    })

    routes.RegisterAnalyzeRoutes(app)
    routes.RegisterWaveformRoutes(app)
    routes.RegisterMetricsRoutes(app)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Starting Sonic Sage backend on port %s\n", port)
    if err := app.Listen(":" + port); err != nil {
        log.Fatal(err)
    }
}
