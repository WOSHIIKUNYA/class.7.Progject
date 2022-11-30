package main

import (
	"class.7.Progject/api"
	"class.7.Progject/dao"
)

func main() {
	dao.Open()
	go api.SetApi()
	api.Message()
}
