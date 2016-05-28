package zed

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	baseURLString  = ".zendesk.com/api/v2/"
	libraryVersion = "0.1"
	userAgent      = "github.com/pzurek/zed/" + libraryVersion
)

var (
	subdomain string
	username  string
	password  string
)

// Client which talks to the Vend API
type Client struct {

	// HTTP client
	client    *http.Client
	UserAgent string

	Tickets       *TicketService
	Users         *UserService
	Articles      *ArticleService
	Labels        *LabelService
	Organizations *OrganizationService
	Search        *SearchService
}

// NewClient creates a new instance of the Client type
func NewClient(dmn, usrname, passwd string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	subdomain = dmn
	username = usrname
	password = passwd

	c := &Client{client: httpClient, UserAgent: userAgent}

	c.Organizations = &OrganizationService{client: c}
	c.Tickets = &TicketService{client: c}
	c.Users = &UserService{client: c}
	c.Articles = &ArticleService{client: c}
	c.Search = &SearchService{client: c}

	return c
}

// NewRequest function creates a new request
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {

	url := ""

	if strings.Contains(urlStr, "http") {
		url = urlStr
	} else {
		url = fmt.Sprintf("https://%s%s%s", subdomain, baseURLString, urlStr)
	}

	buf := &bytes.Buffer{}
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(username, password)
	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

// Do function executes a client request
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	// drain the response body
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()
	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return response, err
}

// Response wraps the http.Response type
type Response struct {
	*http.Response
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

// ErrorResponse wraps an error response
type ErrorResponse struct {
	Response *http.Response
	Message  string `json:"error,omitempty"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method,
		r.Response.Request.URL,
		r.Response.StatusCode,
		r.Message)
}

// CheckResponse checks if a response is valid
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}
