// 코드 4.1 unexported와 private의 차이

package main

import "fmt"

type Person struct {
	firstName string
	lastName string
}

func (p *Person) GetFirstName() string { return p.firstName }

type Book struct {
	Author *Person
}

func (b *Book) AuthorName() string {
	// 같은 패키지 안이므로 Person 구조체의 unexported 필드를 직접 참조할 수 있다.
	return fmt.Sprintf("%s %s", b.Author.firstName, b.Author.lastName)
}