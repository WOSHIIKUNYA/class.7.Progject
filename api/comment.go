package api

import (
	"class.7.Progject/modle"
	"class.7.Progject/service"
	"class.7.Progject/util"
	"fmt"
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
	modle.Comment2 = false
}
func GetComment(m *gin.Context) {
	if modle.LoginUser == "" {
		m.String(200, "请先登录")
		return
	}
	var x modle.Comment3
	m.ShouldBind(&x)
	if x.Message == "" {
		m.String(200, "煞笔你要读的信息呢")
		return
	}
	service.CheckMessage(x.Message)
	if modle.Comment2 == false {
		m.String(200, "没有该留言")
		return
	}
	err := service.GetComment(x.Message)
	if err != nil {
		util.Number6InternalErr(m)
		return
	}
	for a := 0; a < len(modle.CommentBook); a++ {
		m.JSON(200, modle.CommentBook[a])
	}
	modle.Comment2 = false
}
func ChangeComment(m *gin.Context) {
	if modle.LoginUser == "" {
		m.String(200, "请先登录")
		return
	}
	var xx modle.Comment1
	m.ShouldBind(&xx)
	fmt.Println(xx)
	var R modle.Comment
	R.Commenter = modle.LoginUser
	R.Detail = xx.Detail
	R.Message = xx.Message
	if xx.Message == "" {
		m.String(200, "煞笔你要改的信息呢")
		return
	}
	service.CheckMessage1(xx.Message, R.Commenter)
	if modle.Comment4 == false {
		m.String(200, "没有该评论")
		return
	}
	err := service.ChangeComment(R)
	if err != nil {
		util.Number7InternalErr(m)
		return
	}
	m.String(200, "已成功修改")
	modle.Comment4 = false
}
func DeleteComment(m *gin.Context) {
	if modle.LoginUser == "" {
		m.String(200, "请先登录")
		return
	}
	var x modle.Comment1
	m.ShouldBind(&x)
	if x.Message == "" {
		m.String(200, "煞笔你要删的信息呢")
		return
	}
	service.CheckMessage1(x.Message, modle.LoginUser)
	var z modle.Comment
	z.Commenter = modle.LoginUser
	z.Detail = x.Detail
	z.Message = x.Message
	err := service.DeleteComment(z)
	if err != nil {
		util.Number8InternalErr(m)
		return
	}
	m.String(200, "删除成功")
}
func RespondComment(m *gin.Context) {
	if modle.LoginUser == "" {
		m.String(200, "请先登录")
		return
	}
	var R modle.Comment
	var zz modle.Comment1
	m.ShouldBind(&zz)
	R.Commenter = modle.LoginUser
	R.Message = zz.Message
	R.Detail = zz.Detail
	service.CheckComment(zz.Message)
	if modle.Comment2 == false {
		m.String(200, "没有该评论")
		return
	}
	m.JSON(200, R)
	err := service.SendComment(R)
	if err != nil {
		util.Number9InternalErr(m)
		return
	}
	m.String(200, "已成功评论")
	modle.Comment2 = false
}
