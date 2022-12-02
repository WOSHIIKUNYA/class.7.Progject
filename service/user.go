package service

import (
	"class.7.Progject/dao"
	"class.7.Progject/modle"
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
func CheckName(m string) string {
	if len(m) < 6 {
		return "密码太短"
	} else if len(m) > 15 {
		return "密码太长"
	} else {
		return "密码格式正确"
	}
}
