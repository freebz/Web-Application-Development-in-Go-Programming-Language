// 코드 20.7 store.KVS.Save 메서드용 테스트

package store

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/freebz/go_todo_app/entity"
	"github.com/freebz/go_todo_app/testutil"
)

func TestKVS_Save(t *testing.T) {
	t.Parallel()

	cli := testutil.OpenRedisForTest(t)

	sut := &KVS{Cli: cli}
	key := "TestKVS_Save"
	uid := entity.UserID(1234)
	ctx := context.Background()
	cli.Set(ctx, key, int64(uid), 30*time.Minute)
	t.Cleanup(func() {
		cli.Del(ctx, key)
	})
	if err := sut.Save(ctx, key, uid); err != nil {
		t.Errorf("want no error, but got %v", err)
	}
}