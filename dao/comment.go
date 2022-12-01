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
