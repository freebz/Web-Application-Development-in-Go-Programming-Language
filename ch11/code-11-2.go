// 코드 11.2 객체 초기화 시에 DI하는 방법

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

// 다른 언어의 생성자 주입에 해당하는 구현
func NewApplication(os OrderService) *Application {
	return &Application{os: os}
}
func (app *Application) Apply(id int) error {
	return app.os.Apply(id)
}

func main() {
	app := NewApplication(&ServiceImpl{})
	app.Apply(19)
}