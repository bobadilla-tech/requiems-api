package translate

import (
	"context"
	"fmt"

	googletranslate "cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

// Client wraps the Google Cloud Translation v2 API.
type Client struct {
	inner *googletranslate.Client
}

// NewClient creates a Translation API client.
// Authenticates via GOOGLE_APPLICATION_CREDENTIALS environment variable.
func NewClient(ctx context.Context) (*Client, error) {
	c, err := googletranslate.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("create translation client: %w", err)
	}
	return &Client{inner: c}, nil
}

// Close releases resources held by the client.
func (c *Client) Close() error {
	return c.inner.Close()
}

// Translate sends a batch of texts to Google for translation into targetLang.
// Returns translated strings in the same order as texts.
// sourceLang may be "" to let Google auto-detect.
func (c *Client) Translate(ctx context.Context, texts []string, targetLang, sourceLang string) ([]string, error) {
	if len(texts) == 0 {
		return nil, nil
	}

	target, err := language.Parse(targetLang)
	if err != nil {
		return nil, fmt.Errorf("invalid target language %q: %w", targetLang, err)
	}

	opts := &googletranslate.Options{
		Format: googletranslate.Text,
	}
	if sourceLang != "" {
		src, err := language.Parse(sourceLang)
		if err != nil {
			return nil, fmt.Errorf("invalid source language %q: %w", sourceLang, err)
		}
		opts.Source = src
	}

	results, err := c.inner.Translate(ctx, texts, target, opts)
	if err != nil {
		return nil, fmt.Errorf("translation API error: %w", err)
	}
	if len(results) != len(texts) {
		return nil, fmt.Errorf("API returned %d results for %d inputs", len(results), len(texts))
	}

	out := make([]string, len(results))
	for i, r := range results {
		out[i] = r.Text
	}
	return out, nil
}
