// 코드 19.12 store 패키지의 직접 참조를 피하기 위한 인터페이스

package service

import (
	"context"

	"github.com/freebz/go_todo_app/entity"
	"github.com/freebz/go_todo_app/store"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . TaskAdder TaskLister
type TaskAdder interface {
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}
type TaskLister interface {
	ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}