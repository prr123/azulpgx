// conPrepare
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
	"github.com/jackc/pgx/v5/pgconn"
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


	fmt.Println("*** option A: use arguments ***")
	query := "insert into person (first, last) values ($1, $2);"

	_, err = dbcon.Prepare(ctx, "prep", query)
	if err != nil {
		fmt.Printf("error -- prepare failed: %v\n", err)
		os.Exit(1)
	}

	tag, err := dbcon.Exec(ctx, "prep", "peter2", "testInsert2")
	if err != nil {
		fmt.Printf("error -- insert with args failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("tag: %s\n", tag.String())


	fmt.Println("*** success ***")
}


/*
func PrintCmdTag(tag pgconn.CommandTag) {


	return
}
*/

func PrintPrep(ps *pgconn.StatementDescription) {
	fmt.Printf(" **** Statement: %s *****\n", ps.Name)
	fmt.Printf("SQL: %s\n", ps.SQL)

	fmt.Printf("Field desc: %d\n", len(ps.Fields))
	count:= -1
	for _, desc := range(ps.Fields) {
		count++
		fmt.Printf("  field[%d]: %s\n", count+1, desc.Name)
	}

	fmt.Printf(" **** End Statement *****\n")
}
