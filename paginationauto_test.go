// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courtlistenersdk_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/court-listener-sdk-go"
	"github.com/stainless-sdks/court-listener-sdk-go/internal/testutil"
	"github.com/stainless-sdks/court-listener-sdk-go/option"
)

func TestAutoPagination(t *testing.T) {
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
	iter := client.Courts.ListAutoPaging(context.TODO(), courtlistenersdk.CourtListParams{})
	// The mock server isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		court := iter.Current()
		t.Logf("%+v\n", court.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
