package adminUsecase

import (
	"errors"

	"github.com/burabatbold/delivery-auth-service/database"
	"github.com/burabatbold/delivery-auth-service/database/entity"
	adminDto "github.com/burabatbold/delivery-auth-service/modules/admin/dto"
	adminRepository "github.com/burabatbold/delivery-auth-service/modules/admin/repository"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"

	"golang.org/x/crypto/bcrypt"
)

type AdminAuthUsecaseInterface interface {
	Register(dto adminDto.RegisterDto) (*entity.AdminEntity, error)
	Login(dto adminDto.LoginDto) (*adminDto.LoginOutputDto, error)
	VerifyToken(param string) error
}

type adminAuthUsecase struct {
	adminRepository adminRepository.AdminRepositoryInterface
}

func NewAdminAuthUsecase() AdminAuthUsecaseInterface {
	return &adminAuthUsecase{adminRepository: adminRepository.NewAdminRepository(database.DB)}
}

func (u *adminAuthUsecase) Register(dto adminDto.RegisterDto) (*entity.AdminEntity, error) {
	return u.adminRepository.Create(dto.ToEntity())
}

func (u *adminAuthUsecase) Login(dto adminDto.LoginDto) (*adminDto.LoginOutputDto, error) {
	admin, err := u.adminRepository.FindByEmail(dto.Email)
	if err != nil {
		return nil, err
	}

	if admin == nil {
		return nil, errors.New("Хэрэглэгчийн имэйл хаяг эсвэл нууц үг буруу байна")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(dto.Password)); err != nil {
		return nil, errors.New("Хэрэглэгчийн имэйл хаяг эсвэл нууц үг буруу байна")
	}

	return adminDto.ToLoginOutputDto(admin), nil
}

func (u *adminAuthUsecase) VerifyToken(param string) error {

	token, err := jwt.Parse(param, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt.secret")), nil
	})
	if err != nil {
		return errors.New("token is invalid")
	}

	if !token.Valid {
		return errors.New("token is invalid")
	}

	id, ok := token.Claims.(jwt.MapClaims)["admin_id"].(float64)
	if !ok {
		return errors.New("token is invalid")
	}

	if _, err := u.adminRepository.FindByID(uint(id)); err != nil {
		return errors.New("token is invalid")
	}

	return nil
}
