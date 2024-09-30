package book

type Book struct {
	ID        string
	Title     string
	Author    string
	IssueYear int32
}

type GetBooksResult struct {
	Books      []*Book
	Total      int64
	PageNumber int32
}
