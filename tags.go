package mapstructure

import (
	"reflect"
	"strings"
)

type FieldTag struct {
	field     reflect.StructField
	tags      []string

	Name      string
	Squash    bool
	Skip      bool
}

// ContainsOneof checks if any of a slice of substrings are contained in an input string
func ContainsOneof(s string, inputs []string) bool {
	for _, substring := range inputs {
		if strings.Contains(s, substring) {
			return true
		}
	}
	return false
}

func NewFieldTag(f reflect.StructField, tagName string, jsonFallback bool) FieldTag {
	fieldTag := FieldTag{
		field: f,
	}

	tagValue := f.Tag.Get(tagName)
	if tagValue == "" && jsonFallback {
		tagValue = f.Tag.Get("json")
	}

	var fieldName string
	var fieldTags []string

	tagParts := strings.Split(tagValue, ",")
	if len(tagParts) == 0 {
		return fieldTag
	}

	fieldName = tagParts[0]
	if fieldName == "-" {
		fieldTag.Skip = true
		fieldName = ""
	}
	fieldTag.Name = fieldName

	fieldTags = tagParts[1:]
	for _, tag := range fieldTags {
		if ContainsOneof(tag, []string{"squash", "inline"}) {
			fieldTag.Squash = true
		}
	}
	fieldTag.tags = fieldTags

	return fieldTag
}
