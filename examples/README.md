# pgx examples

 - author: prr, azulsoftware

pgx: https://pkg.go.dev/github.com/jackc/pgx/v5@v5.7.1

pgx is a tool kit to interact with the postgresql database in the go language.
The toolkit was developed and is maintained by 

There is a google doc detailing resources an examples collected by me.
https://docs.google.com/document/d/1a-HZX49x8gjtBvPi-IrDjeEL5PBm9q7fQCo7aSemqJg/edit

This package provides an extensive set of examples to facilitate learing how to interact from go with postgresql.

## Connections

Pgx supports two types of connections:
 - single connection 
 - pooled connections (https://pkg.go.dev/github.com/jackc/pgx/v5@v5.7.1/pgxpool)

Examples:
 - single connection via unix domain sockets

https://github.com/prr123/azulpgx/blob/dev/examples/testConn.go

 - creation of a connection pool

https://github.com/prr123/azulpgx/blob/dev/examples/poolTest.go


## Exec Functions

The Exec routine intiates posgresql changes that do not return data from the database.
The alternative for data is a Query routine.


## Query Functions

Query functions are sql requests based on the SELECT keyword, that return data from the database.

## Prepared Statements

## Batch Submissions

## Copy Submitions

Copy functions permit filling datatables with multiple entries in one statement. It is a quicker method than filling datatables one row at a time.

## Large Objects


## Trace


