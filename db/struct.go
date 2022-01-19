package db

type Students struct {
	Name       string `gorm:"column:name"`
	Email      string `gorm:"column:email"`
	Phone      string `gorm:"column:phone"`
	Password   string `gorm:"column:password"`
	AccessType int    `gorm:"column:access_type"`
	Access     string `gorm:"column:access"`
}
