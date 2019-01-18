package serializer

import (
	"fmt"
	"yougame.com/yougame-server/models"
)

const (
	DefaultTransactionSerializeTemplateType = "DefaultTransaction"
)

func NewTransactionTemplate(templateType string) Template {
	return &DefaultTransactionTemplate{}
}

type DefaultTransactionTemplate struct {
	Id      int64      `json:"id"  source_type:"int"`
	Amount  float64    `json:"amount"  source_type:"float"`
	Type    string     `json:"type" source_type:"string"`
	Created string     `json:"created" source:"Created" source_type:"string" converter:"time"`
	UserId  int64      `json:"user_id"  source:"User.Id" source_type:"int"`
	OrderId int        `json:"order_id"`
	Link    []*ApiLink `json:"link"`
}

func (t *DefaultTransactionTemplate) Serialize(model interface{}, context map[string]interface{}) {
	data := model.(*models.Transaction)
	SerializeModelData(model, t)
	if data.Order != nil {
		t.OrderId = data.Order.Id
	}

	site := context["site"].(string)
	t.Link = []*ApiLink{
		{
			Rel:  "user",
			Href: fmt.Sprintf("%s/api/user/%d", site, data.User.Id),
			Type: "GET",
		},
	}
}
