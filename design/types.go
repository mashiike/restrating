package design

import (
	. "goa.design/goa/v3/dsl"
)

var RRNDefinition = func() {
	Description("Rating Resource Name")
	Example("rrn:player:XRQ85mtXnINISH25zfM0m5RlC6L2")
}
var Rating = Type("rating", func() {
	Attribute("strength", Float64, func() {
		Example(1500.0)
	})
	Attribute("lower", Float64, func() {
		Example(1300.0)
	})
	Attribute("upper", Float64, func() {
		Example(1700.0)
	})
	Required("strength", "lower", "upper")
})

var RatingResource = Type("RatingResource", func() {
	Description("RatingResource describes a strength information.")

	Attribute("rrn", String, "RRN is the unique id of the player/team.", RRNDefinition)
	Attribute("rating", Rating, "players strength")

	Required("rrn", "rating")
})
