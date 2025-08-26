package admin

import (
	"github.com/burabatbold/delivery-auth-service/common"
	adminMiddleware "github.com/burabatbold/delivery-auth-service/controllers/middlewares/admin"
	adminDto "github.com/burabatbold/delivery-auth-service/modules/admin/dto"
	adminUsecase "github.com/burabatbold/delivery-auth-service/modules/admin/usecase"
	"github.com/gofiber/fiber/v2"
)

type AdminController struct {
	common.BaseController
}

func NewAdminController(bc common.BaseController) *AdminController {
	return &AdminController{BaseController: bc}
}

func (co *AdminController) Init(app fiber.Router) {
	app.Post("/login", co.Login)
	app.Get("/me", adminMiddleware.AuthMiddleware(), co.Me)
}

func (co *AdminController) Login(c *fiber.Ctx) error {
	defer func() {
		co.GetBody(c)
	}()

	var dto adminDto.LoginDto
	if err := c.BodyParser(&dto); err != nil {
		return co.SetError(fiber.StatusBadRequest, err)
	}

	if err := dto.Validate(); err != nil {
		return co.SetError(fiber.StatusBadRequest, err)
	}

	admin, err := adminUsecase.NewAdminAuthUsecase().Login(dto)
	if err != nil {
		return co.SetError(fiber.StatusUnauthorized, err)
	}

	return co.SetBody(admin)
}

func (co *AdminController) Me(c *fiber.Ctx) error {
	defer func() {
		co.GetBody(c)
	}()

	admin := adminMiddleware.GetAuth(c)

	return co.SetBody(admin)
}
