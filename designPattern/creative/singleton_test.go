package creative

import "testing"

func Test_Singleton(t *testing.T) {
	one := GetInstance()
	two := GetInstance()
	if one == two {
		t.Log("ok")
	} else {
		t.Fatal("different instance")
	}
}
