package clients

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/Matheus-Armando/go-api/config"
)

// JSONServerClient is a client for interacting with the JSON Server
type JSONServerClient struct {
    BaseURL    string
    HTTPClient *http.Client
}

// NewJSONServerClient creates a new JSON Server client
func NewJSONServerClient() *JSONServerClient {
    cfg := config.GetConfig()
    return &JSONServerClient{
        BaseURL: cfg.JSONServerURL,
        HTTPClient: &http.Client{
            Timeout: 10 * time.Second,
        },
    }
}

// Get performs a GET request to the JSON Server
func (c *JSONServerClient) Get(path string, result interface{}) error {
    url := fmt.Sprintf("%s/%s", c.BaseURL, path)
    
    resp, err := c.HTTPClient.Get(url)
    if err != nil {
        return fmt.Errorf("error making request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
        return fmt.Errorf("error decoding response: %w", err)
    }

    return nil
}