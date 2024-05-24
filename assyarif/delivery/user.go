package delivery

import (
	"assyarif-backend-web-go/domain"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserUC domain.UserUseCase
}

func NewUserHandler(c *fiber.App, das domain.UserUseCase) {
	handler := &UserHandler{
		UserUC: das,
	}
	api := c.Group("/user")
	public := api.Group("/public")
	public.Post("/login", handler.Login)
}

func (t *UserHandler) Login(c *fiber.Ctx) error {
	req := new(domain.LoginPayload)
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
	res, token, er := t.UserUC.LoginUser(c.Context(), req)
	if er != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": er,
			"error":   er.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"token":   token,
		"message": "Successfully login",
	})
}
