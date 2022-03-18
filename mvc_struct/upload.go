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
}
