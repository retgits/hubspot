package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	// hubspotBaseURL is the base URL for the HubSpot API
	hubspotBaseURL = "https://api.hubapi.com/"
)

// Client contains all the functions to communicate between HubSpot and your app.
type Client struct {
	// HubSpot's APIs allow for two means of authentication, OAuth and API keys.  API keys are great for rapid prototyping.
	// You can generate a new API key under the Settings -> Integrations -> API key menu
	APIKey string
}

// NewClient returns a new Client pointer that can be chained with builder
// methods to set multiple configuration values inline without using pointers.
func NewClient() *Client {
	return &Client{}
}

// WithAPIKey sets a config API key value returning a Client pointer for
// chaining.
func (c *Client) WithAPIKey(apikey string) *Client {
	c.APIKey = apikey
	return c
}

// Call sends a request to HubSpot and receives the response.
func (c *Client) Call(urlSuffix string, httpMethod string, payload []byte) ([]byte, error) {
	var req *http.Request
	var err error

	if len(payload) > 0 {
		req, err = http.NewRequest(httpMethod, fmt.Sprintf("%s%s", hubspotBaseURL, urlSuffix), strings.NewReader(string(payload)))
	} else {
		req, err = http.NewRequest(httpMethod, fmt.Sprintf("%s%s", hubspotBaseURL, urlSuffix), nil)
	}

	req.Header["Content-Type"] = []string{"application/json"}

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	byteArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, fmt.Errorf(string(byteArray))
	}

	return byteArray, nil
}
