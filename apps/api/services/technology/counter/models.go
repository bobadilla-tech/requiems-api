package counter

// Response model returned by both endpoints.
type Counter struct {
	Namespace string `json:"namespace"`
	Value     int64  `json:"value"`
}

// BatchCounterItem is the result for a single item in a batch counter request.
type BatchCounterItem struct {
	Namespace string `json:"namespace"`
	Value     int64  `json:"value,omitempty"`
	Error     string `json:"error,omitempty"`
}

func redisKey(namespace string) string {
	return "counter:" + namespace
}
