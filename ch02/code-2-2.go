// 코드 2.2 항상 context.Context 타입값을 전달한다

func Handle(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ID int
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		// 오류 처리
	}
	b, err = GetBook(r.Context, body.ID)
	// 남은 처리
}

// 로직 내에 context를 사용할 예정이 없는 메서드
func GetBook(ctx context.Context, id int) (*Book, error{
	// 호출한 함수나 메서드에서 context가 필요한 경우도 있디.
	rows, err := db.QueryContext(ctx, "SELECT id, name, isdn, price FROM books WHERE id=?", id)
	// 남은 처리
})