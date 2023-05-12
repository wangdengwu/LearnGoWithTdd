//@Author: Dennis
//@Date: 2023/5/12

package counter

import "sync"

type Counter struct {
	sync.Mutex
	value int
}

func (c *Counter) IncAsync() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *Counter) IncSync() {
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
