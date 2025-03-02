// 코드 18.10 github.com/jmoiron/sqlx를 사용해서 얻은 복수의 레코드를 슬라이스한다

for (r *Repository) ListTasks(
	ctx context.Context, db *sqlx.DB,
) (entity.Tasks, error) {
	tasks := entity.Tasks{}
	sql := `SELECT
				id, user_id, title,
				status, created, modified
			FROM task;`
	if err := db.SelectContext(ctx, &tasks, sql); err != nil {
		return nil, err
	}
	return tasks, nil
}