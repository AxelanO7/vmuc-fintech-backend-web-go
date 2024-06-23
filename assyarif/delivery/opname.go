package delivery

import (
	"assyarif-backend-web-go/domain"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type OpnameHandler struct {
	OpnameUC domain.OpnameUseCase
}

func NewOpnameHandler(c *fiber.App, das domain.OpnameUseCase) {
	handler := &OpnameHandler{
		OpnameUC: das,
	}
	api := c.Group("/opname")

	// _ = api.Group("/public")

	private := api.Group("/private")
	private.Post("/stuff", handler.AddOpname)
	private.Get("/stuff", handler.GetAllOpname)
	private.Get("/stuff/:id", handler.GetOpnameByID)
	private.Post("/date", handler.GetOpnameByDate)
}

func (t *OpnameHandler) GetAllOpname(c *fiber.Ctx) error {
	res, err := t.OpnameUC.FetchOpnames(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": err,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Successfully get all user",
	})
}

func (t *OpnameHandler) GetOpnameByID(c *fiber.Ctx) error {
	id := c.Params("id")
	strId, erStr := strconv.Atoi(id)
	if erStr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse id",
			"error":   erStr.Error(),
		})
	}
	res, err := t.OpnameUC.FetchOpnameByID(c.Context(), uint(strId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": err,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Successfully get user by id",
	})
}

func (t *OpnameHandler) GetOpnameByDate(c *fiber.Ctx) error {
	req := new(domain.ReqByDate)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  422,
			"success": false,
			"message": "Failed to parse request",
			"error":   err.Error(),
		})
	}
	startDate := req.StartDate
	endDate := req.EndDate
	res, err := t.OpnameUC.FetchOpnameByDate(c.Context(), startDate, endDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": err,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Successfully get user by date",
	})
}

func (t *OpnameHandler) AddOpname(c *fiber.Ctx) error {
	req := new(domain.Opname)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  422,
			"success": false,
			"message": "Failed to parse request",
			"error":   err.Error(),
		})
	}
	err := t.OpnameUC.AddOpname(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": err,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  201,
		"success": true,
		"message": "Successfully add user",
	})
}
