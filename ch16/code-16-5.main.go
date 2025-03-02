// 코드 16.5 run 함수만 호출하도록 되돌아간 main 함수

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}