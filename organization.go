package zd

import (
	"fmt"
)

type OrganizationWrapper struct {
	Organization *Organization `json:"organization"`
}

type OrganizationResponse struct {
	Organization *Organization `json:"organization"`
}

type OrganizationListResponse struct {
	Organizations []Organization `json:"organizations"`
	NextPage      string        `json:"next_page,omitempty"`
	PreviousPage  string        `json:"previous_page,omitempty"`
	Count         int           `json:"count,omitempty"`
}

type Organization struct {
	URL                string           `json:"url,omitempty"`
	ID                 int              `json:"id,omitempty"`
	Name               string           `json:"name,omitempty"`
	SharedTickets      bool             `json:"shared_tickets,omitempty"`
	SharedComments     bool             `json:"shared_comments,omitempty"`
	ExternalID         string           `json:"external_id,omitempty"`
	CreatedAt          string           `json:"created_at,omitempty"`
	UpdatedAt          string           `json:"updated_at,omitempty"`
	DomainNames        []string          `json:"domain_names,omitempty"`
	Details            string           `json:"details,omitempty"`
	Notes              string           `json:"notes,omitempty"`
	GroupID            string           `json:"group_id,omitempty"`
	Tags               []string          `json:"tags,omitempty"`
	OrganizationFields map[string]string `json:"organization_fields,omitempty"`
}

// TicketService struct
type OrganizationService struct {
	client *Client
}

// finds an organization in zendesk byt id
func (s *OrganizationService) GetOrganizationById(organization_id string) (*Organization, *Response, error) {
	org := OrganizationResponse{}

	url := fmt.Sprintf("organizations/%s.json", organization_id)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, &org)
	if err != nil {
		return nil, nil,  err
	}

	return org.Organization, resp, err
}

//// updates and organization by id
func (s *OrganizationService) UpdateOrganizationById(org *Organization) (*Organization, error) {
	var organization *Organization
	var err error

	fmt.Println(org)

	url := fmt.Sprintf("organizations/%v.json", org.ID)

	or := &OrganizationWrapper{Organization: org}

	fmt.Println(or.Organization.Notes)

	req, err := s.client.NewRequest("PUT", url, or)
	if err != nil {
		return organization, err
	}

	result := new(OrganizationWrapper)

	_, err = s.client.Do(req, result)
	if err != nil {
		return organization, err
	}

	organization = result.Organization

	return organization, err
}

//creates a new organization
func (s *OrganizationService) CreateOrganization(org *Organization) (*Organization, error) {
	var organization *Organization
	var err error

	or := &OrganizationWrapper{Organization: org}

	url := fmt.Sprintf("organizations.json")

	req, err := s.client.NewRequest("POST", url, or)
	if err != nil {
		return organization, err
	}

	result := new(OrganizationResponse)
	_, err = s.client.Do(req, result)

	if err != nil {
		return organization, err
	}

	organization = result.Organization

	return organization, nil
}