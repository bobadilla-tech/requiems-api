package counter

// Response model returned by both endpoints.
type Counter struct {
	Namespace string `json:"namespace"`
	Value     int64  `json:"value"`
}

func redisKey(namespace string) string {
	return "counter:" + namespace
}
