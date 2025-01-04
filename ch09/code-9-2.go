// 코드 9.2 익명 함수는 상태를 가질 수 있다.

package main

import "fmt"

func store() func(int) int {
	var x int
	return func(n int) int {
		x += n
		return x
	}
}

func main() {
	s1 := store()
	s2 := store()
	fmt.Printf("s1: %d, s2: %d\n", s1(1), s2(2)) // s1: 1, s2: 2
	fmt.Printf("s1: %d, s2: %d\n", s1(1), s2(2)) // s1: 2, s2: 4
	fmt.Printf("s1: %d, s2: %d\n", s1(1), s2(2)) // s1: 3, s2: 6
}