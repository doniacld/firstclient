package firstclient

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Client holds the necessary information
type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

// ListUsers retrieves all users
func (c *Client) ListUsers() ([]User, error) {
	rel := &url.URL{
		Path: "/users",
	}
	u := u.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)
	return users, err
}
