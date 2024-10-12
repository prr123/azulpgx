// copyFrom example
// examples of the usage of pgx
// simple connection
// perform a query to retrieve data from a table
//
// author prr, azul software
// date: 12 Oct 2024
// copyright 2024 prr, azul softwre
//


package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)


func main() {

	ctx := context.Background()
	dburl :="postgresql://dbuser:dbtest@/testdb"

	dbcon, err := pgx.Connect(ctx, dburl)
	if err != nil {
		fmt.Printf("error -- Unable to create connection: %v\n", err)
		os.Exit(1)
	}
	defer dbcon.Close(ctx)

	var user_id int
	var first, last string
	query := "select user_id, first, last from person;"

	rows, err := dbcon.Query(ctx, query)
	if err != nil {
		fmt.Printf("error -- query failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	count :=0
	for rows.Next() {
		count++
    	err := rows.Scan(&user_id, &first, &last)
    	if err != nil {
      		fmt.Errorf("error row[%d] -- unable to scan row: %w", count, err)
			os.Exit(1)
    	}
		fmt.Printf("row[%d]: %d %s %s\n", count, user_id, first, last)
	}
/*
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
*/
	fmt.Println("*** success ***")
}
