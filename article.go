package zd

import (
	"fmt"
	"log"
)

// Article struct
type Article struct {
	ID               *int64   `json:"id,omitempty"`
	URL              *string  `json:"url,omitempty"`
	HtmlURL          *string  `json:"html_url,omitempty"`
	Title            *string  `json:"title"`
	Name             *string  `json:"name,omitempty"`
	Body             *string  `json:"body,omitempty"`
	Locale           *string  `json:"locale"`
	SourceLocale     *string  `json:"source_locale,omitempty"`
	AuthorID         *int64   `json:"author_id,omitempty"`
	CommentsDisabled *bool    `json:"comments_disabled,omitempty"`
	Outdated         *bool    `json:"outdated,omitempty"`
	LabelNames       []string `json:"label_names,omitempty"`
	Draft            *bool    `json:"draft,omitempty"`
	Promoted         *bool    `json:"promoted,omitempty"`
	Position         *int64   `json:"position,omitempty"`
	VoteSum          *int64   `json:"vote_sum,omitempty"`
	VoteCount        *int64   `json:"vote_count,omitempty"`
	SectionID        *int64   `json:"section_id,omitempty"`
	TranslationIds   []int64  `json:"translation_ids,omitempty"`
	CreatedAt        *string  `json:"created_at,omitempty"`
	UpdatedAt        *string  `json:"updated_at,omitempty"`
}

// ArticleResponse struct
type ArticleResponse struct {
	Results   []Article `json:"articles"`
	Count     *int64    `json:"count"`
	Next      *string   `json:"next_page"`
	Page      *int64    `json:"page"`
	PageCount *int64    `json:"page_count"`
	PerPage   *int64    `json:"per_page"`
	Previous  *string   `json:"previous_page"`
}

// ArticleRequest struct
type ArticleRequest struct {
	Article *Article `json:"article"`
}

// ArticleService struct
type ArticleService struct {
	client *Client
}

// GetAll function
func (s *ArticleService) GetAll() ([]Article, error) {
	var resource []Article
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

	result := new(ArticleResponse)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, nil, resp, err
	}

	next := result.Next
	resource := result.Results
	return &resource, next, resp, err
}

// Create func creates a single new article
func (s *ArticleService) Create(a *Article) error {
	var err error

	if a.SectionID == nil {
		return fmt.Errorf("missing section id")
	}

	if a.Title == nil {
		return fmt.Errorf("missing article title")
	}

	if a.Body == nil {
		return fmt.Errorf("missing article body")
	}

	l := "en-us"

	if a.Locale == nil {
		a.Locale = &l
	}

	ar := &ArticleRequest{Article: a}

	url := fmt.Sprintf("help_center/sections/%v/articles.json", *a.SectionID)

	req, err := s.client.NewRequest("POST", url, ar)
	if err != nil {
		return fmt.Errorf("creating new request failed: %v\n", err)
	}

	result := new(ArticleResponse)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return fmt.Errorf("request failed with: %v\n", err)
	}

	log.Printf("%v %s\n", resp.StatusCode, resp.Status)

	return err
}

// Update func updates a single article
func (s *ArticleService) Update(a *Article) error {
	var err error

	if a.ID == nil {
		return fmt.Errorf("missing article id")
	}

	if a.Title == nil {
		return fmt.Errorf("missing article title")
	}

	if a.Body == nil {
		return fmt.Errorf("missing article body")
	}

	l := "en-us"

	if a.Locale == nil {
		a.Locale = &l
	}

	ar := &ArticleRequest{Article: a}

	url := fmt.Sprintf("help_center/articles/%v.json", *a.ID)

	req, err := s.client.NewRequest("PUT", url, ar)
	if err != nil {
		return fmt.Errorf("creating new request failed: %v\n", err)
	}

	result := new(ArticleResponse)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return fmt.Errorf("request failed with: %v\n", err)
	}

	log.Printf("%v %s\n", resp.StatusCode, resp.Status)

	return err
}

// Delete func deletes a single article
func (s *ArticleService) Delete(id *int64) error {
	var err error

	if id == nil {
		return fmt.Errorf("missing article id")
	}

	url := fmt.Sprintf("help_center/articles/%v.json", *id)

	req, err := s.client.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("creating new request failed: %v\n", err)
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return fmt.Errorf("request failed with: %v\n", err)
	}

	log.Printf("%v %s\n", resp.StatusCode, resp.Status)

	return err
}
