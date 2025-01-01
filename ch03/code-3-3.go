// 코드 3.3 defer 문으로 롤백하기

func (r *Repository) Update(ctx context.Context) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	
	defer tx.Rollback()
	
	_, err := tx.Exec( /* 변경 처리1 */ )
	if err != nil {
		return err
	}

	_, err := tx.Exec( /* 변경 처리2 */ )
	if err != nil {
		return err
	}

	_, err := tx.Exec( /* 변경 처리3 */ )
	if err != nil {
		return err
	}

	return tx.Commit()
}