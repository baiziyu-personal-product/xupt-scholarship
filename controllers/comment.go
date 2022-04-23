package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"xupt-scholarship/model"
	"xupt-scholarship/mvc_struct"
)

type CommentMVC struct {
	BaseController
}

type CommentController interface {
	GetBy(processId int) BaseControllerFmtData
	Post() BaseControllerFmtData
	Delete() BaseControllerFmtData
	DeleteBy(commentId int) BaseControllerFmtData
}

var commentModel model.CommentModel

func UseCommentMVC(app *mvc.Application) {
	app.Register(userSession.Start).Handle(new(CommentMVC))
}

func (c *CommentMVC) GetBy(processId int) BaseControllerFmtData {
	replyId := c.Ctx.URLParamIntDefault("reply_id", 0)
	result := commentModel.GetCommentList(processId, replyId)
	return HandleControllerRes(result, "获取评论信息")
}

func (c *CommentMVC) Post() BaseControllerFmtData {
	var comment mvc_struct.CommentData
	user := GetUserData(c.Session)
	GetRequestParams(c.Ctx, &comment)
	comment.UserId = user.UserId
	result := commentModel.CreateComment(comment)
	return HandleControllerRes(result, "创建评论")
}

func (c *CommentMVC) DeleteBy(commentId int) BaseControllerFmtData {
	result := commentModel.DelComment(commentId)
	return HandleControllerRes(result, "删除评论")
}
