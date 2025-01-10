// 코드 11.6 함수형을 이용한 DI

package main

func CutomApply(id int) error { return nil }

type Application struct {
	Apply func(int) error
}

func (app *Application) Run(id int) error {
	return app.Apply(id)
}

func main() {
	app := &Application{Apply: CutomApply}
	app.Run(19)
}