package mvc_struct

type BaseSignForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginForm struct {
	BaseSignForm
	Remember bool `json:"remember"`
}

type RegisterForm struct {
	BaseSignForm
	Phone     string `json:"phone"`
	StudentId string `json:"student_id"`
	ManagerId string `json:"manager_id"`
	Avatar    string `json:"avatar"`
}
