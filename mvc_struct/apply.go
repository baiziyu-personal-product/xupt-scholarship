package mvc_struct

type BaseApply struct {
	Form      interface{} `json:"form"`
	EditAt    int64       `json:"edit_at"`
	StudentId string      `json:"student_id"`
	Type      string      `json:"type"`
}

type ApplyOfCreate struct {
	BaseApply
	CreateAt int64 `json:"create_at"`
}

type ApplyOfEdit struct {
	BaseApply
}
