package utils

import (
	"sync"
	"testing"
	"time"
)

func TestIDGen(t *testing.T) {
	inst := NewIDInstance(41, 12, 10, 9)
	var wg sync.WaitGroup
	wg.Add(100)
	results := make(chan int64, 10000)
	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				id := inst.NextID()
				// t.Logf("id: %b \t %x \t %d", id, id, id)
				results <- id
			}
			wg.Done()
		}()
	}
	wg.Wait()
	m := make(map[int64]bool)
	tc := make(chan int, 1)
	// var x int
	for i := 0; i < 10000; i++ {
		if i == 10 {
			tc <- i
		}
		select {
		case id := <-results:
			if _, ok := m[id]; ok {
				t.Errorf("Found duplicated id: %x", id)
			} else {
				m[id] = true
			}
		case <-time.After(0 * time.Second):
			t.Errorf("got %v", i)
			return
		}
	}
}

func BenchGenID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewIDInstance(41, 12, 10, 9).NextID()
	}
}

func BenchmarkGenIDP(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			NewIDInstance(41, 12, 10, 9).NextID()
		}
	})
}
