package api

import (
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
