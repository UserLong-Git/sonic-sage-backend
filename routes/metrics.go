package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/daniellong/sonic-sage-backend/models"
)

func RegisterMetricsRoutes(app *fiber.App) {
    app.Post("/metrics", metricsHandler)
}

func metricsHandler(c *fiber.Ctx) error {
    var req models.MetricsRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "invalid request body",
        })
    }

    resp := models.MetricsResponse{
        Message: "Metrics placeholder",
        Metrics: map[string]float64{
            "lufs":        -12.3,
            "rms":         -9.8,
            "peakDb":      -1.2,
            "crestFactor": 3.1,
        },
    }

    return c.JSON(resp)
}
