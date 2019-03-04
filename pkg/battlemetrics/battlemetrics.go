package battlemetrics

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Client for BattleMetrics
type Client struct {
	httpClient *http.Client
	AuthToken  string
}

// New client
func New(token string) *Client {
	return NewClient(token, http.DefaultClient)
}

// NewClient returns a new client
func NewClient(token string, c *http.Client) *Client {
	return &Client{
		httpClient: c,
		AuthToken:  token,
	}
}

func get(c *Client, url string) ([]byte, error) {
	return doHttpRequest(c, url, "GET", nil)
}

func post(c *Client, url string, payload interface{}) ([]byte, error) {
	return doHttpRequest(c, url, "POST", payload)
}

func doHttpRequest(c *Client, url, method string, payload interface{}) ([]byte, error) {
	var buffer *bytes.Buffer
	var err error
	if payload != nil {
		// prep payload
		buffer = &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		if err = encoder.Encode(payload); err != nil {
			return nil, err
		}
	}

	// prep request
	var req *http.Request
	if buffer != nil {
		req, err = http.NewRequest(method, url, buffer)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, err
	}
	req.Close = true
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Auth
	if c.AuthToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AuthToken))
	}

	// send it
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// read response
	var reader io.ReadCloser
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(res.Body)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
	default:
		reader = res.Body
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	data := buf.Bytes()

	// check for errors
	err = handleHTTPResponse(res.StatusCode, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func handleHTTPResponse(statusCode int, data []byte) error {
	var err error
	switch statusCode {
	case 200:
		return nil
	case 403:
		err = errors.New(http.StatusText(statusCode))
	case 500:
		fmt.Println(string(data))
		err = errors.New(http.StatusText(statusCode))
	default:
		err = fmt.Errorf("Unhandled HTTP status code: %v - %v", statusCode, http.StatusText(statusCode))
	}
	return err
}
