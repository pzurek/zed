package zd

import (
	"fmt"
)

// OrganizationWrapper struct
type OrganizationWrapper struct {
	Organization *Organization `json:"organization"`
}

// OrganizationResponse struct
type OrganizationResponse struct {
	Organizations Organization `json:"organizations,omitempty"`
	NextPage      *string      `json:"next_page,omitempty"`
	PreviousPage  *string      `json:"previous_page,omitempty"`
	Count         *int         `json:"count,omitempty"`
}

// Organization struct
type Organization struct {
	URL                *string           `json:"url,omitempty"`
	ID                 *int              `json:"id,omitempty"`
	Name               *string           `json:"name,omitempty"`
	SharedTickets      *bool             `json:"shared_tickets,omitempty"`
	SharedComments     *bool             `json:"shared_comments,omitempty"`
	ExternalID         *string           `json:"external_id,omitempty"`
	CreatedAt          *string           `json:"created_at,omitempty"`
	UpdatedAt          *string           `json:"updated_at,omitempty"`
	DomainNames        []string          `json:"domain_names,omitempty"`
	Details            *string           `json:"details,omitempty"`
	Notes              *string           `json:"notes,omitempty"`
	GroupID            *string           `json:"group_id,omitempty"`
	Tags               []string          `json:"tags,omitempty"`
	OrganizationFields map[string]string `json:"organization_fields,omitempty"`
}

// OrganizationService struct
type OrganizationService struct {
	client *Client
}

// GetOrganizationByID finds an organization in Zendesk by ID
func (s *OrganizationService) GetOrganizationByID(organizationID string) (*Organization, *Response, error) {
	org := OrganizationWrapper{}

	url := fmt.Sprintf("organizations/%s.json", organizationID)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, &org)
	if err != nil {
		return nil, nil, err
	}

	return org.Organization, resp, err
}

// UpdateOrganization updates and organization by id
func (s *OrganizationService) UpdateOrganization(org *Organization) (*Organization, error) {
	organization := &Organization{}

	url := fmt.Sprintf("organizations/%d.json", *org.ID)
	or := &OrganizationWrapper{Organization: org}

	req, err := s.client.NewRequest("PUT", url, or)
	if err != nil {
		return organization, err
	}

	result := OrganizationWrapper{}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return organization, err
	}

	organization = result.Organization
	return organization, err
}

//CreateOrganization creates a new organization
func (s *OrganizationService) CreateOrganization(org *Organization) (*Organization, error) {
	organization := &Organization{}

	or := &OrganizationWrapper{Organization: org}
	url := fmt.Sprintf("organizations.json")

	req, err := s.client.NewRequest("POST", url, or)
	if err != nil {
		return organization, err
	}

	result := OrganizationWrapper{}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return organization, err
	}

	organization = result.Organization
	return organization, nil
}
