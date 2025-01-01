// 코드 3.2 defer 문을 사용하지 않는 롤백

func (r *Repository) Update(ctx context.Context) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err := tx.Exec( /* 변경 처리1 */ )
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err := tx.Exec( /* 변경 처리2 */ )
	if err != nil {
		return err // Rollback 메서드 실행을 빠뜨렸다.
	}

	_, err := tx.Exec( /* 변경 처리 3 */ )
	if err != nil {
		tx.Rollback()
		return err
	}

	// 다른 처리
	return tx.Commit()
}