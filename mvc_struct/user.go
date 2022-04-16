package mvc_struct

type StudentInfo struct {
	Professional string `json:"professional"`
	Grade        string `json:"grade"`
	Class        int    `json:"class"`
	College      string `json:"college"`
	Type         string `json:"type"`
}

type ManagerInfo struct {
	Department string `json:"department"`
	Office     string `json:"office"`
	Position   string `json:"position"`
}

type UpdateUserInfo struct {
	Name    string      `json:"name"`
	Phone   string      `json:"phone"`
	Avatar  string      `json:"avatar"`
	UserId  string      `json:"user_id"`
	Student StudentInfo `json:"student"`
	Manager ManagerInfo `json:"manager"`
}
