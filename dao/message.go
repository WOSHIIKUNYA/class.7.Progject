package dao

import (
	"class.7.Progject/modle"
)

func Send(m modle.Message) error {
	_, err := database.Exec("insert into message values (?,?,?)", m.SendUid, m.ReceiveUid, m.Detail)
	return err
}
