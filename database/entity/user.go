package entity

type (
	CustomerEntity struct {
		BaseEntity
		Username string `gorm:"column:username;index" json:"username"`
		Password string `gorm:"column:password" json:"password"`
		Email    string `gorm:"column:email;index" json:"email"`
		Phone    string `gorm:"column:phone" json:"phone"`
		Address  string `gorm:"column:address" json:"address"`
	}
)

func (CustomerEntity) TableName() string {
	return "customers"
}
