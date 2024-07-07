package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %.3fms", c.StatusCode, c.Req.RequestURI, time.Since(t).Seconds()*1000)
	}
}
