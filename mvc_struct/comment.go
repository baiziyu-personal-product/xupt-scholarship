package mvc_struct

type CommentData struct {
	UserId      string `json:"user_id"`
	ProcedureId int    `json:"procedure_id"`
	Content     string `json:"content"`
	ReplyId     int    `json:"reply_id" default:"0"`
}

type CommentList struct {
	CommentData
	Children   int    `json:"children"`
	Id         int    `json:"comment_id"`
	CreateDate string `json:"create_date"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	Name       string `json:"name"`
}
