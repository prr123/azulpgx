// unit tests

package pgxlib

import (
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

	pglib := InitPgLib(dbcon)

	if pglib == nil {}
}

