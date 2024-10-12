// getTblColsV2
// examples of the usage of pgx
// simple connection
// perform a query to retrieve data from a table
//
// author prr, azul software
// date: 25 Sept 2024
// copyright 2024 prr, azul softwre
//
// V2: retrieves raw values no scanning
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

	tbl := "person"
	query := "select column_name, data_type, character_maximum_length, numeric_precision from information_schema.columns where table_name=$1;"

	rows, err := dbcon.Query(ctx, query, tbl)
	if err != nil {
		fmt.Printf("error -- query failed: %v\n", err)
		os.Exit(1)
	}

	count :=0
	for rows.Next() {
		count++
		res, err  := rows.Values()
    	if err != nil {
      		fmt.Errorf("error row[%d] -- unable to scan row: %w", count, err)
			os.Exit(1)
    	}
		fmt.Printf("res[%d]: %v\n", count, res)
//		fmt.Printf("row[%d]: %s %s %d %d\n", count, colNam, dTyp, nChars, prec)
//		fmt.Printf("row[%d]: %-15s %-20s %d\n", count, ColNam, DTyp, MaxChars)
	}
	rows.Close()

	fmt.Println("*** success ***")
}
