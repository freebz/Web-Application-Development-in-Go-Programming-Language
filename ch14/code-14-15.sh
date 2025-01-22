# 코드 14.15 오류 발생

$ go test -v ./...
2025/01/22 01:54:00 failed to close: listen tcp :18080: bind: address already in use
--- FAIL: TestRun (0.00s)
    main_test.go:42: listen tcp :18080: bind: address already in use
FAIL
exit status 1
FAIL	github.com/freebz/go_todo_app	0.500s