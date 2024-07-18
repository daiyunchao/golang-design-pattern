package simpleFactory

import (
	"testing"
)

func TestNewSF(t *testing.T) {
	sf := NewSF("t1")
	sf2 := NewSF("t2")
	sf3 := NewSF("t3")
	if sf.Name == "t1" {
		t.Logf("success sf Name is t1")
	} else {
		t.Errorf("fail to sf Name")
	}
	if sf2.Name == "t2" {
		t.Logf("success sf2 Name is t2")
	} else {
		t.Errorf("fail to sf2 Name")
	}
	if sf3.Name == "t3" {
		t.Logf("success sf2 Name is t3")
	} else {
		t.Errorf("fail to sf3 Name")
	}
}
