package cmp

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T){
	expected := Person{
		Name: "Eddy",
		Age: 34,
	}
	result := createPerson("Eddy", 34)
	if diff := cmp.Diff(expected, result); diff != ""{
		t.Error(diff)
	}
}

func TestCreatePersonIgnoreDate(t *testing.T){
	expected := Person{
		Name: "Eddy",
		Age: 34,
	}
	result := createPerson("Eddy", 34)
	compare := cmp.Comparer(func(x, y Person)bool{
		return x.Name == y.Name && x.Age == y.Age
	})
	if diff := cmp.Diff(expected, result, compare); diff != ""{
		t.Error(diff)
	}
	if result.DateAdded.IsZero(){
		t.Error("DateAdded wasn't assigned")
	}
}