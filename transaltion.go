package zd

// Translation struct
type Translation struct {
	ID         *float64 `json:"id"`
	URL        *string  `json:"url"`
	SourceID   *float64 `json:"source_id"`
	SourceType *string  `json:"source_type"` // Article, Section, Category
	Locale     *string  `json:"locale"`
	Title      *string  `json:"title"`
	Body       *string  `json:"body"`
	Outdated   *bool    `json:"outdated"`
	Draft      *bool    `json:"draft"`
	Hidden     *bool    `json:"hidden"`
	CreatedAt  *string  `json:"created_at"`
	UpdatedAt  *string  `json:"updated_at"`
}
