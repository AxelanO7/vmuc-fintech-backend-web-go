package delivery

import (
	"assyarif-backend-web-go/domain"

	"github.com/gofiber/fiber/v2"
)

type InHandler struct {
	InUC domain.InUseCase
}

func NewInHandler(c *fiber.App, das domain.InUseCase) {
	handler := &InHandler{
		InUC: das,
	}
	api := c.Group("/stuff")
	in := api.Group("/in")
	in.Get("/", handler.GetIns)
}

func (t *InHandler) GetIns(c *fiber.Ctx) error {
	res, er := t.InUC.GetIns(c.Context())
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
