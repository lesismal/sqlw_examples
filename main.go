package main

import (
	"github.com/lesismal/sqlw_examples/db"
	"github.com/lesismal/sqlw_examples/stmt"
	"github.com/lesismal/sqlw_examples/tx"
)

func main() {
	db.RunTest()
	tx.RunTest()
	stmt.RunTest()
}
