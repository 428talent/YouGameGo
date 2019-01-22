package serializer

const (
	DefaultTagTemplateType = "DefaultTagTemplateType"
)

func NewTagTemplate(templateType string) Template {
	return &DefaultTagTemplate{}
}

type DefaultTagTemplate struct {
	Id   int64  `json:"id" source_type:"int"`
	Name string `json:"name" source_type:"string"`
}

func (t *DefaultTagTemplate) Serialize(model interface{}, context map[string]interface{}) {
	SerializeModelData(model, t)
}
