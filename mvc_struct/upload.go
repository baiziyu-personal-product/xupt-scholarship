package mvc_struct

type StudentItem struct {
	Id           int     `json:"id"`
	Key          int     `json:"key"`
	Name         string  `json:"name"`
	StudentId    string  `json:"student_id"`
	Gender       string  `json:"gender"`
	Professional string  `json:"professional"`
	CourseCredit float32 `json:"course_credit"`
	Phone        string  `json:"phone"`
	Email        string  `json:"email"`
	Password     string  `json:"password"`
	Grade        string  `json:"grade"`
	Class        int     `json:"class"`
	Type         string  `json:"type"`
}

type originFileObj struct {
	Uid string `json:"uid"`
}

type UploadFileItem struct {
	LastModified     int64         `json:"lastModified"`
	LastModifiedDate string        `json:"lastModifiedDate"`
	Name             string        `json:"name"`
	OriginFileObj    originFileObj `json:"originFileObj"`
	Percent          float32       `json:"percent"`
	Size             int           `json:"size"`
	Type             string        `json:"type"`
	Uid              string        `json:"uid"`
	Url              string        `json:"url"`
	ThumbUrl         string        `json:"thumbUrl"`
}

type UploadStudentInfo struct {
	UpdateUserInfo
	SignOfRegister
}
