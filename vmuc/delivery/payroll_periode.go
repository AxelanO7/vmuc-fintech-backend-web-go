package delivery

import (
	"strconv"
	"vmuc-fintech-backend-web-go/domain"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type PayrollPeriodeHandler struct {
	PayrollPeriodeUC domain.PayrollPeriodeUseCase
}

func NewPayrollPeriodeHandler(c *fiber.App, das domain.PayrollPeriodeUseCase) {
	handler := &PayrollPeriodeHandler{
		PayrollPeriodeUC: das,
	}
	api := c.Group("/payroll-periode")

	_ = api.Group("/public")

	private := api.Group("/private")
	private.Get("/employee", handler.GetAllPayrollPeriode)
	private.Get("/employee/:id", handler.GetPayrollPeriodeByID)
	private.Post("/employee", handler.CreatePayrollPeriode)
	private.Post("/employees", handler.CreateBulkPayrollPeriode)
	private.Put("/employee/:id", handler.UpdatePayrollPeriode)
	private.Put("/employees", handler.UpdateBulkPayrollPeriode)
	private.Delete("/employee/:id", handler.DeletePayrollPeriode)
}

func (t *PayrollPeriodeHandler) GetAllPayrollPeriode(c *fiber.Ctx) error {
	res, err := t.PayrollPeriodeUC.FetchPayrollPeriode(c.Context())
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
		"message": "Successfully get all payrollPeriode",
	})
}

func (t *PayrollPeriodeHandler) GetPayrollPeriodeByID(c *fiber.Ctx) error {
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
	res, err := t.PayrollPeriodeUC.FetchPayrollPeriodeByID(c.Context(), uint(strId))
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
		"message": "Successfully get payrollPeriode by id",
	})
}

func (t *PayrollPeriodeHandler) CreatePayrollPeriode(c *fiber.Ctx) error {
	req := new(domain.PayrollPeriode)
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
	res, err := t.PayrollPeriodeUC.AddPayrollPeriode(c.Context(), req)
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
		"message": "Successfully create payrollPeriode",
	})
}

func (t *PayrollPeriodeHandler) CreateBulkPayrollPeriode(c *fiber.Ctx) error {
	req := new([]domain.PayrollPeriode)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	var payrollPeriodes []*domain.PayrollPeriode
	for _, payrollPeriode := range *req {
		payrollPeriode := payrollPeriode
		payrollPeriodes = append(payrollPeriodes, &payrollPeriode)
	}
	res, err := t.PayrollPeriodeUC.AddBulkPayrollPeriode(c.Context(), payrollPeriodes)
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
		"message": "Successfully create bulk payrollPeriode",
	})
}

func (t *PayrollPeriodeHandler) UpdatePayrollPeriode(c *fiber.Ctx) error {
	req := new(domain.PayrollPeriode)
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
	res, err := t.PayrollPeriodeUC.EditPayrollPeriode(c.Context(), req)
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
		"message": "Successfully update payrollPeriode",
	})
}

func (t *PayrollPeriodeHandler) UpdateBulkPayrollPeriode(c *fiber.Ctx) error {
	req := new([]domain.PayrollPeriode)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	var payrollPeriodes []*domain.PayrollPeriode
	for _, payrollPeriode := range *req {
		payrollPeriode := payrollPeriode
		payrollPeriodes = append(payrollPeriodes, &payrollPeriode)
	}
	res, err := t.PayrollPeriodeUC.EditBulkPayrollPeriode(c.Context(), payrollPeriodes)
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
		"message": "Successfully update bulk payrollPeriode",
	})
}

func (t *PayrollPeriodeHandler) DeletePayrollPeriode(c *fiber.Ctx) error {
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
	err := t.PayrollPeriodeUC.DeletePayrollPeriode(c.Context(), uint(strId))
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
		"message": "Successfully delete payrollPeriode",
	})
}
