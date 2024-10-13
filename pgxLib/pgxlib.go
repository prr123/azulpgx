// pgxlib
// library that facilitates certain postgres operations
//
// author: prr, azul software
// date: 7 Oct 2024
// copyright 2024 prr, azul software
//

package pgxlib

import (
    "context"
//    "log"
    "fmt"
//    "os"
    "github.com/jackc/pgx/v5"
)

type PgCol struct {
	ColName string `db:"column_name"`
	ColType string `db:"data_type"`
	ColWidth int `db:"maxchar"`
	Prec	int `db:"prec"`
}

type pgxlib struct {
	con *pgx.Conn
	dbg bool
	ctx context.Context
}

func InitPgLib(con *pgx.Conn) (pglib *pgxlib, err error) {
	if con == nil {return nil, fmt.Errorf("no conn!")}
	pglib = &pgxlib{}
	pglib.con = con
	pglib.dbg = true
	return pglib, nil
}

func (pl *pgxlib) GetColInfo(tbl string)(pgcols []PgCol, err error) {
    sql := `select column_name, data_type, coalesce(character_maximum_length, -1) as maxchar, coalesce(numeric_precision, -1) as prec 
from information_schema.columns where table_name = $1;`

	fmt.Printf("sql: %s\n",sql)
	ctx := pl.ctx
	dbcon := pl.con

	rows, err := dbcon.Query(ctx, sql, tbl)
    if err != nil {
        return pgcols, fmt.Errorf("query failed: %v\n", err)
    }
    defer rows.Close()

/*
    pgcols, err = pgx.CollectRows(rows, pgx.RowToStructByName[PgCol])
    if err != nil {
        return pgcols, fmt.Errorf("failed collecting rows: %v\n", err)
    }
*/
    var ColNam, DTyp string
	var MaxChars, Prec int

	count:= 0
    for rows.Next() {
        count++
        err = rows.Scan(&ColNam, &DTyp, &MaxChars, &Prec)
        if err != nil {
            return pgcols, fmt.Errorf("error row[%d] -- unable to scan row: %w", count, err)
        }
//        fmt.Printf("row[%d]: %s %s %d %d\n", count, ColNam, DTyp, MaxChars, Prec)
        fmt.Printf("row[%d]: %-15s %-20s %d %d\n", count, ColNam, DTyp, MaxChars, Prec)
    }
    rows.Close()

	return pgcols, nil
}


func (pl *pgxlib) GetColInfoV2(tbl string)(pgcols []PgCol, err error) {
    sql := `select column_name, data_type, coalesce(character_maximum_length, -1) as maxchar, coalesce(numeric_precision, -1) as prec 
from information_schema.columns where table_name = $1;`

//	fmt.Printf("sql: %s\n",sql)
	ctx := pl.ctx
	dbcon := pl.con

	rows, err := dbcon.Query(ctx, sql, tbl)
    if err != nil {
        return pgcols, fmt.Errorf("query failed: %v\n", err)
    }
    defer rows.Close()


    pgcols, err = pgx.CollectRows(rows, pgx.RowToStructByName[PgCol])
    if err != nil {
        return pgcols, fmt.Errorf("failed collecting rows: %v\n", err)
    }

	return pgcols, nil
}

func (pl *pgxlib) GetChannels()(listChan []string, err error) {

	sql := "select * FROM pg_listening_channels();"

	ctx := pl.ctx
	dbcon := pl.con

	rows, err := dbcon.Query(ctx, sql)
    if err != nil {
        return listChan, fmt.Errorf("query failed: %v\n", err)
    }
    defer rows.Close()

	var newChan string
	count:= 0
    for rows.Next() {
        count++
        err = rows.Scan(&newChan)
        if err != nil {
            return listChan, fmt.Errorf("error row[%d] -- unable to scan row: %w", count, err)
        }
//        fmt.Printf("row[%d]: %s %s %d %d\n", count, ColNam, DTyp, MaxChars, Prec)
        fmt.Printf("row[%d]: %-15s\n", count, newChan)
    }
    rows.Close()

	return listChan, nil
}
