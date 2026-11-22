package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rafael-vincensi/gitrats-cli/internal/config"
)

type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
}

type APIError struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func NewClient() (*Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	return &Client{
		BaseURL:    cfg.APIURL,
		Token:      cfg.Token,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}, nil

}

func (c *Client) DoRequest(method, endpoint string, payload any, responseTarget any) error {
	var body io.Reader

	if payload != nil {
		jsonBytes, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(jsonBytes)
	}

	req, err := http.NewRequest(method, c.BaseURL+endpoint, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		var apiErr APIError
		json.Unmarshal(respBytes, &apiErr)
		if apiErr.Message != "" {
			return fmt.Errorf("API error (%d): %s", resp.StatusCode, apiErr.Message)
		}
		return fmt.Errorf("HTTP error %d", resp.StatusCode)
	}

	if responseTarget != nil {
		return json.Unmarshal(respBytes, responseTarget)
	}
	return nil
}
