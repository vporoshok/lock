// Package lock provide Lock class implemented mutex.Locker interface
package lock

import (
	"runtime"
	"sync"
)

// Lock is a sync.Locker with batteries.
//
// It is slowest sync.Mutex approx 2 times but more powerful
type Lock struct {
	m sync.Mutex
	b bool
}

// Lock locks l. If the l is already in use, the calling goroutine blocks until the lock is available.
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

// Unlock unlocks l. It is a run-time error if l is not locked on entry to Unlock.
//
// A locked Lock is not associated with a particular goroutine.
// It is allowed for one goroutine to lock a Mutex and then arrange for another goroutine to unlock it.
func (l *Lock) Unlock() {
	l.m.Lock()
	if !l.b {
		panic("lock: unlock of unlocked lock")
	}
	l.b = false
	l.m.Unlock()
}

// Race wait a lock.
//
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
}

// Locked return true if l is locked or false otherwise
func (l *Lock) Locked() bool {

	return l.b
}

// TryLock is try to get lock. Return false if lock is locked
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
