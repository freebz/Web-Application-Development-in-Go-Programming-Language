// 코드 20.12 go:embed 디렉티브를 사용한 파일 첨부

package auth

import (
	_ "embed"
)

//go:embed cert/secret.pem
var rawPrivKey []byte

//go:embed cert/public.pem
var rawPubKey []byte
