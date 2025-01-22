// 코드 14.8 run(ctx.context.Context) error 함수로 처리를 분리시킨 main 함수

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}