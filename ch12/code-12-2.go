// 코드 12.2 미들웨어 패턴 구현

import (
	"log"
	"net/http"
	"time"
)

func MyMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := time.Now()
		h.ServeHTTP(w, r)
		d := time.Now().Sub(s).Millisecond()
		log.Printf("end %s(%d ms)\n", t.format(time.RFC3339), d)
	})
}