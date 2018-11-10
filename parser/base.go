package parser

import "encoding/json"

type RequestBodyParser interface {
	ParseBody([]byte) interface{}
}

func ParseReqeustBody(body []byte,r interface{}) error  {
	err := json.Unmarshal(body, r)
	if err != nil {
		return err
	}
	return nil
}