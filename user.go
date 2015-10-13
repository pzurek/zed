package zd

import (
	"time"
)

type User struct {
	ID                   *int                   `json:"id,omitempty"`
	URL                  *string                `json:"url,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	Email                *string                `json:"email,omitempty"`
	CreatedAt            *time.Time             `json:"created_at,omitempty"`
	UpdatedAt            *time.Time             `json:"updated_at,omitempty"`
	TimeZone             *string                `json:"time_zone,omitempty"`
	Phone                *string                `json:"phone,omitempty"`
	Photo                *Photo                 `json:"photo,omitempty"`
	LocaleID             *int                   `json:"locale_id,omitempty"`
	Locale               *string                `json:"locale,omitempty"`
	OrganizationID       *int                   `json:"organization_id,omitempty"`
	Role                 *string                `json:"role,omitempty"`
	Verified             *bool                  `json:"verified,omitempty"`
	ExternalID           *string                `json:"external_id,omitempty"`
	Tags                 []string               `json:"tags,omitempty"`
	Alias                *string                `json:"alias,omitempty"`
	Active               *bool                  `json:"active,omitempty"`
	Shared               *bool                  `json:"shared,omitempty"`
	SharedAgent          *bool                  `json:"shared_agent,omitempty"`
	LastLoginAt          *time.Time             `json:"last_login_at,omitempty"`
	TwoFactorAuthEnabled *bool                  `json:"two_factor_auth_enabled,omitempty"`
	Signature            *string                `json:"signature,omitempty"`
	Details              *string                `json:"details,omitempty"`
	Notes                *string                `json:"notes,omitempty"`
	CustomRoleID         *int                   `json:"custom_role_id,omitempty"`
	Moderator            *bool                  `json:"moderator,omitempty"`
	TicketRestriction    *string                `json:"ticket_restriction,omitempty"`
	OnlyPrivateComments  *bool                  `json:"only_private_comments,omitempty"`
	RestrictedAgent      *bool                  `json:"restricted_agent,omitempty"`
	Suspended            *bool                  `json:"suspended,omitempty"`
	ChatOnly             *bool                  `json:"chat_only,omitempty"`
	UserFields           map[string]interface{} `json:"user_fields,omitempty"`
}

type UserResponse struct {
	Users    []User  `json:"users,omitempty"`
	Next     *string `json:"next_page,omitempty"`
	Previous *string `json:"previous_page,omitempty"`
	Count    *int    `json:"count,omitempty"`
}

type UserService struct {
	client *Client
}

type Photo struct {
	URL              *string     `json:"url,omitempty"`
	ID               *int        `json:"id,omitempty"`
	FileName         *string     `json:"file_name,omitempty"`
	ContentURL       *string     `json:"content_url,omitempty"`
	MappedContentURL *string     `json:"mapped_content_url,omitempty"`
	ContentType      *string     `json:"content_type,omitempty"`
	Size             *int        `json:"size,omitempty"`
	Inline           *bool       `json:"inline,omitempty"`
	Thumbnails       []Thumbnail `json:"thumbnails,omitempty"`
}

type Thumbnail struct {
	URL              *string `json:"url,omitempty"`
	ID               *int    `json:"id,omitempty"`
	FileName         *string `json:"file_name,omitempty"`
	ContentURL       *string `json:"content_url,omitempty"`
	MappedContentURL *string `json:"mapped_content_url,omitempty"`
	ContentType      *string `json:"content_type,omitempty"`
	Size             *int    `json:"size,omitempty"`
	Inline           *bool   `json:"inline,omitempty"`
}

func (s *UserService) Get() ([]User, error) {
	resource := []User{}

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

func (s *UserService) getPage(url string) ([]User, *string, *Response, error) {

	if url == "" {
		url = "users.json"
	}

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	response := UserResponse{}
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, nil, resp, err
	}

	next := response.Next
	resource := response.Users
	return resource, next, resp, err
}
