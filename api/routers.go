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
		f.POST("/GiveMessage", GiveMessage)
		f.POST("/GetMessage", GetMessage)
	}

	r.Run()
}
