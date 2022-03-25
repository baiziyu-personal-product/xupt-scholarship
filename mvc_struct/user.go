package mvc_struct

type UpdateUserInfo struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Avatar string `json:"avatar"`
	UserId string `json:"user_id"`
}
