// 코드 8.3 errors.new 함수와 fmt.Errorf 함수

func GetAuthor(id AuthorID) (*Author, error) {
	if !id.Valid() {
		return nil, errors.New("GetAuthor: id is invalid")
		// 오류문에 id값을 사용할 때는 Errorf 함수를 사용한다.
		// return nil, fmt.Errorf("GetAuthor: id(%d) is invalid", id)
	}
	// do anything...
}

func GetAuthor(b *Book) (string, error) {
	a, err := GetAuthor(b.AuthorID)
	if err != nil {
		return "", fmt.Errorf("GetAuthorName: %v", err)
	}
	return a.Name(), nil
}