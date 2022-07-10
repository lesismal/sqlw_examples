package util

import (
	"fmt"
	"log"
	"sync/atomic"

	"github.com/lesismal/sqlw"
)

var (
	db *sqlw.DB

	DebugLog   = false
	SqlConnStr = "test:123qwe@tcp(localhost:3306)/mysql"
)

const (
	sqlDropDatabase   = `drop database if exists sqlw_test`
	sqlCreateDatabase = `create database sqlw_test`
	sqlDropTable      = `drop table if exists sqlw_test.sqlw_test`
	sqlCreateTable    = `
		create table sqlw_test.sqlw_test (
		id bigint primary key auto_increment,
		i  bigint not null default 0,
		s  varchar(64) not null default ''
	)`
)

func ConnectDB() *sqlw.DB {
	var err error
	db, err = sqlw.Open("mysql", SqlConnStr, "db")
	if err != nil {
		log.Panic(err)
	}
	return db
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

func CheckResult(testName string, ret sqlw.Result, err error) {
	if DebugLog {
		fmt.Println("-------------------------------------")
		fmt.Printf("%v\n", testName)
		lastInsertId, _ := ret.LastInsertId()
		rowsAffected, _ := ret.RowsAffected()
		fmt.Printf("Sql         : %v\nLastInsertId: %v\nRowsAffected: %v\nError       : %v\n", ret.Sql(), lastInsertId, rowsAffected, err)
	}
	if err != nil {
		log.Panic(err)
	}
}

func PrintModels(ms []Model) {
	if DebugLog {
		fmt.Printf("Models      : %v\n", ms)
	}
}
