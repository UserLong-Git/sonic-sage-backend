package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/daniellong/sonic-sage-backend/models"
)

func RegisterAnalyzeRoutes(app *fiber.App) {
    app.Post("/analyze", analyzeHandler)
}

func analyzeHandler(c *fiber.Ctx) error {
    var req models.AnalyzeRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "invalid request body",
        })
    }

    resp := models.AnalyzeResponse{
        Message: "Sonic Sage analysis placeholder",
        Insights: map[string]any{
            "clarity":   0.87,
            "warmth":    0.72,
            "tightness": 0.81,
        },
    }

    return c.JSON(resp)
}
