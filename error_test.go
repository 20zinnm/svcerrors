package svcerrors

import "testing"

func TestNew(t *testing.T) {
	var se Error = New(200, "ok")
	_, ok := se.(error)
	if !ok {
		t.Fatalf("")
	}
	if se.Status() != 200 {
		t.Fatalf()
	}
}
