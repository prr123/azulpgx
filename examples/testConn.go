// testConn
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

	var greeting string
	err = dbcon.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Printf("error -- QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
