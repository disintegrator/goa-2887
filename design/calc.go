package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("calc", func() {
	Method("add", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(Int)

		HTTP(func() {
			GET("/add/{a}/{b}")
			Response(StatusOK)
		})
	})

	Files("/swagger.json", "./gen/http/openapi.json")
})
