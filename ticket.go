package zed

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Ticket struct
type Ticket struct {
	ID                  *float64            `json:"id,omitempty"`
	URL                 *string             `json:"url,omitempty"`
	ExternalID          *string             `json:"external_id,omitempty"`
	Type                *string             `json:"type,omitempty"`
	Subject             *string             `json:"subject,omitempty"`
	Description         *string             `json:"description,omitempty"`
	Priority            *string             `json:"priority,omitempty"`
	Status              *string             `json:"status,omitempty"`
	Recipient           *string             `json:"recipient,omitempty"`
	RequesterID         *float64            `json:"requester_id,omitempty"`
	Requester           *User               `json:"requester,omitempty"`
	SubmitterID         *float64            `json:"submitter_id,omitempty"`
	AssigneeID          *float64            `json:"assignee_id,omitempty"`
	OrganizationID      *float64            `json:"organization_id,omitempty"`
	GroupID             *float64            `json:"group_id,omitempty"`
	CollaboratorIDs     []float64           `json:"collaborator_ids,omitempty"`
	ForumTopicID        *float64            `json:"forum_topic_id,omitempty"`
	ProblemID           *float64            `json:"problem_id,omitempty"`
	HasIncidents        *bool               `json:"has_incidents,omitempty"`
	DueAt               *string             `json:"due_at,omitempty"`
	Tags                []string            `json:"tags,omitempty"`
	Via                 *Via                `json:"via,omitempty"`
	CustomFields        []CustomField       `json:"custom_fields,omitempty"`
	SatisfactionRating  *SatisfactionRating `json:"satisfaction_rating,omitempty"`
	SharingAgreementIds []float64           `json:"sharing_agreement_ids,omitempty"`
	FollowupIds         []float64           `json:"followup_ids,omitempty"`
	TicketFormID        *float64            `json:"ticket_form_id,omitempty"`
	BrandID             *float64            `json:"brand_id,omitempty"`
	CreatedAt           *string             `json:"created_at,omitempty"`
	UpdatedAt           *string             `json:"updated_at,omitempty"`
	Comment             *Comment            `json:"comment,omitempty"`
}

// Via struct
type Via struct {
	Channel *string      `json:"channel,omitempty"`
	Source  *interface{} `json:"source,omitempty"`
}

// Comment on ticket
type Comment struct {
	Body     *string `json:"body,omitempty"`
	Public   *bool   `json:"private,omitempty"`
	AuthorID *int    `json:"author_id,omitemtpy"`
}

