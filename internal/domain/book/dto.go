package book

type BookDto struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	AuthorID   int    `json:"author_id"`
	CategoryID int    `json:"category_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
