// 코드 16.14 main_test.go를 server_test.go로 변경했다

				"golang.org/x/sync/errgroup"
)

-func TestRun(t *testing.T) {
-				t.Skip("리팩터리 중")
-
+func TestServer_Run(t *testing.T) {
				l, err := net.Listen("tcp", "localhost:0")
				if err != nil {
					t.Fatalf("failed to listen port %v", err)
				}
				ctx, cancel := context.WithCancel(context.Background())
				eg, ctx := errgroup.WithContext(ctx)
+				mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
+					fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
+				})
+
				eg.Go(func() error {
-					return run(ctx)
+					s := NewServer(l, mux)
+					return s.Run(ctx)