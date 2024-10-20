package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func (c *Client) HardOpWithTimeout() (int, error) {
	url := fmt.Sprintf("%s://%s:%s/hard-op", c.Scheme, c.Host, c.Port)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return 0, fmt.Errorf("request timed out")
		}
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
