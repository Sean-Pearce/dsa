package mutex

import (
	"sync/atomic"
)

type Mutex struct {
	state int32 // 1 -> locked
}

func (m *Mutex) Lock() {
	for !atomic.CompareAndSwapInt32(&m.state, 0, 1) {
	}
}

func (m *Mutex) Unlock() {
	m.state = 0
}
