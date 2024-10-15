// conQuery V3
// examples of the usage of pgx
// simple connection
// perform a query to retrieve data from a table
//
// author prr, azul software
// date: 25 Sept 2024
// copyright 2024 prr, azul softwre
//
// v3: trying to pass table name as parameter --- not possible
//

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type pers struct {
	Uid int `db:"user_id"`
	First string `db:"first"`
	Last string `db:"last"`
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
      		fmt.Printf("error row[%d] -- unable to scan row: %v", count, err)
			os.Exit(1)
    	}
		fmt.Printf("row[%d]: %d %s %s\n", count, user_id, first, last)
	}

	// alternative
	fmt.Println("*** issuing Query again ***")
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
