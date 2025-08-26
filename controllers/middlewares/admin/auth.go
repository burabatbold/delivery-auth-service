package adminMiddleware

import (
	"github.com/burabatbold/delivery-auth-service/common"
	"github.com/burabatbold/delivery-auth-service/database/entity"
	adminUsecase "github.com/burabatbold/delivery-auth-service/modules/admin/usecase"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

const AdminKey = "admin"

func AuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(viper.GetString("jwt.secret")),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(common.Response{
				Code:    fiber.StatusUnauthorized,
				Message: "Нэвтрэх нэр эсвэл нууц үг буруу байна",
			})
		},
		SuccessHandler: func(c *fiber.Ctx) error {

			token := c.Locals("user").(*jwt.Token)

			if !token.Valid {
				return c.Status(fiber.StatusUnauthorized).JSON(common.Response{
					Code:    fiber.StatusUnauthorized,
					Message: "Таны идэвхитэй байх хугацаа дууссан байна",
				})
			}

			id, ok := token.Claims.(jwt.MapClaims)["admin_id"].(float64)
			if !ok {
				return c.Status(fiber.StatusUnauthorized).JSON(common.Response{
					Code:    fiber.StatusUnauthorized,
					Message: "Нэвтрэх нэр эсвэл нууц үг буруу байна",
				})
			}

			admin, err := adminUsecase.NewAdminUsecase().Get(uint(id))
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(common.Response{
					Code:    fiber.StatusUnauthorized,
					Message: "Нэвтрэх нэр эсвэл нууц үг буруу байна",
				})
			}

			c.Locals(AdminKey, admin)

			return c.Next()
		},
	})
}

func GetAuth(c *fiber.Ctx) *entity.AdminEntity {
	return c.Locals(AdminKey).(*entity.AdminEntity)
}
