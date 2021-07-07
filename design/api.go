package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("superapi", func() {
	Server("calc", func() {
		Services("calc")
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})

	Server("zoo", func() {
		Services("zoo")
		Host("localhost", func() {
			URI("http://localhost:8001")
		})
	})

	HTTP(func() {
		Produces("application/json")
	})
})
