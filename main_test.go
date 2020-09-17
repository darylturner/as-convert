package main

import (
	"testing"
)

func TestConv(t *testing.T) {
	x, err := convertASplain("200077")
	if err != nil {
		t.Errorf("unexpected error in asplain->asdot conversion: %w", err)
	}

	y, err := convertASdot(x)
	if err != nil {
		t.Errorf("unexpected error in asdot->asplain conversion: %w", err)
	}

	if y != 200077 {
		t.Errorf("%v does not equal starting value %v", y, 200077)
	}
}
