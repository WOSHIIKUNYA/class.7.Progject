package service

import (
	"test/dao"
	"test/modle"
)

func CheckUser(m modle.User) bool {
	x := dao.Repeat(m)
	return x
}
func Login(m modle.User) bool {
	x := dao.Login(m)
	return x
}

func CreatUser(m modle.User) error {
	err := dao.InitMysql(m)
	return err
}
