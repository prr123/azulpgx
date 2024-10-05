// conBatch
// examples of the usage of pgx
// simple connection
// batch execution

// author prr, azul software
// date: 4 Octt 2024
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

type user struct {
	first string
	last string
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

	userList := []user {
		user {
			first: "jim",
			last: "Doe",
		},
		user {
			first: "jane",
			last: "Second",
		},
		user {
			first: "john",
			last: "Third",
		},
	}

	query := "insert into person (first, last) values ($1, $2);"

	batch := &pgx.Batch{}

	for i:=0; i< len(userList); i++ {
		batch.Queue(query, userList[i].first, userList[i].last)
	}

	bLen := batch.Len()
	fmt.Printf("info -- blen: %d\n", bLen)

	bRes := dbcon.SendBatch(ctx, batch)

	for i:=0; i< len(userList); i++ {
		tag, err := bRes.Exec()
		if err != nil {
			fmt.Printf("error -- cmd[%d]: %v\n",i ,err)
			continue
		}
		fmt.Printf("info -- cmd[%d]: %s\n",i , tag.String())
	}
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
