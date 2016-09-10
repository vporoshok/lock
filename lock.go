// Package lock provide Lock class implemented mutex.Locker interface
package lock

import (
	"runtime"
	"sync"
)

// Lock is a `sync.Locker` with battaries.
//
// It is slowest sync.mutex approx 2 times but more powerfull
type Lock struct {
	m sync.Mutex
	b bool
}

// Similar `sync.Locker.Lock`
func (l *Lock) Lock() {
	for {
		l.m.Lock()
		if !l.b {
			l.b = true
			l.m.Unlock()

			return
		}
		l.m.Unlock()
		runtime.Gosched()
	}
}

// Similar to `sync.Locker.Unlock`
func (l *Lock) Unlock() {
	l.m.Lock()
	l.b = false
	l.m.Unlock()
}

// Wait a lock.
// If input chan will close early stop waiting and return false.
// Return true if locked before channel close.
func (l *Lock) Race(in <-chan struct{}) bool {
	for {
		select {
		case <-in:

			return false

		default:
			l.m.Lock()
			if !l.b {
				l.b = true
				l.m.Unlock()

				return true
			}
			l.m.Unlock()
			runtime.Gosched()
		}
	}

	return true
}

// Return true if lock is locked or false otherwise
func (l *Lock) Locked() bool {

	return l.b
}

// Try to get lock. Return false if lock is locked other process
func (l *Lock) TryLock() bool {
	l.m.Lock()
	if !l.b {
		l.b = true
		l.m.Unlock()

		return true
	}
	l.m.Unlock()

	return false
}
