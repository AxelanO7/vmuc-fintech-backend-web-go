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
	out.Get("/", handler.ShowOuts)
	out.Get("/:id", handler.ShowOutById)
	out.Post("/", handler.AddOut)
	out.Put("/:id", handler.EditOutById)
	out.Delete("/:id", handler.DeleteOutById)

	last := api.Group("/last")
	outLast := last.Group("/out")
	outLast.Get("/", handler.GetLastOutNumber)
}

func (t *OutHandler) ShowOuts(c *fiber.Ctx) error {
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

func (t *OutHandler) ShowOutById(c *fiber.Ctx) error {
	id := c.Params("id")
	res, er := t.OutUC.ShowOutById(c.Context(), id)
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

func (t *OutHandler) GetLastOutNumber(c *fiber.Ctx) error {
	res, er := t.OutUC.ShowOutLastNumber(c.Context())
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

func (t *OutHandler) AddOut(c *fiber.Ctx) error {
	var out domain.Out
	c.BodyParser(&out)
	res, er := t.OutUC.AddOut(c.Context(), out)
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

func (t *OutHandler) EditOutById(c *fiber.Ctx) error {
	var out domain.Out
	c.BodyParser(&out)
	res, er := t.OutUC.EditOutById(c.Context(), out)
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

func (t *OutHandler) DeleteOutById(c *fiber.Ctx) error {
	id := c.Params("id")
	er := t.OutUC.DeleteOutById(c.Context(), id)
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
		"data":    nil,
		"message": "Success delete data",
	})
}
