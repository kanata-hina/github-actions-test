package main

import "testing"

func TestAdd(t *testing.T) {
	t.Run("add test", func(t *testing.T) {
		if got := add(10, 16); got != 25 {
			t.Errorf("add() = %v, want 25", got)
		}
	})
}
