package goutils

import (
	"fmt"
	"testing"
)

func TestPointerOf(t *testing.T) {
	str_var := "string"
	int_var := 10
	bool_var := true

	str_point := PointerOf[string](str_var)
	int_point := PointerOf[int](int_var)
	bool_point := PointerOf[bool](bool_var)

	if typeOf(str_point) != "*string" || *str_point != "string" {
		t.Errorf("error while returning a pointer to string")
	}
	if typeOf(int_point) != "*int" || *int_point != 10 {
		t.Errorf("error while returning a pointer to int")
	}
	if typeOf(bool_point) != "*bool" || !*bool_point {
		t.Errorf("error while returning a pointer to bool")
	}
}

func typeOf(t interface{}) string {
	return fmt.Sprintf("%T", t)
}
