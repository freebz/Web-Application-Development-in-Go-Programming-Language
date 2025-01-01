// 코드 2.4 WithDeadline 함수와 WithTimeout 함수

func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)