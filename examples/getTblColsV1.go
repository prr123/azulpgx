// getTblColsV1
// examples of the usage of pgx
// simple connection
// perform a query to retrieve data from a table
//
// author prr, azul software
// date: 25 Sept 2024
// copyright 2024 prr, azul softwre
//
// v1: use coalesce

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

	var ColNam, DTyp string
	var MaxChars, Prec int

	tbl := "person"
	query := "select column_name, data_type, coalesce(character_maximum_length, -1), coalesce(numeric_precision, -1) from information_schema.columns where table_name=$1;"

	rows, err := dbcon.Query(ctx, query, tbl)
	if err != nil {
		fmt.Printf("error -- query failed: %v\n", err)
		os.Exit(1)
	}
//	fmt.Printf("query success\n%v\n", rows)

	count :=0
	for rows.Next() {
//		fmt.Printf("row[%d]\n", count)
		count++
//    	err = rows.Scan(&ColNam, &dTyp, &nChars, &prec)
    	err = rows.Scan(&ColNam, &DTyp, &MaxChars, &Prec)
    	if err != nil {
      		fmt.Errorf("error row[%d] -- unable to scan row: %w", count, err)
			os.Exit(1)
    	}
//		fmt.Printf("row[%d]: %s %s %d %d\n", count, ColNam, DTyp, MaxChars, Prec)
		fmt.Printf("row[%d]: %-15s %-20s %d %d\n", count, ColNam, DTyp, MaxChars, Prec)
	}
	rows.Close()

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
