package serializer

const (
	DefaultGameCollectionTemplateType = "DefaultGameCollection"
)

func NewGameCollectionTemplate(templateType string) Template {
	return &DefaultGameCollection{}
}

type DefaultGameCollection struct {
	Id      int        `json:"id" source_type:"int"`
	Name    string     `json:"name" source_type:"string"`
	Title   string     `json:"title" source_type:"string"`
	Created string     `json:"created" source:"Created" source_type:"string" converter:"time"`
	Updated string     `json:"updated" source:"Updated" source_type:"string" converter:"time"`
	Link    []*ApiLink `json:"link"`
}

func (t *DefaultGameCollection) Serialize(model interface{}, context map[string]interface{}) {
	SerializeModelData(model, t)
	t.Link = []*ApiLink{

	}
}
