package db

// ApplicationForm 奖学金申请表
type ApplicationForm struct {
	Id string `gorm:"column:id"`
	CreateTime string `gorm:"column:create_time"`
	StudentId string `gorm:"column:student_id"`
	Info string `gorm:"column: info"`
	Editable int `gorm:"column: editable"`
	UpdateTime string `gorm:"column: update_time"`
	Year int `gorm:"column: year"`
	Archive int `gorm:"column: archive"`
	ProcessId int `gorm:"column: process_id"`
	ListId int `gorm:"column: list_id"`
}

// ApplicationList 申请列表，学生申请记录
type ApplicationList struct {
	Id string `gorm:"column:id"`
	CreateTime string `gorm:"column:create_time"`
	ProcessId int `gorm:"column: process_id"`
	StudentId int  `gorm:"column: student_id"`
	UserId int  `gorm:"column: user_id"`
	Status int  `gorm:"column: status"`
	Process string  `gorm:"column: process"`
	EndTime string  `gorm:"column: end_time"`
	StatusInfo string  `gorm:"column: status_info"`
}