package cwb

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	libraryVersion   = "0.0.1"
	defaultBaseURL   = "https://opendata.cwb.gov.tw/"
	defaultUserAgent = "go-cwb/" + libraryVersion
)

type service struct {
	client *Client
}

// A Client manages communication with the CWB API.
type Client struct {
	client *http.Client
	token  string

	BaseURL   *url.URL
	UserAgent string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services
	Dataset    *DatasetService
	Forecasts  *ForecastsService
	StationObs *StationObsService
}

// NewClient returns a new CWB API client. The token are required for authentication.
// If a nil httpClient is provided, http.DefaultClient will be used.
func NewClient(token string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:    httpClient,
		token:     token,
		BaseURL:   baseURL,
		UserAgent: defaultUserAgent,
	}
	c.common.client = c
	c.Dataset = (*DatasetService)(&c.common)
	c.Forecasts = (*ForecastsService)(&c.common)
	c.StationObs = (*StationObsService)(&c.common)
	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.token)

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Do sends an API request, and returns the API response.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer resp.Body.Close()

	if err := checkResponse(resp); err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, _ = io.Copy(w, resp.Body)
		} else {
			err = decodeResponse(resp.Body, v)
		}
	}
	return resp, err
}

// Get method make a GET HTTP request.
func (c *Client) Get(ctx context.Context, url string, v interface{}) (*http.Response, error) {
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(ctx, req, v)
}

func (c *Client) generateURL(dataId string, options url.Values) string {
	u, _ := url.Parse(fmt.Sprintf("api/v1/rest/datastore/%v", dataId))
	u.RawQuery = options.Encode()
	return u.String()
}

// ErrorResponse reports error caused by an API request.
type ErrorResponse struct {
	*http.Response
	Message string
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("error: %s", e.Message)
}

// checkResponse checks the API response for errors.
func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r, Message: "unknown"}
	switch r.StatusCode {
	case http.StatusUnauthorized, http.StatusNotFound:
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errorResponse.Message = fmt.Sprintf("reading body, %s", err.Error())
		} else {
			errorResponse.Message = string(data)
		}
	}
	return errorResponse
}

// decodeResponse decodes the API response.
func decodeResponse(body io.Reader, to interface{}) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, to)
	if err != nil {
		return fmt.Errorf("error decoding body: %s", err.Error())
	}
	return nil
}
