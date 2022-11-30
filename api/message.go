package api

import (
	"class.7.Progject/modle"
	"class.7.Progject/service"
	"class.7.Progject/util"
	"github.com/gin-gonic/gin"
)

func GetMessage(m *gin.Context) {
	if modle.LoginUser == "" {
		m.String(200, "请先登录")
		return
	}
	var R modle.Message
	R.SendUid = modle.LoginUser
	R.ReceiveUid = m.PostForm("ReceiveUid")
	R.Detail = m.PostForm("Detail")
	err := service.Send(R)
	if err != nil {
		util.Number3InternalErr(m)
		return
	}
	m.String(200, "已登录")
}
func SendMessage(m *gin.Context) {

}
