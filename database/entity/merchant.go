package entity

type (
	MerchantEntity struct {
		BaseEntity
		Email      string `gorm:"column:email;index" json:"email"`
		Password   string `gorm:"column:password" json:"password"`
		Phone      string `gorm:"column:phone" json:"phone"`
		IsVerified bool   `gorm:"column:is_verified" json:"is_verified"`
	}
)

func (MerchantEntity) TableName() string {
	return "merchants"
}
