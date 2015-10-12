package zd

type OrganizationResponse struct {
	Organizations []Organization `json:"organizations,omitempty"`
	NextPage      *string        `json:"next_page,omitempty"`
	PreviousPage  *string        `json:"previous_page,omitempty"`
	Count         *int           `json:"count,omitempty"`
}

type Organization struct {
	URL                *string           `json:"url,omitempty"`
	ID                 *int              `json:"id,omitempty"`
	Name               *string           `json:"name,omitempty"`
	SharedTickets      *bool             `json:"shared_tickets,omitempty"`
	SharedComments     *bool             `json:"shared_comments,omitempty"`
	ExternalID         *string           `json:"external_id,omitempty"`
	CreatedAt          *time.Time        `json:"created_at,omitempty"`
	UpdatedAt          *time.Time        `json:"updated_at,omitempty"`
	DomainNames        []string          `json:"domain_names,omitempty"`
	Details            *string           `json:"details,omitempty"`
	Notes              *string           `json:"notes,omitempty"`
	GroupID            *string           `json:"group_id,omitempty"`
	Tags               []string          `json:"tags,omitempty"`
	OrganizationFields map[string]string `json:"organization_fields,omitempty"`
}
