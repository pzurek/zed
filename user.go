package zd

type User struct {
	ID        *float64 `json:"id,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Active    *bool    `json:"active,omitempty"`
	CreatedAt *string  `json:"created_at,omitempty"`
	UpdatedAt *string  `json:"updated_at,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

type UsersResponse struct {
	Users    []User  `json:"users,omitempty"`
	Next     *string `json:"next_page,omitempty"`
	Previous *string `json:"previous_page,omitempty"`
	Count    *int    `json:"count,omitempty"`
}

type UserService struct {
	client *Client
}

func (s *UserService) Get() ([]User, error) {
	var resource []User

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

	response := new(UsersResponse)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, nil, resp, err
	}

	next := response.Next
	resource := response.Users
	return resource, next, resp, err
}
