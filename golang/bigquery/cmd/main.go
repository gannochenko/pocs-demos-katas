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
	"google.golang.org/api/option"
)

const (
	QueryLimit = 10
)

type Record struct {
	SKU   string `bigquery:"sku"`
	Title string `bigquery:"title"`
}

type DB struct {
	*dbr.Connection
}

func main() {
	projectName := os.Getenv("PROJECT_NAME")
	credentialsFile := os.Getenv("CREDENTIALS_FILE")

	ctx := context.Background()

	sess, closeConn, err := makeFakeSession()
	if err != nil {
		log.Fatalf("creation error: %v", err)
	}
	defer closeConn()

	query, err := interpolateBQQuery(
		sess.Select("sku, title").
			From(prefTable("products")),
		QueryLimit,
	)
	if err != nil {
		log.Fatalf("interpolation failure: %v", err)
	}

	fmt.Printf("EXECUTING: %s\n", query)

	client, err := bigquery.NewClient(ctx, projectName, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		log.Fatalf("new bigquery client creation error: %v", err)
	}
	defer client.Close()

	result, err := client.Query(query).Read(ctx)
	if err != nil {
		log.Fatalf("bigquery read error: %v", err)
	}

	for {
		var row Record
		err := result.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("error iterating through results: %v", err)
		}

		fmt.Printf("%v\n", row)
		fmt.Println("--------==============================--------")
	}
}

func prefTable(sql string) string {
	tablePrefix := os.Getenv("TABLE_PREFIX")

	return fmt.Sprintf("%s%s", tablePrefix, sql)
}

func makeFakeSession() (*dbr.Session, func() error, error) {
	sql.Register("instrumented-postgres", instrumentedsql.WrapDriver(&pq.Driver{}))
	masterDB, err := sql.Open("instrumented-postgres", "host= port= user= password= dbname= sslmode=disable")
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
