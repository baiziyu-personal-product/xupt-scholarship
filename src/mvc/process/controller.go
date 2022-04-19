package process

import (
	"github.com/kataras/iris/v12/mvc"
	"xupt-scholarship/src/common"
)

type Response common.BaseResponse

type MVC common.BaseMVC

func (c *MVC) BeforeActivation(m *mvc.Application) {
	m
}

func (c *MVC) UseProcessMVC(m *mvc.Application) {
	m.Register(userSess)
}
