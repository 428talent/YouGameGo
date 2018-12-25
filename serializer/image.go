package serializer

type ImageTemplate struct {
	Path string     `json:"path" source_type:"string"`
	Link []*ApiLink `json:"link"`
}

func (t *ImageTemplate) Serialize(model interface{}, context map[string]interface{}) {
	SerializeModelData(model, t)
	t.Link = []*ApiLink{}
}
