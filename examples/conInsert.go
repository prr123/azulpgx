// conInsert
// examples of the usage of pgx
// simple connection
// insert

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
//	"github.com/jackc/pgx/v5/pgconn"
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


	fmt.Println("*** option A: direct ***")
	query := "insert into person (first, last) values ('peterA2', 'testA2' );"

	tag, err := dbcon.Exec(ctx, query)
	if err != nil {
		fmt.Printf("error -- insert with args failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("tag: %s\n", tag.String())

	fmt.Println("*** option B: use arguments ***")
	query = "insert into person (first, last) values ($1, $2);"

	tag, err = dbcon.Exec(ctx, query, "peterB2", "testInsertB2")
	if err != nil {
		fmt.Printf("error -- insert with args failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("tag: %s\n", tag.String())

	fmt.Println("*** option C: use named arguments ***")
	query = "insert into person (first, last) values (@first, @last);"
	args := pgx.NamedArgs{
		"first": "peterC2",
		"last": "testC2",
	}

	tag, err = dbcon.Exec(ctx, query, args)
	if err != nil {
		fmt.Printf("error -- insert with args failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("tag: %s\n", tag.String())


	fmt.Println("*** success ***")
}


