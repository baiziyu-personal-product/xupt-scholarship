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

// ApplicationProcess 奖学金流程
type ApplicationProcess struct {
	Id int `gorm:"column:id"`
	CreateTime string `gorm:"column:create_time"`
	Status int `gorm:"column:status"`
	EndTime string  `gorm:"column:end_time"`
	ManageId int `gorm:"column:manage_id"`
	Year int  `gorm:"column: year"`
	Info string  `gorm:"column:info"`
	Reviewers string  `gorm:"column:reviewers"`
	Students string  `gorm:"column:students"`
	ListIds string  `gorm:"column: list_ids"`
}

// Quantify 得分量化表
type Quantify struct {
	Id int `gorm:"column:id"`
	Group string `gorm:"column:group"`
	Category string `gorm:"column:category"`
	Type string `gorm:"column:type"`
	Score int `gorm:"column:score"`
	Item string `gorm:"column:item"`
	Info string `gorm:"column:info"`
	CreateTime string `gorm:"column:create_time"`
	ReviewerId int `gorm:"column:reviewer_id"`
	UpdateTime string `gorm:"column: update_time"`
	Effect int `gorm:"column: effect"`
}

// Reviewers 参与评审的人员
type Reviewers struct {
	Id int `gorm:"column:id"`
	UserId int `gorm:"column:user_id"`
	Name string `gorm:"column: name"`
	Professional string `gorm:"column: professional"`
	Institute string `gorm:"column: institute"`
}

// Students 学生表
type Students struct {
	Id int 	`gorm:"column:id"`
	UserId int `gorm:"column:user_id"`
	Name int `gorm:"column: name"`
	StudentId string `gorm:"column: student_id"`
	Gender string `gorm:"column: gender"`
	Professional string `gorm:"column: professional"`
	Class int `gorm:"column:class"`
	Session int `gorm:"column:session"`
	ReceivingStatus int `gorm:"column: receiving_status"`
}
