// 코드 7.2 Modem 타입의 Recv 작업에만 주목한 인터페이스 정의

type Modem struct {}
func (Modem) Dial() {}
func (Modem) Hangup() {}
func (Modem) Sender() {}
func (Modem) Recv() {}

// Recv 메서드에만 주목한 인터페이스 정의
type Receiver interface {
	Recv()
}