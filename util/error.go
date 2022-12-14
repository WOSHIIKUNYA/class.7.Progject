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
var Number3 = reTemplate{
	Status: 500,
	Info:   "Number3",
}
var Number4 = reTemplate{
	Status: 500,
	Info:   "Number4",
}
var Number5 = reTemplate{
	Status: 500,
	Info:   "Number5",
}
var Number6 = reTemplate{
	Status: 500,
	Info:   "Number6",
}
var Number7 = reTemplate{
	Status: 500,
	Info:   "Number6",
}
var Number8 = reTemplate{
	Status: 500,
	Info:   "Number6",
}
var Number9 = reTemplate{
	Status: 500,
	Info:   "Number6",
}

func Number1InternalErr(m *gin.Context) {
	m.JSON(500, Number1)
}
func Number2InternalErr(m *gin.Context) {
	m.JSON(500, Number2)
}
func Number3InternalErr(m *gin.Context) {
	m.JSON(500, Number3)
}
func Number4InternalErr(m *gin.Context) {
	m.JSON(500, Number4)
}
func Number5InternalErr(m *gin.Context) {
	m.JSON(500, Number5)
}
func Number6InternalErr(m *gin.Context) {
	m.JSON(500, Number6)
}
func Number7InternalErr(m *gin.Context) {
	m.JSON(500, Number7)
}
func Number8InternalErr(m *gin.Context) {
	m.JSON(500, Number8)
}
func Number9InternalErr(m *gin.Context) {
	m.JSON(500, Number9)
}
