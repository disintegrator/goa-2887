package main

import (
	calccli "bugrepro/gen/http/cli/calc"
	zoocli "bugrepro/gen/http/cli/zoo"
	"fmt"
)

func main() {
	fmt.Println(calccli.UsageExamples())
	fmt.Println(zoocli.UsageExamples())
}
