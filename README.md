# UniContext

A context whose timeout can be reset. So one context can be used across multiple calls in a row.

e.g.

```go

ctx := unicontext.WithTimeout(context.Background(), 1 * time.Second)
defer ctx.Cancel()

call1(ctx)
call2(ctx.Reset())      // Reset cancels previous context automatically and return a new context
call3(ctx.ResetTimeout(2 * time.Second))

```