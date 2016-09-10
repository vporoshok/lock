package lock

import (
	"sync"
	"testing"
	"time"

	"golang.org/x/net/context"
)

func TestLock_LockUnlock(t *testing.T) {
	l := Lock{}

	res := make([]int, 0, 100)

	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		a := i
		go func() {

			wg.Add(1)

			defer wg.Done()

			l.Lock()

			defer l.Unlock()

			for j := 0; j < 10; j++ {
				time.Sleep(time.Millisecond)
				res = append(res, a)
			}
		}()
	}

	wg.Wait()
	a := -1
	for i, x := range res {
		if i%10 == 0 {
			if a == x {
				t.Error("Order corupted")
			}
			a = x
		} else {
			if a != x {
				t.Error("Order corupted")
			}
		}
	}
}

func TestLock_Race(t *testing.T) {
	l := Lock{}
	l.Lock()

	go func() {
		defer l.Unlock()

		time.Sleep(10 * time.Millisecond)
	}()

	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond)
	if l.Race(ctx.Done()) {
		t.Error("Context should be expire before lock")
	}

	ctx, _ = context.WithTimeout(context.Background(), 12*time.Millisecond)
	if !l.Race(ctx.Done()) {
		t.Error("Lock should passing before context")
	}
}

func TestLock_Locked(t *testing.T) {
	l := Lock{}
	if l.Locked() {
		t.Error("New Lock should be unlocked")
	}
	l.Lock()
	if !l.Locked() {
		t.Error("Lock must be locked after lock")
	}
	l.Unlock()
	if l.Locked() {
		t.Error("Lock must be unlocked after unlock")
	}
}

func TestLock_TryLock(t *testing.T) {
	l := Lock{}
	if !l.TryLock() {
		t.Error("New Lock should be unlocked")
	}
	if l.TryLock() {
		t.Error("Lock must be locked after lock")
	}
	l.Unlock()
	if !l.TryLock() {
		t.Error("Lock must be unlocked after unlock")
	}
}
