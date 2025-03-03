// 코드 19.17 수정 전후의 NewMux 함수

// 중략
				"github.com/freebz/go_todo_app/clock"
				"github.com/freebz/go_todo_app/config"
				"github.com/freebz/go_todo_app/handler"
+				"github.com/freebz/go_todo_app/service"
				"github.com/freebz/go_todo_app/store"
				"github.com/go-chi/chi/v5"
				"github.com/go-playground/validator/v10"
// 중략
								return nil, cleanup, err
				}
				r := store.Repository{Clocker: clock.RealClocker{}}
-				at := &handler.AddTask{DB: db, Repo: r, Validator: v}
+				at := &handler.AddTask{
+								Service: &service.AddTask{DB: db, Repo: &r},
+								Validator: v,
+				}
				mux.Post("/tasks", at.ServeHTTP)
-				lt := &handler.ListTask{DB: db, Repo: r}
+				lt := &handler.ListTask{
+								Service: &service.ListTask{DB: db, Repo: &r},
+				}
				mux.Get("/tasks", lt.ServeHTTP)
				return mux, cleanup, nil
}