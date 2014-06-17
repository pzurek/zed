package zd

// Category struct
type Category struct {
	ID             float64       `json:"id"`
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	Locale         string        `json:"locale"`
	SourceLocale   string        `json:"source_locale"`
	URL            string        `json:"url"`
	HtmlURL        string        `json:"html_url"`
	Outdated       bool          `json:"outdated"`
	Position       float64       `json:"position"`
	TranslationIDs []interface{} `json:"translation_ids"`
	CreatedAt      string        `json:"created_at"`
	UpdatedAt      string        `json:"updated_at"`
}

// CategoryService struct
type CategoryService struct {
	client *Client
}

// Get function
// func (*CategoryService) Get() []Category {

// }
