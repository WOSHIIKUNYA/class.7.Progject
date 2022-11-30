package api

import (
	"class.7.Progject/modle"
	"class.7.Progject/service"
	"class.7.Progject/util"
	"github.com/gin-gonic/gin"
)

func register(m *gin.Context) {
	var z modle.User
	m.ShouldBind(&z)
	if z.Name == "" {
		m.String(200, "傻逼你账号都没输入")
		return
	}
	if z.Password == "" {
		m.String(200, "你的密码呢？")
		return
	}
	x := modle.User{Password: z.Password, Name: z.Name}
	p := service.CheckUser(x)
	if p == false {
		m.String(200, "用户名重复")
		return
	}
	err := service.CreatUser(x)
	if err != nil {
		util.Number1InternalErr(m)
	}
}
func login(m *gin.Context) {
	var x modle.User
	err := m.ShouldBind(&x)
	if x.Name == "" {
		m.String(200, "傻逼你账号都没输入")
		return
	}
	if x.Password == "" {
		m.String(200, "你的密码呢？")
		return
	}
	if err != nil {
		util.Number2InternalErr(m)
	}
	p := service.Login(x)
	if p == false {
		m.String(200, "用户名或密码错误")
		return
	}
	m.JSON(200, x.Name)
	m.SetCookie("Name", x.Password, 0, "", "/", false, false)
	modle.LoginUser = x.Name
	m.String(200, "已登入，欢迎:", x.Name)
}
