package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("zoo", func() {
	Method("petAnimal", func() {
		Payload(func() {
			Field(1, "animal", String, "The animal to pet")
			Field(2, "duration", Int, "How long to pet animal")
			Required("animal", "duration")
		})

		Result(Int)

		HTTP(func() {
			POST("/pet")
			Response(StatusOK)
		})
	})
})
