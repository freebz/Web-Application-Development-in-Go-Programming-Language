// 코드 5.3 v2 이상의 모듈 사용 방법

package main

// go.mod의 module에 작성돼 있는 module명으로 임포트한다.
import "github.com/labstack/echo/v4"

func main() {
	// 호출할 때는 v4.New()가 아니다.
	e := echo.New()

	// ...
}