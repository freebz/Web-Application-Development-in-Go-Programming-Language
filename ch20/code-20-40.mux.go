// 코드 20.40 /tasks 엔드포인트에 AuthMiddleware를 적용한다

at := &handler.AddTask{
	Service: &service.AddTask{DB: db, Repo: &r},
	Validator: v,
}
mux.Post("/tasks", at.ServeHTTP)
lt := &handler.ListTask{
	Service: &service.ListTask{DB: db, Repo: &r},
}
mux.Get("/tasks", lt.ServeHTTP)
mux.Route("/tasks", func(r chi.Router) {
	r.Use(handler.AuthMiddleware(jwter))
	r.Post("/", at.ServeHTTP)
	r.Get("/", lt.ServeHTTP)
})