// Package unicontext contains a context whose timeout can be reset. 
// So one context can be used across multiple calls in a row.
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

// Value returns val from the parent context
func (this *UniContext) Value(key any) any {
	this.Lock()
	defer this.Unlock()

	return this.parent.Value(key)
}

// Reset cancels previous wrapped context automatically and wrap a new context of the same timeout
// parent context is unchanged
func (this *UniContext) Reset() *UniContext {
	this.Lock()
	defer this.Unlock()

	this.cancel()
	this.ctx, this.cancel = context.WithTimeout(this.parent, this.timeout)

	return this
}


// ResetTimeout cancels previous wrapped context automatically and wrap a new context of a new timeout
// parent context is unchanged
func (this *UniContext) ResetTimeout(tm time.Duration) *UniContext {
	this.Lock()
	defer this.Unlock()

	this.cancel()
	this.ctx, this.cancel = context.WithTimeout(this.parent, tm)
	this.timeout = tm

	return this
}


// Cancel the underlying context
func (this *UniContext) Cancel() {
	this.Lock()
	defer this.Unlock()

	this.cancel()
}


// WithTimeOut creates a UniContext which wraps a context of the specified timeout, bounded by the parent context.
func WithTimeOut(parent context.Context, timeout time.Duration) *UniContext {
	c, cancel := context.WithTimeout(parent, timeout)

	return &UniContext{
		parent: parent,
		ctx: c,
		cancel: cancel,
		timeout: timeout,
	}
}