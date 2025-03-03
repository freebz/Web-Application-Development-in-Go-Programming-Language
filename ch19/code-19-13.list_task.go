// 코드 19.13 handler.ListTask 타입의 리팩터링 전후 내용

				"net/http"

				"github.com/freebz/go_todo_app/entity"
-				"github.com/freebz/go_todo_app/store"
-				"github.com/jmoiron/sqlx"
 )

 type ListTask struct {
-				DB *sqlx.DB
-				Repo store.Repository
+				Service ListTasksService
 }

type task struct {
// 중략

 func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
-				tasks, err := lt.Repo.ListTasks(ctx, lt.DB)
+				tasks, err := lt.Service.ListTasks(ctx)
				if err != nil {
					RespondJSON(ctx, w, &ErrResponse{
						Message: err.Error(),