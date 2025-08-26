package entity

type (
	AdminEntity struct {
		BaseEntity
		Username string `gorm:"column:username;index" json:"username"`
		Password string `gorm:"column:password" json:"password"`
		Email    string `gorm:"column:email;index" json:"email"`
	}
)

func (AdminEntity) TableName() string {
	return "admins"
}
