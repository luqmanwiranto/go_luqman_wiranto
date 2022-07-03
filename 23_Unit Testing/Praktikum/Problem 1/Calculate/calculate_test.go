package calculate

import "testing"

func TestAddition(t *testing.T) {
	if Addition(1, 2) != 3 {
		t.Error("Expected 1 (+) 2 to equel 3")
	}
}

func TestSubtraction(t *testing.T) {
	if Subtraction(6, 3) != 1 {
		t.Error("Expected 6 (-) 3 to equel 3")
	}
}

func TestDivision(t *testing.T) {
	if Division(6, 2) != 3 {
		t.Error("Expected 6 (/) 2 to equel 3")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(3, 1) != 3 {
		t.Error("Expected 3 (*) 1 to equel 3")
	}
}
