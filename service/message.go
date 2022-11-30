package service

import (
	"class.7.Progject/dao"
	"class.7.Progject/modle"
)

func Send(message modle.Message) error {
	err := dao.Send(message)
	return err
}
