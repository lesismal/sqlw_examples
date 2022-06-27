package stmt

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	util "github.com/lesismal/sqlw_examples/util"
)

var db = util.ConnectDB()

func RunTest() {
	util.Reset()

	util.CreateDatabase()
	defer util.DropDatabase()

	util.CreateTable()
	defer util.DropTable()

	insertTest()
	deleteTest()
	updateTest()
	selectTest()

	log.Printf("Test Pass [stmt]")
}

func insertTest() {
	// insert struct ptr
	stmt, err := db.Prepare(`insert into sqlw_test.sqlw_test(i,s) values(?,?)`)
	if err != nil {
		log.Panic(err)
	}
	ret, err := stmt.Insert(util.NewModel())
	if err != nil {
		log.Panic(err)
	}
	util.CheckResult("[insert struct ptr]", ret, err)

	// insert struct
	stmt, err = db.Prepare(`insert into sqlw_test.sqlw_test(i,s) values(?,?)`)
	if err != nil {
		log.Panic(err)
	}
	ret, err = stmt.Insert(*util.NewModel())
	if err != nil {
		log.Panic(err)
	}
	util.CheckResult("[insert struct]", ret, err)

	// insert []*struct
	stmt, err = db.Prepare(`insert into sqlw_test.sqlw_test(i,s) values(?,?),(?,?)`)
	if err != nil {
		log.Panic(err)
	}
	ret, err = stmt.Insert([]*util.Model{util.NewModel(), util.NewModel()})
	if err != nil {
		log.Panic(err)
	}
	util.CheckResult("[insert []*struct]", ret, err)

	// insert []struct
	stmt, err = db.Prepare(`insert into sqlw_test.sqlw_test(i,s) values(?,?),(?,?)`)
	if err != nil {
		log.Panic(err)
	}
	ret, err = stmt.Insert([]util.Model{*util.NewModel(), *util.NewModel()})
	if err != nil {
		log.Panic(err)
	}
	util.CheckResult("[insert []struct]", ret, err)

	// insert raw
	stmt, err = db.Prepare(`insert into sqlw_test.sqlw_test(i,s) values(?,?)`)
	if err != nil {
		log.Panic(err)
	}
	ret, err = stmt.Insert(util.NextInt(), util.NextString())
	util.CheckResult("[insert raw]", ret, err)
}

func deleteTest() {
	// delete
	stmt, err := db.Prepare(`delete from sqlw_test.sqlw_test where id=?`)
	if err != nil {
		log.Panic(err)
	}
	ret, err := stmt.Delete(1)
	util.CheckResult("[delete raw]", ret, err)
}

func updateTest() {
	// update by struct
	var m = util.Model{I: 20, S: "str_20"}
	stmt, err := db.Prepare(`update sqlw_test.sqlw_test set i=?, s=? where id=2`)
	if err != nil {
		log.Panic(err)
	}
	ret, err := stmt.Update(m)
	util.CheckResult("[update by struct]", ret, err)

	// update by struct ptr
	m = util.Model{I: 30, S: "str_30"}
	stmt, err = db.Prepare(`update sqlw_test.sqlw_test set i=?, s=? where id=3`)
	if err != nil {
		log.Panic(err)
	}
	ret, err = stmt.Update(&m)
	util.CheckResult("[update by struct ptr]", ret, err)

	// update by struct and raw
	m = util.Model{I: 40, S: "str_40"}
	stmt, err = db.Prepare(`update sqlw_test.sqlw_test set i=?, s=? where id=?`)
	if err != nil {
		log.Panic(err)
	}
	ret, err = stmt.Update(m, 4)
	util.CheckResult("[update by struct and raw]", ret, err)

	// update by struct ptr and raw
	m = util.Model{I: 50, S: "str_50"}
	stmt, err = db.Prepare(`update sqlw_test.sqlw_test set i=?, s=? where id=?`)
	if err != nil {
		log.Panic(err)
	}
	ret, err = stmt.Update(&m, 5)
	util.CheckResult("[update by struct ptr and raw]", ret, err)

	// update by raw
	stmt, err = db.Prepare(`update sqlw_test.sqlw_test set i=?, s=? where id=?`)
	if err != nil {
		log.Panic(err)
	}
	ret, err = stmt.Update(60, "str_60", 6)
	util.CheckResult("[update by raw]", ret, err)
}

func selectTest() {
	// select one
	var one util.Model
	stmt, err := db.Prepare(`select * from sqlw_test.sqlw_test`)
	if err != nil {
		log.Panic(err)
	}
	ret, err := stmt.Select(&one)
	util.CheckResult("[select one]", ret, err)

	one = util.Model{}
	stmt, err = db.Prepare(`select * from sqlw_test.sqlw_test order by id asc`)
	if err != nil {
		log.Panic(err)
	}
	ret, err = stmt.Select(&one)
	util.CheckResult("[select one]", ret, err)
	if one.Id != 2 || one.I != 20 || one.S != "str_20" {
		log.Panic(fmt.Errorf("invalid record: %v", one))
	}

	one = util.Model{}
	stmt, err = db.Prepare(`select * from sqlw_test.sqlw_test order by id desc`)
	if err != nil {
		log.Panic(err)
	}
	ret, err = stmt.Select(&one)
	util.CheckResult("[select one]", ret, err)
	if one.Id != 7 || one.I != 7 || one.S != "str_7" {
		log.Panic(fmt.Errorf("invalid record: %v", one))
	}

	one = util.Model{}
	stmt, err = db.Prepare(`select id,i from sqlw_test.sqlw_test order by id asc`)
	if err != nil {
		log.Panic(err)
	}
	ret, err = stmt.Select(&one)
	util.CheckResult("[select one]", ret, err)
	if one.I != one.Id*10 || one.S != "" {
		log.Panic(fmt.Errorf("invalid record: %v", one))
	}

	one = util.Model{}
	stmt, err = db.Prepare(`select id,s from sqlw_test.sqlw_test order by id desc`)
	if err != nil {
		log.Panic(err)
	}
	ret, err = stmt.Select(&one)
	util.CheckResult("[select one]", ret, err)
	if one.I != 0 || one.S != fmt.Sprintf("str_%v", one.Id) {
		log.Panic(fmt.Errorf("invalid record: %v", one))
	}

	// select all struct ptr
	var allPtr []*util.Model
	stmt, err = db.Prepare(`select * from sqlw_test.sqlw_test`)
	if err != nil {
		log.Panic(err)
	}
	ret, err = stmt.Select(&allPtr)
	util.CheckResult("[select all struct ptr]", ret, err)
	if len(allPtr) != 6 {
		log.Panic(fmt.Errorf("invalid records num: %v", len(allPtr)))
	}

	// select all struct
	var allStruct []util.Model
	stmt, err = db.Prepare(`select * from sqlw_test.sqlw_test`)
	if err != nil {
		log.Panic(err)
	}
	ret, err = stmt.Select(&allStruct)
	util.CheckResult("[select all struct]", ret, err)
	if len(allStruct) != 6 {
		log.Panic(fmt.Errorf("invalid records num: %v", len(allStruct)))
	}
	for _, v := range allStruct {
		if v.Id == 7 {
			if v.I != 7 || v.S != "str_7" {
				log.Panic(fmt.Errorf("invalid record: %v", v))
			}
		} else if v.I != v.Id*10 || v.S != fmt.Sprintf("str_%v", v.I) {
			log.Panic(fmt.Errorf("invalid record: %v", v))
		}
	}
	util.PrintModels(allStruct)
}
