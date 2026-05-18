package advice

type Advice struct {
	ID   int    `json:"id"`
	Text string `json:"advice"`
}

func (Advice) IsData() {}


type BatchRequest struct {
	Count int `json:"count"`
}

type BatchResponse[T any] struct {
	Results []T `json:"results"`
}

func (BatchResponse[T]) IsData() {}