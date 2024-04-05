package structs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/matryer/is"
)

func TestInvalidNewInput(t *testing.T) {
	is := is.New(t)
	data := "abc"
	_, err := NewStruct(data)
	is.True(err != nil) // should not be nil
}

func TestStructPtrNewInput(t *testing.T) {
	is := is.New(t)
	var data = struct {
		A string
		B int
		C bool
	}{
		A: "a-value",
		B: 2,
		C: true,
	}

	_, err := NewStruct(&data)
	is.NoErr(err)
}

func TestFieldsNoTag(t *testing.T) {
	is := is.New(t)

	var data = struct {
		A string
		B int
		C bool
	}{
		A: "a-value",
		B: 2,
		C: true,
	}
	s, err := NewStruct(data)
	is.NoErr(err)

	f := s.Fields("any")
	is.Equal(len(f), 0)
}

func TestFieldsNonMatchingTag(t *testing.T) {
	is := is.New(t)

	var data = struct {
		A string `tag:"A"`
		B int    `tag:"B"`
		C bool   `tag:"C"`
	}{
		A: "a-value",
		B: 2,
		C: true,
	}
	s, err := NewStruct(data)
	is.NoErr(err)

	f := s.Fields("any")
	is.Equal(len(f), 0)
}

func TestFieldsMatchingTag(t *testing.T) {
	is := is.New(t)

	var data = struct {
		A string `tag:"A"`
		B int    `tag:"B"`
		C bool   `tag:"C"`
	}{
		A: "a-value",
		B: 2,
		C: true,
	}
	s, err := NewStruct(data)
	is.NoErr(err)

	f := s.Fields("tag")
	is.Equal(len(f), 3)
}

func TestFieldsWithStructPtr(t *testing.T) {
	is := is.New(t)

	var data = struct {
		A string `tag:"A"`
		B int    `tag:"B"`
		C bool   `tag:"C"`
	}{
		A: "a-value",
		B: 2,
		C: true,
	}

	v := reflect.ValueOf(&data)
	fmt.Println("v", v.Kind())
	s := Struct{
		raw:   data,
		value: v,
	}

	f := s.Fields("tag")
	is.Equal(len(f), 3)
}
