package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lesismal/sqlw"
)

// init table
// create table sqlw_test.sqlw_test (
// 	id bigint primary key auto_increment,
// 	i  bigint not null default 0,
// 	s  varchar(64) not null default ''
// )

type Model struct {
	Id int64  `db:"id"`
	I  int64  `db:"i"`
	S  string `db:"s"`
}

func main() {
	db, err := sqlw.Open("mysql", "test:123qwe@tcp(localhost:3306)/mysql", "db")
	if err != nil {
		log.Panic(err)
	}

	m := Model{
		I: 1,
		S: "str_1",
	}

	updateId := 1
	result, err := db.Update("update sqlw_test.sqlw_test set i=?, s=? where id=?", &m, updateId)
	if err != nil {
		log.Panic(err)
	}
	log.Println("sql:", result.Sql())
}
