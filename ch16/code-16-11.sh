# 코드 16.11 동작 차이 확인

# signal.NotifyContext 사용하지 않은 경우
$ time curl -i localhost:28000/hello
curl: (52) Empty reply from server
curl -i localhost:28000/hello  0.00s user 0.01s system 0% cpu 2.317 total