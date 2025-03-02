// 코드 18.17 RDBMS에 태스크를 등록하는 메서드

package store

import (
	"context"

	"github.com/freebz/go_todo_app/entity"
)

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
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = entity.TaskID(id)
	return nil
}