package delivery

import (
	"strconv"
	"vmuc-fintech-backend-web-go/domain"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type LedgerHandler struct {
	LedgerUC domain.LedgerUseCase
}

func NewLedgerHandler(c *fiber.App, das domain.LedgerUseCase) {
	handler := &LedgerHandler{
		LedgerUC: das,
	}
	api := c.Group("/general-ledger")

	_ = api.Group("/public")

	private := api.Group("/private")
	private.Get("/employee", handler.GetAllGeneralLedger)
	private.Get("/employee/:id", handler.GetGeneralLedgerByID)
	private.Get("/employee-report/:id", handler.GetGeneralLedgerByIDReport)
	private.Post("/employee", handler.CreateGeneralLedger)
	private.Post("/employees", handler.CreateBulkGeneralLedger)
	private.Put("/employee/:id", handler.UpdateGeneralLedger)
	private.Put("/employees", handler.UpdateBulkGeneralLedger)
	private.Delete("/employee/:id", handler.DeleteGeneralLedger)
}

func (t *LedgerHandler) GetAllGeneralLedger(c *fiber.Ctx) error {
	res, err := t.LedgerUC.FetchLedgers(c.Context())
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

func (t *LedgerHandler) GetGeneralLedgerByID(c *fiber.Ctx) error {
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
	res, err := t.LedgerUC.FetchLedgerByID(c.Context(), uint(strId), false)
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

func (t *LedgerHandler) GetGeneralLedgerByIDReport(c *fiber.Ctx) error {
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
	res, err := t.LedgerUC.FetchLedgerByID(c.Context(), uint(strId), true)
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

func (t *LedgerHandler) CreateGeneralLedger(c *fiber.Ctx) error {
	req := new(domain.Ledger)
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
	res, err := t.LedgerUC.AddLedger(c.Context(), req)
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

func (t *LedgerHandler) CreateBulkGeneralLedger(c *fiber.Ctx) error {
	req := new([]domain.Ledger)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	var payrolls []*domain.Ledger
	for _, payroll := range *req {
		payroll := payroll
		payrolls = append(payrolls, &payroll)
	}
	res, err := t.LedgerUC.AddBulkLedger(c.Context(), payrolls)
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

func (t *LedgerHandler) UpdateGeneralLedger(c *fiber.Ctx) error {
	req := new(domain.Ledger)
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
	res, err := t.LedgerUC.EditLedger(c.Context(), req)
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

func (t *LedgerHandler) UpdateBulkGeneralLedger(c *fiber.Ctx) error {
	req := new([]domain.Ledger)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	var payrolls []*domain.Ledger
	for _, payroll := range *req {
		payroll := payroll
		payrolls = append(payrolls, &payroll)
	}
	res, err := t.LedgerUC.EditBulkLedger(c.Context(), payrolls)
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

func (t *LedgerHandler) DeleteGeneralLedger(c *fiber.Ctx) error {
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
	err := t.LedgerUC.DeleteLedger(c.Context(), uint(strId))
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
