package main

import (
	"fmt"
	"os"

	"github.com/jtreutel/costscript/ddcost"
)

// Checks whether env vars in the array are set.  Does not check whether value is valid.
func checkEnvVars(moduleName string, envArray []string) bool {

	notSet := 0
	runModule := false

	//Count how many env vars are not set
	for i := 0; i < len(envArray); i++ {
		if _, exists := os.LookupEnv(envArray[i]); !exists {
			notSet += 1
		}
	}

	if notSet == len(envArray) {
		// No env vars set --> Skip this module
		fmt.Println("No env vars set for", moduleName, ", skipping.")
	} else if notSet == 0 {
		// All env vars set --> Run this module
		fmt.Println("All env vars set for", moduleName, ", running.")
		runModule = true
	} else {
		// Only some env vars set --> Skip this module
		fmt.Println("Env vars for", moduleName, "only partially set, check your configuration!  Skipping.")
	}

	return runModule

}

func main() {

	var ddEnv = []string{"DD_APP_KEY", "DD_API_KEY", "DD_SITE"}

	//Check if DD env vars are configured properly; if so, run Datadog
	if configured := checkEnvVars("Datadog", ddEnv); configured {
		ddDate, ddCost := ddcost.Execute()
		ddDateStr := ddDate.Format("2006-01")
		fmt.Printf("Datadog Costs for %s: %g", ddDateStr, ddCost)
	}

}
