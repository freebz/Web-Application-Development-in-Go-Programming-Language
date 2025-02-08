# 코드 15.5 컨테이너 빌드

docker compose build --no-cache
[+] Building 12.6s (7/7) FINISHED                          docker:desktop-linux
 => [app internal] load build definition from Dockerfile                   0.0s
 => => transferring dockerfile: 770B                                       0.0s
 => [app internal] load .dockerignore                                      0.0s
 => => transferring context: 110B                                          0.0s
 => [app internal] load metadata for docker.io/library/golang:1.18.2       0.9s
 => [app dev 1/3] FROM docker.io/library/golang:1.18.2@sha256:04fab5aaf4f  0.0s
 => CACHED [app dev 2/3] WORKDIR /app                                      0.0s
 => [app dev 3/3] RUN go install github.com/cosmtrek/air@v1.50.0          10.6s
 => [app] exporting to image                                               1.0s
 => => exporting layers                                                    1.0s
 => => writing image sha256:ccd8225de41f6d6f104c43373439a2aa9b374bbba9965  0.0s 
 => => naming to docker.io/library/gotodo                                  0.0s

 Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them