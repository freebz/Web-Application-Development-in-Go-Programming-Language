// 코드 19.15 handler/list_task_test.go의 수정 내용

-				sut := ListTask{Repo: &store.TaskStore{Tasks: tt.tasks}}
+				moq := &ListTasksServiceMock{}
+				moq.ListTasksFunc = func(ctx context.Context) (entity.Tasks, error) {
+								if tt.tasks != nil {
+												return tt.tasks, nil
+								}
+								return nil, errors.New("error from mock")
+				}
+				sut := ListTask{Service: moq}