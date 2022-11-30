package api

import (
	"class.7.Progject/modle"
	"class.7.Progject/service"
	"class.7.Progject/util"
	"github.com/gin-gonic/gin"
)

func GiveMessage(m *gin.Context) {
	m.String(200, "sda")
	if modle.LoginUser == "" {
		m.String(200, "请先登录")
		return
	}
	var R modle.Message
	var x modle.Add
	m.ShouldBind(&x)
	R.SendUid = modle.LoginUser
	R.Detail = x.Detail
	R.ReceiveUid = x.ReceiveUid
	m.JSON(200, R)
	err := service.Send(R)
	if err != nil {
		util.Number3InternalErr(m)
		return
	}
	m.String(200, "已登录")
}
func GetMessage(m *gin.Context) {

}
