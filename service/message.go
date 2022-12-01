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
func Change(message modle.Change) error {
	err := dao.Change(message)
	return err
}
func DeleteMessage(message modle.Add3) error {
	err := dao.DeleteMessage(message)
	return err
}
