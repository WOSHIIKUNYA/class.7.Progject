package service

import (
	"class.7.Progject/dao"
	"class.7.Progject/modle"
)

func Send(message modle.Message) error {
	err := dao.Send(message)
	return err
}
func Get(m string) (error, []modle.Add1) {
	err, l := dao.Get(m)
	return err, l
}
