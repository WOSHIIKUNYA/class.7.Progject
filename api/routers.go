package api

import "github.com/gin-gonic/gin"

func SetApi() {
	r := gin.Default()
	x := r.Group("/First")
	{
		x.POST("/login", login)

		x.POST("/register", register)
	}
	f := r.Group("/Second")
	{
		f.GET("/SendMessage", GiveMessage)
		f.POST("/GetMessage", GetMessage)
		f.PUT("/ChangeMessage", ChangeMessage)
		f.DELETE("/DeleteMessage", DeleteMessage)
	}
	z := r.Group("/Thread")
	{
		z.GET("/PresentComment", SendComment)
		z.POST("/Comment")
		z.PUT("/changeComment")
		z.DELETE("/ChangeComment")
	}
	r.Run()
}
