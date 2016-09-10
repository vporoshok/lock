package lock

import (
	"sync"
	"testing"
)

func BenchmarkLock(b *testing.B) {
	l := NewLock()
	for i := 0; i < b.N; i++ {
		l.Lock()
		l.Unlock()
	}
}

func BenchmarkMutex(b *testing.B) {
	m := sync.Mutex{}
	for i := 0; i < b.N; i++ {
		m.Lock()
		m.Unlock()
	}
}
