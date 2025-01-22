# 코드 14.14 테스트가 실패하는 것을 확인한다

$ go test -v ./...
=== RUN   TestRun
    main_test.go:34: want "Hi, message!", but got "Hello, message!"
--- FAIL: TestRun (0.00s)
FAIL
FAIL	github.com/freebz/go_todo_app	0.500s
FAIL