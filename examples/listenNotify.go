// listen notify example
// examples of the usage of pgx
// creates a listen channel and destroys a listen channel
//
// author prr, azul software
// date: 12 Oct 2024
// copyright 2024 prr, azul softwre
//
// V2
// use pool 

package main

import (
	"context"
	"fmt"
	"os"
//	"time"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
)


func main() {

	ctx := context.Background()
	dburl :="postgresql://dbuser:dbtest@/testdb"

	dbpool, err := pgxpool.New(ctx, dburl)
	if err != nil {
		fmt.Printf("error -- Unable to create connection: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	// create a listen channel
	query := "listen new;"
    tag, err := dbpool.Exec(ctx, query)
    if err != nil {
        log.Fatalf("error -- exec failed: %v\n", err)
    }

    fmt.Printf("tag: %s\n", tag.String())

	chlist, err := getChanInfo(ctx, dbpool)
	if err !=nil {log.Fatalf("error -- getChanInfo: %v\n", err)}
	fmt.Printf("*** channels: %d ****\n", len(chlist))
	for i:=0; i< len(chlist); i++ {
		fmt.Printf(" --%d: %s\n", i, chlist[i])
	}


	log.Println("getting a connection")

	dbcon, err := dbpool.Acquire(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error acquiring connection:", err)
		os.Exit(1)
	}
	defer dbcon.Release()

	log.Println("ready to receive notification!")

	for {


		notification, err := dbcon.Conn().WaitForNotification(ctx)
//	notification, err := dbcon.WaitForNotification(ctx)
		if err != nil {log.Fatalf("error -- Wait for Notification: %v\n", err)}

		fmt.Println("PID:", notification.PID, "Channel:", notification.Channel, "Payload:", notification.Payload)

		if notification.Payload == "exit" {
			fmt.Printf("exiting!\n")
			break
		}
	}

	fmt.Println("*** success ***")
}

func getChanInfo(ctx context.Context, dbpool *pgxpool.Pool) (chList []string, err error) {

	var lChan string


	nquery := "select * from pg_listening_channels();"
	rows, err := dbpool.Query(ctx, nquery)
	if err != nil {
		return chList, fmt.Errorf("query failed: %v\n", err)
	}
	defer rows.Close()

	count :=0
	for rows.Next() {
		count++
    	err := rows.Scan(&lChan)
    	if err != nil {
      		return chList, fmt.Errorf("row[%d] -- unable to scan row: %w", count, err)
    	}
		chList = append(chList, lChan)
	}

	return chList, nil
}

func recNotify(ctx context.Context, dbpool *pgxpool.Pool) {

	log.Println("getting a connection")
	dbcon, err := dbpool.Acquire(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error acquiring connection:", err)
		os.Exit(1)
	}
	defer dbcon.Release()

	log.Println("ready to receive notification!")

	notification, err := dbcon.Conn().WaitForNotification(ctx)
//	notification, err := dbcon.WaitForNotification(ctx)
	if err != nil {log.Fatalf("error -- Wait for Notification: %v\n", err)}

	fmt.Println("PID:", notification.PID, "Channel:", notification.Channel, "Payload:", notification.Payload)


	fmt.Println("*** end recNotify ***")
	return
}
