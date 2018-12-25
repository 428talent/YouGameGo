package serializer

import (
	"errors"
	"fmt"
	"yougame.com/yougame-server/models"
)

type SerializeUser struct {
	Id        int               `json:"id"`
	Username  string            `json:"username"`
	LastLogin int64             `json:"last_login"`
	CreateAt  int64             `json:"create_at"`
	Profile   *UserProfileModel `json:"profile"`
}
type UserProfileModel struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	UpdateAt int64  `json:"update_at"`
}

func SerializeUserObject(data models.User, template interface{}) (interface{}, error) {
	switch template.(type) {
	case SerializeUser:
		models.ReadProfile(&data)
		return SerializeUser{
			Id:        data.Id,
			LastLogin: data.LastLogin.Unix(),
			Username:  data.Username,
			CreateAt:  data.Created.Unix(),
			Profile: &UserProfileModel{
				Nickname: data.Profile.Nickname,
				Email:    data.Profile.Email,
				Avatar:   data.Profile.Avatar,
				UpdateAt: data.Profile.Updated.Unix(),
			},
		}, nil
	}
	return nil, errors.New("template not match")
}

func SerializeUserList(data []*models.User, template interface{}) ([]*interface{}, error) {
	var result []*interface{}
	for _, user := range data {
		item, err := SerializeUserObject(*user, template)
		if err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

type UserLoginResponseBody struct {
	UserId int64
	Sign   string
}

func (p *UserLoginResponseBody) Serialize(sign string, user models.User) CommonApiResponseBody {
	p.UserId = int64(user.Id)
	p.Sign = sign
	return CommonApiResponseBody{
		Success: true,
		Payload: p,
	}
}

type UserSerializerModel struct {
	Id        int        `json:"id"`
	Username  string     `json:"username"`
	LastLogin int64      `json:"last_login"`
	CreateAt  int64      `json:"create_at"`
	Link      []*ApiLink `json:"link"`
}

func (s *UserSerializerModel) Serialize(model models.User, site string) *UserSerializerModel {
	return &UserSerializerModel{
		Id:        model.Id,
		Username:  model.Username,
		LastLogin: model.LastLogin.Unix(),
		CreateAt:  model.Created.Unix(),
		Link: []*ApiLink{
			&ApiLink{
				Rel:  "profile",
				Href: fmt.Sprintf("%s/api/user/%d/profile", site, model.Id),
				Type: "GET",
			},
		},
	}
}

func (p *UserProfileModel) Serialize(model models.Profile, site string) *UserProfileModel {
	return &UserProfileModel{
		Nickname: model.Nickname,
		Email:    model.Email,
		UpdateAt: model.Updated.Unix(),
		Avatar:   model.Avatar,
	}
}

const (
	DefaultUserTemplateType = "DefaultUserTemplateType"
)

func NewUserTemplate(templateType string) Template {
	return &UserTemplate{}
}

type UserTemplate struct {
	Id        int        `json:"id" source:"Id" source_type:"int"`
	Username  string     `json:"username" source:"Username" source_type:"string"`
	LastLogin string     `json:"last_login" source:"LastLogin" source_type:"string" converter:"time"`
	CreateAt  string     `json:"create_at" source:"Created" source_type:"string" converter:"time"`
	Link      []*ApiLink `json:"link"`
}

func (t *UserTemplate) Serialize(model interface{}, context map[string]interface{}) {
	data := model.(*models.User)
	SerializeModelData(data, t)
	site := context["site"].(string)
	t.Link = []*ApiLink{
		&ApiLink{
			Rel:  "profile",
			Href: fmt.Sprintf("%s/api/user/%d/profile", site, data.Profile.Id),
			Type: "GET",
		},
	}
}

const (
	DefaultProfileTemplateType = "DefaultProfileTemplateType"
)

func NewProfileTemplate(templateType string) Template {
	return &UserProfileTemplate{}
}

type UserProfileTemplate struct {
	Nickname string `json:"nickname" source:"Nickname" source_type:"string"`
	Email    string `json:"email" source:"Email" source_type:"string"`
	Avatar   string `json:"avatar" source:"Avatar" source_type:"string" `
	UpdateAt int64  `json:"update_at" source:"Updated" source_type:"string" converter:"time"`
}

func (t *UserProfileTemplate) Serialize(model interface{}, context map[string]interface{}) {
	SerializeModelData(model, t)
}
