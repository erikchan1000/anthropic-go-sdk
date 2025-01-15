package anthropic

import (
  "net/http"
)

const (
	DefaultAPIVersion = "2023-06-01"
	DefaultBaseURL    = "https://api.anthropic.com/v1"
)

type Client struct {
  APIKey  string
  BaseURL string
  APIVersion string
  HTTPClient *http.Client
}
type Tool struct {
  Type string `json:"type"`
  Name string `json:"name"`
  DisplayWidthPx int `json:"display_width_px,omitempty"`
	DisplayHeightPx int `json:"display_height_px,omitempty"`
	DisplayNumber int `json:"display_number,omitempty"`
}

type Message struct {
  Role string `json:"role"`
  Content string `json:"content"`
}

type ComputerRequest struct {
	Model      string    `json:"model"`
	MaxTokens  int       `json:"max_tokens"`
	Tools      []Tool    `json:"tools"`
	Messages   []Message `json:"messages"`
}

type ComputerResponse struct {
	ID           string    `json:"id"`
	Type         string    `json:"type"`
	Role         string    `json:"role"`
	Content      string    `json:"content"`
	Model        string    `json:"model"`
	StopReason   string    `json:"stop_reason"`
	StopSequence string    `json:"stop_sequence"`
}

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
