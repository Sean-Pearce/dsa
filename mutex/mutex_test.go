package mutex

import (
	"sync"
	"testing"
)

var (
	m, l = 10000, 10
)

func TestNoLock(t *testing.T) {
	var n int

	wg := sync.WaitGroup{}
	for i := 0; i < l; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < m; j++ {
				n++
			}
			wg.Done()
		}()
	}
	wg.Wait()

	if n != m*l {
		t.Skipf("n = %d, want %d", n, m*l)
	}
}

func TestSync(t *testing.T) {
	var n int

	mu := &sync.Mutex{}
	var wg sync.WaitGroup
	for i := 0; i < l; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < m; j++ {
				mu.Lock()
				n++
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	if n != m*l {
		t.Errorf("n = %d, want %d", n, m*l)
	}
}

func TestMutex(t *testing.T) {
	var n int

	mu := &Mutex{}
	var wg sync.WaitGroup
	for i := 0; i < l; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < m; j++ {
				mu.Lock()
				n++
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	if n != m*l {
		t.Errorf("n = %d, want %d", n, m*l)
	}
}

func BenchmarkMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var n int64
		mu := &Mutex{}
		var wg sync.WaitGroup
		for i := 0; i < l; i++ {
			wg.Add(1)
			go func() {
				for j := 0; j < m; j++ {
					mu.Lock()
					n++
					mu.Unlock()
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var n int64
		mu := &sync.Mutex{}
		var wg sync.WaitGroup
		for i := 0; i < l; i++ {
			wg.Add(1)
			go func() {
				for j := 0; j < m; j++ {
					mu.Lock()
					n++
					mu.Unlock()
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
