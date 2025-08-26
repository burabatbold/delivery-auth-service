package adminRepository

import (
	"strings"

	"github.com/burabatbold/delivery-auth-service/database/entity"

	"gorm.io/gorm"
)

type AdminRepositoryInterface interface {
	Create(dto *entity.AdminEntity) (*entity.AdminEntity, error)
	FindByEmail(email string) (*entity.AdminEntity, error)
	FindByID(id uint) (*entity.AdminEntity, error)
}
type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepositoryInterface {
	return &adminRepository{db: db}
}

func (r *adminRepository) FindByID(id uint) (*entity.AdminEntity, error) {
	var admin entity.AdminEntity
	err := r.db.Where("id = ?", id).First(&admin).Error
	return &admin, err
}

func (r *adminRepository) Create(admin *entity.AdminEntity) (*entity.AdminEntity, error) {

	err := r.db.Create(&admin).Error

	return admin, err
}

func (r *adminRepository) FindByEmail(email string) (*entity.AdminEntity, error) {
	var admin entity.AdminEntity
	err := r.db.Where("LOWER(email) = ?", strings.ToLower(email)).First(&admin).Error
	return &admin, err
}
