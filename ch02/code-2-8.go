// 코드 2.8 context 패키지를 사용한 인증/권한 처리 아이디어

import "example.com/auth"
type MyDB struct {
	db *sql.DB
}

// ExecContext는 *sql.DB와 같은 명칭의 메서드를 context 객체를 통해
// 인증 기능으로 래핑(wrapping)한 메서드
func (mydb *MyDB) ExecContext(ctx context.Context, query string, args ...any)
(Result, error) {
	if !auth.Writeable(ctx) {
		return errors.New("작성 권한이 없는 사용자가 실행했다.")
	}
	return mydb.db.ExecContext(ctx, query, args...)
}