// 코드 7.3 fmt.Println 함수가 false를 출력한다

package main

import "fmt"

type MyErr struct{}

func (me *MyErr) Error() string { return "" }

func Apply() error {
	var err *MyErr = nil
	// 봔환값에 타입 정보가 포함돼 있으므로 예상 밖의 처리를 한다.
	return err
}

func Apply2() error {
	var err error = nil
	// 결과적으로 문제가 없지만,
	// 명시적으로 nil을 return해서 실수를 방지하는 것이 좋다.
	return err
}

func main() {
	fmt.Println(Apply() == nil)  // false
	fmt.Println(Apply2() == nil) // true
}