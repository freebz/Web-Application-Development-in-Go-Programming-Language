// 코드 20.47 admin 권한만 접속할 수 있는 앤드포인트

mux.Route("/admin", func(r chi.Router) {
	r.Use(handler.AuthMiddleware(jwter), handler.AdminMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"message": "admin only"}`))
	})
})