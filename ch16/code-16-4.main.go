// 코드 16.4 config 패키지를 사용하도록 수정한 run 함수

-func fun(ctx context.Context, l net.Listen) error {
+func run(ctx context.Context) error {
+			 cfg, err := config.New()
+			 if err != nil {
+							 return err
+			 }
+			 l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
+			 if err != nil {
+							 log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
+			 }
+			 url := fmt.Sprintf("http://%s", l.Addr().String())
+			 log.Printf("start with: %v", url)
			 s := &http.Server{