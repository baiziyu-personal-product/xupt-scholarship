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

// Application 申请表单
type Application struct {
	Form     interface{} `gorm:"column:form"`
	History  interface{} `gorm:"column:history"`
	CreateAt int64       `gorm:"column:create_at"`
	EditAt   int64       `gorm:"column:edit_at"`
	Score    interface{} `gorm:"column:score"`
	Creator  string      `gorm:"column:creator"`
	Status   int         `gorm:"column:status"`
	Step     interface{} `gorm:"column:step"`
	Year     int         `gorm:"column:year"`
}

// Process 流程
type Process struct {
	Info        interface{} `gorm:"column:form"`
	CurrentStep interface{} `gorm:"column:current_step"`
	CreateAt    int64       `gorm:"column:create_at"`
	Creator     string      `gorm:"column:creator"`
}
