// Package f1 provides utilties for interfacing
// with the F1 Fantasy Game API.
package f1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL = "https://fantasy-api.formula1.com/f1/2022"
)

type service struct {
	client *Client
}

type Client struct {
	http    *http.Client // HTTP client used to communicate with the API.
	baseURL *url.URL     // Base URL for API requests

	shared service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the API.
	Drivers  *DriversService
	Teams    *TeamsService
	Circuits *CircuitsService
}

// NewClient returns a new F1 Fantasy Gamex API client.
// If a nil httpClient is provided, a new http.Client will be used.
//
// To use API methods which require authentication, you must call
// the Client.Authenticate method with valid credentials.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: time.Minute}
	}
	baseURL, _ := url.Parse(baseURL)

	c := &Client{http: httpClient, baseURL: baseURL}

	c.shared.client = c
	c.Drivers = (*DriversService)(&c.shared)
	c.Teams = (*TeamsService)(&c.shared)
	c.Circuits = (*CircuitsService)(&c.shared)
	return c
}

// Error represents an error returned by the API.
type Error struct {
	Code    int    `json:"code"`    // The HTTP status code.
	Status  string `json:"status"`  // The HTTP response status (error or success).
	Message string `json:"message"` // A short description of the error.
}

// decodeError decodes an error from an io.Reader.
func (c *Client) decodeError(resp *http.Response) error {
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(respBody) == 0 {
		return fmt.Errorf("HTTP %d: %s (body empty)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	buf := bytes.NewBuffer(respBody)
	var e Error
	err = json.NewDecoder(buf).Decode(&e)
	if err != nil {
		return fmt.Errorf("couldn't decode error: [%s]", respBody)
	}

	return fmt.Errorf("error %v: %s ", e.Code, e.Message)
}

// newRequest creates a new API request with context. If specified,
// the value pointed to by body is JSON encoded and included in the request body.
func (c *Client) newRequest(ctx context.Context, method, url string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(url)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// send makes a request to the API, the response body will be
// unmarshalled into v.
func (c *Client) send(req *http.Request, v interface{}) error {
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return c.decodeError(resp)
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

// get makes a GET request to the given url. The response body will be
// unmarshalled into v.
func (c *Client) get(ctx context.Context, url string, v interface{}) error {
	req, err := c.newRequest(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	err = c.send(req, v)
	if err != nil {
		return err
	}

	return nil
}
