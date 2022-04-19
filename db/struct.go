package db

type BaseDataModel struct {
	ID       int   `gorm:"primaryKey;column:id;autoIncrement;type:int;" json:"id"`
	CreateAt int64 `gorm:"autoCreateTime;type:bigint;column:create_at" json:"create_at"`
	UpdateAt int64 `gorm:"autoUpdateTime;type:bigint;column:update_at" json:"update_at"`
}

// User 用户表
type User struct {
	BaseDataModel
	Name     string `gorm:"column:name;type:varchar(45)" json:"name"`
	Avatar   string `gorm:"column:avatar;type:longtext" json:"avatar"`
	Email    string `gorm:"primaryKey;column:email;type:varchar(320);unique" json:"email"`
	Phone    string `gorm:"primaryKey;column:phone;type:varchar(45);unique" json:"phone"`
	Password string `gorm:"column:password;type:varchar(45)" json:"password"`
	Identity string `gorm:"column:identity;type:set('student', 'manager');default:student;" json:"identity"`
	UserId   string `gorm:"primaryKey;column:user_id;type:varchar(20);unique" json:"user_id"`
	Info     []byte `gorm:"column:info;type:json" json:"info"`
}

// Application 申请表单
type Application struct {
	BaseDataModel
	UserId      string  `gorm:"primaryKey;column:user_id;type:varchar(20)" json:"user_id"`
	Status      string  `gorm:"column:status;type:varchar(30);default:save" json:"status"`
	Score       float32 `gorm:"column:score;type:float;default:0" json:"score"`
	ProcedureId int     `gorm:"primaryKey;column:procedure_id;type:int;" json:"procedure_id"`
	ScoreInfo   []byte  `gorm:"column:score_info;type:json" json:"score_info"`
	Info        []byte  `gorm:"column:info;type:json" json:"info"`
	Step        []byte  `gorm:"column:step;type:json" json:"step"`
	History     []byte  `gorm:"column:history;type:json" json:"history"`
}

// Procedure 流程
type Procedure struct {
	BaseDataModel
	UserId      string `gorm:"primaryKey;column:user_id;type:varchar(20)" json:"user_id"`
	CurrentStep []byte `gorm:"column:current_step;type:json" json:"current_step"`
	Info        []byte `gorm:"column:info;type:json" json:"info"`
	History     []byte `gorm:"column:history;type:json" json:"history"`
}

// Action 评定操作记录
type Action struct {
	BaseDataModel
	UserId  string `gorm:"primaryKey;column:user_id;type:varchar(20)" json:"user_id"`
	ApplyId int    `gorm:"primaryKey;column:apply_id;type:int;" json:"apply_id"`
	Info    []byte `gorm:"column:info;type:json" json:"info"`
}

// Log 日志表
type Log struct {
	BaseDataModel
	Data string `gorm:"column:data;type:json" json:"data"`
	Info []byte `gorm:"column:info;type:json" json:"info"`
}
