package parser

import "encoding/json"

type Parser interface {
	Parse(body []byte) error
}

func ParseReqeustBody(body []byte, r interface{}) error {
	err := json.Unmarshal(body, r)
	if err != nil {
		return err
	}
	return nil
}
