package author

import (
	"context"

	"gorm.io/gorm"
)

const (
	TableName = "authors"
)

type Repository struct {
	Session *gorm.DB
}

func (r *Repository) RefreshHasBooksFlag(ctx context.Context, condition interface{}) (err error) {
	_, err = r.Session.Exec(`
		UPDATE
			authors
		SET has_books = bookinfo.book_count > 0
		FROM
			(
				SELECT
					authors.id AS author_id,
					count(books.*) AS book_count
				FROM
					authors
						LEFT JOIN books
								  ON books.author_id = authors.id
				WHERE
				    books.deleted_at is null
				    AND
				    authors.id in (?)
				GROUP BY
					authors.id
			) AS bookinfo
		WHERE
			authors.id = bookinfo.author_id
	`, condition).Rows()

	return err
}
