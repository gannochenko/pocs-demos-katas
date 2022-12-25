package copier

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/luna-duclos/instrumentedsql"
)

const (
	TableName = "elements"
)

type Options struct {
	SrcDBPort     int32
	SrcDBUser     string
	SrcDBPassword string
	DstDBPort     int32
	DstDBUser     string
	DstDBPassword string
}

type Element struct {
	ID    string `db:"id"`
	Title string `db:"title"`
}

type Copier struct {
	options *Options
}

func New(options *Options) *Copier {
	return &Copier{
		options: options,
	}
}

func (c *Copier) CopyElements(ids []string) error {
	srcDB, dstDB, err := c.makeConnections()
	if err != nil {
		return err
	}
	defer srcDB.Close()
	defer dstDB.Close()

	srcDBSession := srcDB.NewSession(nil)
	dstDBSession := dstDB.NewSession(nil)

	dstDBTransaction, err := dstDBSession.Begin()
	if err != nil {
		return err
	}
	defer dstDBTransaction.RollbackUnlessCommitted()

	ctx := context.Background()

	processed := 0
	for _, id := range ids {
		element, err := c.getElementByID(ctx, srcDBSession, id)
		if err != nil {
			fmt.Printf("Could not read element with id %s\n", id)
			continue
		}
		if element.ID == "" {
			fmt.Printf("Could not read element with id %s: element not found\n", id)
			continue
		}

		err = c.createElement(ctx, dstDBTransaction, element)
		if err != nil {
			fmt.Printf("Could not create an element with id %s\n", element.ID)
			continue
		}

		processed += 1
	}

	err = dstDBTransaction.Commit()
	if err != nil {
		return err
	}

	fmt.Printf("Done. Items processed: %d\n", processed)

	return nil
}

func (c *Copier) getElementByID(ctx context.Context, session *dbr.Session, id string) (element *Element, err error) {
	_, err = session.
		Select("*").
		From(TableName).
		Where("id = ?", id).
		LoadContext(ctx, &element)

	if err != nil {
		return nil, err
	}

	return element, nil
}

func (c *Copier) createElement(ctx context.Context, session *dbr.Tx, element *Element) error {
	columns := []string{"id", "title"}

	element.ID = uuid.New().String()
	err := session.InsertInto(TableName).Columns(columns...).Record(element).LoadContext(ctx, &element.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c *Copier) makeConnections() (srcDBConn *dbr.Connection, dstDBConn *dbr.Connection, err error) {
	sql.Register("instrumented-postgres", instrumentedsql.WrapDriver(&pq.Driver{}))

	srcDB, err := c.makeConnection(c.makeDBInfo(c.options.SrcDBPort, c.options.SrcDBUser, c.options.SrcDBPassword))
	if err != nil {
		return nil, nil, err
	}

	dstDB, err := c.makeConnection(c.makeDBInfo(c.options.SrcDBPort, c.options.SrcDBUser, c.options.SrcDBPassword))
	if err != nil {
		srcDB.Close()
		return nil, nil, err
	}

	return srcDB, dstDB, nil
}

func (c *Copier) makeDBInfo(port int32, user string, password string) string {
	return fmt.Sprintf(
		"host=localhost port=%d user=%s password=%s dbname=test sslmode=disable",
		port,
		user,
		password,
	)
}

func (c *Copier) makeConnection(dbInfo string) (*dbr.Connection, error) {
	dbConnection, err := sql.Open("instrumented-postgres", dbInfo)
	if err != nil {
		return nil, err
	}

	err = dbConnection.Ping()
	if err != nil {
		return nil, err
	}

	return &dbr.Connection{
		DB:            dbConnection,
		Dialect:       dialect.PostgreSQL,
		EventReceiver: &dbr.NullEventReceiver{},
	}, nil
}
