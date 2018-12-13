package b_test

import (
	"testing"

	"github.com/jhonghee/b"
)

func TestVersion(t *testing.T) {
	expected := "B v1.2->D v1.3->E v1.2"
	if b.Version() != expected {
		t.Error("Expected", expected, "but got", b.Version())
	}
}
