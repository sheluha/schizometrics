package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/sheluha/schizometrics/internal/config"
)

type Client struct {
	client *http.Client
	config config.Llm
}

func NewClient(llm config.Llm) *Client {
	return &Client{
		config: llm,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) SendRequest(ctx context.Context, state string, questions map[string]Question) (*Response, error) {
	payload := Request{
		Model:     "clef-flash",
		State:     state,
		Questions: questions,
	}

	pretty, _ := json.MarshalIndent(payload, "", "  ")
	log.Printf("Request: %s\n", string(pretty))

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.config.Url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("llm creating request error: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.config.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("llm request error: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4<<10))
		return nil, fmt.Errorf("llm request status code %d: %s", resp.StatusCode, body)
	}

	var result Response

	respBody, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, fmt.Errorf("llm request status code != 200: %w", err)
	}

	prettyresp, _ := json.MarshalIndent(result, "", "  ")
	log.Printf("Response: %s\n", prettyresp)

	return &result, nil
}
