package model

import (
	"xupt-scholarship/db"
	"xupt-scholarship/mvc_struct"
	"xupt-scholarship/utils"
)

type CommentModel struct {
}

func (c *CommentModel) CreateComment(data mvc_struct.CommentData) BaseModelFmtData {
	comment := db.Comment{
		Content:     data.Content,
		ReplyId:     data.ReplyId,
		ProcedureId: data.ProcedureId,
		UserId:      data.UserId,
	}
	result := db.Mysql.Create(&comment)
	return HandleDBData(result, comment.ID)
}

func (c *CommentModel) DelComment(id int) BaseModelFmtData {
	var comment db.Comment
	result := db.Mysql.Delete(&comment, id)
	return HandleDBData(result, comment.ID)
}

func (c *CommentModel) GetCommentList(processId int, replyId int) BaseModelFmtData {
	var comments []db.Comment
	var commentList []mvc_struct.CommentList
	result := db.Mysql.Where("procedure_id = ? AND reply_id = ?", processId, replyId).Find(&comments)
	if result.Error == nil {
		for _, v := range comments {
			var user db.User
			var childComments []db.Comment
			res := db.Mysql.Where("user_id = ?", v.UserId).Find(&user)
			db.Mysql.Where("procedure_id = ? AND reply_id = ?", processId, v.ID).Find(&childComments)
			if res.Error == nil {
				commentList = append(commentList, mvc_struct.CommentList{
					CommentData: mvc_struct.CommentData{
						UserId:      v.UserId,
						ProcedureId: v.ProcedureId,
						Content:     v.Content,
						ReplyId:     v.ReplyId,
					},
					Children:   len(childComments),
					Id:         v.ID,
					CreateDate: utils.FmtTimeByUnix(v.CreateAt),
					Avatar:     user.Avatar,
					Email:      user.Email,
					Name:       user.Name,
				})
			}
		}
	}
	return HandleDBData(result, commentList)
}
