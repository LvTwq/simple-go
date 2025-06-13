package models

type IHttpCode interface {
	Code() string
	Desc() string
}

type AjaxResult[T any] struct {
	Code     int    `json:"code,omitempty"`
	Data     T      `json:"data,omitempty"`
	Messages string `json:"messages,omitempty"`
}

func NewAjaxResult[T any](code int, data T) *AjaxResult[T] {
	return &AjaxResult[T]{
		Code: code,
		Data: data,
	}
}
