// 코드 20.31 store/user.go 파일에 추가한 *store.Repository.GetUser 메서드

func (r *Repository) GetUser(
	ctx context.Context, db Queryer, name string,
) (*entity.User, error) {
	u := &entity.User{}
	sql := `SELECT
		id, name, password, role, created, modified
		FROM user WHERE name = ?`
	if err := db.GetContext(ctx, u, sql, name); err != nil {
		return nil, err
	}
	return u, nil
}