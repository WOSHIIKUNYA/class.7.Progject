package util

import "github.com/gin-gonic/gin"

type reTemplate struct {
	Status int    `json:"Status"`
	Info   string `json:"Info"`
}

var OK = reTemplate{
	Status: 200,
	Info:   "success",
}
var Number1 = reTemplate{
	Status: 500,
	Info:   "Number1",
}
var Number2 = reTemplate{
	Status: 500,
	Info:   "Number2",
}

func Number1InternalErr(m *gin.Context) {
	m.JSON(500, Number1)
}
func Number2InternalErr(m *gin.Context) {
	m.JSON(500, Number2)
}
