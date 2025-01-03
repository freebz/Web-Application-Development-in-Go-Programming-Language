// 코드 8.13 error 타입값을 반환할 때는 부가 정보를 추가한다

func AddBook(ctx context.Cotext, b *Book) error {
	if err := repo.SaveBook(ctx, b); err != nil {
		// 앞선 오류의 타입 정보를 저장하려면 %w, 그렇지 않다면 %v
		return fmt.Errorf("AddBook: %v", err)
	}
	return nil
}