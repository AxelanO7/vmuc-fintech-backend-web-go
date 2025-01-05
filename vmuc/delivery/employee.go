package delivery

import (
	"strconv"
	"vmuc-fintech-backend-web-go/domain"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct {
	EmployeeUC domain.EmployeeUseCase
}

func NewEmployeeHandler(c *fiber.App, das domain.EmployeeUseCase) {
	handler := &EmployeeHandler{
		EmployeeUC: das,
	}
	api := c.Group("/employee")

	_ = api.Group("/public")

	private := api.Group("/private")
	private.Get("/account", handler.GetAllEmployee)
	private.Get("/account/:id", handler.GetEmployeeByID)
	private.Post("/account", handler.CreateEmployee)
	private.Put("/account/:id", handler.UpdateEmployee)
	private.Delete("/account/:id", handler.DeleteEmployee)

	private.Get("/last", handler.ShowEmployeeLastNumber)
}

func (t *EmployeeHandler) GetAllEmployee(c *fiber.Ctx) error {
	res, err := t.EmployeeUC.FetchEmployees(c.Context())
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

func (t *EmployeeHandler) GetEmployeeByID(c *fiber.Ctx) error {
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
	res, err := t.EmployeeUC.FetchEmployeeByID(c.Context(), uint(strId))
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

func (t *EmployeeHandler) CreateEmployee(c *fiber.Ctx) error {
	req := new(domain.Employee)
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
	res, err := t.EmployeeUC.CreateEmployee(c.Context(), req)
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
		"message": "Successfully create user",
	})
}

func (t *EmployeeHandler) UpdateEmployee(c *fiber.Ctx) error {
	req := new(domain.Employee)
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
	res, err := t.EmployeeUC.UpdateEmployee(c.Context(), req)
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
		"message": "Successfully update user",
	})
}

func (t *EmployeeHandler) DeleteEmployee(c *fiber.Ctx) error {
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
	err := t.EmployeeUC.DeleteEmployee(c.Context(), uint(strId))
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
		"message": "Successfully delete user",
	})
}

func (t *EmployeeHandler) ShowEmployeeLastNumber(c *fiber.Ctx) error {
	res, err := t.EmployeeUC.ShowEmployeeLastNumber(c.Context())
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
		"message": "Successfully get last number",
	})
}
