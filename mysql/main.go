package main

import (
	"github.com/lesismal/sqlw_examples/mysql/db"
	"github.com/lesismal/sqlw_examples/mysql/stmt"
	"github.com/lesismal/sqlw_examples/mysql/tx"
)

func main() {
	db.RunTest()
	tx.RunTest()
	stmt.RunTest()
}
