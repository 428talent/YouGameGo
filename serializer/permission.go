package serializer

var (
	DefaultPermissionTemplateType = "DefaultPermissionTemplateType"
)
func NewPermissionTemplate(templateType string) Template {
	return &PermissionTemplate{}
}
type PermissionTemplate struct {
	Id     int    `json:"id" source:"Id" source_type:"int"`
	Name   string `json:"name" source:"Name" source_type:"string"`
	Enable bool   `json:"enable" source_type:"bool"`
}

func (t *PermissionTemplate) Serialize(model interface{}, context map[string]interface{}) {
	SerializeModelData(model, t)
}