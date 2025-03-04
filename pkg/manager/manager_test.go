package manager

import "testing"

func TestManager(t *testing.T) {
	m := NewManager()
	num := m.Add(2)
	if num != 2 {
		t.Error(num)
	}
}
