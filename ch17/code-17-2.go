// 코드 17.2 Defined Type을 사용해 잘못된 대입 방지

func main() {
	var id int = 1
	// TaskID 타입에 변환한 후에 대입하고 있으므로 문제없다.
	_ = Task{ID: TaskID(id)}
	// 빌드 오류
	// cannot use id (variable of type int) as type TaskID in struct literal
	_ = Task{ID: id}

	// 타입 추론으로 TaskID 타입의 1이 되므로 빌드 오류가 발생하지 않는다.
	_ = Task{ID: 1}
}