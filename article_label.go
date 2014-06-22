package zd

import "fmt"

// ArticleLabel struct
type ArticleLabel struct {
	ID        *float64 `json:"id"`
	URL       *string  `json:"url"`
	Name      *string  `json:"name"`
	CreatedAt *string  `json:"created_at"`
	UpdatedAt *string  `json:"updated_at"`
}

// LabelWrapper struct
type LabelWrapper struct {
	Label *ArticleLabel `json:"label"`
}

// LabelListResponse struct
type LabelListResponse struct {
	Results   []ArticleLabel `json:"labels"`
	Count     *int64         `json:"count"`
	Next      *string        `json:"next_page"`
	Page      *int64         `json:"page"`
	PageCount *int64         `json:"page_count"`
	PerPage   *int64         `json:"per_page"`
	Previous  *string        `json:"previous_page"`
}

// LabelService struct
type LabelService struct {
	client *Client
}

// Create func creates a single new article
func (s *LabelService) Create(id *int64, l *ArticleLabel) (*ArticleLabel, error) {
	var label *ArticleLabel
	var err error

	if l.Name == nil {
		return label, fmt.Errorf("missing required field: label name")
	}

	lw := &LabelWrapper{Label: l}

	url := fmt.Sprintf("help_center/articles/%v/labels.json", *id)

	req, err := s.client.NewRequest("POST", url, lw)
	if err != nil {
		return label, err
	}

	result := new(LabelWrapper)
	_, err = s.client.Do(req, result)
	if err != nil {
		return label, err
	}
	label = result.Label

	return label, err
}

// GetAll function lists labels used in all articles
func (s *LabelService) GetAll() ([]ArticleLabel, error) {
	var resource []ArticleLabel

	rp, next, _, err := s.getPage("")
	if err != nil {
		return resource, err
	}
	resource = append(resource, rp...)

	for next != nil {
		rp, nx, _, err := s.getPage(*next)
		if err != nil {
			return resource, err
		}
		next = nx
		resource = append(resource, rp...)
	}

	return resource, err
}

// GetByArticleID function lists lablels used in an article with a given id
func (s *LabelService) GetByArticleID(id *int64) ([]ArticleLabel, error) {
	var resource []ArticleLabel

	url := fmt.Sprintf("help_center/articles/%v/labels.json", *id)
	rp, next, _, err := s.getPage(url)
	if err != nil {
		return resource, err
	}
	resource = append(resource, rp...)

	for next != nil {
		rp, nx, _, err := s.getPage(*next)
		if err != nil {
			return resource, err
		}
		next = nx
		resource = append(resource, rp...)
	}

	return resource, err
}

func (s *LabelService) getPage(url string) ([]ArticleLabel, *string, *Response, error) {

	if url == "" {
		url = "help_center/articles/labels.json"
	}

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	result := new(LabelListResponse)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, nil, resp, err
	}

	next := result.Next
	resource := result.Results
	return resource, next, resp, err
}

// Delete func deletes a single article
func (s *LabelService) Delete(articleId, id *int64) error {
	var err error

	if articleId == nil {
		return fmt.Errorf("missing required field: article id")
	}

	if id == nil {
		return fmt.Errorf("missing required field: id")
	}

	url := fmt.Sprintf("help_center/articles/%v/labels/%v.json", *articleId, *id)

	req, err := s.client.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("creating new request failed: %v\n", err)
	}

	_, err = s.client.Do(req, nil)
	if err != nil {
		return fmt.Errorf("request failed with: %v\n", err)
	}

	return err
}
