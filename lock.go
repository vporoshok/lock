// Package lock provide Lock class implemented mutex.Locker interface
package lock

import (
	"runtime"
	"sync"
)

// Lock is a `sync.Locker` with battaries
type Lock interface {
	// Similar `sync.Locker.Lock`
	Lock()
	// Similar to `sync.Locker.Unlock`
	Unlock()
	// Wait a lock.
	// If input chan will close early stop waiting and return false.
	// Return true if locked before channel close.
	Race(<-chan struct{}) bool
	// Return true if lock is locked or false otherwise
	Locked() bool
	// Try to get lock. Return false if lock is locked other process
	TryLock() bool
}

// It is slowest sync.mutex more than 3 times but more powerfull
type lock struct {
	m sync.Mutex
	b bool
}

func NewLock() Lock {

	return &lock{}
}

func (l *lock) Lock() {
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

func (l *lock) Unlock() {
	l.m.Lock()
	l.b = false
	l.m.Unlock()
}

func (l *lock) Race(in <-chan struct{}) bool {
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

func (l *lock) Locked() bool {

	return l.b
}

func (l *lock) TryLock() bool {
	l.m.Lock()
	if !l.b {
		l.b = true
		l.m.Unlock()

		return true
	}
	l.m.Unlock()

	return false
}
