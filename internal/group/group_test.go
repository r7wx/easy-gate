package group

import "testing"

func TestIsAllowed(t *testing.T) {
	if !IsAllowed([]Group{{
		Name:   "test",
		Subnet: "127.0.0.1/32",
	}}, []string{"test"}, "127.0.0.1") {
		t.Fail()
	}

	if IsAllowed([]Group{{
		Name:   "test",
		Subnet: "127.0.0.1/32",
	}}, []string{"test"}, "xxxxxx") {
		t.Fail()
	}

	if IsAllowed([]Group{{
		Name:   "test",
		Subnet: "xxxxxxxxxxx",
	}}, []string{"test"}, "127.0.0.1") {
		t.Fail()
	}
}
