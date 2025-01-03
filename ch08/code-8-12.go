// 코드 8.12 자체 오류 구조체에서 연속 오류를 저장하기 위한 메서드 정의

func (e *MyError) Unwrap() error
func (e *MyError) As(target interface{}) bool
func (e *MyError) Is(target error) bool