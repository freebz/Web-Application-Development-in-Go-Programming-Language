// 코드 11.4 메서드(함수) 호출 시에 DI하는 방법

package main

// 상세 구현
type ServiceImpl struct{}

func (s *ServiceImpl) Apply(id int) error { return nil }

// 상위 계층이 정의하는 추상
type OrderService interface {
	Apply(int) error
}

// 상위 계층 사용자측 타입
type Application struct {
	os OrderService
}



func (app *Application) Apply(os OrderService, id int) error {
	return os.Apply(id)
}

func main() {
	app := &Application{}
	app.Apply(&ServiceImpl{}, 19)
}