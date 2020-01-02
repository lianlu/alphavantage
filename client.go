// Copyright 2019 Miles Barr <milesbarr2@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package alphavantage

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/tradyfinance/csvext"
	"github.com/tradyfinance/httpext"
)

// BaseURL is the base URL for the API.
const BaseURL = "https://www.alphavantage.co"

// A Client is an Alpha Vantage API client. Its zero value (DefaultClient) is a
// usable client that uses http.DefaultClient.
type Client struct {
	client *http.Client
	APIKey string
}

// NewClient returns a new Client given a HTTP client and API key. The HTTP
// client will default to http.DefaultClient when nil. The API key may be left
// empty, but not all endpoints will be accessible.
//
// An Alpha Vantage API key can be claimed at
// https://www.alphavantage.co/support/#api-key.
func NewClient(c *http.Client, apiKey string) *Client {
	if c == nil {
		c = http.DefaultClient
	}
	return &Client{c, apiKey}
}

// DefaultClient is the default client.
var DefaultClient = NewClient(nil, "")

func (c *Client) getJSON(path string, query url.Values, v interface{}) error {
	// Use DefaultClient when nil.
	if c == nil {
		c = DefaultClient
	}

	// Build the query string.
	if c.APIKey != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("apikey", c.APIKey)
	}

	// Build the URL.
	url := BaseURL + path
	if query != nil {
		if queryString := query.Encode(); queryString != "" {
			url += "?" + queryString
		}
	}

	// Create a HTTP request.
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")

	// Send the HTTP request.
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	if !httpext.IsSuccessStatus(resp.StatusCode) {
		return httpext.StatusError{URL: url, StatusCode: resp.StatusCode}
	}

	// Decode JSON from the HTTP response.
	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		resp.Body.Close()
		return err
	}
	return resp.Body.Close()
}

// ErrRateLimitExceeded indicates that the rate limit was exceeded.
var ErrRateLimitExceeded = errors.New("alphavantage: rate limit exceeded")

func (c *Client) getCSV(path string, query url.Values, f func(header, record []string) error) error {
	// Use DefaultClient when nil.
	if c == nil {
		c = DefaultClient
	}

	// Build the query string.
	if query == nil {
		query = url.Values{}
	}
	if c.APIKey != "" {
		query.Set("apikey", c.APIKey)
	}
	query.Set("datatype", "csv")

	// Build the URL.
	url := BaseURL + path + "?" + query.Encode()

	// Create a HTTP request.
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	//req.Header.Set("Accept", "text/csv")

	// Send the HTTP request.
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	if !httpext.IsSuccessStatus(resp.StatusCode) {
		return httpext.StatusError{URL: url, StatusCode: resp.StatusCode}
	}
	if resp.Header.Get("Content-Type") == "application/json" {
		resp.Body.Close()
		return ErrRateLimitExceeded
	}

	// Read a CSV table from the HTTP response.
	if err := csvext.ReadTable(resp.Body, f); err != nil {
		resp.Body.Close()
		return err
	}
	return resp.Body.Close()
}
