// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courtlistenersdk_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/court-listener-sdk-go"
	"github.com/stainless-sdks/court-listener-sdk-go/internal/testutil"
	"github.com/stainless-sdks/court-listener-sdk-go/option"
)

func TestClusterGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Clusters.Get(
		context.TODO(),
		0,
		courtlistenersdk.ClusterGetParams{
			Fields: courtlistenersdk.String("fields"),
			Format: courtlistenersdk.ClusterGetParamsFormatJson,
			Omit:   courtlistenersdk.String("omit"),
		},
	)
	if err != nil {
		var apierr *courtlistenersdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestClusterListWithOptionalParams(t *testing.T) {
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
	_, err := client.Clusters.List(context.TODO(), courtlistenersdk.ClusterListParams{
		ID:                 courtlistenersdk.Int(0),
		Citation:           courtlistenersdk.String("citation"),
		Count:              courtlistenersdk.ClusterListParamsCountOn,
		Cursor:             courtlistenersdk.String("cursor"),
		DateCreated:        courtlistenersdk.Time(time.Now()),
		DateCreatedGte:     courtlistenersdk.Time(time.Now()),
		DateCreatedLte:     courtlistenersdk.Time(time.Now()),
		DateFiled:          courtlistenersdk.Time(time.Now()),
		DateFiledGte:       courtlistenersdk.Time(time.Now()),
		DateFiledLte:       courtlistenersdk.Time(time.Now()),
		DateModified:       courtlistenersdk.Time(time.Now()),
		DateModifiedGte:    courtlistenersdk.Time(time.Now()),
		DateModifiedLte:    courtlistenersdk.Time(time.Now()),
		Docket:             courtlistenersdk.Int(0),
		DocketCourt:        courtlistenersdk.String("docket__court"),
		DocketDocketNumber: courtlistenersdk.String("docket__docket_number"),
		Fields:             courtlistenersdk.String("fields"),
		Format:             courtlistenersdk.ClusterListParamsFormatJson,
		IDGt:               courtlistenersdk.Int(0),
		IDGte:              courtlistenersdk.Int(0),
		IDLt:               courtlistenersdk.Int(0),
		IDLte:              courtlistenersdk.Int(0),
		IDRange:            courtlistenersdk.String("id__range"),
		Judges:             courtlistenersdk.String("judges"),
		Omit:               courtlistenersdk.String("omit"),
		OrderBy:            courtlistenersdk.String("order_by"),
		Page:               courtlistenersdk.Int(1),
	})
	if err != nil {
		var apierr *courtlistenersdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
