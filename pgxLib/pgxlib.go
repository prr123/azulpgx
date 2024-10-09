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
    "fmt"
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

func InitPgLib(con *pgx.Conn) (pglib *pgxlib, err error) {
	if con == nil {return nil, fmt.Errorf("no conn!")}
	pglib = &pgxlib{}
	pglib.con = con
	return pglib, nil
}

func (pl *pgxlib) GetColInfo(tbl string)(pgcols []PgCol, err error) {


	return pgcols, nil
}
