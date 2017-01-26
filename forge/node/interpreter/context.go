package interpreter

type Context struct {
	parent      *Context
	interpreter *Interpreter
	data        map[string]interface{}
}

func (c *Context) Parent() *Context {
	return c.parent
}

func (c *Context) Parents(done <-chan struct{}) <-chan *Context {
	r := make(chan *Context)

	go func() {
		defer close(r)

		for ctx := c; ctx != nil; ctx = ctx.parent {
			select {
			case r <- ctx:
			case <-done:
				return
			}
		}
	}()

	return r
}
