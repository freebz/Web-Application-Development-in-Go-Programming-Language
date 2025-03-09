// 코드 20.39 UserID 추가

type Task struct {
	ID       TaskID     `json:"id" db:"id"`
	UserID   UserID     `json:"user_id" db:"user_id"`
	Title    string     `json:"title" db:"title"`
	Status   TaskStatus `json:"status" db:"status"`
	Created  time.Time  `json:"create" db:"created"`
	Modified time.Time  `json:"modified" db:"modified"`
}