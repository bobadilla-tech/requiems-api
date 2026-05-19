package useragent

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_Parse(t *testing.T) {
	t.Parallel()
	svc := NewService()

	tests := []struct {
		name           string
		ua             string
		wantBrowser    string
		wantBrowserVer string
		wantOS         string
		wantOSVer      string
		wantDevice     string
		wantBot        bool
	}{
		{
			name:           "Chrome on Windows",
			ua:             "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			wantBrowser:    "Chrome",
			wantBrowserVer: "120.0",
			wantOS:         "Windows",
			wantOSVer:      "10/11",
			wantDevice:     "desktop",
			wantBot:        false,
		},
		{
			name:           "Firefox on Windows",
			ua:             "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:121.0) Gecko/20100101 Firefox/121.0",
			wantBrowser:    "Firefox",
			wantBrowserVer: "121.0",
			wantOS:         "Windows",
			wantOSVer:      "10/11",
			wantDevice:     "desktop",
			wantBot:        false,
		},
		{
			name:           "Safari on macOS",
			ua:             "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_0) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Safari/605.1.15",
			wantBrowser:    "Safari",
			wantBrowserVer: "17.0",
			wantOS:         "macOS",
			wantOSVer:      "14.0",
			wantDevice:     "desktop",
			wantBot:        false,
		},
		{
			name:           "Edge on Windows",
			ua:             "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.2210.133",
			wantBrowser:    "Edge",
			wantBrowserVer: "120.0",
			wantOS:         "Windows",
			wantOSVer:      "10/11",
			wantDevice:     "desktop",
			wantBot:        false,
		},
		{
			name:           "Opera on Windows",
			ua:             "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 OPR/106.0.0.0",
			wantBrowser:    "Opera",
			wantBrowserVer: "106.0",
			wantOS:         "Windows",
			wantOSVer:      "10/11",
			wantDevice:     "desktop",
			wantBot:        false,
		},
		{
			name:           "Chrome on Android mobile",
			ua:             "Mozilla/5.0 (Linux; Android 14; Pixel 8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36",
			wantBrowser:    "Chrome",
			wantBrowserVer: "120.0",
			wantOS:         "Android",
			wantOSVer:      "14",
			wantDevice:     "mobile",
			wantBot:        false,
		},
		{
			name:           "Safari on iPhone",
			ua:             "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1",
			wantBrowser:    "Safari",
			wantBrowserVer: "17.0",
			wantOS:         "iOS",
			wantOSVer:      "17.0",
			wantDevice:     "mobile",
			wantBot:        false,
		},
		{
			name:           "Safari on iPad",
			ua:             "Mozilla/5.0 (iPad; CPU OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1",
			wantBrowser:    "Safari",
			wantBrowserVer: "17.0",
			wantOS:         "iOS",
			wantOSVer:      "17.0",
			wantDevice:     "tablet",
			wantBot:        false,
		},
		{
			name:           "Googlebot",
			ua:             "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			wantBrowser:    "",
			wantBrowserVer: "",
			wantOS:         "",
			wantOSVer:      "",
			wantDevice:     "bot",
			wantBot:        true,
		},
		{
			name:           "empty UA",
			ua:             "",
			wantBrowser:    "",
			wantBrowserVer: "",
			wantOS:         "",
			wantOSVer:      "",
			wantDevice:     "unknown",
			wantBot:        false,
		},
		{
			name:           "Linux desktop",
			ua:             "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			wantBrowser:    "Chrome",
			wantBrowserVer: "120.0",
			wantOS:         "Linux",
			wantOSVer:      "",
			wantDevice:     "desktop",
			wantBot:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := svc.Parse(tt.ua)

			assert.Equal(t, tt.wantBrowser, got.Browser)
			assert.Equal(t, tt.wantBrowserVer, got.BrowserVersion)
			assert.Equal(t, tt.wantOS, got.OS)
			assert.Equal(t, tt.wantOSVer, got.OSVersion)
			assert.Equal(t, tt.wantDevice, got.Device)
			assert.Equal(t, tt.wantBot, got.IsBot)
		})
	}
}

func TestService_BatchParse(t *testing.T) {
	t.Parallel()
	svc := NewService()

	userAgents := []string{
		"Mozilla/5.0 (Linux; Android 14; Pixel 8) AppleWebKit/537.36",
		"PostmanRuntime/7.36.0",
		"curl/8.4.0",
	}

	results := svc.ParseBatch(context.Background(), userAgents)

	require.Len(t, results, 3)

	require.Equal(t, "Android Browser", results[0].Data.Browser)
	require.Equal(t, "mobile", results[0].Data.Device)
	assert.False(t, results[0].Data.IsBot)

	require.Equal(t, "unknown", results[1].Data.Device)
	assert.False(t, results[1].Data.IsBot)

	require.Equal(t, "bot", results[2].Data.Device)
	assert.True(t, results[2].Data.IsBot)
}

func TestService_BatchParse_PreservedOrder(t *testing.T) {
	t.Parallel()
	svc := NewService()

	userAgents := []string{
		"Mozilla/5.0 (Linux; Android 14; Pixel 8) AppleWebKit/537.36",
		"PostmanRuntime/7.36.0",
		"curl/8.4.0",
	}

	results := svc.ParseBatch(context.Background(), userAgents)

	require.Len(t, results, 3)

	for i, ua := range userAgents {
		assert.Equal(t, ua, results[i].UserAgent)
	}
}
