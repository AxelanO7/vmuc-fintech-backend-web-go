package delivery

import (
	"strconv"
	"vmuc-fintech-backend-web-go/domain"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type PeriodeHandler struct {
	PeriodeUC domain.PeriodeUseCase
}

func NewPeriodeHandler(c *fiber.App, das domain.PeriodeUseCase) {
	handler := &PeriodeHandler{
		PeriodeUC: das,
	}
	api := c.Group("/periode")

	_ = api.Group("/public")

	private := api.Group("/private")
	private.Get("/payroll-employee", handler.GetAllPayrollPeriode)
	private.Get("/adjusment-entries", handler.GetAllAdjusmentEntriesPeriode)
	private.Get("/general-journal", handler.GetAllGeneralJournalPeriode)
	private.Get("/trial-balance", handler.GetAllTrialBalancePeriode)
	private.Get("/get-report-trial-balance/:periode", handler.GetTrialBalanceReportByPeriode)

	general := private.Group("/general")
	general.Get("/:id", handler.GetPeriodeByID)
	general.Post("/", handler.CreatePeriode)
	general.Post("/", handler.CreateBulkPeriode)
	general.Put("/:id", handler.UpdatePeriode)
	general.Put("/", handler.UpdateBulkPeriode)
	general.Delete("/:id", handler.DeletePeriode)

	payroll := private.Group("/payroll")
	payroll.Post("/", handler.CreatePayrollWithPeriode)
}

func (t *PeriodeHandler) GetAllPayrollPeriode(c *fiber.Ctx) error {
	res, err := t.PeriodeUC.FetchPayrollPeriode(c.Context())
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
		"message": "Successfully get all Periode",
	})
}

func (t *PeriodeHandler) GetAllAdjusmentEntriesPeriode(c *fiber.Ctx) error {
	res, err := t.PeriodeUC.FetchAdjusmentEntriesPeriode(c.Context())
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
		"message": "Successfully get all Periode",
	})
}

func (t *PeriodeHandler) GetAllGeneralJournalPeriode(c *fiber.Ctx) error {
	res, err := t.PeriodeUC.FetchGeneralJournalPeriode(c.Context())
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
		"message": "Successfully get all Periode",
	})
}

func (t *PeriodeHandler) GetAllTrialBalancePeriode(c *fiber.Ctx) error {
	res, err := t.PeriodeUC.FetchTrialBalancePeriode(c.Context())
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
		"message": "Successfully get all Periode",
	})
}

func (t *PeriodeHandler) GetTrialBalanceReportByPeriode(c *fiber.Ctx) error {
	periode := c.Params("periode")
	res, err := t.PeriodeUC.GetTrialBalanceReportByPeriode(c.Context(), periode)
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
		"message": "Successfully get Periode by id",
	})
}

func (t *PeriodeHandler) GetPeriodeByID(c *fiber.Ctx) error {
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
	res, err := t.PeriodeUC.FetchPeriodeByID(c.Context(), uint(strId))
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
		"message": "Successfully get Periode by id",
	})
}

func (t *PeriodeHandler) CreatePeriode(c *fiber.Ctx) error {
	req := new(domain.Periode)
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
	res, err := t.PeriodeUC.AddPeriode(c.Context(), req)
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
		"message": "Successfully create Periode",
	})
}

func (t *PeriodeHandler) CreateBulkPeriode(c *fiber.Ctx) error {
	req := new([]domain.Periode)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	var Periodes []*domain.Periode
	for _, Periode := range *req {
		Periode := Periode
		Periodes = append(Periodes, &Periode)
	}
	res, err := t.PeriodeUC.AddBulkPeriode(c.Context(), Periodes)
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
		"message": "Successfully create bulk Periode",
	})
}

func (t *PeriodeHandler) UpdatePeriode(c *fiber.Ctx) error {
	req := new(domain.Periode)
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
	res, err := t.PeriodeUC.EditPeriode(c.Context(), req)
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
		"message": "Successfully update Periode",
	})
}

func (t *PeriodeHandler) UpdateBulkPeriode(c *fiber.Ctx) error {
	req := new([]domain.Periode)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	var Periodes []*domain.Periode
	for _, Periode := range *req {
		Periode := Periode
		Periodes = append(Periodes, &Periode)
	}
	res, err := t.PeriodeUC.EditBulkPeriode(c.Context(), Periodes)
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
		"message": "Successfully update bulk Periode",
	})
}

func (t *PeriodeHandler) DeletePeriode(c *fiber.Ctx) error {
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
	err := t.PeriodeUC.DeletePeriode(c.Context(), uint(strId))
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
		"message": "Successfully delete Periode",
	})
}

func (t *PeriodeHandler) CreatePayrollWithPeriode(c *fiber.Ctx) error {
	req := new(domain.PayrollPeriode)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err.Error(),
		})
	}
	valRes, er := govalidator.ValidateStruct(req)
	if !valRes {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to validate body",
			"error":   er.Error(),
		})
	}
	res, err := t.PeriodeUC.AddPayrollPeriode(c.Context(), req)
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
		"message": "Successfully create Payroll Periode",
	})
}
