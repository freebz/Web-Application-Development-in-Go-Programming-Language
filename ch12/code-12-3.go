// 코드 12.3 미들웨어 패턴을 반환한다.

import "net/http"

// VersionAdder 함수는 다음과 같이 사용할 수 있는 미들웨어 함수를 반환한다.
// vwm := VersionAdder("1.0.1")
// http.Handle("/users", vwm(userHandler))
func VersionAdder(v AppVersion) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Add("App-Version", v)
			next.ServeHTTP(w, r)
		})
	}
}