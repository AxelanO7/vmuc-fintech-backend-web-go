package delivery

import (
	"strconv"
	"vmuc-fintech-backend-web-go/domain"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type GeneralJournalHandler struct {
	GeneralJournalUC domain.GeneralJournalUseCase
}

func NewGeneralJournalHandler(c *fiber.App, das domain.GeneralJournalUseCase) {
	handler := &GeneralJournalHandler{
		GeneralJournalUC: das,
	}
	api := c.Group("/journal")

	_ = api.Group("/public")

	private := api.Group("/private")
	private.Get("/general", handler.GetAllGeneralJournal)
	private.Get("/general/:id", handler.GetGeneralJournalByID)
	private.Post("/general", handler.CreateGeneralJournal)
	private.Post("/generals", handler.CreateBulkGeneralJournal)
	private.Put("/general/:id", handler.UpdateGeneralJournal)
	private.Put("/generals", handler.UpdateBulkGeneralJournal)
	private.Delete("/general/:id", handler.DeleteGeneralJournal)
}

func (t *GeneralJournalHandler) GetAllGeneralJournal(c *fiber.Ctx) error {
	res, err := t.GeneralJournalUC.FetchGeneralJournals(c.Context())
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

func (t *GeneralJournalHandler) GetGeneralJournalByID(c *fiber.Ctx) error {
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
	res, err := t.GeneralJournalUC.FetchGeneralJournalByID(c.Context(), uint(strId))
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

func (t *GeneralJournalHandler) CreateGeneralJournal(c *fiber.Ctx) error {
	req := new(domain.GeneralJournal)
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
	res, err := t.GeneralJournalUC.AddGeneralJournal(c.Context(), req)
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

func (t *GeneralJournalHandler) CreateBulkGeneralJournal(c *fiber.Ctx) error {
	req := new([]domain.GeneralJournal)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	var payrolls []*domain.GeneralJournal
	for _, payroll := range *req {
		payroll := payroll
		payrolls = append(payrolls, &payroll)
	}
	res, err := t.GeneralJournalUC.AddBulkGeneralJournal(c.Context(), payrolls)
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

func (t *GeneralJournalHandler) UpdateGeneralJournal(c *fiber.Ctx) error {
	req := new(domain.GeneralJournal)
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
	res, err := t.GeneralJournalUC.EditGeneralJournal(c.Context(), req)
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

func (t *GeneralJournalHandler) UpdateBulkGeneralJournal(c *fiber.Ctx) error {
	req := new([]domain.GeneralJournal)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	var payrolls []*domain.GeneralJournal
	for _, payroll := range *req {
		payroll := payroll
		payrolls = append(payrolls, &payroll)
	}
	res, err := t.GeneralJournalUC.EditBulkGeneralJournal(c.Context(), payrolls)
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

func (t *GeneralJournalHandler) DeleteGeneralJournal(c *fiber.Ctx) error {
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
	err := t.GeneralJournalUC.DeleteGeneralJournal(c.Context(), uint(strId))
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
