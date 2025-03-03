// 코드 19.24 POST /register 엔드포인트 추가

mux.Get("/tasks", lt.ServeHTTP)
ru := &handler.RegisterUser{
	Service: &service.RegisterUser{DB: db, Repo: &r},
	Validator: v,
}
mux.Post("/register", ru.ServeHTTP)