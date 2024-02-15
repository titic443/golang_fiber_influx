package validators

import "fmt"

type IError struct {
	Field string
	Tag   string
	Value interface{}
}

func (i IError) Error() string {
	r := fmt.Sprintf("Field: %v, Tag: %v, Value: %v", i.Field, i.Tag, i.Value)
	return r
}
