// 코드 19.23 store/repository.go 파일에 정의하는 범용 오류 정의

const (
	// ErrCodeMySQLDuplicateEntry는 MySQL의 데이터 중복 오류 코드
	// https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html
	// Error number: 1062; Symbol: ER_DUP_ENTRY; SQLSTATE: 23000
	ErrCodeMySQLDuplicateEntry = 1062
)

var (
	ErrAlreadyEntry = errors.New("duplicate entry")
)