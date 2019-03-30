package lib

import "testing"

func TestSayHello(t *testing.T) {
	if SayHello() != "Hello World" {
		t.Logf("Wrong Message")
		t.Fail()
	}
}
