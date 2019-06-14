package design

import (
	. "goa.design/goa/v3/dsl"
)

var RRNResponse = ResultType("application/vnd.restrating.rrn+json", func() {
	Description("RRN Only Response")
	ContentType("application/json")

	Attributes(func() {
		Attribute("rrn", String, RRNDefinition)
		Required("rrn")
	})

	View("default", func() {
		Attribute("rrn")
	})
})

var MatchResponse = ResultType("application/vnd.restrating.match+json", func() {
	Description("Match Response")
	ContentType("application/json")

	Attributes(func() {
		Attribute("participants", ArrayOf(RatingResource))
		Required("participants")
	})

	View("default", func() {
		Attribute("participants")
	})
})
