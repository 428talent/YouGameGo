package serializer

type TagTemplate struct {
	Id   int64  `json:"id" source_type:"int"`
	Name string `json:"name" source_type:"string"`
}

func (*TagTemplate) CustomSerialize(convertTag string, value interface{}) interface{} {
	return value
}

func (t *TagTemplate) Serialize(model interface{}, context map[string]interface{}) {
	SerializeModelData(model, t)
}
