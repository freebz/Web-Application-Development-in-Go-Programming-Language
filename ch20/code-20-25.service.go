// 코드 20.25 LoginService 인터페이스 정의

package handler

import (
	"context"

	"github.com/freebz/go_todo_app/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . ListTasksService AddTaskService RegisterUserService LoginService
// 다를 인터페이스 정의는 생략

type LoginService interface {
	Login(ctx context.Context, name, pw string) (string, error)
}