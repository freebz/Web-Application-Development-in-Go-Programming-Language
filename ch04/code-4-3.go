// 코드 4.3 Book 타입만 존재하는 book 패키지

package book

import (
	"fmt"

	"person"
)

type Book struct {
	Author *person.Person
}

func (b *Book) AuthorName() string {
	// 패키지가 다르므로 Person, FirstName을 직접 참조할 수 없다.
	return fmt.Sprintf("%s %s", b.Author.GetFirstName(), b.Author.GetLastName())
}