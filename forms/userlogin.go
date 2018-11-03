package forms

type UserLoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
