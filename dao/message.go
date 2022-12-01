package dao

import (
	"class.7.Progject/modle"
	"fmt"
)

func Send(m modle.Message) error {
	_, err := database.Exec("insert into message (Senduid,ReceiveUid,Detail) values (?,?,?)", m.SendUid, m.Detail)
	return err
}
func Get(m string) (error, []modle.Add1) {
	x, err := database.Query("select SendUid,Detail from message")
	var z []modle.Add1
	var l modle.Add1
	fmt.Println(err)
	for x.Next() {
		x.Scan(&l.SendUid, &l.Detail)
		z = append(z, l)
	}
	return err, z
}
func Change(m modle.Message) error {
	_, err := database.Exec("update message set ReceiveUid=? where SendUid=? ", m.ReceiveUid, m.SendUid)
	fmt.Println(err)
	_, err = database.Exec("update message set Detail=? where SendUid=? ", m.Detail, m.SendUid)
	fmt.Println(err)
	return err
}
