package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
	"github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
	"github.com/lib/pq"
	"github.com/luna-duclos/instrumentedsql"
	"google.golang.org/api/iterator"
)

type Product struct {
	SKU   string `bigquery:"sku"`
	Title string `bigquery:"title"`
}

type DB struct {
	*dbr.Connection
}

func main() {
	projectName := os.Getenv("PROJECT_NAME")
	ctx := context.Background()

	sess, closeConn, err := makeFakeSession()
	if err != nil {
		log.Fatalf("creation error: %v", err)
	}
	defer closeConn()

	query, err := interpolateBQQuery(
		sess.Select("sku, title").
			From(prefTable("products")),
		10,
	)
	if err != nil {
		log.Fatalf("interpolation failure: %v", err)
	}

	client, err := bigquery.NewClient(ctx, projectName)
	if err != nil {
		log.Fatalf("new bigquery client: %v", err)
	}
	defer client.Close()

	result, err := client.Query(query).Read(ctx)
	if err != nil {
		log.Fatalf("bigquery read: %v", err)
	}

	for {
		var row Product
		err := result.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("error iterating through results: %w", err)
		}

		fmt.Printf("%v\n", row)
	}
}

func prefTable(sql string) string {
	tablePrefix := os.Getenv("TABLE_PREFIX")

	return fmt.Sprintf("%s%s", tablePrefix, sql)
}

func makeFakeSession() (*dbr.Session, func() error, error) {
	masterDBInfo := "host= port= user= password= dbname= sslmode=disable"
	sql.Register("instrumented-postgres", instrumentedsql.WrapDriver(&pq.Driver{}))

	// establishing master connection
	masterDB, err := sql.Open("instrumented-postgres", masterDBInfo)
	if err != nil {
		return nil, nil, err
	}

	masterConn := &DB{
		Connection: &dbr.Connection{
			DB:            masterDB,
			Dialect:       dialect.PostgreSQL,
			EventReceiver: &dbr.NullEventReceiver{},
		},
	}

	sess := masterConn.NewSession(nil)

	return sess, masterDB.Close, nil
}

func interpolateSelectStatement(statement *dbr.SelectStmt) (string, error) {
	buf := dbr.NewBuffer()
	err := statement.Build(statement.Dialect, buf)
	if err != nil {
		return "", err
	}

	return dbr.InterpolateForDialect(buf.String(), buf.Value(), statement.Dialect)
}

func interpolateBQQuery(statement *dbr.SelectStmt, limit uint64) (string, error) {
	statement.Limit(limit)
	sqlQuery, err := interpolateSelectStatement(statement)
	if err != nil {
		return "", err
	}

	return sqlQuery, nil
}
