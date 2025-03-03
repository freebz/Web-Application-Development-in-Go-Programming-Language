// 코드 19.9 리팩터링 전후의 handler.AddTask 타입 수정 내용

				"net/http"

				"github.com/freebz/go_todo_app/entity"
-				"github.com/freebz/go_todo_app/store"
				"github.com/go-playground/validator/v10"
-				"github.com/jmoiron/sqlx"
)

type AddTask struct {
-				DB *sqlx.DB
-				Repo store.Repository
+				Service AddTaskService
				Validator *validator.Validate
}

// 중략
								return
				}

-				t := &entity.Task{
-							 Tible: b.Title,
-							 Status: entity.TaskStatusTodo,
-				}
-				err := at.Repo.AddTask(ctx, at.DB, t)
+				t, err := at.Service.AddTask(ctx, b.Title)