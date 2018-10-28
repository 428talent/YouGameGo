package serializer

type Serializer interface {
	SerializeList() []*interface{}
	SerializeSerialize(data interface{}, output interface{})
}

type CommonApiResponseBody struct {
	Success bool `json:"success"`
	Payload interface{} `json:"payload"`
}
