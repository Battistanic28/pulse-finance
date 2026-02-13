package client

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	APIKey     string
	HTTPClient *http.Client
}

func (c *Client) httpClient() *http.Client {
	if c.HTTPClient != nil {
		return c.HTTPClient
	}
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

func (c *Client) Fetch(ticker string, lookbackHours int) ([]byte, error) {
	from := time.Now().UTC().Add(-time.Duration(lookbackHours) * time.Hour)
	timeFromStr := from.Format("20060102T1504")

	endpoint := fmt.Sprintf(
		"https://www.alphavantage.co/query?function=NEWS_SENTIMENT&tickers=%s&time_from=%s&apikey=%s",
		ticker,
		timeFromStr,
		c.APIKey,
	)

	resp, err := c.httpClient().Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}
