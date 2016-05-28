package zed

// Category struct
type Category struct {
	ID             float64       `json:"id,omitempty"`
	Name           string        `json:"name,omitempty"`
	Description    string        `json:"description,omitempty"`
	Locale         string        `json:"locale,omitempty"`
	SourceLocale   string        `json:"source_locale,omitempty"`
	URL            string        `json:"url,omitempty"`
	HtmlURL        string        `json:"html_url,omitempty"`
	Outdated       bool          `json:"outdated,omitempty"`
	Position       float64       `json:"position,omitempty"`
	TranslationIDs []interface{} `json:"translation_ids,omitempty"`
	CreatedAt      string        `json:"created_at,omitempty"`
	UpdatedAt      string        `json:"updated_at,omitempty"`
}

// CategoryService struct
type CategoryService struct {
	client *Client
}
