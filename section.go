package zed

// Section struct
type Section struct {
	ID                         *float64 `json:"id,omitempty"`
	Name                       *string  `json:"name,omitempty"`
	Description                *string  `json:"description,omitempty"`
	Locale                     *string  `json:"locale,omitempty"`
	SourceLocale               *string  `json:"source_locale,omitempty"`
	URL                        *string  `json:"url,omitempty"`
	HtmlURL                    *string  `json:"html_url,omitempty"`
	CategoryID                 *float64 `json:"category_id,omitempty"`
	Visibility                 *string  `json:"visibility,omitempty"`
	Outdated                   *bool    `json:"outdated,omitempty"`
	RestrictToManager          *bool    `json:"restrict_to_manager,omitempty"`
	InternalRestrictionTags    []string `json:"internal_restriction_tags,omitempty"`
	RestrictedRestrictionTags  []string `json:"restricted_restriction_tags,omitempty"`
	RestrictionGroupIds        []string `json:"restriction_group_ids,omitempty"`
	RestrictionOrganizationIds []string `json:"restriction_organization_ids,omitempty"`
	Position                   *float64 `json:"position,omitempty"`
	Sorting                    *string  `json:"sorting,omitempty"`
	TranslationIds             []string `json:"translation_ids,omitempty"`
	CreatedAt                  *string  `json:"created_at,omitempty"`
	UpdatedAt                  *string  `json:"updated_at,omitempty"`
}

// SectionService struct
type SectionService struct {
	client *Client
}
