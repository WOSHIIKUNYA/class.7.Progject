package dao

import (
	"class.7.Progject/modle"
)

func SendComment(m modle.Comment) error {

	_, err := database.Exec("insert into comment (message,Detail,Commenter) values (?,?,?)", m.Message, m.Detail, m.Commenter)
	return err
}
