// 코드 19.4 run 함수 변경 내용

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/freebz/go_todo_app/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)
	mux, cleanup, err := NewMux(ctx, cfg)
	// 오류가 반환돼도 cleanup 함수를 실행한다.
	defer cleanup()
	if err != nil {
		return err
	}
	s := NewServer(l, mux)
	return s.Run(ctx)
}
