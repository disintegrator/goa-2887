// Code generated by goa v3.4.3, DO NOT EDIT.
//
// zoo HTTP client CLI support package
//
// Command:
// $ goa gen bugrepro/design

package cli

import (
	zooc "bugrepro/gen/http/zoo/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `calc add
zoo pet-animal
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` calc add --a 1485553568483242564 --b 1900791871082358048` + "\n" +
		os.Args[0] + ` zoo pet-animal --body '{
      "animal": "Ut distinctio molestiae quis velit at qui.",
      "duration": 344612583194791941
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		calcFlags = flag.NewFlagSet("calc", flag.ContinueOnError)

		calcAddFlags = flag.NewFlagSet("add", flag.ExitOnError)
		calcAddAFlag = calcAddFlags.String("a", "REQUIRED", "Left operand")
		calcAddBFlag = calcAddFlags.String("b", "REQUIRED", "Right operand")

		zooFlags = flag.NewFlagSet("zoo", flag.ContinueOnError)

		zooPetAnimalFlags    = flag.NewFlagSet("pet-animal", flag.ExitOnError)
		zooPetAnimalBodyFlag = zooPetAnimalFlags.String("body", "REQUIRED", "")
	)
	calcFlags.Usage = calcUsage
	calcAddFlags.Usage = calcAddUsage

	zooFlags.Usage = zooUsage
	zooPetAnimalFlags.Usage = zooPetAnimalUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "calc":
			svcf = calcFlags
		case "zoo":
			svcf = zooFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "calc":
			switch epn {
			case "add":
				epf = calcAddFlags

			}

		case "zoo":
			switch epn {
			case "pet-animal":
				epf = zooPetAnimalFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "calc":
			c := calcc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = calcc.BuildAddPayload(*calcAddAFlag, *calcAddBFlag)
			}
		case "zoo":
			c := zooc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "pet-animal":
				endpoint = c.PetAnimal()
				data, err = zooc.BuildPetAnimalPayload(*zooPetAnimalBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// calcUsage displays the usage of the calc command and its subcommands.
func calcUsage() {
	fmt.Fprintf(os.Stderr, `Service is the calc service interface.
Usage:
    %s [globalflags] calc COMMAND [flags]

COMMAND:
    add: Add implements add.

Additional help:
    %s calc COMMAND --help
`, os.Args[0], os.Args[0])
}
func calcAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] calc add -a INT -b INT

Add implements add.
    -a INT: Left operand
    -b INT: Right operand

Example:
    `+os.Args[0]+` calc add --a 1485553568483242564 --b 1900791871082358048
`, os.Args[0])
}

// zooUsage displays the usage of the zoo command and its subcommands.
func zooUsage() {
	fmt.Fprintf(os.Stderr, `Service is the zoo service interface.
Usage:
    %s [globalflags] zoo COMMAND [flags]

COMMAND:
    pet-animal: PetAnimal implements petAnimal.

Additional help:
    %s zoo COMMAND --help
`, os.Args[0], os.Args[0])
}
func zooPetAnimalUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] zoo pet-animal -body JSON

PetAnimal implements petAnimal.
    -body JSON: 

Example:
    `+os.Args[0]+` zoo pet-animal --body '{
      "animal": "Ut distinctio molestiae quis velit at qui.",
      "duration": 344612583194791941
   }'
`, os.Args[0])
}
