package cp_go

import "time"

type Config struct {
	ApiSecret string
	PublicId  string
	Timeout   time.Duration
}

const (
	DefaultRequestTimeout = 30 * time.Second
)

type Client struct {
	config Config
}

func NewClient(cfg ...Config) *Client {
	client := &Client{
		// Create config
		config: Config{},
	}
	// Override config if provided
	if len(cfg) > 0 {
		client.config = cfg[0]
	}
	// Override default values
	if client.config.Timeout == 0 {
		client.config.Timeout = DefaultRequestTimeout
	}

	return client
}
