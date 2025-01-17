// 코드 12.1 GO의 일반적인 미들웨어 패턴

import "net/http"

func(h http.Handler) http.Handler