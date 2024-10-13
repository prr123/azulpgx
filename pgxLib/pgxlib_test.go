// unit tests

package pgxlib

import (
	"fmt"
	"testing"
	"context"
	"github.com/jackc/pgx/v5"
)

func TestInit(t *testing.T) {

	   // open a connection
    ctx := context.Background()
    dburl :="postgresql://dbuser:dbtest@/testdb"
    dbcon, err := pgx.Connect(ctx, dburl)
    if err != nil {t.Fatalf("error -- Unable to create connection: %v\n", err)}
    defer dbcon.Close(ctx)

	pglib, err := InitPgLib(dbcon)
	if err != nil {t.Fatalf("error -- InitPgLib: %v\n", err)}

	if pglib.con == nil {t.Fatalf("error -- Unable to init pflib\n")}
}

func TestGetCols(t *testing.T) {

    ctx := context.Background()
    dburl :="postgresql://dbuser:dbtest@/testdb"
    dbcon, err := pgx.Connect(ctx, dburl)
    if err != nil {t.Fatalf("error -- Unable to create connection: %v\n", err)}
    defer dbcon.Close(ctx)

    pglib, err := InitPgLib(dbcon)
    if err != nil {t.Fatalf("error -- InitPgLib: %v\n", err)}
//	fmt.Println("init")
	pglib.ctx = ctx

	pgcol, err := pglib.GetColInfo("person")
    if err != nil {t.Fatalf("error -- GetColInfo: %v\n", err)}

	fmt.Printf("columns: %d\n", len(pgcol))
	for i:=0; i<len(pgcol); i++ {
		fmt.Printf("col[%d]: %s %d %d\n",i, pgcol[i].ColName, pgcol[i].ColWidth, pgcol[i].Prec)

	}
}

func TestGetCols2(t *testing.T) {

    ctx := context.Background()
    dburl :="postgresql://dbuser:dbtest@/testdb"
    dbcon, err := pgx.Connect(ctx, dburl)
    if err != nil {t.Fatalf("error -- Unable to create connection: %v\n", err)}
    defer dbcon.Close(ctx)

    pglib, err := InitPgLib(dbcon)
    if err != nil {t.Fatalf("error -- InitPgLib: %v\n", err)}
//	fmt.Println("init")
	pglib.ctx = ctx

	pgcol, err := pglib.GetColInfoV2("person")
    if err != nil {t.Fatalf("error -- GetColInfo: %v\n", err)}

	fmt.Printf("columns: %d\n", len(pgcol))
	for i:=0; i<len(pgcol); i++ {
		fmt.Printf("col[%d]: %-15s %-20s %d %d\n",i, pgcol[i].ColName, pgcol[i].ColType, pgcol[i].ColWidth, pgcol[i].Prec)

	}
}


func TestListChannels(t *testing.T) {

    ctx := context.Background()
    dburl :="postgresql://dbuser:dbtest@/testdb"
    dbcon, err := pgx.Connect(ctx, dburl)
    if err != nil {t.Fatalf("error -- Unable to create connection: %v\n", err)}
    defer dbcon.Close(ctx)

    pglib, err := InitPgLib(dbcon)
    if err != nil {t.Fatalf("error -- InitPgLib: %v\n", err)}
//	fmt.Println("init")
	pglib.ctx = ctx

	chanList, err := pglib.GetChannels()
    if err != nil {t.Fatalf("error -- GetChannels: %v\n", err)}


}
