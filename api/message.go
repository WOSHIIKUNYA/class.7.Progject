package api

import (
	"class.7.Progject/modle"
	"github.com/gin-gonic/gin"
)

func GetMessage(m *gin.Context) {
	var R modle.Message
	m.ShouldBind(&R)
}
func SendMessage(m *gin.Context) {

}
