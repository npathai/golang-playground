package unittest_test

import "testing"
import "../unittest"

func TestGreet(t *testing.T) {
	actual := unittest.Greet("Narendra")
	expected := "Hello, Narendra"
	if actual != expected {
		t.Errorf("Expected: %v, Got: %v\n", expected, actual)
	}
}