// 코드 8.6 Go 1.13부터 도입된 래핑

func (r *Repo) GetBook(t BookTitle)(*Book, error) {
	rows, err := db.QueryContext(ctx,
		"SELECT id, title, author_id FROM books WHERE title=?", title,
	)
	if err != nil {
		return nil, fmt.Errorf("GetBook: %w", err)
	}
	defer rows.Close()
	// ...
}

func (r *Repo) GetBook(t BookTitle)(*Book, error) {
	b, err := r.GetBook(t)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("GetAuthor: unknown %v", err)
		}
		// ...
}