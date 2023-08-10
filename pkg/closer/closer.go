package closer

import (
	"context"
	"fmt"
	"sync"
)

type Func func(ctx context.Context) error

type Closer struct {
	mu    sync.Mutex
	funcs []Func
}

func (c *Closer) Add(f Func) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.funcs = append(c.funcs, f)
}

func (c *Closer) Close(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var (
		wg       = sync.WaitGroup{}
		msg      = make(chan error, len(c.funcs))
		complete = make(chan struct{}, 1)
	)

	go func() {
		for _, f := range c.funcs {
			wg.Add(1)
			go func(f Func) {
				defer wg.Done()

				if err := ctx.Err(); err != nil {
					return
				}

				if err := f(ctx); err != nil {
					msg <- err
				}
			}(f)
		}
	}()

	go func() {
		wg.Wait()
		close(msg)
		complete <- struct{}{}
	}()

	select {
	case <-complete:
		var msgErrors []string
		for val := range msg {
			if val != nil {
				msgErrors = append(msgErrors, val.Error())
			}
		}

		if len(msgErrors) > 0 {
			return fmt.Errorf("failed to shutdown funcs: %v", msgErrors)
		}
		return nil

	case <-ctx.Done():
		return fmt.Errorf("shutdown canceled: %v", ctx.Err())
	}
}
