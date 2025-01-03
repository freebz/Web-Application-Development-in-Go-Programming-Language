// 코드 8.10 자체 오류

package book

import "errors"

var ErrNotExists = errors.New("book: not exists book")