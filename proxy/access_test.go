package proxy

import "testing"

func TestAccess(t *testing.T) {
	access := Proxy{
		PersonA{Name: "B"},
	}
	access.Visit()
	access = Proxy{
		PersonA{Name: "A"},
	}
	access.Visit()
}
