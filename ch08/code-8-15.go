// 코드 8.15 명명된 반환값과 recover를 조합할 수 있다

// Parse가 입력값인 문자열을 스페이스 단위로 단어를 분리한 후 integer로 반환한다.
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields)
	return
}