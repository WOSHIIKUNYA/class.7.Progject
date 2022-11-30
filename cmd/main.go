package main

import (
	"class.7.Progject/api"
	"class.7.Progject/dao"
)

func main() {
	dao.Open()
	api.SetApi()
}
