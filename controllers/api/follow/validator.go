package follow

import (
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/service"
)

type DuplicateFollowValidator struct {
}

func (*DuplicateFollowValidator) Check(context map[string]interface{}) bool {
	requestParser := context["parser"].(*parser.CreateFollowRequestBody)
	queryBuilder := service.FollowQueryBuilder{}
	queryBuilder.InUser(requestParser.UserId)
	queryBuilder.InFollowing(requestParser.FollowingId)
	count, _, err := queryBuilder.Query()
	if err != nil {
		return false
	}
	if *count != 0 {
		return false
	}
	return true
}

func (*DuplicateFollowValidator) GetMessage() string {
	return "already followed"
}
