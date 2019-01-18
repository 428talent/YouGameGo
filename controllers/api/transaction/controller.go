package transaction

import (
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
)

type Controller struct {
	api.ApiController
}

func (c *Controller) GetTransactionList() {
	c.WithErrorContext(func() {
		listView := api.ListView{
			Controller: &c.ApiController,
			Init: func() {
				c.GetAuth()
			},
			QueryBuilder:  &service.TransactionQueryBuilder{},
			ModelTemplate: serializer.NewTransactionTemplate(serializer.DefaultTransactionSerializeTemplateType),
			SetFilter: func(builder service.ApiQueryBuilder) {
				transactionQueryBuilder := builder.(*service.TransactionQueryBuilder)
				transactionQueryBuilder.InUser(c.User.Id)
			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}
	})
}
