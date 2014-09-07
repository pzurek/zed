package zd

// Translation struct
type Translation struct {
	ID         *float64 `json:"id,omitempty"`
	URL        *string  `json:"url,omitempty"`
	SourceID   *float64 `json:"source_id,omitempty"`
	SourceType *string  `json:"source_type,omitempty"` // Article, Section, Category
	Locale     *string  `json:"locale,omitempty"`
	Title      *string  `json:"title,omitempty"`
	Body       *string  `json:"body,omitempty"`
	Outdated   *bool    `json:"outdated,omitempty"`
	Draft      *bool    `json:"draft,omitempty"`
	Hidden     *bool    `json:"hidden,omitempty"`
	CreatedAt  *string  `json:"created_at,omitempty"`
	UpdatedAt  *string  `json:"updated_at,omitempty"`
}
