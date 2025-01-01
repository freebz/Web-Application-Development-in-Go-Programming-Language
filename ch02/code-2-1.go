// 코드 2.1 context.Context 인터페이스 정의

type Context interface {
  Deadline(deadline time.Time, ok bool)
  Done() <-chan struct{}
  Err() error
  Value(key interface{}) interface{}
}