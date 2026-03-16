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

func TestDocketGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Dockets.Get(
		context.TODO(),
		0,
		courtlistenersdk.DocketGetParams{
			Fields: courtlistenersdk.String("fields"),
			Format: courtlistenersdk.DocketGetParamsFormatJson,
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

func TestDocketListWithOptionalParams(t *testing.T) {
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
	_, err := client.Dockets.List(context.TODO(), courtlistenersdk.DocketListParams{
		ID:                      courtlistenersdk.Int(0),
		Blocked:                 courtlistenersdk.Bool(true),
		CaseName:                courtlistenersdk.String("case_name"),
		Cause:                   courtlistenersdk.String("cause"),
		Count:                   courtlistenersdk.DocketListParamsCountOn,
		Court:                   courtlistenersdk.String("court"),
		QueryCourtJurisdiction1: courtlistenersdk.String("court__jurisdiction"),
		QueryCourtJurisdiction2: courtlistenersdk.String("court__jurisdiction!"),
		Cursor:                  courtlistenersdk.String("cursor"),
		DateCreated:             courtlistenersdk.Time(time.Now()),
		DateCreatedGte:          courtlistenersdk.Time(time.Now()),
		DateCreatedLte:          courtlistenersdk.Time(time.Now()),
		DateFiled:               courtlistenersdk.Time(time.Now()),
		DateFiledGte:            courtlistenersdk.Time(time.Now()),
		DateFiledLte:            courtlistenersdk.Time(time.Now()),
		DateModified:            courtlistenersdk.Time(time.Now()),
		DateModifiedGte:         courtlistenersdk.Time(time.Now()),
		DateModifiedLte:         courtlistenersdk.Time(time.Now()),
		DateTerminated:          courtlistenersdk.Time(time.Now()),
		DateTerminatedGte:       courtlistenersdk.Time(time.Now()),
		DateTerminatedLte:       courtlistenersdk.Time(time.Now()),
		DocketNumber:            courtlistenersdk.String("docket_number"),
		Fields:                  courtlistenersdk.String("fields"),
		Format:                  courtlistenersdk.DocketListParamsFormatJson,
		IDGt:                    courtlistenersdk.Int(0),
		IDGte:                   courtlistenersdk.Int(0),
		IDLt:                    courtlistenersdk.Int(0),
		IDLte:                   courtlistenersdk.Int(0),
		IDRange:                 courtlistenersdk.String("id__range"),
		NatureOfSuit:            courtlistenersdk.String("nature_of_suit"),
		Omit:                    courtlistenersdk.String("omit"),
		OrderBy:                 courtlistenersdk.String("order_by"),
		Page:                    courtlistenersdk.Int(1),
		Source:                  courtlistenersdk.Int(0),
	})
	if err != nil {
		var apierr *courtlistenersdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
