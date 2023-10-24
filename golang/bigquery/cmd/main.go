package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
	"github.com/lib/pq"
	"github.com/luna-duclos/instrumentedsql"
)

type DB struct {
	*dbr.Connection
}

func main() {
	//projectID := os.Getenv("PROJECT_ID")
	//
	//ctx := context.Background()
	//
	//client, err := bigquery.NewClient(ctx, projectID)
	//if err != nil {
	//	log.Fatalf("bigquery.NewClient: %v", err)
	//}
	//defer client.Close()

	sess, closeConn, err := makeFakeSession()
	if err != nil {
		log.Fatalf("creation error: %v", err)
	}
	defer closeConn()

	query, err := interpolateBQQuery(
		sess.Select("foo").
			From(prefTable("table")).
			Where("bar = ?", true),
		10,
	)

	if err != nil {
		log.Fatalf("interpolation failure: %v", err)
	}

	fmt.Printf("%v\n", query)

	//query := client.Query(
	//	`SELECT
	//		CONCAT(
	//			'https://stackoverflow.com/questions/',
	//			CAST(id as STRING)) as url,
	//		view_count
	//	FROM ` + "`bigquery-public-data.stackoverflow.posts_questions`" + `
	//	WHERE tags like '%google-bigquery%'
	//	ORDER BY view_count DESC
	//	LIMIT 10;`)
	//return query.Read(ctx)
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
