package coinbasepro

import (
	"fmt"
)

type Profile struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	Active    bool   `json:"active"`
	IsDefault bool   `json:"is_default"`
	CreatedAt Time   `json:"created_at,string"`
}

type ProfileTransfer struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

// Client Funcs

// GetProfiles retrieves a list of profiles
func (c *Client) GetProfiles() ([]Profile, error) {
	var profiles []Profile

	url := fmt.Sprintf("/profiles")
	_, err := c.Request("GET", url, nil, &profiles)
	return profiles, err
}

// GetProfile retrieves a single profile
func (c *Client) GetProfile(id string) (Profile, error) {
	var profile Profile

	url := fmt.Sprintf("/profiles/%s", id)
	_, err := c.Request("GET", url, nil, &profile)
	return profile, err
}

// CreateProfileTransfer transfers a currency amount from one profile to another
func (c *Client) CreateProfileTransfer(newTransfer *ProfileTransfer) error {
	url := fmt.Sprintf("/profiles/transfer")
	_, err := c.Request("POST", url, newTransfer, nil)

	return err
}
