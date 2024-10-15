// getTblCols
// examples of the usage of pgx
// simple connection
// perform a query to retrieve column attributes from a table
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

type PgCol struct {
    ColName string `db:"column_name"`
    ColType string `db:"data_type"`
    ColWidth int `db:"character_maximum_length"`
    Prec    int `db:"numeric_precision"`
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

	var ColNam, DTyp string
	var MaxChars int

	tbl := "person"
	query := "select column_name, data_type, character_maximum_length from information_schema.columns where table_name=$1;"

	rows, err := dbcon.Query(ctx, query, tbl)
	if err != nil {
		fmt.Printf("error -- query failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("*** query success ***\n")
	fmt.Printf("*** reading query rows ***\n")

	count :=0
	for rows.Next() {
		fmt.Printf("*** scanning row[%d]: ***\n", count)
		count++
    	err = rows.Scan(&ColNam, &DTyp, &MaxChars)
    	if err != nil {
      		fmt.Printf("error row[%d] -- unable to scan row: %v\n", count, err)
			os.Exit(1)
    	}
		fmt.Printf("  row[%d]: %-15s %-20s %d\n", count, ColNam, DTyp, MaxChars)
	}
	rows.Close()

	fmt.Println("*** success ***")
}
