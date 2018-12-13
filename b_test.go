package b_test

import (
	"testing"

	"github.com/jhonghee/b"
)

func TestVersion(t *testing.T) {
	expected := "B v1.1->D v1.1->E v1.1"
	if b.Version() != expected {
		t.Error("Expected", expected, "but got", b.Version())
	}
}
