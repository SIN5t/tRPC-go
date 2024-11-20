package topic

type TopicItem struct {
	Id        int64  `db:"id"`
	Title     string `db:"title"`
	CreatedAt string `db:"created_at"`
}
