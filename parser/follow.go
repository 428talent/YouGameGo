package parser

type CreateFollowRequestBody struct {
	UserId      int `json:"user_id"`
	FollowingId int `json:"following_id"`
}
