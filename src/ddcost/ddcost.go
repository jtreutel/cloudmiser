package ddcost

//DD_SITE="us5.datadoghq.com" DD_API_KEY="<DD_API_KEY>" DD_APP_KEY="<DD_APP_KEY>"
//https://docs.datadoghq.com/api/latest/usage-metering/?code-lang=go#get-estimated-cost-across-your-account
// Get estimated cost across your account returns "OK" response

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func GetEstimatedDatadogCosts() (time.Time, float64) {

	//Set up Datadog client
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewUsageMeteringApi(apiClient)

	//Create empty request param struct
	reqParams := datadogV2.NewGetEstimatedCostByOrgOptionalParameters()

	// Set time params
	currentTime := time.Now()
	year, month, _ := currentTime.Date()
	startMonth := time.Date(year, (month + 1), 1, 0, 0, 0, 0, time.Local)
	reqParams.StartMonth = &startMonth

	resp, r, err := api.GetEstimatedCostByOrg(ctx, *reqParams)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetEstimatedCostByOrg`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	//A slice containing data for all DD orgs
	costByOrgs := resp.GetData()
	//We only have one org for now, so just grab the first (and only) one
	costByOrg := costByOrgs[0]

	//Get data.attributes
	costByOrgAttributes := costByOrg.GetAttributes()

	//Get data.attributes.total_cost/date
	totalCost := *costByOrgAttributes.TotalCost
	date := *costByOrgAttributes.Date

	return date, totalCost
}

func GetHistoricalDatadogCosts() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewUsageMeteringApi(apiClient)
	//
	resp, r, err := api.GetHistoricalCostByOrg(ctx, time.Now().AddDate(0, -2, 0), *datadogV2.NewGetHistoricalCostByOrgOptionalParameters().WithView("summary"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetHistoricalCostByOrg`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.GetHistoricalCostByOrg`:\n%s\n", responseContent)
}

func Execute() (time.Time, float64) {

	//Get DD costs for current month
	date, totalCost := GetEstimatedDatadogCosts()

	return date, totalCost

}
