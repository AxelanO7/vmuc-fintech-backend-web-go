package delivery

import (
	"assyarif-backend-web-go/domain"

	"github.com/gofiber/fiber/v2"
)

type OutHandler struct {
	OutUC domain.OutUseCase
}

func NewOutHandler(c *fiber.App, das domain.OutUseCase) {
	handler := &OutHandler{
		OutUC: das,
	}
	api := c.Group("/stuff")
	out := api.Group("/out")
	out.Get("/", handler.GetOuts)
}

func (t *OutHandler) GetOuts(c *fiber.Ctx) error {
	res, er := t.OutUC.GetOuts(c.Context())
	if er != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": er.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Success get data",
	})
}
