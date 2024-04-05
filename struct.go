package structs

import (
	"errors"
	"fmt"
	"reflect"
)

type Struct struct {
	raw   any
	value reflect.Value
}

func (s *Struct) Fields(tagName string) []*Field {
	v := s.value
	if s.value.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()
	fields := []*Field{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if tag := field.Tag.Get(tagName); tag == "" {
			continue
		}

		fields = append(fields, &Field{
			field: field,
			value: v.FieldByName(field.Name),
		})
	}

	return fields
}

type Field struct {
	value reflect.Value
	field reflect.StructField
}

func (f Field) String() string {
	return fmt.Sprintf("%s %s `%s`", f.field.Name, f.field.Type, f.field.Tag)
}

func NewStruct(s any) (Struct, error) {
	v, ok := strctVal(s)
	if !ok {
		return Struct{}, errors.New("input must be of type struct")

	}
	return Struct{
		raw:   s,
		value: v,
	}, nil

}

func strctVal(s interface{}) (reflect.Value, bool) {
	v := reflect.ValueOf(s)

	// if pointer get the underlying element≤
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return v, false
	}

	return v, true
}
