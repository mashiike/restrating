package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/cors/dsl"
)

// API describes the global properties of the API server.
var _ = API("restrating", func() {
	Title("REST Rating Service")
	Description("HTTP service for rating")
	Version("1.0")
	Contact(func() {
		Name("mashiike")
		URL("https://github.com/mashiike/rating/issues")
	})
	Docs(func() {
		Description("github")
		URL("https://github.com/mashiike/restrating")
	})

	cors.Origin("/.*localhost.*/", func() {
		cors.Headers("content-type")
		cors.Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
		cors.MaxAge(600)
	})
	Server("restrating", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
	HTTP(func() {
		Path("/v1")
		Consumes("application/json")
		Produces("application/json")
	})
})

// Service describes a service
var _ = Service("Players", func() {
	Description("Player serves an indicator of strength using Rating")

	HTTP(func() {
		Path("/players")
	})

	Method("create player", func() {
		Description("Add new player and return its RRN(Rating Resource Name).")

		Payload(CreatePlayerPayload)
		Result(RRNResponse)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})
})

var _ = Service("Matches", func() {
	Description("Matches serves a learning function of Rating from match results")

	HTTP(func() {
		Path("/matches")
	})

	Method("Apply match", func() {
		Description("Apply match and return RatingResources")

		Payload(ApplyMatchPayload)
		Result(MatchResponse)

		HTTP(func() {
			POST("/")
			Response(StatusOK)
		})
	})
})
