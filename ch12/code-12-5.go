// 코드 12.5 요청 바디를 지정하는 미들웨어 구현 예

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"go.uber.org/zap"
)

func RequestBodyLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to log request body", zap.Error(err))
			http.Error(w, "Failed to get request body", http.StatusBadRequest)
			return
		}
		defer e.Body.Close()
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		next.ServeHTTP(w, r)
	})
}