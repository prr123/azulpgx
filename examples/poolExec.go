// conExec
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

//	"github.com/jackc/pgx/v5"
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

	query := "drop schema acnt;"

	tag, err := dbpool.Exec(ctx, query)
	if err != nil {
		fmt.Printf("error -- exec failed: %v\n", err)
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
