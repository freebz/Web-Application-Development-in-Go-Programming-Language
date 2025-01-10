// 코드 11.3 setter를 준비해서 DI하는 방법

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



func (app *Application) Apply(id int) error {
	return app.os.Apply(id)
}


func (app *Application) SetService(os OrderService) {
	app.os = os
}

func main() {
	app := &Application{}
	app.SetService(&ServiceImpl{})
	app.Apply(19)
}