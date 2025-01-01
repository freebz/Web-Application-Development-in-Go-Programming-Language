// 코드 6.4 Go로 코드 6.1과 같은 부모 자식 관계를 표현한 경우

type Person struct {
	Name string
	Age  int
}

// Person을 내장한 Korean 타입
type Korean struct {
	Person
	MyNumber int
}

func Hello(p Person) {
	fmt.Println("Hello " + p.Name)
}