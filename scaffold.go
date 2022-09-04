package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type Scaffold struct {
	vecty.Core

	Child vecty.ComponentOrHTML
}

func (c Scaffold) Render() vecty.ComponentOrHTML {
	return elem.Body(c.Child)
}
