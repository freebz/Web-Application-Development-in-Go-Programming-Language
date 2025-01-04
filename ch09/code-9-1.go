// 코드 9.1 http.HandlerFunc 타입

type HandlerFunc func(ResponseWrite, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWrite, r *Request)