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
	Identity string `json:"identity"`
}

type SignOfForget struct {
	Phone  string `json:"phone"`
	UserId string `json:"user_id"`
	Email  string `json:"email"`
}
