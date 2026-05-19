package useragent

// ParseRequest holds the query parameters for the user agent parse endpoint.
type ParseRequest struct {
	UA string `query:"ua" validate:"required"`
}

// Result holds parsed user agent information.
type Result struct {
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	OS             string `json:"os"`
	OSVersion      string `json:"os_version"`
	Device         string `json:"device"`
	IsBot          bool   `json:"is_bot"`
}

// BatchParseRequest holds the body for validating multiple user agents at once.
type BatchParseRequest struct {
	UserAgents []string `json:"user_agents" validate:"required,min=1,max=50,dive,required"`
}

// BatchParseItem represents the parsed result for a single user agent in a batch request.
type BatchParseItem struct {
	UserAgent string `json:"user_agent"`
	Data      Result `json:"data"`
}

func (Result) IsData() {}
