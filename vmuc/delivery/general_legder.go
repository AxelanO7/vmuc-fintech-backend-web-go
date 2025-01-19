package delivery

import (
	"strconv"
	"vmuc-fintech-backend-web-go/domain"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type GeneralLedgerHandler struct {
	GeneralLedgerUC domain.GeneralLedgerUseCase
}

func NewGeneralLedgerHandler(c *fiber.App, das domain.GeneralLedgerUseCase) {
	handler := &GeneralLedgerHandler{
		GeneralLedgerUC: das,
	}
	api := c.Group("/general-ledger")

	_ = api.Group("/public")

	private := api.Group("/private")
	private.Get("/employee", handler.GetAllGeneralLedger)
	private.Get("/employee/:id", handler.GetGeneralLedgerByID)
	private.Post("/employee", handler.CreateGeneralLedger)
	private.Post("/employees", handler.CreateBulkGeneralLedger)
	private.Put("/employee/:id", handler.UpdateGeneralLedger)
	private.Put("/employees", handler.UpdateBulkGeneralLedger)
	private.Delete("/employee/:id", handler.DeleteGeneralLedger)
}

func (t *GeneralLedgerHandler) GetAllGeneralLedger(c *fiber.Ctx) error {
	res, err := t.GeneralLedgerUC.FetchGeneralLedgers(c.Context())
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

func (t *GeneralLedgerHandler) GetGeneralLedgerByID(c *fiber.Ctx) error {
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
	res, err := t.GeneralLedgerUC.FetchGeneralLedgerByID(c.Context(), uint(strId))
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

func (t *GeneralLedgerHandler) CreateGeneralLedger(c *fiber.Ctx) error {
	req := new(domain.GeneralLedger)
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
	res, err := t.GeneralLedgerUC.AddGeneralLedger(c.Context(), req)
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

func (t *GeneralLedgerHandler) CreateBulkGeneralLedger(c *fiber.Ctx) error {
	req := new([]domain.GeneralLedger)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	var payrolls []*domain.GeneralLedger
	for _, payroll := range *req {
		payroll := payroll
		payrolls = append(payrolls, &payroll)
	}
	res, err := t.GeneralLedgerUC.AddBulkGeneralLedger(c.Context(), payrolls)
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

func (t *GeneralLedgerHandler) UpdateGeneralLedger(c *fiber.Ctx) error {
	req := new(domain.GeneralLedger)
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
	res, err := t.GeneralLedgerUC.EditGeneralLedger(c.Context(), req)
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

func (t *GeneralLedgerHandler) UpdateBulkGeneralLedger(c *fiber.Ctx) error {
	req := new([]domain.GeneralLedger)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	var payrolls []*domain.GeneralLedger
	for _, payroll := range *req {
		payroll := payroll
		payrolls = append(payrolls, &payroll)
	}
	res, err := t.GeneralLedgerUC.EditBulkGeneralLedger(c.Context(), payrolls)
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

func (t *GeneralLedgerHandler) DeleteGeneralLedger(c *fiber.Ctx) error {
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
	err := t.GeneralLedgerUC.DeleteGeneralLedger(c.Context(), uint(strId))
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
