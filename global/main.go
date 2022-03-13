package global

import (
	"github.com/kataras/iris/v12/sessions"
	"time"
	"xupt-scholarship/config"
)

// Settings 项目设置
var Settings config.GlobalConfig

const SuccessCode = 0
const ErrorCode = 1
const SessionExpires = 24 * 2 * time.Hour

const XUPT_SESSION_ID = "xupt_session_id"

var (
	SessionId   = XUPT_SESSION_ID
	UserSession = sessions.New(sessions.Config{
		Cookie:  SessionId,
		Expires: SessionExpires,
	})
)
