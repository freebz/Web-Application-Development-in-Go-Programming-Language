// 코드 9.3 고루틴과 변수 참조

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			// 대부분 "i: 5"가 5회 출력된다.
			fmt.Printf("i: %d\n", i)
			wg.Done()
		}()
	}
	wg.Wait()
	for j := 0; j < 5; j++ {
		wg.Add(1)
		go func(j int) {
			// 0부터 4까지 출력된다.
			fmt.Printf("j: %d\n", j)
			wg.Done()
		}(j)
	}
	wg.Wait()
	for k := 0; k < 5; k++ {
		k := k
		wg.Add(1)
		go func() {
			// 0부터 4까지 출력된다.
			fmt.Printf("k: %d\n", k)
			wg.Done()
		}()
	}
	wg.Wait()
}