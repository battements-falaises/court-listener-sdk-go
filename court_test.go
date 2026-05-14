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

func TestCourtGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Courts.Get(
		context.TODO(),
		"id",
		courtlistenersdk.CourtGetParams{
			Fields: courtlistenersdk.String("fields"),
			Format: courtlistenersdk.CourtGetParamsFormatJson,
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

func TestCourtListWithOptionalParams(t *testing.T) {
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
	_, err := client.Courts.List(context.TODO(), courtlistenersdk.CourtListParams{
		ID:                 courtlistenersdk.String("id"),
		Count:              courtlistenersdk.CourtListParamsCountOn,
		Cursor:             courtlistenersdk.String("cursor"),
		DateModified:       courtlistenersdk.Time(time.Now()),
		DateModifiedGte:    courtlistenersdk.Time(time.Now()),
		DateModifiedLte:    courtlistenersdk.Time(time.Now()),
		Fields:             courtlistenersdk.String("fields"),
		Format:             courtlistenersdk.CourtListParamsFormatJson,
		FullName:           courtlistenersdk.String("full_name"),
		FullNameStartswith: courtlistenersdk.String("full_name__startswith"),
		IDIn:               courtlistenersdk.String("id__in"),
		Jurisdiction:       courtlistenersdk.String("jurisdiction"),
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
