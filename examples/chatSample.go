// chatSample
// test listen and notify
// based on chatsample
//
// author prr, azul software
// date: 13 Oct 2024
// copyright 2024 prr, azul softwre
//
// https://github.com/jackc/pgx/blob/master/examples/chat/main.go

package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func main() {

    ctx := context.Background()
    dburl :="postgresql://dbuser:dbtest@/testdb"

    dbpool, err := pgxpool.New(ctx, dburl)
    if err != nil {
        fmt.Printf("error -- Unable to create connection pool: %v\n", err)
        os.Exit(1)
    }
    defer dbpool.Close()

	go listen(ctx, dbpool)

	for {
		time.Sleep(500*time.Millisecond)
		fmt.Printf("> ")
	    reader := bufio.NewReader(os.Stdin)
    	msgb, err := reader.ReadBytes('\n')
    	if err != nil {
        	fmt.Printf("error -- reading input: %v\n",err)
			os.Exit(1)
    	}
		msglen:= len(msgb)
		fmt.Printf("msg: %s\n", string(msgb[:msglen-1]))
		msg := string(msgb[:msglen-1])
		if msg == "exit" {
			fmt.Println("exiting!")
			os.Exit(0)
		}

		_, err = dbpool.Exec(ctx, "select pg_notify('chat', $1)", msg)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error -- sending notification:", err)
			os.Exit(1)
		}
	}
}

func listen(ctx context.Context, dbpool *pgxpool.Pool) {
	conn, err := dbpool.Acquire(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error acquiring connection: %v\n", err)
		os.Exit(1)
	}
	defer conn.Release()
	fmt.Fprintf(os.Stderr, "obtained conn!\n")

	_, err = conn.Exec(context.Background(), "listen chat")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error -- listening to chat channel: %v\n", err)
		os.Exit(1)
	}

	for {
		notification, err := conn.Conn().WaitForNotification(context.Background())
		if err != nil {
			fmt.Fprintf(os.Stderr, "error -- waiting for notification: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("PID:", notification.PID, "Channel:", notification.Channel, "Payload:", notification.Payload)
	}
}
