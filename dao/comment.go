package dao

import (
	"class.7.Progject/modle"
	"fmt"
)

func SendComment(m modle.Comment) error {

	_, err := database.Exec("insert into comment (message,Detail,Commenter) values (?,?,?)", m.Message, m.Detail, m.Commenter)
	return err
}
func GetComment(m string) error {

	x, err := database.Query("select message,detail,commenter from comment")
	fmt.Println(err)
	var f modle.Comment
	for x.Next() {
		x.Scan(&f.Message, &f.Detail, &f.Commenter)
		if m == f.Message {
			modle.CommentBook = append(modle.CommentBook, f)
		}
	}
	return err
}
func ChangeComment(m modle.Comment) error {
	x, err := database.Query("select* from comment")
	if err != nil {
		fmt.Println(err)
	}
	var p int
	for x.Next() {
		var l modle.Comment5
		x.Scan(&l.Id, &l.Message, &l.Detail, &l.Commenter)
		if m.Commenter == l.Commenter && m.Message == l.Message {
			p = l.Id
			break
		}
	}
	fmt.Println(p)
	_, err = database.Exec("update comment set detail=? where id=?", m.Detail, p)
	fmt.Println(err)
	return err
}
