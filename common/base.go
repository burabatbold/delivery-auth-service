package common

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BaseController struct {
	DB       *gorm.DB
	Response Response
}

// SetBody successfully response
func (co *BaseController) SetBody(body interface{}) error {
	co.Response.Code = fiber.StatusOK
	co.Response.Body = body

	return nil
}

// SetError error response
func (co *BaseController) SetError(code int, err error) error {
	co.Response.Code = code
	co.Response.Body = nil
	co.Response.Message = err.Error()

	return nil
}

func (co *BaseController) GetBody(c *fiber.Ctx) error {
	return c.Status(co.Response.Code).JSON(co.Response)
}
