package zd

import "fmt"

// Article struct
type Article struct {
	ID               *float64 `json:"id,omitempty"`
	URL              *string  `json:"url,omitempty"`
	HtmlURL          *string  `json:"html_url,omitempty"`
	Title            *string  `json:"title,omitempty"`
	Name             *string  `json:"name,omitempty"`
	Body             *string  `json:"body,omitempty"`
	Locale           *string  `json:"locale,omitempty"`
	SourceLocale     *string  `json:"source_locale,omitempty"`
	AuthorID         *float64 `json:"author_id,omitempty"`
	CommentsDisabled *bool    `json:"comments_disabled,omitempty"`
	Outdated         *bool    `json:"outdated,omitempty"`
	Labels           []string `json:"label_names,omitempty"`
	Draft            *bool    `json:"draft,omitempty"`
	Promoted         *bool    `json:"promoted,omitempty"`
	Position         *int64   `json:"position,omitempty"`
	VoteSum          *int64   `json:"vote_sum,omitempty"`
	VoteCount        *int64   `json:"vote_count,omitempty"`
	SectionID        *float64 `json:"section_id,omitempty"`
	TranslationIds   []int64  `json:"translation_ids,omitempty"`
	CreatedAt        *string  `json:"created_at,omitempty"`
	UpdatedAt        *string  `json:"updated_at,omitempty"`
}

// ArticleWrapper struct
type ArticleWrapper struct {
	Article *Article `json:"article"`
}

// ArticleListResponse struct
type ArticleListResponse struct {
	Results   []Article `json:"articles,omitempty"`
	Count     *int64    `json:"count,omitempty"`
	Next      *string   `json:"next_page,omitempty"`
	Page      *int64    `json:"page,omitempty"`
	PageCount *int64    `json:"page_count,omitempty"`
	PerPage   *int64    `json:"per_page,omitempty"`
	Previous  *string   `json:"previous_page,omitempty"`
}

// ArticleService struct
type ArticleService struct {
	client *Client
}

// List function
func (s *ArticleService) List() ([]Article, error) {
	resource := []Article{}
	rp, next, _, err := s.getPage("")
	if err != nil {
		return resource, err
	}
	resource = append(resource, *rp...)

	for next != nil {
		rp, nx, _, err := s.getPage(*next)
		if err != nil {
			return resource, err
		}
		next = nx
		resource = append(resource, *rp...)
	}

	return resource, err
}

func (s *ArticleService) getPage(url string) (*[]Article, *string, *Response, error) {

	if url == "" {
		url = "help_center/articles.json"
	}

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	result := ArticleListResponse{}
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, nil, resp, err
	}

	next := result.Next
	resource := result.Results
	return &resource, next, resp, err
}

// Create func creates a single new article
func (s *ArticleService) Create(a *Article) (*Article, error) {
	article := &Article{}

	if a.SectionID == nil {
		return article, fmt.Errorf("missing required field: section id")
	}

	if a.Title == nil {
		return article, fmt.Errorf("missing required field: article title")
	}

	if a.Body == nil {
		return article, fmt.Errorf("missing required field: article body")
	}

	l := "en-us"

	if a.Locale == nil {
		a.Locale = &l
	}

	ar := &ArticleWrapper{Article: a}

	url := fmt.Sprintf("help_center/sections/%v/articles.json", int(*a.SectionID))

	req, err := s.client.NewRequest("POST", url, ar)
	if err != nil {
		return article, err
	}

	result := ArticleWrapper{}
	_, err = s.client.Do(req, result)
	if err != nil {
		return article, err
	}

	article = result.Article
	return article, err
}

// Update func updates a single article
func (s *ArticleService) Update(a *Article) (*Article, error) {
	article := &Article{}

	if a.ID == nil {
		return article, fmt.Errorf("missing required field: article id")
	}

	if a.Title == nil {
		return article, fmt.Errorf("missing required field: article title")
	}

	if a.Body == nil {
		return article, fmt.Errorf("missing required field: article body")
	}

	l := "en-us"

	if a.Locale == nil {
		a.Locale = &l
	}

	ar := &ArticleWrapper{Article: a}

	url := fmt.Sprintf("help_center/articles/%v.json", int(*a.ID))

	req, err := s.client.NewRequest("PUT", url, ar)
	if err != nil {
		return article, err
	}

	result := ArticleWrapper{}
	_, err = s.client.Do(req, result)
	if err != nil {
		return article, err
	}

	article = result.Article
	return article, err
}

// Delete func deletes a single article
func (s *ArticleService) Delete(id *int64) error {
	if id == nil {
		return fmt.Errorf("missing article id")
	}

	url := fmt.Sprintf("help_center/articles/%v.json", int(*id))

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
