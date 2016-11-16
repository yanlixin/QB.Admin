package utils

import (
	"testing"
)

func TestPrintGuid(t *testing.T) {

	var except error = nil
	actual := PrintGuid()
	if except != actual {

		t.Errorf("expected %v, actual %v", except, actual)
	}
}
