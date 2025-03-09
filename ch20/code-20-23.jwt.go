// 코드 20.23 전달된 context.Context 타입값에서 관리자 권한 유무를 확인한다

func IsAdmin(ctx context.Context) bool {
	role, ok := GetRole(ctx)
	if !ok {
		return false
	}
	return role == "admin"
}