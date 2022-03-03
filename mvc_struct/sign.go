package mvc_struct

type BaseSignForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignOfLogin struct {
	BaseSignForm
	Remember bool `json:"remember"`
}

type SignOfRegister struct {
	BaseSignForm
	Phone     string `json:"phone"`
	StudentId string `json:"student_id"`
	ManageId  string `json:"manage_id"`
	Avatar    string `json:"avatar"`
}

type SignOfForget struct {
	Phone     string `json:"phone"`
	StudentId string `json:"student_id"`
	ManagerId string `json:"manager_id"`
	Email     string `json:"email"`
}
