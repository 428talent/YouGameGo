package api_admin_game

import "errors"

type GameCreateValidate struct {
	requestData CreateGameRequest
}

func (v *GameCreateValidate) CheckValidation() (error) {
	if len(v.requestData.Name) == 0 {
		return errors.New("名字不能为空")
	}
	return nil
}