// CustomField struct
type CustomField struct {
	ID    *float64    `json:"id,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// SatisfactionRating struct
type SatisfactionRating struct {
	ID      *float64 `json:"id,omitempty"`
	Score   *string  `json:"score,omitempty"`
	Comment *string  `json:"comment,omitempty"`
}

// TicketCollectionResponse struct
type TicketCollectionResponse struct {
	Results  []Ticket `json:"tickets"`
	Next     *string  `json:"next_page,omitempty"`
	Previous *string  `json:"previous_page,omitempty"`
	Count    *int     `json:"count,omitempty"`
}

// TicketResponse struct
type TicketResponse struct {
	Ticket Ticket `json:"ticket"`
}

// TicketUserGroupResponse struct
type TicketUserGroupResponse struct {
	Tickets  []Ticket `json:"tickets,omitempty"`
	Users    []User   `json:"users,omitempty"`
	Groups   []Group  `json:"groups,omitempty"`
	Next     *string  `json:"next_page,omitempty"`
	Previous *string  `json:"previous_page,omitempty"`
	Count    *int     `json:"count,omitempty"`
}

// TicketService struct
type TicketService struct {
	client *Client
}

// List returns a slice of all products
func (s *TicketService) List() ([]Ticket, error) {
	resource := []Ticket{}

	rp, next, _, err := s.getPage("")

	if err != nil {
		return nil, err
	}

	resource = append(resource, rp...)

	for next != nil {
		rp, nx, _, err := s.getPage(*next)
		if err != nil {
			return nil, err
		}
		next = nx
		resource = append(resource, rp...)
	}

	return resource, err
}

// ListByView function
func (s *TicketService) ListByView(id string) ([]Ticket, error) {
	resource := []Ticket{}

	url := fmt.Sprintf("views/%s/tickets.json", id)
	rp, next, _, err := s.getPage(url)
	if err != nil {
		return nil, err
	}

	resource = append(resource, rp...)

	for next != nil {
		rp, nx, _, err := s.getPage(*next)
		if err != nil {
			return nil, err
		}
		next = nx
		resource = append(resource, rp...)
	}

	return resource, err
}

// ListByViewUG function
func (s *TicketService) ListByViewUG(id string) ([]Ticket, []User, []Group, error) {

	tickets := []Ticket{}
	users := []User{}
	groups := []Group{}

	url := fmt.Sprintf("views/%s/tickets.json?include=users,groups", id)
	tkts, usrs, grps, next, _, err := s.getPageUG(url)
	if err != nil {
		return nil, nil, nil, err
	}

	tickets = append(tickets, tkts...)
	users = append(users, usrs...)
	groups = append(groups, grps...)

	for next != nil {
		tkts, usrs, grps, nx, _, err := s.getPageUG(*next)
		if err != nil {
			return nil, nil, nil, err
		}
		next = nx
		tickets = append(tickets, tkts...)
		users = append(users, usrs...)
		groups = append(groups, grps...)
	}

	return tickets, users, groups, err
}

// GetProblemIncidents gets all problem tickets
func (s *TicketService) GetProblemIncidents(id string) ([]Ticket, error) {
	resource := []Ticket{}

	url := fmt.Sprintf("tickets/%s/incidents.json", id)

	rp, next, _, err := s.getPage(url)
	if err != nil {
		return nil, err
	}

	resource = append(resource, rp...)

	for next != nil {
		rp, nx, _, err := s.getPage(*next)
		if err != nil {
			return nil, err
		}
		next = nx
		resource = append(resource, rp...)
	}

	return resource, err
}

// GetProblemIncidentsCount gets only the first page of tickets which includes the count field
func (s *TicketService) GetProblemIncidentsCount(id string) (int, error) {

	url := fmt.Sprintf("tickets/%s/incidents.json", id)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	response := TicketCollectionResponse{}
	_, err = s.client.Do(req, response)
	if err != nil {
		return 0, err
	}

	resource := response.Count
	return *resource, err
}

func (s *TicketService) getPage(url string) ([]Ticket, *string, *Response, error) {

	if url == "" {
		url = "tickets.json"
	}

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	response := TicketCollectionResponse{}
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, nil, resp, err
	}

	next := response.Next
	resource := response.Results
	return resource, next, resp, err
}

func (s *TicketService) getPageUG(url string) ([]Ticket, []User, []Group, *string, *Response, error) {

	if url == "" {
		url = "tickets.json?include=users,groups"
	}

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	response := TicketUserGroupResponse{}
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, nil, nil, nil, resp, err
	}

	next := response.Next
	tickets := response.Tickets
	users := response.Users
	groups := response.Groups
	return tickets, users, groups, next, resp, err
}

// Get method
func (s *TicketService) Get(id string) (*Ticket, *Response, error) {
	url := fmt.Sprintf("tickets/%s.json", id)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	response := &TicketResponse{}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return &response.Ticket, resp, err
}

// Create a new Zendesk Ticket
func (s *TicketService) Create(ticket *Ticket) (*Ticket, *Response, error) {
	url := "tickets.json"

	body, err := json.Marshal(&ticket)

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	response := TicketResponse{}
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, resp, err
	}

	return ticket, resp, err
}

// Update a Zendesk Ticket
func (s *TicketService) Update(ticket *Ticket) (*Ticket, *Response, error) {
	if ticket.ID == nil {
		// no ticket id so return and error.
		return nil, nil, errors.New("Please supply a ticket with an ID to update")
	}

	url := fmt.Sprintf("tickets/%v.json", *ticket.ID)

	payload := TicketResponse{
		Ticket: *ticket,
	}

	req, err := s.client.NewRequest("PUT", url, payload)
	if err != nil {
		return nil, nil, err
	}

	response := TicketResponse{}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return ticket, resp, err
}
