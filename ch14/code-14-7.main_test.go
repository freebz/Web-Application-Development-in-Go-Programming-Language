// 코드 14.7 main 함수용 테스트

// main.test.go
package main

import "testing"

func TestMainFunc(t *testing.T) {
	go main()
	// 실행은 할 수 있지만 종료를 지시할 수 없다.
}