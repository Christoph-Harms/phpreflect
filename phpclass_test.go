package phpreflect

import (
	"reflect"
	"testing"
)

func TestNewPhpClass(t *testing.T) {
	phpClass := NewPhpClass()

	if reflect.TypeOf(phpClass) != reflect.TypeOf(PhpClass{}) {
		t.Error("NewPhpClass() did not return a PhpClass object.")
	}
}
