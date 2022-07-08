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

	var model Model
	selectId := 1
	result, err := db.Select(&model, "select * from sqlw_test.sqlw_test where id=?", selectId)
	// result, err := db.SelectOne(&model, "select (i,s) from sqlw_test.sqlw_test where id=?", selectId) // select some fields
	if err != nil {
		log.Panic(err)
	}
	log.Println("model:", model)
	log.Println("sql:", result.Sql())

	var models []*Model // type []Model is also fine
	result, err = db.Select(&models, "select * from sqlw_test.sqlw_test")
	// result, err = db.SelectOne(&models, "select (i,s) from sqlw_test.sqlw_test") // select some fields
	if err != nil {
		log.Panic(err)
	}
	for i, v := range models {
		log.Printf("models[%v]: %v", i, v)
	}
	log.Println("sql:", result.Sql())
}
