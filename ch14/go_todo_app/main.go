// 코드 14.18 run 함수에 맞춰 main 함수 수정

package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("need port number\n")
		os.Exit(1)
	}
	p := os.Args[1]
	l, err := net.Listen("tcp", ":"+p)
	if err != nil {
		log.Fatalf("failed to listen port %s: %v", p, err)
	}
	if err := run(context.Background(), l); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

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
		return nil
	})

	// 채널로부터의 알림(종료 알림)을 기다린다.
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}
	// Go 메서드로 실행한 다른 고루틴의 종료를 기다린다.
	return eg.Wait()
}