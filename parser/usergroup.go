package parser

type CreateUserGroupRequestBody struct {
	Name string `json:"name"`
}

type AddPermissionRequestBody struct {
	Ids []int `json:"ids"`
}
