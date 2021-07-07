// Code generated by goa v3.4.3, DO NOT EDIT.
//
// zoo HTTP client CLI support package
//
// Command:
// $ goa gen bugrepro/design

package client

import (
	zoo "bugrepro/gen/zoo"
	"encoding/json"
	"fmt"
)

// BuildPetAnimalPayload builds the payload for the zoo petAnimal endpoint from
// CLI flags.
func BuildPetAnimalPayload(zooPetAnimalBody string) (*zoo.PetAnimalPayload, error) {
	var err error
	var body PetAnimalRequestBody
	{
		err = json.Unmarshal([]byte(zooPetAnimalBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"animal\": \"Ut distinctio molestiae quis velit at qui.\",\n      \"duration\": 344612583194791941\n   }'")
		}
	}
	v := &zoo.PetAnimalPayload{
		Animal:   body.Animal,
		Duration: body.Duration,
	}

	return v, nil
}
