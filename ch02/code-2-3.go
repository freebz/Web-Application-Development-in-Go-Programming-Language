// 코드 2.3 WithCancel 함수를 사용한 중단할 수 있는 context.Context 타입값

package main

import (
	"context"
	"fmt"
)

func child(ctx context.Context) {
	// 함수 처리를 시작하기 전에 context.Context 상태를 검증한다.
	if err := ctx.Err(); err != nil {
		return
	}
	fmt.Println("중단되지 않음")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	child(ctx)
	cancel()
	child(ctx)
}