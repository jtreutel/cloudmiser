package main

import (
	"fmt"

	"github.com/jtreutel/costscript/ddcost"
)

func main() {

	ddDate, ddCost := ddcost.Execute()

	ddDateStr := ddDate.Format("2006-01")

	fmt.Printf("Datadog Costs for %s: %g", ddDateStr, ddCost)
}
