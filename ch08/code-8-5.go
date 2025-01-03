// 코드 8.5 sql.ErrNoRows와 비교해도 참이 되지 않는다

var db *sql.DB

func (r *Repo) GetBook(t BookTitle)(*Book, error) {
	rows, err := db.QueryContext(ctx,
		"SELECT id, title, author_id FROM books WHERE title=?", title,
	)
	if err != nil {
		return nil, fmt.Errorf("GetBook: %v", err)
	}
	defer rows.Close()
	// ...
}

func GetAuthorName(t BookTitle) (string, error) {
	b, err := r.GetBook(t)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("GetAuthor: unknown book %v", err)
		}
		// ...
	}
}