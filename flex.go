package main

import "github.com/hexops/vecty"

func Flex(children vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("div", children, vecty.Markup(
		vecty.Class("flex"),
	))
}
