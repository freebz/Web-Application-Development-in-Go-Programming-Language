// 코드 18.19 ListTasks 메서드가 기대한 데이터를 반환하는지 검증

package store

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/freebz/go_todo_app/clock"
	"github.com/freebz/go_todo_app/entity"
	"github.com/freebz/go_todo_app/testutil"
	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
)

func TestRepository_ListTasks(t *testing.T) {
	ctx := context.Background()
	// entity.Task를 저장하는 다른 테스트 케이스와 섞이면 테스트가 실패한다.
	// 이를 위해 트랜잭션을 적용해서 테스트 케이스 내로 한정된 테이블 상태를 만든다.
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	// 이 테스트 케이스가 끝나면 원래 상태로 되돌린다.
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}
	wants := prepareTasks(ctx, t, tx)

	sut := &Repository{}
	gots, err := sut.ListTasks(ctx, tx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if d := cmp.Diff(gots, wants); len(d) != 0 {
		t.Errorf("differs: (-got +want)\n%s", d)
	}
}