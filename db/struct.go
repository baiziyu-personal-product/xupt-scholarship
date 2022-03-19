package db

type BaseDataModel struct {
	ID       int   `gorm:"primaryKey;column:id;autoIncrement;type:int;"`
	CreateAt int64 `gorm:"autoCreateTime:nano;type:bigint;column:create_at"`
	UpdateAt int64 `gorm:"autoUpdateTime:nano;type:bigint;column:update_at"`
}

// User 用户表
type User struct {
	BaseDataModel
	Name         string      `gorm:"column:name;type:varchar(45)"`
	Avatar       string      `gorm:"column:avatar;type:longtext"`
	Email        string      `gorm:"column:email;type:varchar(320);unique"`
	Phone        string      `gorm:"column:phone;type:varchar(45);unique"`
	Password     string      `gorm:"column:password;type:varchar(45)"`
	Identity     string      `gorm:"column:identity;type:varchar(60)"`
	ManageId     string      `gorm:"column:manage_id;type:varchar(20);unique"`
	StudentId    string      `gorm:"column:student_id;type:varchar(20);unique"`
	CourseCredit interface{} `gorm:"column:course_credit;type:json"`
}

// Application 申请表单
type Application struct {
	BaseDataModel
	Creator string      `gorm:"column:creator;type:varchar(20)"`
	Status  string      `gorm:"column:status;type:varchar(30);default:save"`
	Form    interface{} `gorm:"column:form;type:json"`
	Step    string      `gorm:"column:step;type:varchar(320)"`
	History interface{} `gorm:"column:history;type:json"`
}

// Procedure 流程
type Procedure struct {
	BaseDataModel
	CurrentStep interface{} `gorm:"column:current_step;type:json"`
	Creator     string      `gorm:"column:creator;type:varchar(20);primaryKey"`
	Steps       interface{} `gorm:"column:step;type:json"`
}
