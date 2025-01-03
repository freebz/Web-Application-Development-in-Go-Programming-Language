// 코드 8.14 panic을 catch하는 간단한 recover

package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("보완한 panic: %v", r)
		}
	}()
	fmt.Println("출력된다")
	panic("happening!")
	fmt.Println("출력되지 않는다")
}

// Output:
// 출력된다
// 보완한 panic: happening!