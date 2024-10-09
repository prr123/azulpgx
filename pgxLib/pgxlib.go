// pgxlib
// library that facilitates certain postgres operations
//
// author: prr, azul software
// date: 7 Oct 2024
// copyright 2024 prr, azul software
//

package pgxlib

import (
//    "context"
//    "log"
//    "fmt"
//    "os"
    "github.com/jackc/pgx/v5"
)

type PgCol struct {
	ColName string
	ColType string
	ColWidth int
	Prec	int
}

type pgxlib struct {
	con *pgx.Conn

}

func InitPglib(con *pgx.Conn) (pglib *pgxlib) {
	pglib = &pgxlib{}
	pglib.con = con
	return pglib
}

func (pl *pgxlib) GetColInfo(tbl string)(pgcols []PgCol, err error) {


	return pgcols, nil
}
