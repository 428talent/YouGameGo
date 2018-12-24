package serializer


type ImageTemplate struct {
	Path string `json:"path" source_type:"string"`
	Link []*ApiLink
}

func (*ImageTemplate) CustomSerialize(convertTag string, value interface{}) interface{} {
	return value
}

func (t *ImageTemplate) Serialize(model interface{}, context map[string]interface{}) {
	SerializeModelData(model, t)
	t.Link = []*ApiLink{}
}
