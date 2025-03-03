// 코드 19.1 handler 패키지에서 store 패키지를 사용한다

    "net/http"
-   "time"
    "github.com/freebz/go_todo_app/entity"
    "github.com/freebz/go_todo_app/store"
    "github.com/go-playground/validator/v10"
+   "github.com/jmoiron/sqlx"
)

type AddTask struct {
-      Store     *store.TaskStore
+      DB *sqlx.DB
+      Repo *store.Repository
       Validator *validator.Validate
}

// 중략
       }

       t := &entity.Task{
-             Title: b.Title,
-             Status: entity.TaskStatusTodo,
-             Created: time.Now(),
+             Title: b.Title,
+             Status: entity.TaskStatusTodo,
  }
-     id, err := at.Store.Add(t)
+     err := at.Repo.AddTask(ctx, at.DB, t)
// 중략

-   }{ID: id}
+   }{ID: t.ID}
    RespondJSON(ctx, w, rsp, http.StatusOK)
  }