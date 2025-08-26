package adminDto

import (
	"errors"

	"github.com/burabatbold/delivery-auth-service/database/entity"

	"golang.org/x/crypto/bcrypt"
)

type (
	LoginDto struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RegisterDto struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func (dto *LoginDto) Validate() error {
	if dto.Email == "" {
		return errors.New("email is required")
	}
	if dto.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func (dto *RegisterDto) Validate() error {
	if dto.Email == "" {
		return errors.New("email is required")
	}
	if dto.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func (dto *RegisterDto) ToEntity() *entity.AdminEntity {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	return &entity.AdminEntity{
		Email:    dto.Email,
		Password: string(hashedPassword),
	}
}
