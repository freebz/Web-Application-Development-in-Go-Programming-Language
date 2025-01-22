// 코드 14.16 함수 외부에서 네트워크 리스너를 받는다

func run(ctx context.Context, l net.Listener) error {
	s := &http.Server{
		// 인수로 받은 net.Listener를 이용하므로 Addr 필드는 지정하지 않는다.
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		// ListenAndServe 메서드가 아닌 Serve 메서드로 변경한다.
		if err := s.Serve(l); err != nil &&
			// http.ErrServerClosed는
			// http.Server.Shutdown()가 정상 종료됐다고 표시하므로 문제없다.
			err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}

// 이후 코드는 동일하다.