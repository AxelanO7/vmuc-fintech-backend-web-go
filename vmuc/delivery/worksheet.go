package delivery

import (
	"strconv"
	"vmuc-fintech-backend-web-go/domain"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type WorksheetHandler struct {
	WorksheetUC domain.WorksheetUseCase
}

func NewWorksheetHandler(c *fiber.App, das domain.WorksheetUseCase) {
	handler := &WorksheetHandler{
		WorksheetUC: das,
	}
	api := c.Group("/journal")

	_ = api.Group("/public")

	private := api.Group("/private")
	private.Get("/work-sheet", handler.GetAllWorksheet)
	private.Get("/work-sheet/:id", handler.GetWorksheetByID)
	private.Get("/report-worksheet/:id", handler.GetWorksheetByID)
	private.Post("/work-sheet", handler.CreateWorksheet)
	private.Post("/work-sheets", handler.CreateBulkWorksheet)
	private.Put("/work-sheet/:id", handler.UpdateWorksheet)
	private.Put("/work-sheets", handler.UpdateBulkWorksheet)
	private.Delete("/work-sheet/:id", handler.DeleteWorksheet)
}

func (t *WorksheetHandler) GetAllWorksheet(c *fiber.Ctx) error {
	res, err := t.WorksheetUC.FetchWorksheets(c.Context(), true)
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
		"message": "Successfully get all payroll",
	})
}

func (t *WorksheetHandler) GetWorksheetByID(c *fiber.Ctx) error {
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
	res, err := t.WorksheetUC.FetchWorksheetByID(c.Context(), uint(strId), false)
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
		"message": "Successfully get payroll by id",
	})
}

func (t *WorksheetHandler) GetWorksheetByIDForReport(c *fiber.Ctx) error {
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
	res, err := t.WorksheetUC.FetchWorksheetByID(c.Context(), uint(strId), true)
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
		"message": "Successfully get payroll by id",
	})
}

func (t *WorksheetHandler) CreateWorksheet(c *fiber.Ctx) error {
	req := new(domain.Worksheet)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	valRes, er := govalidator.ValidateStruct(req)
	if !valRes {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   er.Error(),
		})
	}
	res, err := t.WorksheetUC.AddWorksheet(c.Context(), req)
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
		"data":    res,
		"message": "Successfully create payroll",
	})
}

func (t *WorksheetHandler) CreateBulkWorksheet(c *fiber.Ctx) error {
	req := new([]domain.Worksheet)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	var payrolls []*domain.Worksheet
	for _, payroll := range *req {
		payroll := payroll
		payrolls = append(payrolls, &payroll)
	}
	res, err := t.WorksheetUC.AddBulkWorksheet(c.Context(), payrolls)
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
		"data":    res,
		"message": "Successfully create bulk payroll",
	})
}

func (t *WorksheetHandler) UpdateWorksheet(c *fiber.Ctx) error {
	req := new(domain.Worksheet)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	valRes, er := govalidator.ValidateStruct(req)
	if !valRes {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   er.Error(),
		})
	}
	res, err := t.WorksheetUC.EditWorksheet(c.Context(), req)
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
		"message": "Successfully update payroll",
	})
}

func (t *WorksheetHandler) UpdateBulkWorksheet(c *fiber.Ctx) error {
	req := new([]domain.Worksheet)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	var payrolls []*domain.Worksheet
	for _, payroll := range *req {
		payroll := payroll
		payrolls = append(payrolls, &payroll)
	}
	res, err := t.WorksheetUC.EditBulkWorksheet(c.Context(), payrolls)
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
		"message": "Successfully update bulk payroll",
	})
}

func (t *WorksheetHandler) DeleteWorksheet(c *fiber.Ctx) error {
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
	err := t.WorksheetUC.DeleteWorksheet(c.Context(), uint(strId))
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
		"message": "Successfully delete payroll",
	})
}
