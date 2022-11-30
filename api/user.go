package api

import (
	"github.com/gin-gonic/gin"
	"test/modle"
	"test/service"
	"test/util"
)

func register(m *gin.Context) {
	Name := m.PostForm("Name")
	Password := m.PostForm("Password")
	if Name == "" {
		m.String(200, "傻逼你账号都没输入")
		return
	}
	if Password == "" {
		m.String(200, "你的密码呢？")
		return
	}
	x := modle.User{Password: Password, Name: Name}
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
	m.String(200, "已登入，欢迎:")
	m.JSON(200, x.Name)
	m.SetCookie("Name", x.Password, 0, "", "/", false, false)
}
