package svcerror

import "testing"

func TestNew(t *testing.T) {
	var se Error = New(200, "ok")
	_, ok := se.(error)
	if !ok {
		t.Fatalf("could not cast svcerrors.Error to ")
	}
	if se.Status() != 200 {
		t.Fatalf()
	}
}
