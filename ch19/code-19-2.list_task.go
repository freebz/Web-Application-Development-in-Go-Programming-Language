// 코드 19.2 handler.ListTask 타입의 수정 내용

				"github.com/freebz/go_todo_app/store"
+				"github.com/jmoiron/sqlx"
 )

 type ListTask struct {
-				Store *store.TaskStore
+				DB *sqlx.DB
+				Repo *store.Repository
 }

 type task struct {
// 중략

 func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
-				tasks := lt.Store.All()
+				tasks, err := lt.Repo.ListTasks(ctx, lt.DB)
+				if err != nil {
+						RespondJSON(ctx, w, &ErrResponse{
+										Message: err.Error(),
+						}, http.StatusInternalServerError)
+						return
+		}
		rsp := []task{}
		for _, t := range tasks {