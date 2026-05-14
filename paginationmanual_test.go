// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courtlistenersdk_test

import (
	"context"
	"os"
	"testing"

	"github.com/battements-falaises/court-listener-sdk-go"
	"github.com/battements-falaises/court-listener-sdk-go/internal/testutil"
	"github.com/battements-falaises/court-listener-sdk-go/option"
)

func TestManualPagination(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := courtlistenersdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	page, err := client.Courts.List(context.TODO(), courtlistenersdk.CourtListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, court := range page.Results {
		t.Logf("%+v\n", court.ID)
	}
	// The mock server isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, court := range page.Results {
			t.Logf("%+v\n", court.ID)
		}
	}
}
