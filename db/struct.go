package db

type BaseDataModel struct {
	ID       int   `gorm:"primaryKey;column:id;autoIncrement;type:int;" json:"id"`
	CreateAt int64 `gorm:"autoCreateTime;type:bigint;column:create_at" json:"create_at"`
	UpdateAt int64 `gorm:"autoUpdateTime;type:bigint;column:update_at" json:"update_at"`
}

// User 用户表
type User struct {
	BaseDataModel
	Name         string `gorm:"column:name;type:varchar(45)"`
	Avatar       string `gorm:"column:avatar;type:longtext"`
	Email        string `gorm:"column:email;type:varchar(320);unique"`
	Phone        string `gorm:"column:phone;type:varchar(45);unique"`
	Password     string `gorm:"column:password;type:varchar(45)"`
	Identity     string `gorm:"column:identity;type:varchar(60)"`
	ManageId     string `gorm:"column:manage_id;type:varchar(20);unique"`
	StudentId    string `gorm:"column:student_id;type:varchar(20);unique"`
	CourseCredit []byte `gorm:"column:course_credit;type:json"`
}

// Application 申请表单
type Application struct {
	BaseDataModel
	Creator string  `gorm:"column:creator;type:varchar(20)" json:"creator"`
	Status  string  `gorm:"column:status;type:varchar(30);default:save" json:"status"`
	Form    []byte  `gorm:"column:form;type:json" json:"form"`
	Step    string  `gorm:"column:step;type:varchar(320)" json:"step"`
	History []byte  `gorm:"column:history;type:json" json:"history"`
	Score   float32 `gorm:"column:score;type:float;default:0" json:"score"`
}

// Procedure 流程
type Procedure struct {
	BaseDataModel
	CurrentStep []byte `gorm:"column:current_step;type:json" json:"current_step"`
	Creator     string `gorm:"column:creator;type:varchar(20);primaryKey" json:"creator"`
	Steps       []byte `gorm:"column:step;type:json" json:"steps"`
}
