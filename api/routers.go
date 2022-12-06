package api

import "github.com/gin-gonic/gin"

func SetApi() {
	r := gin.Default()

	x := r.Group("/First")
	{
		x.POST("register", register)
		x.POST("/login", login)
	}

	f := r.Group("/Second", CheckLogin())
	{
		f.GET("/SendMessage", GiveMessage)
		f.POST("/GetMessage", GetMessage)
		f.PUT("/ChangeMessage", ChangeMessage)
		f.DELETE("/DeleteMessage", DeleteMessage)
	}
	z := r.Group("/Thread", CheckLogin())
	{
		z.GET("/PresentComment", SendComment)
		z.GET("/RespondComment", RespondComment)
		z.POST("/GetComment", GetComment)
		z.PUT("/ChangeComment", ChangeComment)
		z.DELETE("/DeleteComment", DeleteComment)
	}
	r.Run()
}
