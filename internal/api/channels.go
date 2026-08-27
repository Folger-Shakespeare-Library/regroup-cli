package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type Channel struct {
	ID                           int      `json:"id"`
	Slug                         string   `json:"slug"`
	Name                         string   `json:"name"`
	Description                  string   `json:"description"`
	Emergency                    bool     `json:"emergency"`
	AdminPrivacy                 string   `json:"admin_privacy"`
	SubscribersPrivacy           string   `json:"subscribers_privacy"`
	Permission                   string   `json:"permission"`
	AdminApproval                bool     `json:"admin_approval"`
	NotifyAdmins                 bool     `json:"notify_admins"`
	SMSSubscriptionKeyword       string   `json:"sms_subscription_keyword"`
	Subscribers                  int      `json:"subscribers"`
	RequiresSubscriptionApproval bool     `json:"requires_subscription_approval"`
	ProcessingStyle              string   `json:"processing_style"`
	SyncSubscribersToGroups      []string `json:"sync_subscribers_to_groups"`
}

type ChannelsResponse struct {
	Count   json.RawMessage `json:"count"`
	Results []Channel       `json:"results"`
}

type ListChannelsParams struct {
	Name     string
	Page     int
	Count    int
	SortBy   string
	SortType string
	All      bool
}

func (c *Client) ListChannels(p ListChannelsParams) (*ChannelsResponse, error) {
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

	body, err := c.get("/api/v3/channels", params)
	if err != nil {
		return nil, err
	}

	var resp ChannelsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}
	for i := range resp.Results {
		resp.Results[i].Slug = Slugify(resp.Results[i].Name)
	}
	return &resp, nil
}
