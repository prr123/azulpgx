# pgx examples

 - author: prr, azulsoftware

pgx: https://pkg.go.dev/github.com/jackc/pgx/v5@v5.7.1

pgx is a tool kit to interact with the postgresql database in the go language.  
The toolkit was developed and is maintained by Jack Christensen (jack@jackchristensen.com).  

There is a google document detailing resources and the listed examples collected by me.  
https://docs.google.com/document/d/1a-HZX49x8gjtBvPi-IrDjeEL5PBm9q7fQCo7aSemqJg/edit

The purpose of this package is providing an extensive set of examples to facilitate learing how to interact from go with postgresql.  

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
https://github.com/prr123/azulpgx/blob/dev/examples/conExec.go

## Query Functions

Query functions are sql requests based on the SELECT keyword, that return data from the database.  
https://github.com/prr123/azulpgx/blob/dev/examples/conQuery.go

## Prepared Statements

A prepared statement is a sql statement that is submitted without parameters. After the submission it can be executed multiple times with different parameters by referring to the name of the prepared statement.  
https://github.com/prr123/azulpgx/blob/dev/examples/conPrepare.go

## Batch Submissions

Batch is an implicit transaction. Batch prepares a number of sql statements that are submitted to the server. The execution of these statements occurs with the use of subsequent function calls.  
https://github.com/prr123/azulpgx/blob/dev/examples/conBatch.go

## Copy Submitions

Copy functions permit filling datatables with multiple entries in one statement. It is a quicker method than filling datatables one row at a time.

## Large Objects


## Trace


## pgtype 

### json

https://github.com/prr123/azulpgx/blob/dev/examples/typJson.go
