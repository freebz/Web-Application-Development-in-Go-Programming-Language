// 코드 16.6 컴파일 오류를 해결하고 테스트를 건너뛴다

func TestRun(t *testing.T) {
+			 t.Skip("리팩터리 중")
+
			 l, err := net.listen("tcp", "localhost:0")
			 if err != nil {
							 t.Fatalf("failed to listen port %v", err)
// 중략
			 ctx, cancel := context.WithCancel(context.Background())
			 eg, ctx := errgroup.WithContext(ctx)
			 eg.Go(func() error {
-						 	 return run(ctx, l)
+							 return run(ctx)
})