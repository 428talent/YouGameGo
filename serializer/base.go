package serializer

type Serializer interface {
	SerializeList() []*interface{}
	SerializeSerialize(data interface{}, output interface{})
}