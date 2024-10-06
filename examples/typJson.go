// typJson
// examples of the usage of pgx
// simple connection
// author prr, azul software
// date: 6 Oct 2024
// copyright 2024 prr, azul softwre
//
// modified example from
// https://godocs.io/github.com/jackc/pgx/v5/pgtype
//

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
//	"github.com/jackc/pgx/v5/pgconn"
)

	type person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
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

	input := person{
		Name: "John",
		Age:  42,
	}

	output := &person{}

	err = dbcon.QueryRow(ctx, "select $1::json", input).Scan(&output)
	if err != nil {
		fmt.Printf("error -- queryrow: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(output.Name, output.Age)
	// Output:
	// John 42


	fmt.Println("*** success ***")
}


/*
func PrintCmdTag(tag pgconn.CommandTag) {


	return
}
*/
