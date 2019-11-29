package test

import (
	"bytes"
	"function-golang-helloworld/internal/openfaasfn"
	"testing"
)

func TestHandler(t *testing.T) {
	say := "Oops"
	input := bytes.NewBufferString(say)
	expected := "Hi,I'm OpenFaaS. I have received your message '" + say + "'"
	actual := openfaasfn.Handler(input.Bytes())
	if actual != expected {
		t.Errorf("actual %s; expected %s", actual, expected)
	}
}
