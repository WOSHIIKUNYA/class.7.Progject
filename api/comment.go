package api

import (
	"class.7.Progject/modle"
	"class.7.Progject/service"
	"class.7.Progject/util"
	"github.com/gin-gonic/gin"
)

func SendComment(m *gin.Context) {
	if modle.LoginUser == "" {
		m.String(200, "请先登录")
		return
	}
	var R modle.Comment
	var zz modle.Comment1
	m.ShouldBind(&zz)
	service.CheckMessage(zz.Message)
	if modle.Comment2 == false {
		m.String(200, "没有该评论")
		return
	}
	R.Commenter = modle.LoginUser
	R.Message = zz.Message
	R.Detail = zz.Detail
	m.JSON(200, R)
	err := service.SendComment(R)
	if err != nil {
		util.Number5InternalErr(m)
		return
	}
	m.String(200, "已成功评论")
}
