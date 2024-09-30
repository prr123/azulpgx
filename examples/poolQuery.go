// poolQuery
// examples of the usage of pgx
//
// pool connection
// using a query with select
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

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	ctx := context.Background()
	dburl :="postgresql://dbuser:dbtest@/testdb"

    dbpool, err := pgxpool.New(ctx, dburl)
    if err != nil {
        fmt.Printf("error -- Unable to create connection pool: %v\n", err)
        os.Exit(1)
    }
    defer dbpool.Close()

	var schema, table, owner string
	query := "select schemaname, tablename, tableowner from pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema';"
	rows, err := dbpool.Query(ctx, query)
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
