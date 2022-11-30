package dao

import (
	"test/modle"
)

func Repeat(m modle.User) bool {
	k, _ := database.Query("select Name from user")
	var j []string
	for k.Next() {
		var x string
		_ = k.Scan(&x)
		j = append(j, x)
	}
	for P := 0; P < len(j); P++ {
		if m.Name == j[P] {
			return false
		}
	}
	return true
}
func Login(m modle.User) bool {
	k, _ := database.Query("select Name,Password from user")
	var j []modle.User
	for k.Next() {
		var x modle.User
		_ = k.Scan(&x.Name, &x.Password)
		j = append(j, x)
	}
	for P := 0; P < len(j); P++ {
		if m.Name == j[P].Name || m.Password == j[P].Password {
			return true
		}
	}
	return false
}

func InitMysql(m modle.User) error {
	_, err := database.Exec("insert into user values (?,?)", m.Name, m.Password)
	return err
}
