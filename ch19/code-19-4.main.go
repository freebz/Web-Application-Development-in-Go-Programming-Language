// 코드 19.4 run 함수 변경 내용

				}
				url := fmt.Sprintf("http://%s", l.Addr().String())
				log.Printf("start with: %v", url)
-				mux := NewMux()
+				mux, cleanup, err := NewMux(ctx, cfg)
+				// 오류가 반환돼도 cleanup 함수를 실행한다.
+				defer cleanup()
+				if err != nil {
+
+				}
				s := NewServer(l, mux)