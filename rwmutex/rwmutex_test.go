package rwmutex

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestReader(t *testing.T) {
	num := 0
	rwlock := RWMutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			rwlock.RLock()
			log.Println(num)
			time.Sleep(time.Second)
			rwlock.RUnlock()
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestWriter(t *testing.T) {
	num := 0
	rwlock := RWMutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			rwlock.Lock()
			num++
			rwlock.Unlock()
			wg.Done()
		}()
	}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			rwlock.Lock()
			num--
			rwlock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	if num != 0 {
		t.Errorf("num = %v\n", num)
	}
}

// TODO: read write test
