// 코드 19.10 handler/add_task_test.go 수정

+									moq := &AddTaskServiceMock{}
+									moq.AddTaskFunc = func(
+										ctx context.Context, title string,
+									) (*entity.Task, error) {
+										if tt.want.status == http.StatusOK {
+											return &entity.Task{ID: 1}, nil
+										}
+										return nil, errors.New("error from mock")
+									}
									sut := AddTask{
-										Repo: &store.TaskStore{
-											Tasks: map[entity.TaskID]*entity.Task{},
-										},
+										Service: moq,