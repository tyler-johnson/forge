package interpreter

type Scope struct {
	parent *Scope
	data   map[string]interface{}
}

func NewScope() *Scope {
	return &Scope{}
}

func (c *Scope) Parent() *Scope {
	return c.parent
}

func (c *Scope) Parents(done <-chan struct{}) <-chan *Scope {
	r := make(chan *Scope)

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

func (c *Scope) Spawn() *Scope {
	return &Scope{
		parent: c,
	}
}

// func (c *Scope) Get(key string) interface{} {
// 	for parent := range c.Parents(nil) {

// 	}
// }

func (c *Scope) Set(key string, val interface{}) {
	c.data[key] = val
}
