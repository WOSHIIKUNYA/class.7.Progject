package dao

import (
	"class.7.Progject/modle"
	"fmt"
)

func Send(m modle.Message) error {
	fmt.Println(m.Detail)
	_, err := database.Exec("insert into message (Senduid,ReceiveUid,Detail) values (?,?,?)", m.SendUid, m.ReceiveUid, m.Detail)
	return err
}
