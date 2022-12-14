package dao

import (
	"class.7.Progject/modle"
	"fmt"
)

func Send(m modle.Message) error {
	_, err := database.Exec("insert into message (Senduid,ReceiveUid,Detail) values (?,?,?)", m.SendUid, m.ReceiveUid, m.Detail)
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
func Change(m modle.Change) error {
	x, err := database.Query("select Mid,SendUid,ReceiveUid from message")
	if err != nil {
		fmt.Println(err)
	}
	var p int
	for x.Next() {
		var l modle.Add2
		x.Scan(&l.Id, &l.SendUid, &l.ReceiveUid)
		if m.SendUid == l.SendUid && m.ReceiveUid == l.ReceiveUid {
			p = l.Id
			break
		}
	}
	fmt.Println(p)
	_, err = database.Exec("update message set Detail=? where Mid=?", m.Detail, p)
	fmt.Println(err)
	return err
}
func DeleteMessage(r modle.Add3) error {
	x, err := database.Query("select Mid,SendUid,ReceiveUid from message")
	if err != nil {
		fmt.Println(err)
	}
	var p int
	for x.Next() {
		var l modle.Add2
		x.Scan(&l.Id, &l.SendUid, &l.ReceiveUid)
		if r.SendUid == l.SendUid && r.ReceiveUid == l.ReceiveUid {
			p = l.Id
			break
		}
	}
	fmt.Println(p)
	_, err = database.Exec("update  message set Detail=? where Mid=?", "has been delete", p)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
func CheckMessage(m string) error {
	x, err := database.Query("select Detail from message")
	if err != nil {
		fmt.Println(err)
	}
	var f string
	fmt.Println(m)
	for x.Next() {
		x.Scan(&f)
		if f == m {
			modle.Comment2 = true
			break
		}
	}
	return err
}
func CheckMessage1(m string, r string) error {
	x, err := database.Query("select message,commenter from comment")
	if err != nil {
		fmt.Println(err)
	}
	var f string
	var z string
	fmt.Println(m)
	for x.Next() {
		x.Scan(&f, &z)
		if f == m && z == r {
			modle.Comment4 = true
			break
		}
	}
	return err
}
