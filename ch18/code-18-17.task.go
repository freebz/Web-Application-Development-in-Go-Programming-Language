// 코드 18.17 RDBMS에 태스크를 등록하는 메서드

func (r *Repository) AddTask(
	ctx context.Context, db Execer, t *entity.Task,
) error {
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	sql := `INSERT INTO task
		(title, status, created, modified)
	VALUES (?, ?, ?, ?)`
	result, err := db.ExecContext(
		ctx, sql, t.Title, t.Status,
		t.Created, t.Modified,
	)
	if err != nil {
		return err
	}
	t.ID = entity.TaskID(id)
	return nil
}