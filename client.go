package anthropic

import (
  "bytes"
  "encoding/json"
  "fmt"
  "net/http"
)

func NewClient(apiKey string) *Client {
  return &Client {
    APIKey: apiKey,
    BaseURL: DefaultBaseURL,
    APIVersion: DefaultAPIVersion,
    HTTPClient: &http.Client{},
  }
}

func (c *Client) CreateComputer(req *ComputerRequest) (*ComputerResponse, error) {
  body, err := json.Marshal(req)

  if err != nil {
    return nil, fmt.Errorf("Error marshaling request")
  }

  httpReq, err := http.NewRequest("Post", fmt.Sprintf("%s/messages", c.BaseURL), bytes.NewReader(body))

  if err != nil {
    return nil, fmt.Errorf("Error creating request")
  }

  httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-API-Key", c.APIKey)
	httpReq.Header.Set("Anthropic-Version", c.APIVersion)
	httpReq.Header.Set("Anthropic-Beta", "computer-use-2024-10-22")

  resp, err := c.HTTPClient.Do(httpReq)

  if err != nil {
    return nil, fmt.Errorf("Error sending response")
  }

  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    var apiErr Error
    if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
      return nil, fmt.Errorf("error decoding error response: %v", err)
    }
		return nil, fmt.Errorf("API error: %s - %s", apiErr.Type, apiErr.Message)
  }

	var result ComputerResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &result, nil
}
