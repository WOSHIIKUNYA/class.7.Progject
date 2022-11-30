package api

import "github.com/gin-gonic/gin"

func SetApi() {
	r := gin.Default()
	x := r.Group("/First")
	{
		x.POST("/register", register)
		x.POST("/login", login)
	}
	r.Run()
}
func Message() {
	r := gin.Default()
	x := r.Group("Second")
	{
		x.GET("/Send", SendMessage)
		x.POST("/GetMessage", GetMessage)

	}
	r.Run()

}
