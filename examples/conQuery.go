// conQuery
// examples of the usage of pgx
// simple connection
// perform a query to retrieve data from a table
//
// author prr, azul software
// date: 25 Sept 2024
// copyright 2024 prr, azul softwre
//


package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type pers struct {
	Schemaname string `db:"schemaname"`
	Tablename string `db:"tablename"`
	Tableowner string `db:"tableowner"`
}

func main() {

	ctx := context.Background()
	dburl :="postgresql://dbuser:dbtest@/testdb"

	dbcon, err := pgx.Connect(ctx, dburl)
	if err != nil {
		fmt.Printf("error -- Unable to create connection: %v\n", err)
		os.Exit(1)
	}
	defer dbcon.Close(ctx)

	var schema, table, owner string
	query := "select schemaname, tablename, tableowner from pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema';"
	rows, err := dbcon.Query(ctx, query)
	if err != nil {
		fmt.Printf("error -- query failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	count :=0
	for rows.Next() {
		count++
    	err := rows.Scan(&schema, &table, &owner)
    	if err != nil {
      		fmt.Errorf("error row[%d] -- unable to scan row: %w", count, err)
			os.Exit(1)
    	}
		fmt.Printf("row[%d]: %s %s %s\n", count, schema, table, owner)
	}

	// alternative
    nrows, err := dbcon.Query(ctx, query)
    if err != nil {
		fmt.Printf("error -- query failed: %v\n", err)
		os.Exit(1)
	}
	defer nrows.Close()

	fmt.Println("query success!")

	persList, err := pgx.CollectRows(nrows, pgx.RowToStructByName[pers])
	if err != nil {
		fmt.Printf("error -- failed collecting rows: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("pers list: %d\n", len(persList))
	for _, per := range persList {
		fmt.Printf("%#v\n", per)
	}

	fmt.Println("*** success ***")
}
