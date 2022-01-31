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
