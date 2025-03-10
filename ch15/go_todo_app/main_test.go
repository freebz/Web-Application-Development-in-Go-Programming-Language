// 코드 14.17 run 함수용 테스트 수정

package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net"
	"net/http"
	"testing"
)

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
	if err != nil {
		t.Errorf("failed to get: %v+v", err)
	}
	defer rsp.Body.Close()
	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	// HTTP 서버의 반환값을 검증한다.
	want := fmt.Sprintf("Hello, %s!", in)
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}

	// run 함수에 종료 알림을 전송한다.
	cancel()

	// run 함수의 반환값을 검증한다.
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}