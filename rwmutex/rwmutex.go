package rwmutex

import (
	"context"
	"sync"
	"sync/atomic"

	"golang.org/x/sync/semaphore"
)

const maxReaders = 1 << 30

type RWMutex struct {
	mu            sync.Mutex
	readerCounter int32
	readerWait    int32
	readerSema    semaphore.Weighted
	writerSema    semaphore.Weighted
}

func (r *RWMutex) RLock() {
	if atomic.AddInt32(&r.readerCounter, 1) < 0 {
		r.readerSema.Acquire(context.Background(), 1)
	}
}

func (r *RWMutex) RUnlock() {
	if c := atomic.AddInt32(&r.readerCounter, -1); c < 0 {
		if atomic.AddInt32(&r.readerWait, -1) == 0 {
			r.writerSema.Release(1)
		}
	}
}

func (r *RWMutex) Lock() {
	r.mu.Lock()
	c := atomic.AddInt32(&r.readerCounter, -maxReaders) + maxReaders
	if c != 0 && atomic.AddInt32(&r.readerWait, c) != 0 {
		r.writerSema.Acquire(context.Background(), 1)
	}
}

func (r *RWMutex) Unlock() {
	c := atomic.AddInt32(&r.readerCounter, maxReaders)
	r.readerSema.Release(int64(c))
	r.mu.Unlock()
}
