// 코드 14.9 http.Server 타입을 사용한 HTTP 서버 실행

s := &http.Server{
	Addr: ":18080",
	Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	}),
}
s.ListenAndServe()