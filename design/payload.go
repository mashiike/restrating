package design

import (
	. "goa.design/goa/v3/dsl"
)

var CreatePlayerPayload = Type("CreatePlayerPayload", func() {
	Attribute("name", String, func() { Example("XRQ85mtXnINISH25zfM0m5RlC6L2") })
	Required("name")
})

var ApplyMatchPayload = Type("ApplyMatchPayload", func() {
	Attribute("scores", MapOf(String, Float64), func() {
		Key(RRNDefinition)
		Elem(func() { Example(1.0) })
	})
	Required("scores")
})
