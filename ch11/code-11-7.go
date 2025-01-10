// 코드 11.7 컴파일러를 사용한 구현 확인

type Jedi interface{
	HasForce() bool
}

type Knight struct {}

// 이 상태에서는 컴파일 오류가 발생하므로 구현이 불충분하다는 것을 알 수 있다.
var _ Jedi = (*Knight)(nil)