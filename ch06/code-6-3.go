// 코드 6.3 내장을 사용한 개(Dog) 구현

package main

import "fmt"

type Dog struct{}

func (d *Dog) Bark() string { return "Bow" }

// *Dog의 Bark 메서드를 사용할 수 있는 구조체
type BullDog struct{ Dog }

type ShibaInu struct{ Dog }

// *Dog의 Bark 메서드를 덮어쓰기 하는 메서드
func (s *ShibaInu) Bark() string { return "멍" }

func DogVoice(d *Dog) string     { return d.Bark() }

func main() {
	bd := &BullDog{}
	fmt.Println(bd.Bark())
	si := &ShibaInu{}
	fmt.Println(si.Bark())

	// cannot use si (type *ShibaInu) as type *Dog in argument to DogVoice
	// fmt.Println(DogVoice(si))
}