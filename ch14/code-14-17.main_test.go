// 코드 14.17 run 함수용 테스트 수정

func TestRun(t *testing.T) {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen port %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx, l)
	})
	in := "message"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)
	// 어떤 포트 번호로 리슨하고 있는지 확인
	t.Logf("try request to %q", url)
	rsp, err := http.Get(url)

// 이후 코드는 동일하다.