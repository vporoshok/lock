package lock

import (
	"time"

	"golang.org/x/net/context"
)

func ExampleLock_Race() {
	l := Lock{}
	l.Lock()

	go func() {
		defer l.Unlock()

		time.Sleep(10 * time.Millisecond)
	}()

	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond)
	if l.Race(ctx.Done()) {
		panic("Context should be expire before lock")
	}

	ctx, _ = context.WithTimeout(context.Background(), 12*time.Millisecond)
	if !l.Race(ctx.Done()) {
		panic("Lock should passing before context")
	}
}
