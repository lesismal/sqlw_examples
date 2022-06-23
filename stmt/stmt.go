package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lesismal/sqlw"
	examples "github.com/lesismal/sqlw_examples"
)

func main() {
	db, err := sqlw.Open("mysql", "test:123qwe@tcp(localhost:3306)/sqlw_test", "db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	{
		// Exec as std sql.
		stmt, err := db.Prepare(`insert into sqlw_test.sqlw_test(field_bool, field_int) values(?, ?)`)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < 2; i++ {
			ret, err := stmt.Exec(i%2 == 1, i)
			if err != nil {
				log.Fatal(err)
			}
			lastInsertId, err := ret.LastInsertId()
			if err != nil {
				log.Fatal(err)
			}
			rowsAffected, err := ret.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("stmt.Exec insert [%d] LastInsertId: %d, RowsAffected: %d", i, lastInsertId, rowsAffected)
		}
	}

	{
		// QueryRow like std sql, but pass one more arg to receive the row result.
		stmt, err := db.Prepare(`select * from sqlw_test.sqlw_test order by id asc`)
		if err != nil {
			log.Fatal(err)
		}
		var dstMinId examples.ModelForTest
		err = stmt.QueryRow(&dstMinId)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("stmt.QueryRow result, MinId: %v, FieldBool: %v, FieldInt: %v", dstMinId.Id, dstMinId.FieldBool, dstMinId.FieldInt)

		stmt, err = db.Prepare(`select * from sqlw_test.sqlw_test order by id desc`)
		if err != nil {
			log.Fatal(err)
		}
		var dstMaxId examples.ModelForTest
		err = stmt.QueryRow(&dstMaxId)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("stmt.QueryRow result, MaxId: %v, FieldBool: %v, FieldInt: %v", dstMaxId.Id, dstMaxId.FieldBool, dstMaxId.FieldInt)
	}

	{
		// Query like std sql, but pass one more arg to receive the row result.
		stmt, err := db.Prepare(`select * from sqlw_test.sqlw_test`)
		if err != nil {
			log.Fatal(err)
		}

		var dstArr []*examples.ModelForTest
		err = stmt.Query(&dstArr)
		if err != nil {
			log.Fatal(err)
		}
		for i, v := range dstArr {
			log.Printf("stmt.Query result[%v]: Id: %v, FieldBool: %v, FieldInt: %v", i, v.Id, v.FieldBool, v.FieldInt)
		}
	}
}
