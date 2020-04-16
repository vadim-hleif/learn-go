package main

import (
	"testing"
)

func TestToJSON(t *testing.T) {
	_, err := toJSON("asdasdas")

	if err != nil {
		t.Error("shouldn't convert string to json")
	}

	value, err := toJSON(person{
		First:   "test",
		Last:    "test",
		Sayings: []string{"one", "blabla", "test"},
	})

	if err != nil {
		t.Error("shouldn't fail on struct case")
	}

	if len(value) == 0 {
		t.Error("result shouldn't be empty")
	}



}
