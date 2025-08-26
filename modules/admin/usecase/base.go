package adminUsecase

import (
	"github.com/burabatbold/delivery-auth-service/database"
	"github.com/burabatbold/delivery-auth-service/database/entity"
	adminRepository "github.com/burabatbold/delivery-auth-service/modules/admin/repository"
)

type AdminUsecaseInterface interface {
	Get(id uint) (*entity.AdminEntity, error)
}

type adminUsecase struct {
	adminRepository adminRepository.AdminRepositoryInterface
}

func NewAdminUsecase() AdminUsecaseInterface {
	return &adminUsecase{adminRepository: adminRepository.NewAdminRepository(database.DB)}
}

func (u *adminUsecase) Get(id uint) (*entity.AdminEntity, error) {
	return u.adminRepository.FindByID(id)
}
