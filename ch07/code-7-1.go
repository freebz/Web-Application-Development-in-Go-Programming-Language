// 코드 7.1 어떤 인터페이스를 따르고 있는지 모른다

// 별도로 명시하지 않지만 Read 메서드가 있는 *File 객체는
// io.Reader 인터페이스를 따르고 있다.
type File struct{}

func (f *File) Read(p []byte) (n int, err error){ return 0, nil}