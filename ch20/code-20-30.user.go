// 코드 20.30 entity/user.go 파일에 추가한 패스워드 검증 로직

func (u *User) ComparePassword(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw))
}