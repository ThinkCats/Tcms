package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func getDb() *sql.DB {
	db,err := sql.Open("mysql", "root:123456@/test")
	if err != nil{
		panic(err.Error())
	}
	return db
}

//Query user info
func QueryUser(){
	db := getDb()
	defer db.Close()
	stmt,err := db.Prepare("select name from user")
	checkError(err)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	rows,err := stmt.Query()
	checkError(err)
	for rows.Next() {
		var name string
		err:= rows.Scan(&name)
		checkError(err)
		fmt.Println("Name :",name)
	}
}

func checkError(err error){
	if err != nil{
		panic(err.Error())
	}
}


