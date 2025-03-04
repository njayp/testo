package manager

import "sync/atomic"

type Manager struct {
	count atomic.Int32
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Add(num int32) int32 {
	return m.count.Add(num)
}
