package wallet

import (
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
)

type Controller struct {
	api.ApiController
}

func (c *Controller) GetWallet() {
	c.WithErrorContext(func() {
		objectView := api.ObjectView{
			Controller:    &c.ApiController,
			QueryBuilder:  &service.WalletQueryBuilder{},
			ModelTemplate: serializer.NewWalletTemplate(serializer.DefaultWalletSerializerTemplateType),
			LookUpField:   "-",
			SetFilter: func(builder service.ApiQueryBuilder) {
				walletQueryBuilder := builder.(*service.WalletQueryBuilder)
				if security.CheckUserGroup(c.User, security.UserGroupAdmin) {
					walletQueryBuilder.InUser(c.Ctx.Input.Param(":id"))
				}else{
					walletQueryBuilder.InUser(c.User.Id)
				}

			},
		}
		err := objectView.Exec()
		if err != nil {
			panic(err)
		}
	})
}
