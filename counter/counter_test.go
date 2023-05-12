//@Author: Dennis
//@Date: 2023/5/12

package counter

import (
	"sync"
	"testing"
)

func TestCount(t *testing.T) {
	t.Run("sync inc", func(t *testing.T) {
		counter := Counter{}
		counter.IncSync()
		counter.IncSync()
		counter.IncSync()
		got := counter.Value()
		want := 3
		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})
	t.Run("async inc", func(t *testing.T) {
		counter := Counter{}
		var wg sync.WaitGroup
		wg.Add(1000)
		for i := 0; i < 1000; i++ {
			go func() {
				counter.IncAsync()
				wg.Done()
			}()
		}
		wg.Wait()
		got := counter.Value()
		want := 1000
		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})
}

func BenchmarkSync(b *testing.B) {
	counter := Counter{}
	for i := 0; i < b.N; i++ {
		counter.IncSync()
	}
}

func BenchmarkAsync(b *testing.B) {
	counter := Counter{}
	for i := 0; i < b.N; i++ {
		go func() {
			counter.IncAsync()
		}()
	}
}
