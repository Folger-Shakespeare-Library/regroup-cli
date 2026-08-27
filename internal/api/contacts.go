package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type SemicolonList []string

func (s *SemicolonList) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	if raw == "" {
		*s = []string{}
		return nil
	}
	parts := strings.Split(raw, ";")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	*s = out
	return nil
}

type Contact struct {
	FirstName        string        `json:"first_name"`
	LastName         string        `json:"last_name"`
	Email            string        `json:"email"`
	Username         string        `json:"username"`
	DatabaseID       string        `json:"databaseid"`
	OtherEmails      SemicolonList `json:"other_emails"`
	PhoneNumbers     SemicolonList `json:"phone_numbers"`
	CustomAttributes SemicolonList `json:"custom_attributes"`
	Groups           SemicolonList `json:"groups"`
	Address          string        `json:"address"`
	AutoDeleteAt     string        `json:"auto_delete_at"`
	PostsLanguage    string        `json:"posts_language"`
	CustomField      string        `json:"customfield"`
	PreferredMethod  string        `json:"preferred_method"`
}

func (c *Contact) UnmarshalJSON(data []byte) error {
	type Alias Contact
	aux := &struct {
		*Alias
		GroupName SemicolonList `json:"groupname"`
		Emails    SemicolonList `json:"emails"`
	}{Alias: (*Alias)(c)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	c.Groups = aux.GroupName
	c.OtherEmails = aux.Emails

	if c.OtherEmails == nil {
		c.OtherEmails = []string{}
	}
	if c.PhoneNumbers == nil {
		c.PhoneNumbers = []string{}
	}
	if c.CustomAttributes == nil {
		c.CustomAttributes = []string{}
	}
	if c.Groups == nil {
		c.Groups = []string{}
	}
	return nil
}

type ContactsResponse struct {
	Count   int       `json:"count"`
	Results []Contact `json:"results"`
}

type ScopedContact struct {
	FirstName        string        `json:"first_name"`
	LastName         string        `json:"last_name"`
	Email            string        `json:"email"`
	Username         string        `json:"username"`
	DatabaseID       string        `json:"databaseid"`
	OtherEmails      SemicolonList `json:"other_emails"`
	PhoneNumbers     SemicolonList `json:"phone_numbers"`
	CustomAttributes SemicolonList `json:"custom_attributes"`
	Address          string        `json:"address"`
	AutoDeleteAt     string        `json:"auto_delete_at"`
	PostsLanguage    string        `json:"posts_language"`
	CustomField      string        `json:"customfield"`
	PreferredMethod  string        `json:"preferred_method"`
}

type ScopedContactsResponse struct {
	Count   int             `json:"count"`
	Results []ScopedContact `json:"results"`
}

func (r *ContactsResponse) Scoped() *ScopedContactsResponse {
	out := &ScopedContactsResponse{Count: r.Count}
	for _, c := range r.Results {
		out.Results = append(out.Results, ScopedContact{
			FirstName:        c.FirstName,
			LastName:         c.LastName,
			Email:            c.Email,
			Username:         c.Username,
			DatabaseID:       c.DatabaseID,
			OtherEmails:      c.OtherEmails,
			PhoneNumbers:     c.PhoneNumbers,
			CustomAttributes: c.CustomAttributes,
			Address:          c.Address,
			AutoDeleteAt:     c.AutoDeleteAt,
			PostsLanguage:    c.PostsLanguage,
			CustomField:      c.CustomField,
			PreferredMethod:  c.PreferredMethod,
		})
	}
	return out
}

type ListContactsParams struct {
	Email         string
	Username      string
	DatabaseID    string
	Phone         string
	ShowPhoneType bool
	Page          int
	Count         int
	SortBy        string
	SortType      string
	All           bool
	Group         string
	Channel       string
}

func (c *Client) ListContacts(p ListContactsParams) (*ContactsResponse, error) {
	params := url.Values{}
	if p.Email != "" {
		params.Set("email", p.Email)
	}
	if p.Username != "" {
		params.Set("username", p.Username)
	}
	if p.DatabaseID != "" {
		params.Set("databaseid", p.DatabaseID)
	}
	if p.Phone != "" {
		params.Set("phone", p.Phone)
	}
	if p.ShowPhoneType {
		params.Set("show_phone_type", "true")
	}
	if p.Page > 0 {
		params.Set("page", strconv.Itoa(p.Page))
	}
	if p.Count > 0 {
		params.Set("count", strconv.Itoa(p.Count))
	}
	if p.SortBy != "" {
		params.Set("sort_by", p.SortBy)
	}
	if p.SortType != "" {
		params.Set("sort_type", p.SortType)
	}
	if p.All {
		params.Set("all", "true")
	}

	path := "/api/v3/contacts"
	if p.Group != "" {
		path = "/api/v3/groups/" + url.PathEscape(p.Group) + "/contacts"
	} else if p.Channel != "" {
		path = "/api/v3/channels/" + url.PathEscape(p.Channel) + "/contacts"
	}

	body, err := c.get(path, params)
	if err != nil {
		return nil, err
	}

	var resp ContactsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}
	return &resp, nil
}

type AddContactParams struct {
	FirstName  string
	LastName   string
	Email      string
	Phone      string
	Username   string
	DatabaseID string
	GroupIDs   string
	UserType   string
}

type importUser struct {
	FirstName  string `json:"firstname,omitempty"`
	LastName   string `json:"lastname,omitempty"`
	Email      string `json:"email"`
	Phone      string `json:"phone,omitempty"`
	Username   string `json:"username,omitempty"`
	DatabaseID string `json:"databaseid,omitempty"`
	GroupID    string `json:"groupid"`
	UserType   string `json:"usertype"`
}

type importUsersRequest struct {
	Users []importUser `json:"users"`
}

type ImportUsersResponse struct {
	SuccessMessages []string `json:"success_messages"`
	ErrorMessages   []string `json:"error_messages"`
}

func (c *Client) AddContact(p AddContactParams) (*ImportUsersResponse, error) {
	req := importUsersRequest{
		Users: []importUser{
			{
				FirstName:  p.FirstName,
				LastName:   p.LastName,
				Email:      p.Email,
				Phone:      p.Phone,
				Username:   p.Username,
				DatabaseID: p.DatabaseID,
				GroupID:    p.GroupIDs,
				UserType:   p.UserType,
			},
		},
	}

	body, err := c.post("/api/v3/orgs/import_users", req)
	if err != nil {
		return nil, err
	}

	var resp ImportUsersResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}
	return &resp, nil
}
