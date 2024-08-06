package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sqlboiler_test/models"

	mssql "github.com/microsoft/go-mssqldb"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

//go:generate sqlboiler mssql

func main() {

	pass := os.Getenv("PASSWORD")

	conn := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;", "192.168.86.50", "sa", pass, "sqlboiler_test")

	db, err := sql.Open("sqlserver", conn)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	uid := mssql.UniqueIdentifier{}
	uid.Scan("5AB0412A-2B2F-430F-8830-002A42125148")

	// this works
	all, err := models.TestTables().All(ctx, db)
	if err != nil {
		panic(err) // not reached
	} else {
		for _, t := range all {
			fmt.Println(t.Date)
		}
	}

	test, err := models.TestTables(Where("ID = ?", uid)).One(ctx, db)
	if err != nil {
		// models: failed to execute a one query for TestTable: bind failed to execute query: mssql: Operand type clash: uniqueidentifier is incompatible with money
		fmt.Println(err)
	} else {
		fmt.Println(test.Date)
	}

	test2, err := models.FindTestTable(ctx, db, uid)
	if err != nil {
		// models: unable to select from TestTable: bind failed to execute query: mssql: Operand type clash: uniqueidentifier is incompatible with money
		fmt.Println(err)
	} else {
		fmt.Println(test2.Date)
	}
}
