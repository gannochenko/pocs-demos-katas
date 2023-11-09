package book

type Book struct {
	ID        string
	Title     string
	AuthorID  string
	IssueYear int32
}

type GetBooksResult struct {
	Books      []*Book
	Total      int64
	PageNumber int32
}

type DeleteBookResult struct {
}
