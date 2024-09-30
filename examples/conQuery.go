// conSelectQuery
// examples of the usage of pgx
// simple connection
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
		fmt.Printf("roe[%d]: %s %s %s\n", count, schema, table, owner)
	}
	fmt.Println("*** success ***")
}
