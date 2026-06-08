package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/daniellong/sonic-sage-backend/models"
)

func RegisterWaveformRoutes(app *fiber.App) {
    app.Post("/waveform", waveformHandler)
}

func waveformHandler(c *fiber.Ctx) error {
    var req models.WaveformRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "invalid request body",
        })
    }

    resp := models.WaveformResponse{
        Message:    "Waveform placeholder",
        Samples:    []float64{0.0, 0.2, 0.5, -0.3, -0.1, 0.4},
        SampleRate: 44100,
    }

    return c.JSON(resp)
}
