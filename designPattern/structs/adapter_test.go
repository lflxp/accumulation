package structs

import "testing"

var expect = "adaptee method"

func Test_Adapter(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	if res != expect {
		t.Fatalf("expect : %s, actual: %s", expect, res)
	}
}
