// 코드 20.44 사용자 ID를 이용해 태스크를 검색한다

func (r *Repository) ListTasks(
	ctx context.Context, db Queryer, id entity.UserID,
) (entity.Tasks, error) {
	tasks := entity.Tasks{}
	sql := `SELECT
				id, user_id, title,
				status, created, modified
			FROM task
			WHERE user_id = ?;`
	if err := db.SelectContext(ctx, &tasks, sql, id); err != nil {
		return nil, err
	}
	return tasks, nil
}