# 코드 16.9 make build 명령 실행

$ make build
docker build -t freebz/gotodo:latest \
--target deploy ./

[+] Building 7.5s (17/17) FINISHED

$ docker run -p 28000:80 freebz/gotodo:latest
2022/05/28 08:02:29 start with: http://[::]:80