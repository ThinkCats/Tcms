package dao

import (
	"database/sql"

	//just use mysql init
	_ "github.com/go-sql-driver/mysql"
)

//User user info
type User struct {
	Name     string
	Password string
}

//GetDb get db connect
func GetDb() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@/test")
	if err != nil {
		panic(err.Error())
	}
	return db
}

//QueryUser Query user info
func QueryUser(name string) (User, error) {
	var user User
	db := GetDb()
	defer db.Close()
	stmt, err := db.Prepare("select name,password from user where name = ?")
	defer stmt.Close()
	rows, err := stmt.Query(name)
	for rows.Next() {
		err := rows.Scan(&user.Name, &user.Password)
		checkError(err)
	}
	return user, err
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
