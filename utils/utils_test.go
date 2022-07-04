package utils

import "testing"

func TestSmallUint(t *testing.T) {
	result := Base62(11157)
	if result != "2TX" {
		t.Fatalf(`Incorrect result %+v != "2TX"`, result)
	}
}

func TestUint64(t *testing.T) {
	result := Base62(2009215674938)
	if result != "zn9edcu" {
		t.Fatalf(`Incorrect result %+v != "zn9edcu"`, result)
	}
}
