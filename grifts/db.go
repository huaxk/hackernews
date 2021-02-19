package grifts

import (
	"fmt"

	. "github.com/markbates/grift/grift"
)

var _ = Desc("db", "Task Description")
var _ = Add("db", func(c *Context) error {
	fmt.Println("Hello world")
	return nil
})
