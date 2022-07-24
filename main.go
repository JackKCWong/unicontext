package unicontext 

import (
	"context"
	"sync"
	"time"
)


type UniContext struct {
	sync.Mutex
	parent context.Context
	ctx context.Context
	cancel context.CancelFunc
	timeout time.Duration
}

func (this *UniContext) Deadline() (deadline time.Time, ok bool) {
	this.Lock()
	defer this.Unlock()

	return this.ctx.Deadline()
}

func (this *UniContext) Done() <-chan struct{} {
	this.Lock()
	defer this.Unlock()

	return this.ctx.Done()
}

func (this *UniContext) Err() error {
	this.Lock()
	defer this.Unlock()

	return this.ctx.Err()
}

func (this *UniContext) Value(key any) any {
	this.Lock()
	defer this.Unlock()

	return this.parent.Value(key)
}

// Reset cancels previous context automatically and return a new context with the same timeout
func (this *UniContext) Reset() *UniContext {
	this.Lock()
	defer this.Unlock()

	this.cancel()
	this.ctx, this.cancel = context.WithTimeout(this.parent, this.timeout)

	return this
}


// ResetTimeout cancels previous context automatically and return a new context with a new timeout
func (this *UniContext) ResetTimeout(tm time.Duration) *UniContext {
	this.Lock()
	defer this.Unlock()

	this.cancel()
	this.ctx, this.cancel = context.WithTimeout(this.parent, tm)
	this.timeout = tm

	return this
}

func (this *UniContext) Cancel() {
	this.Lock()
	defer this.Unlock()

	this.cancel()
}


func WithTimeOut(parent context.Context, timeout time.Duration) *UniContext {
	c, cancel := context.WithTimeout(parent, timeout)

	return &UniContext{
		parent: parent,
		ctx: c,
		cancel: cancel,
		timeout: timeout,
	}
}