package serializer

type ImageTemplate struct {
	Id   int     `json:"id" source_type:"int"`
	Path string     `json:"path" source_type:"string"`
	Enable bool `json:"enable" source_type:"bool"`
	Link []*ApiLink `json:"link"`
}

func (t *ImageTemplate) Serialize(model interface{}, context map[string]interface{}) {
	SerializeModelData(model, t)
	t.Link = []*ApiLink{}
}

