//@Author: Dennis
//@Date: 2023/5/12

package service

import "testing"

func TestService(t *testing.T) {
	t.Run("test sync service", func(t *testing.T) {
		got := RemoteCallSync()
		want := "product+shipping->order"
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})
	t.Run("test async service", func(t *testing.T) {
		got := RemoteCallAsync()
		want := "product+shipping->order"
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})
}

func BenchmarkSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoteCallSync()
	}
}

func BenchmarkAsync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoteCallAsync()
	}
}
