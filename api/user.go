package api

import (
	"class.7.Progject/modle"
	"class.7.Progject/service"
	"class.7.Progject/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	a := service.CheckName(z.Password)
	if a != "密码格式正确" {
		m.String(200, a)
		return
	}
	m.String(200, a)
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

	m.SetCookie("Name", "good", 3600, "/", "localhost", false, true)
	m.JSON(200, x.Name)
}
func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("Name")
		fmt.Println(err)
		if cookie == "good" {
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		return
	}
}
