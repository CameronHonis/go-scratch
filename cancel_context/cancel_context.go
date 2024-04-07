package cancel_context

import (
	"context"
	"fmt"
	"time"
)

func Main() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	time.Sleep(time.Second)
	asdf(ctx)
}

func asdf(ctx context.Context) {
	// Q: How can this function determine if a context has been cancelled?
	select {
	case <-ctx.Done():
		fmt.Println("context has been cancelled")
	}
	// A: Even after the context is cancelled, the select statement gets triggered
}
