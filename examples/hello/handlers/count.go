package handlers

import (
	"fmt"
	"github.com/SkylakeCoder/go-web/web"
	"sync"
)

type Count struct {
	visitCount int64
	lock       sync.Mutex
}

func (c *Count) HandleRequest(req *web.Request, res *web.Response) {
	c.lock.Lock()
	c.visitCount++
	c.lock.Unlock()

	res.WriteString(fmt.Sprintf("count: %d\n", c.visitCount))
	res.Flush()
}
