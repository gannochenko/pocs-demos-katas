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
	// TableName holds the name of the table which contains elements we copy
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

// Element describes a structure of an element we copy.
// The field types could be as well just interface{}, because simply copy-paste it without manipulation.
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

// CopyElements copies elements under given IDs from the source database to a destination database
func (c *Copier) CopyElements(ids []string) error {
	srcDB, dstDB, err := c.makeConnections()
	if err != nil {
		return err
	}
	defer srcDB.Close()
	defer dstDB.Close()

	srcDBSession := srcDB.NewSession(nil)
	dstDBSession := dstDB.NewSession(nil)

	// for the destination table session we create a transaction,
	// so in case if there is a problem, the entire operation gets rolled back
	dstDBTransaction, err := dstDBSession.Begin()
	if err != nil {
		return err
	}
	defer dstDBTransaction.RollbackUnlessCommitted()

	ctx := context.Background()

	processed := 0
	for _, id := range ids {
		// for every ID we get an element from the source database
		// absence of en element is not an error
		element, err := c.getElementByID(ctx, srcDBSession, id)
		if err != nil {
			fmt.Printf("Could not read element with id %s\n", id)
			return err
		}
		if element.ID == "" {
			fmt.Printf("Could not read element with id %s: element not found\n", id)
			continue
		}

		fmt.Printf("Copying element %s\n", id)

		err = c.saveElement(ctx, dstDBTransaction, element)
		if err != nil {
			fmt.Printf("Could not create an element with id %s\n", element.ID)
			return err
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

// getElementByID returns an element from the source database, by it's ID
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

// saveElement saves a given element to the destination database
func (c *Copier) saveElement(ctx context.Context, session *dbr.Tx, element *Element) error {
	columns := []string{"id", "title"}

	// to avoid primary key collisions, the ID column must be different
	element.ID = uuid.New().String()
	err := session.InsertInto(TableName).Columns(columns...).Record(element).LoadContext(ctx, &element.ID)
	if err != nil {
		return err
	}

	return nil
}

// makeConnections creates two connections: for a database to read from, and for a database to write to
func (c *Copier) makeConnections() (srcDBConn *dbr.Connection, dstDBConn *dbr.Connection, err error) {
	sql.Register("instrumented-postgres", instrumentedsql.WrapDriver(&pq.Driver{}))

	// making a connection to an instance we read from
	srcDB, err := c.makeConnection(c.makeDBInfo(c.options.SrcDBPort, c.options.SrcDBUser, c.options.SrcDBPassword))
	if err != nil {
		return nil, nil, err
	}

	// making a connection to an instance we write to
	dstDB, err := c.makeConnection(c.makeDBInfo(c.options.DstDBPort, c.options.DstDBUser, c.options.DstDBPassword))
	if err != nil {
		srcDB.Close()
		return nil, nil, err
	}

	return srcDB, dstDB, nil
}

// makeDBInfo returns a connection string for a chosen instance
func (c *Copier) makeDBInfo(port int32, user string, password string) string {
	return fmt.Sprintf(
		"host=localhost port=%d user=%s password=%s dbname=test sslmode=disable",
		port,
		user,
		password,
	)
}

// makeConnection makes a connection using a provided connection string
func (c *Copier) makeConnection(dbInfo string) (*dbr.Connection, error) {
	// make a connection, using a specific driver
	dbConnection, err := sql.Open("instrumented-postgres", dbInfo)
	if err != nil {
		return nil, err
	}

	// check if the connection is really there
	err = dbConnection.Ping()
	if err != nil {
		return nil, err
	}

	// create a dbr wrapping connection
	return &dbr.Connection{
		DB:            dbConnection,
		Dialect:       dialect.PostgreSQL,
		EventReceiver: &dbr.NullEventReceiver{},
	}, nil
}
