package zd

import "fmt"

// OrganizationSearchResponse struct
type OrganizationSearchResponse struct {
	Organizations []Organization `json:"results,omitempty"`
	NextPage      *string        `json:"next_page,omitempty"`
	PreviousPage  *string        `json:"previous_page,omitempty"`
	Count         *int           `json:"count,omitempty"`
}

// SearchService struct
type SearchService struct {
	client *Client
}

// OrganizationByName searches the organization by name
func (s *SearchService) OrganizationByName(name string) (*OrganizationSearchResponse, error) {
	result := &OrganizationSearchResponse{}

	url := fmt.Sprintf("search?query=type:organization+name:%s", name)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	_, err = s.client.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
