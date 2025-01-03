// 코드 8.11 간단한 자체 오류 구조체

type ErrCode int
type MyError struct {
	Code ErrCode
}

func (e *MyError) Error() string {
	return fmt.Sprintf("code: %d", e.Code)
}