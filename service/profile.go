package service

import (
	"errors"
	"yougame.com/yougame-server/models"
)

type UserProfileQueryBuilder struct {
	ids []interface{}
}

func (b *UserProfileQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	var count int64
	count = 0
	user := GetUserById(b.ids[0].(int))
	if user == nil {
		return &count, nil, errors.New("no found profile")
	}
	err := user.ReadProfile()
	if err != nil {
		return &count, nil, err
	}
	count = 1
	profiles := make([]*models.Profile, 0)
	profiles = append(profiles, user.Profile)
	return &count, profiles, nil
}

func (b *UserProfileQueryBuilder) SetPage(page int64, pageSize int64) {

}

func (b *UserProfileQueryBuilder) InId(id ...interface{}) {
	b.ids = append(b.ids, id...)
}
