package main

import (
	"fmt"
	"log"
	"sync/atomic"

	"github.com/lesismal/sqlw"
	_ "github.com/lib/pq"
)

var (
	db *sqlw.DB
)

const (
	sqlDropDatabase   = `drop database if exists sqlw_test`
	sqlCreateDatabase = `create database sqlw_test`
	sqlDropTable      = `drop table if exists sqlw_test`
	sqlCreateTable    = `
		create table sqlw_test (
		id serial not null,
		i  bigint not null default 0,
		s  varchar(64) not null default ''
	)`
)

func main() {
	var err error
	db, err = sqlw.Open("postgres", "user=postgres password=123qwe dbname=postgres sslmode=disable", "db")
	if err != nil {
		log.Panic(err)
	}

	CreateDatabase()
	defer DropDatabase()

	CreateTable()
	defer DropTable()

	insertTest()
	deleteTest()
	updateTest()
	selectTest()

	log.Printf("Test Pass [db]")
}

func insertTest() {
	// insert struct ptr
	ret, err := db.Insert(`insert into sqlw_test(i, s)`, NewModel())
	if err != nil {
		log.Panic(err)
	}
	CheckResult("[insert struct ptr]", ret, err)

	// insert struct
	ret, err = db.Insert(`insert into sqlw_test`, *NewModel())
	if err != nil {
		log.Panic(err)
	}
	CheckResult("[insert struct]", ret, err)

	// insert []*struct
	ret, err = db.Insert(`insert into sqlw_test`, []*Model{NewModel(), NewModel()})
	if err != nil {
		log.Panic(err)
	}
	CheckResult("[insert []*struct]", ret, err)

	// insert []struct
	ret, err = db.Insert(`insert into sqlw_test`, []Model{*NewModel(), *NewModel()})
	if err != nil {
		log.Panic(err)
	}
	CheckResult("[insert []struct]", ret, err)

	// insert raw
	ret, err = db.Insert(`insert into sqlw_test(i,s) values($1,$2)`, NextInt(), NextString())
	CheckResult("[insert raw]", ret, err)
}

func deleteTest() {
	// delete
	ret, err := db.Delete(`delete from sqlw_test where id=$1`, 1)
	CheckResult("[delete raw]", ret, err)
}

func updateTest() {
	// update by struct
	var m = Model{I: 20, S: "str_20"}
	ret, err := db.Update(`update sqlw_test set i=$1, s=$2 where id=2`, m)
	CheckResult("[update by struct]", ret, err)

	// update by struct ptr
	m = Model{I: 30, S: "str_30"}
	ret, err = db.Update(`update sqlw_test set i=$1, s=$2 where id=3`, &m)
	CheckResult("[update by struct ptr]", ret, err)

	// update by struct and raw
	m = Model{I: 40, S: "str_40"}
	ret, err = db.Update(`update sqlw_test set i=$1, s=$2 where id=$3`, m, 4)
	CheckResult("[update by struct and raw]", ret, err)

	// update by struct ptr and raw
	m = Model{I: 50, S: "str_50"}
	ret, err = db.Update(`update sqlw_test set i=$1, s=$2 where id=$3`, &m, 5)
	CheckResult("[update by struct ptr and raw]", ret, err)

	// update by raw
	ret, err = db.Update(`update sqlw_test set i=$1, s=$2 where id=$3`, 60, "str_60", 6)
	CheckResult("[update by raw]", ret, err)
}

func selectTest() {
	// select one
	var one Model
	ret, err := db.Select(&one, "select * from sqlw_test")
	CheckResult("[select one]", ret, err)

	one = Model{}
	ret, err = db.Select(&one, "select * from sqlw_test order by id asc")
	CheckResult("[select one]", ret, err)

	one = Model{}
	ret, err = db.Select(&one, "select * from sqlw_test order by id desc")
	CheckResult("[select one]", ret, err)

	one = Model{}
	ret, err = db.Select(&one, "select id,i from sqlw_test order by id asc")
	CheckResult("[select one]", ret, err)

	one = Model{}
	ret, err = db.Select(&one, "select id,s from sqlw_test order by id desc")
	CheckResult("[select one]", ret, err)

	// select all struct ptr
	var allPtr []*Model
	ret, err = db.Select(&allPtr, "select * from sqlw_test")
	CheckResult("[select all struct ptr]", ret, err)
	if len(allPtr) != 6 {
		log.Panic(fmt.Errorf("invalid records num: %v", len(allPtr)))
	}

	// select all struct
	var allStruct []Model
	ret, err = db.Select(&allStruct, "select * from sqlw_test")
	CheckResult("[select all struct]", ret, err)
	if len(allStruct) != 6 {
		log.Panic(fmt.Errorf("invalid records num: %v", len(allStruct)))
	}
	PrintModels(allStruct)
}

func CreateDatabase() {
	_, err := db.Exec(sqlDropDatabase)
	if err != nil {
		log.Panic(err)
	}

	_, err = db.Exec(sqlCreateDatabase)
	if err != nil {
		log.Panic(err)
	}
}

func DropDatabase() {
	_, err := db.Exec(sqlDropDatabase)
	if err != nil {
		log.Panic(err)
	}
}

func CreateTable() {
	_, err := db.Exec(sqlDropTable)
	if err != nil {
		log.Panic(err)
	}

	_, err = db.Exec(sqlCreateTable)
	if err != nil {
		log.Panic(err)
	}
}

func DropTable() {
	_, err := db.Exec(sqlDropTable)
	if err != nil {
		log.Panic(err)
	}
}

func CheckResult(testName string, ret sqlw.Result, err error) {
	fmt.Println("-------------------------------------")
	fmt.Printf("%v\n", testName)
	lastInsertId, _ := ret.LastInsertId()
	rowsAffected, _ := ret.RowsAffected()
	fmt.Printf("Sql         : %v\nLastInsertId: %v\nRowsAffected: %v\nError       : %v\n", ret.Sql(), lastInsertId, rowsAffected, err)
	if err != nil {
		log.Panic(err)
	}
}

var cnt int64

type Model struct {
	Id int64  `db:"id"`
	I  int64  `db:"i"`
	S  string `db:"s"`
}

func NewModel() *Model {
	return &Model{
		I: NextInt(),
		S: NextString(),
	}
}

func Reset() {
	cnt = 0
}

func NextInt() int64 {
	return atomic.AddInt64(&cnt, 1)
}

func NextString() string {
	return fmt.Sprintf("str_%d", cnt)
}

func PrintModels(ms []Model) {
	fmt.Printf("Models      : %v\n", ms)
}
