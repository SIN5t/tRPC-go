package topic

type Topic struct {
	ID        int64  `db:"id"`
	Title     string `db:"title"`
	CreatedAt string `db:"created_at"`
}
