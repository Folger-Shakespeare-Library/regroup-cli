package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type GroupAddress struct {
	Street       string `json:"street"`
	City         string `json:"city"`
	AddressState string `json:"address_state"`
	Zip          string `json:"zip"`
	Country      string `json:"country"`
	AddressType  string `json:"address_type"`
	Lat          string `json:"lat"`
	Lng          string `json:"lng"`
}

type Group struct {
	ID                int          `json:"id"`
	Slug              string       `json:"slug"`
	Name              string       `json:"name"`
	Description       string       `json:"description"`
	ParentGroup       string       `json:"parent_group"`
	Access            string       `json:"access"`
	AlertModeration   bool         `json:"alert_moderation"`
	AllowMembersLeave int          `json:"allow_members_leave"`
	Contacts          int          `json:"contacts"`
	Members           int          `json:"members"`
	Address           GroupAddress `json:"address"`
}

type GroupsResponse struct {
	Count   json.RawMessage `json:"count"`
	Results []Group         `json:"results"`
}

type ListGroupsParams struct {
	Name     string
	Page     int
	Count    int
	SortBy   string
	SortType string
	All      bool
}

func (c *Client) ListGroups(p ListGroupsParams) (*GroupsResponse, error) {
	params := url.Values{}
	if p.Name != "" {
		params.Set("name", p.Name)
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

	body, err := c.get("/api/v3/groups", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}
	for i := range resp.Results {
		resp.Results[i].Slug = Slugify(resp.Results[i].Name)
	}
	return &resp, nil
}

func (c *Client) ResolveGroupID(slug string) (int, error) {
	resp, err := c.ListGroups(ListGroupsParams{All: true})
	if err != nil {
		return 0, err
	}
	for _, g := range resp.Results {
		if g.Slug == slug {
			return g.ID, nil
		}
	}
	return 0, fmt.Errorf("group not found: %s", slug)
}
