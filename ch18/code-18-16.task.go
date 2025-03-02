// 코드 18.16 RDBMS로부터 태스크를 가져오는 메서드

func (r *Repository) ListTasks(
	ctx context.Context, db Queryer,
) (entity.Tasks, error) {
	tasks := entity.Tasks{}
	sql := `SELECT
			id, title,
			status, created, modified
		FROM task;`
	if err := db.SelectContext(ctx, &tasks, sql); err != nil {
		return nil, err
	}
	return tasks, nil
}