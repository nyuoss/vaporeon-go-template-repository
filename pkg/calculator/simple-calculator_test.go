package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	c := NewCalculator(2,2)
	want := 4
	answer := c.Add()
	if want != answer  {
		t.Fatalf(`2 + 2 = %d, want %d`, answer, want)
	}
}
