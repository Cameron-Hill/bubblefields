package bubbleform

import (
	"fmt"
	"reflect"
)

type Form struct {
	data   *any
	fields map[string]Field
}

func NewForm(data *any) Form {
	return Form{
		data: data,
	}
}

func (f *Form) Bind(key string, field Field) error {
	_, ok := reflect.TypeOf(f.data).Elem().FieldByName(key)
	if !ok {
		return fmt.Errorf("no such field: %s", key)
	}
	f.fields[key] = field
	return nil
}
