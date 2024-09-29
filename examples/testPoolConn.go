// testPoolConn
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

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbpool, err := pgxpool.New(context.Background(), "postgresql://dbuser:dbtest@/testdb")
	if err != nil {
		fmt.Printf("error -- Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Printf("error -- QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
