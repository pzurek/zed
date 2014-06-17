package zd

// Section struct
type Section struct {
	ID                         float64       `json:"id"`
	Name                       string        `json:"name"`
	Description                string        `json:"description"`
	Locale                     string        `json:"locale"`
	SourceLocale               string        `json:"source_locale"`
	URL                        string        `json:"url"`
	HtmlURL                    string        `json:"html_url"`
	CategoryID                 float64       `json:"category_id"`
	Visibility                 string        `json:"visibility"`
	Outdated                   bool          `json:"outdated"`
	RestrictToManager          bool          `json:"restrict_to_manager"`
	InternalRestrictionTags    []interface{} `json:"internal_restriction_tags"`
	RestrictedRestrictionTags  []interface{} `json:"restricted_restriction_tags"`
	RestrictionGroupIds        []interface{} `json:"restriction_group_ids"`
	RestrictionOrganizationIds []interface{} `json:"restriction_organization_ids"`
	Position                   float64       `json:"position"`
	Sorting                    string        `json:"sorting"`
	TranslationIds             []interface{} `json:"translation_ids"`
	CreatedAt                  string        `json:"created_at"`
	UpdatedAt                  string        `json:"updated_at"`
}

// SectionService struct
type SectionService struct {
	client *Client
}

// Get function
// func (*SectionService) Get() []Category {

// }
