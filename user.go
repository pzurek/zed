package zd

type User struct {
	Id        *float64  `json:"id" gorethink:"id"`
	Name      *string   `json:"name" gorethink:"name"`
	Active    *bool     `json:"active" gorethink:"active"`
	CreatedAt *string   `json:"created_at" gorethink:"created_at"`
	UpdatedAt *string   `json:"updated_at" gorethink:"updated_at"`
	Tags      *[]string `json:"tags" gorethink:"tags"`
	// UserFields *map[string]interface{} `json:"user_fields" gorethink:"-"`
	// Phone      bool                    `json:"-" gorethink:"phone"`
}

type UsersResponse struct {
	Users    *[]User `json:"users"`
	Next     *string `json:"next_page,omitempty"`
	Previous *string `json:"previous_page,omitempty"`
	Count    *int    `json:"count,omitempty"`
}

type UserService struct {
	client *Client
}

func (s *UserService) Get() ([]User, error) {
	resource := make([]User, 0)

	rp, next, _, err := s.getPage("")
	if err != nil {
		return nil, err
	}

	resource = append(resource, *rp...)

	for next != nil {
		rp, nx, _, err := s.getPage(*next)
		if err != nil {
			return nil, err
		}
		next = nx
		resource = append(resource, *rp...)
	}

	return resource, err
}

func (s *UserService) getPage(url string) (*[]User, *string, *Response, error) {

	if url == "" {
		url = "users.json?role=agent"
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
