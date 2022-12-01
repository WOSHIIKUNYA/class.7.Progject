package api

import (
	"class.7.Progject/modle"
	"class.7.Progject/service"
	"class.7.Progject/util"
	"github.com/gin-gonic/gin"
)

func GiveMessage(m *gin.Context) {
	if modle.LoginUser == "" {
		m.String(200, "请先登录")
		return
	}
	var R modle.Message
	var xx modle.Add
	m.ShouldBind(&xx)
	m.JSON(200, xx)
	R.SendUid = modle.LoginUser
	R.Detail = xx.Detail
	R.ReceiveUid = xx.ReceiveUid
	m.JSON(200, R)
	err := service.Send(R)
	if err != nil {
		util.Number3InternalErr(m)
		return
	}
	m.String(200, "已成功发送")
}
func GetMessage(m *gin.Context) {
	if modle.LoginUser == "" {
		m.String(200, "请先登录")
		return
	}
	err, l := service.Get(modle.LoginUser)
	if err != nil {
		util.Number4InternalErr(m)
		return
	}
	if l == nil {
		m.String(200, "还没有信息")
	}
	for x := 0; x < len(l); x++ {
		m.JSON(200, l[x])
	}
}
func ChangeMessage(m *gin.Context) {
	if modle.LoginUser == "" {
		m.String(200, "请先登录")
		return
	}
	var R modle.Change
	var xx modle.Add
	m.ShouldBind(&xx)
	R.SendUid = modle.LoginUser
	R.Detail = xx.Detail
	R.ReceiveUid = xx.ReceiveUid
	err := service.Change(R)
	if err != nil {
		util.Number3InternalErr(m)
		return
	}
	m.String(200, "已成功修改")
}
func DeleteMessage(m *gin.Context) {
	if modle.LoginUser == "" {
		m.String(200, "请先登录")
		return
	}
	var x modle.Add3
	m.ShouldBind(&x)
	x.SendUid = modle.LoginUser
	err := service.DeleteMessage(x)
	if err != nil {
		util.Number5InternalErr(m)
	}
	m.String(200, "删除成功")
}
