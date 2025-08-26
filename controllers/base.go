package controllers

import (
	"github.com/burabatbold/delivery-auth-service/common"
	"github.com/burabatbold/delivery-auth-service/controllers/admin"
	"github.com/burabatbold/delivery-auth-service/database"
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {

	bc := common.BaseController{
		DB:       database.DB,
		Response: common.Response{},
	}

	admin.NewAdminController(bc).Init(app.Group("/admin"))
}
