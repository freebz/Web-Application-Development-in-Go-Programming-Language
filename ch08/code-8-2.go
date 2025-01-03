// 코드 8.2 오류 관련 표준 패키지의 함수

// errors 패키지
func As(err error, target interface{}) bool
func Is(err, target error) bool
func New(text string) error
func Unwrap(err error) error

// fmt 패키지
func Errorf(format string, a ...interface{}) error