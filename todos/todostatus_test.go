package todos

import (
	"testing"
)

func TestIsEqual(t *testing.T) {
	if Complete != Complete {
		t.Errorf("Should be equal")
	}
}

func TestIsNotEqual(t *testing.T) {
	if Complete == Incomplete {
		t.Errorf("Should not be equal")
	}
}
