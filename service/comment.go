package service

import (
	"class.7.Progject/dao"
	"class.7.Progject/modle"
)

func SendComment(message modle.Comment) error {
	err := dao.SendComment(message)
	return err
}
