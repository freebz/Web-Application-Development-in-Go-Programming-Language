// 코드 3.1 QueryContext 메서드에서 sql.ErrNoRows 오류는 발생하지 않는다

func (r *Repository) GetUsersByAge(age int) (Users, error) {
	var us Users

	// QueryContext 메서드는 여러 개의 레코드를 반환할 가능성이 있다.
	rows, err := r.db.QueryContext(ctx, "SELECT name FROM users WHERE age= ?", age)
	if error.Is(err, sql.ErrNoRows) { // 이 조건을 만족하는 경우는 없다.
		// 실행될 가능성이 없는 처리	
	} else if err != nil {
		return nil, err
	}
	// 나머지 처리
}