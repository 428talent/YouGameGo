package parser

type RequestBodyParser interface {
	ParseBody([]byte) interface{}
}
