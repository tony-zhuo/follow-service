package pager

type PagerResp struct {
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
}
