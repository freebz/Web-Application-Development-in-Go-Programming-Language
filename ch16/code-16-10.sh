# 코드 16.10 시그널 핸들링 구현 후

# signal.NotifyContext 사용한 경우
$ time curl -i localhost:28000/hello
HTTP/1.1 200 OK
Date: Sun, 29 May 2022 01:09:47 GMT
Content-Length: 13
Content-Type: text/plain; charset=utf-8
Connection: close

Hello, hello!curl -i localhost:28000/hello  0.00s user 0.01s system 0% cpu 5.025 total