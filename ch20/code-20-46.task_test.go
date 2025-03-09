// 코드 20.46 *store.Repository.ListTasks 메서드용 테스트

func TestRepository_ListTasks(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	// entity.Task를 작성하는 다른 테스트 케이스와 섞이면 테스트가 실패한다.
	// 하지만 트랜잭션을 적용해서 이 테스트 케이스 내에서는 테이블 상태가 유지된다.
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	// 이 테스트 케이스가 완료되면 원래대로 되돌린다.
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}
	wantUserID, wants := prepareTasks(ctx, t, tx)

	sut := &Repository{}
	gots, err := sut.ListTasks(ctx, tx, wantUserID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if d := cmp.Diff(gots, wants); len(d) != 0 {
		t.Errorf("differs: (-got +want)\n%s", d)
	}
}