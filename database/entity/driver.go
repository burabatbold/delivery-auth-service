package entity

type DriverStatus string

const (
	DriverStatusActive   DriverStatus = "active"
	DriverStatusInactive DriverStatus = "inactive"
)

type (
	DriverEntity struct {
		BaseEntity
		Email      string       `gorm:"column:email;index" json:"email"`
		Password   string       `gorm:"column:password" json:"password"`
		Phone      string       `gorm:"column:phone" json:"phone"`
		Status     DriverStatus `gorm:"column:status;index" json:"status"`
		IsVerified bool         `gorm:"column:is_verified" json:"is_verified"`
	}
)

func (DriverEntity) TableName() string {
	return "drivers"
}
