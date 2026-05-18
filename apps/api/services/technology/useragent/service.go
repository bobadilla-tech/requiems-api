package useragent

import (
	"regexp"
	"strings"

	ua "github.com/medama-io/go-useragent"
)

// Result holds parsed user agent information.
type Result struct {
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	OS             string `json:"os"`
	OSVersion      string `json:"os_version"`
	Device         string `json:"device"`
	IsBot          bool   `json:"is_bot"`
}

func (Result) IsData() {}

// osVersionRegexes maps OS names (as returned by the library) to regexes that
// extract the version string from the raw UA.
var osVersionRegexes = map[string]*regexp.Regexp{
	"Windows": regexp.MustCompile(`Windows NT (\d+\.\d+)`),
	"MacOS":   regexp.MustCompile(`Mac OS X (\d+[._]\d+)`),
	"Android": regexp.MustCompile(`Android (\d+(?:\.\d+)?)`),
	"iOS":     regexp.MustCompile(`OS (\d+[._]\d+)`),
}

// windowsVersions maps NT version strings to marketing names.
var windowsVersions = map[string]string{
	"10.0": "10/11",
	"6.3":  "8.1",
	"6.2":  "8",
	"6.1":  "7",
	"6.0":  "Vista",
	"5.2":  "XP x64",
	"5.1":  "XP",
}

// Service wraps the go-useragent parser. Initialize once at startup.
type Service struct {
	parser *ua.Parser
}

// NewService constructs a new Service. The underlying parser is initialized once.
func NewService() *Service {
	return &Service{parser: ua.NewParser()}
}

// Parse extracts browser, OS, device, and bot information from a UA string.
func (s *Service) Parse(uaStr string) Result {
	if uaStr == "" {
		return Result{Device: "unknown"}
	}

	agent := s.parser.Parse(uaStr)

	osName := normalizeOS(string(agent.OS()))
	device := normalizeDevice(string(agent.Device()))

	return Result{
		Browser:        string(agent.Browser()),
		BrowserVersion: buildBrowserVersion(agent),
		OS:             osName,
		OSVersion:      extractOSVersion(uaStr, string(agent.OS())),
		Device:         device,
		IsBot:          agent.IsBot(),
	}
}

// normalizeOS converts library OS names to the canonical form used in responses.
// The library uses "MacOS" internally; we return "macOS" to match the casing
// used by Apple and the endpoint response spec.
func normalizeOS(os string) string {
	if os == "MacOS" {
		return "macOS"
	}
	return os
}

// normalizeDevice lowercases the title-case device string returned by the library,
// and falls back to "unknown" when no device was detected.
func normalizeDevice(device string) string {
	if device == "" {
		return "unknown"
	}
	return strings.ToLower(device)
}

// buildBrowserVersion returns "major.minor" (e.g. "120.0") or an empty string.
func buildBrowserVersion(agent ua.UserAgent) string {
	major := agent.BrowserVersionMajor()
	if major == "" {
		return ""
	}
	minor := agent.BrowserVersionMinor()
	if minor == "" {
		return major
	}
	return major + "." + minor
}

// extractOSVersion extracts a version string from the raw UA using a lightweight
// regex, since go-useragent only tracks browser versions.
func extractOSVersion(uaStr, libOS string) string {
	re, ok := osVersionRegexes[libOS]
	if !ok {
		return ""
	}
	m := re.FindStringSubmatch(uaStr)
	if len(m) < 2 {
		return ""
	}
	v := strings.ReplaceAll(m[1], "_", ".")
	if libOS == "Windows" {
		if friendly, ok := windowsVersions[v]; ok {
			return friendly
		}
	}
	return v
}
