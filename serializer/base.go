package serializer

type Model interface {
	SerializeData()
}

type CommonApiResponseBody struct {
	Success bool `json:"success"`
	Payload interface{} `json:"payload"`
}

