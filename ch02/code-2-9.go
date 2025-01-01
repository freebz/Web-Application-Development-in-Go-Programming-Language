// 코드 2.9 context.TODO를 사용한 단계적인 코드 변환

// context 미적용 함수, 오류 처리는 생략
func GetCompanyUsecase(userID UserID) (*Company, error) {
	// context 적용 완료 함수
	u, _ := GetUser(context.TODO(), userID)
	// context 미적용 함수
	c, _ := GetCompanyByUser(u)
	return c
}

func GetUser(ctx context.Context, id UserID) (*User, error) {
	// 어떤 처리
}

func GetCompanyByUserID(u *User) (*Company, error) {
	// 어떤 처리
}