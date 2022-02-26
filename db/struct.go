package db

// User 用户表
type User struct {
	Email     string `gorm:"column:email"`
	Phone     string `gorm:"column:phone"`
	Password  string `gorm:"column:password"`
	Avatar    string `gorm:"column:avatar"`
	Identity  string `gorm:"column:identity"`
	CreateAt  int64  `gorm:"column:create_at"`
	ManageId  string `gorm:"column:manage_id"`
	StudentId string `gorm:"column:student_id"`
}
