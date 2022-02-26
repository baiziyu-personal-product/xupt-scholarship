package db

// User 用户表
type User struct {
	Id        int    `gorm:"column: id"`
	Email     string `gorm:"column: email"`
	Phone     string `gorm:"column: phone"`
	Password  string `gorm:"column: password"`
	Avatar    string `gorm:"column: avatar"`
	Identity  int    `gorm:"column: identity"`
	CreateAt  int64  `gorm:"column: create_at"`
	ManageId  string `gorm:"column: manage_id"`
	StudentId string `gorm:"column: student_id"`
}
