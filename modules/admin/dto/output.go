package adminDto

import (
	"time"

	"github.com/burabatbold/delivery-auth-service/database/entity"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type (
	LoginOutputDto struct {
		Token string              `json:"token"`
		Admin *entity.AdminEntity `json:"admin"`
	}
)

func ToLoginOutputDto(admin *entity.AdminEntity) *LoginOutputDto {

	if admin == nil {
		return nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin_id": admin.ID,
		"email":    admin.Email,
		"role":     "admin",
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		return nil
	}

	return &LoginOutputDto{
		Token: tokenString,
		Admin: admin,
	}
}
