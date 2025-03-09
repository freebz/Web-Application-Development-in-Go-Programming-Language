// 코드 20.6 로컬 환경과 CI 환경에서의 접속 방법이 다르다

package testutil

import (
	"context"
	"fmt"
	"os"
	"testing"
	"github.com/go-redis/redis/v8"
)

func OpenRedisForTest(t *testing.T) *redis.Client {
	t.Helper()

	host := "127.0.0.1"
	port := 36379
	if _, defined := os.LookupEnv("CI"); defined {
		// https://docs.github.com/ja/actions/using-containerized-services/creating-redis-service-containers#configuring-the-runner-job
		port = 6379
	}
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "",
		DB:       0, // default database number
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		t.Fatalf("failed to connect redis: %s", err)
	}
	return client
}
