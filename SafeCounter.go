package goreco

import "sync"

type SafeCounter struct {
	m sync.Mutex
	n int
}

func (c *SafeCounter) Inc() {
	c.m.Lock()
	c.n++
	c.m.Unlock()
}

func (c *SafeCounter) Dec() {
	c.m.Lock()
	c.n--
	c.m.Unlock()
}

func (c *SafeCounter) Get() int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.n
}
