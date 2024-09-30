// testPoolConn
// examples of the usage of pgx
// create a connection pool
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

	ctx :=context.Background()
	dburl := "postgresql://dbuser:dbtest@/testdb"

	dbpool, err := pgxpool.New(ctx, dburl)
	if err != nil {
		fmt.Printf("error -- Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var greeting string
	sql :="select 'Hello, world!'"
	err = dbpool.QueryRow(ctx, sql).Scan(&greeting)
	if err != nil {
		fmt.Printf("error -- QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
