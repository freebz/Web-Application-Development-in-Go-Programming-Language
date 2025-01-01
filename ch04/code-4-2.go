// 코드 4.2 Person 타입만 존재하는 person 패키지

package person

type Person struct {
	firstName string
	lastName string
}

func (p *Person) FirstName() string { return p.firstName }

func (p *Person) LastName() string { return p.lastName }