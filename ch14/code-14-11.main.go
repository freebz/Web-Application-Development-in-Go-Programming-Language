// 코드 14.11 외부의 종료 지시를 받아서 서버를 종료하는 run 함수

package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
)

// main 함수 정의는 생략

func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}

	eg, ctx := errgroup.WithContext(ctx)
	// 다른 고루틴에서 HTTP 서버를 실행한다.
	eg.Go(func() error {
		// http.ErrServerClosed 는
		// http.Server.Shutdonw()가 정상 종료된 것을 나타내므로 이상 처리가 아니다.
		if err := s.ListenAndServe(); err != nil &&
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