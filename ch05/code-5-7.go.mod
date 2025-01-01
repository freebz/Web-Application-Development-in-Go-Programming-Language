<!-- 코드 5.7 replace 명령으로 지정한 go.mod 파일 -->

module example.com/pkg
go 1.17
require example.com/othermodule v1.2.3
replace example.com/othermodule => ../othermodule