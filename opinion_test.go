// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courtlistenersdk_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/battements-falaises/court-listener-sdk-go"
	"github.com/battements-falaises/court-listener-sdk-go/internal/testutil"
	"github.com/battements-falaises/court-listener-sdk-go/option"
)

func TestOpinionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Opinions.Get(
		context.TODO(),
		0,
		courtlistenersdk.OpinionGetParams{
			Fields: courtlistenersdk.String("fields"),
			Format: courtlistenersdk.OpinionGetParamsFormatJson,
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

func TestOpinionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Opinions.List(context.TODO(), courtlistenersdk.OpinionListParams{
		ID:                        courtlistenersdk.Int(0),
		CitedOpinion:              courtlistenersdk.Int(0),
		Cluster:                   courtlistenersdk.Int(0),
		ClusterDocketCourt:        courtlistenersdk.String("cluster__docket__court"),
		ClusterDocketDocketNumber: courtlistenersdk.String("cluster__docket__docket_number"),
		Count:                     courtlistenersdk.OpinionListParamsCountOn,
		Cursor:                    courtlistenersdk.String("cursor"),
		DateCreated:               courtlistenersdk.Time(time.Now()),
		DateCreatedGte:            courtlistenersdk.Time(time.Now()),
		DateCreatedLte:            courtlistenersdk.Time(time.Now()),
		DateModified:              courtlistenersdk.Time(time.Now()),
		DateModifiedGte:           courtlistenersdk.Time(time.Now()),
		DateModifiedLte:           courtlistenersdk.Time(time.Now()),
		Fields:                    courtlistenersdk.String("fields"),
		Format:                    courtlistenersdk.OpinionListParamsFormatJson,
		IDGt:                      courtlistenersdk.Int(0),
		IDGte:                     courtlistenersdk.Int(0),
		IDLt:                      courtlistenersdk.Int(0),
		IDLte:                     courtlistenersdk.Int(0),
		IDRange:                   courtlistenersdk.String("id__range"),
		Omit:                      courtlistenersdk.String("omit"),
		OrderBy:                   courtlistenersdk.String("order_by"),
		Page:                      courtlistenersdk.Int(1),
		Type:                      courtlistenersdk.String("type"),
	})
	if err != nil {
		var apierr *courtlistenersdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
