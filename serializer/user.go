package serializer

import (
	"errors"
	"yougame.com/yougame-server/models"
)

type SerializeUser struct {
	Id        int                    `json:"id"`
	Username  string                 `json:"username"`
	LastLogin int64                  `json:"last_login"`
	CreateAt  int64                  `json:"create_at"`
	Profile   *SerializerUserProfile `json:"profile"`
}
type SerializerUserProfile struct {
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
			Username:data.Username,
			CreateAt:  data.Created.Unix(),
			Profile: &SerializerUserProfile{
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
