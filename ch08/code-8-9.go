// 코드 8.9 errors.As 함수를 사용해 RDBMS 오류 정보를 가져온다

package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// MySQLDuplicateEntryErrorCode 중복 데이터가 있는 경우의 MySQL 오류 코드
// ref: https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html
const MySQLDuplicateEntryErrorCode uint16 = 1062

var ErrAlreadyExists = errors.New("duplicate entry")

func (repo *Repo) SaveBook(ctx context.Context, book *Book) error {
	r, err := repo.db.ExecContext(ctx /* ... */)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr); mysqlErr.Number == MySQLDuplicateEntryErrorCode {
			return fmt.Errorf("store: cannot save book_id %d: %w", book.ID, ErrAlreadyExists)
		}
		// *mysql.MySQLError 이외의 오류인 경우 처리
	}
	// ...
}