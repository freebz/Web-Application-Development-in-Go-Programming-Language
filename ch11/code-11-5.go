// 코드 11.5 내장을 이용해서 의존성 주입하는 방법

package main

type OrderService interface {
	Apply(int) error
}

type ServiceImpl struct{}

func (s *ServiceImpl) Apply(id int) error { return nil }

type Application struct {
	OrderService // 내장 인터페이스
}

func (app *Application) Run(id int) error {
	return app.Apply(id)
}

func main() {
	// 초기화 시의 선언은 객체 초기화 시에 DI하는 방법과 같다.
	app := &Application{OrderService: &ServiceImpl{}}
	app.Run(19)
}