// 코드 16.8 종료 시그널을 기다리는 run 함수

package main

import (
				"context"
				"fmt"
				"log"
				"net"
				"net/http"
				"os"
+				"os/signal"
+				"syscall"
+				"time"

				"github.com/freebz/go_todo_app/config"
				"golang.org/x/sync/errgroup"
)

// main 함수 선언은 생략

func run(ctx context.Context) error {
+				ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
+				defer stop()
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
				s := &http.Server{
								// 인수로 받은 net.Listener를 사용하므로
								// Addr 필드는 지정하지 않는다.
								Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
+												// 명령줄에서 테스트하기 위한 로직
+												time.Sleep(5 * time.Second)
												fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])