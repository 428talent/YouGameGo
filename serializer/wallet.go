package serializer

const (
	DefaultWalletSerializerTemplateType = "DefaultWalletTemplate"
)

func NewWalletTemplate(templateType string) Template {
	return &DefaultWalletTemplate{}
}

type DefaultWalletTemplate struct {
	Balance float64    `json:"balance"  source_type:"float"`
	Updated string     `json:"updated"  source_type:"string" converter:"time"`
	Link    []*ApiLink `json:"link"`
}

func (t *DefaultWalletTemplate) Serialize(model interface{}, context map[string]interface{}) {
	//data := model.(*models.Wallet)
	SerializeModelData(model, t)
	//site := context["site"].(string)
	t.Link = []*ApiLink{}
}
