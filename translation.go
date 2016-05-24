package zd

import "time"

// Translation struct
type Translation struct {
	ID          *int       `json:"id"`
	URL         *string    `json:"url"`
	HTMLURL     *string    `json:"html_url"`
	SourceID    *int       `json:"source_id"`
	SourceType  *string    `json:"source_type"`
	Locale      *string    `json:"locale"`
	Title       *string    `json:"title"`
	Body        *string    `json:"body"`
	Outdated    *bool      `json:"outdated"`
	Draft       *bool      `json:"draft"`
	Hidden      *bool      `json:"hidden"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	UpdatedByID *int       `json:"updated_by_id"`
	CreatedByID *int       `json:"created_by_id"`
}

// TranslationResponse struct
type TranslationResponse struct {
	Translation *Translation `json:"translation"`
}
