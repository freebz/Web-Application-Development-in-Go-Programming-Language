// 코드 19.14 handler/list_task_test.go

			tests := map[string]struct {
-							tasks map[entity.TaskID]*entity.Task
+							tasks []*entity.Task
							want want
			}{
							"ok": {
-										 tasks: map[entity.TaskID]*entity.Task{
-														 1: {
+										 tasks: []*entity.Task{
+														 {
																		 ID: 1,
																		 Title: "test1",
																		 Status: entity.TaskStatusTodo,
														 },
-														 2: {
+														 {
																	 	 ID: 2,
																		 Title: "test2",
																		 Status: entity.TaskStatusDone,
// 중략									
											},
							},
							"empty": {
-										 tasks: map[entity.TaskID]*entity.Task{},
+										 tasks: []*entity.Task{},