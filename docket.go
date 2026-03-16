// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courtlistenersdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/battements-falaises/court-listener-sdk-go/internal/apijson"
	"github.com/battements-falaises/court-listener-sdk-go/internal/apiquery"
	"github.com/battements-falaises/court-listener-sdk-go/internal/requestconfig"
	"github.com/battements-falaises/court-listener-sdk-go/option"
	"github.com/battements-falaises/court-listener-sdk-go/packages/pagination"
	"github.com/battements-falaises/court-listener-sdk-go/packages/param"
	"github.com/battements-falaises/court-listener-sdk-go/packages/respjson"
)

// Case-level metadata sitting at the top of the object hierarchy.
//
// DocketService contains methods and other services that help with interacting
// with the court-listener API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocketService] method instead.
type DocketService struct {
	Options []option.RequestOption
}

// NewDocketService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDocketService(opts ...option.RequestOption) (r DocketService) {
	r = DocketService{}
	r.Options = opts
	return
}

// Retrieve a single docket
func (r *DocketService) Get(ctx context.Context, id int64, query DocketGetParams, opts ...option.RequestOption) (res *Docket, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("dockets/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of dockets. Dockets sit at the top of the case law
// hierarchy, linking to clusters of opinions.
//
// **Note**: The response does not inline docket entries, parties, or attorneys
// (this does not scale). Use the PACER/RECAP APIs for those.
func (r *DocketService) List(ctx context.Context, query DocketListParams, opts ...option.RequestOption) (res *pagination.CursorURLPage[Docket], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "dockets/"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a paginated list of dockets. Dockets sit at the top of the case law
// hierarchy, linking to clusters of opinions.
//
// **Note**: The response does not inline docket entries, parties, or attorneys
// (this does not scale). Use the PACER/RECAP APIs for those.
func (r *DocketService) ListAutoPaging(ctx context.Context, query DocketListParams, opts ...option.RequestOption) *pagination.CursorURLPageAutoPager[Docket] {
	return pagination.NewCursorURLPageAutoPager(r.List(ctx, query, opts...))
}

type Docket struct {
	ID int64 `json:"id"`
	// Relative URL path on CourtListener.
	AbsoluteURL string `json:"absolute_url"`
	// API URL of the court this case was appealed from.
	AppealFrom                   string `json:"appeal_from" api:"nullable" format:"uri"`
	AppealFromStr                string `json:"appeal_from_str"`
	AppellateCaseTypeInformation string `json:"appellate_case_type_information"`
	AppellateFeeStatus           string `json:"appellate_fee_status"`
	// API URL of the assigned judge.
	AssignedTo string `json:"assigned_to" api:"nullable" format:"uri"`
	// Name of the assigned judge (string, not normalized).
	AssignedToStr string `json:"assigned_to_str"`
	// API URLs of related oral argument audio files.
	AudioFiles []string `json:"audio_files" format:"uri"`
	// Bankruptcy-specific information, if applicable.
	BankruptcyInformation map[string]any `json:"bankruptcy_information" api:"nullable"`
	Blocked               bool           `json:"blocked"`
	// Current case name. May change over time (e.g. if a named party changes). See
	// also the cluster's case_name which is frozen at decision time.
	CaseName      string `json:"case_name"`
	CaseNameFull  string `json:"case_name_full"`
	CaseNameShort string `json:"case_name_short"`
	Cause         string `json:"cause" api:"nullable"`
	// API URLs of related opinion clusters.
	Clusters []string `json:"clusters" format:"uri"`
	// API URL of the court.
	Court string `json:"court" format:"uri"`
	// Court identifier string.
	CourtID              string    `json:"court_id"`
	DateArgued           time.Time `json:"date_argued" api:"nullable" format:"date"`
	DateBlocked          time.Time `json:"date_blocked" api:"nullable" format:"date"`
	DateCertDenied       time.Time `json:"date_cert_denied" api:"nullable" format:"date"`
	DateCertGranted      time.Time `json:"date_cert_granted" api:"nullable" format:"date"`
	DateCreated          time.Time `json:"date_created" format:"date-time"`
	DateFiled            time.Time `json:"date_filed" api:"nullable" format:"date"`
	DateLastFiling       time.Time `json:"date_last_filing" api:"nullable" format:"date"`
	DateLastIndex        time.Time `json:"date_last_index" api:"nullable" format:"date-time"`
	DateModified         time.Time `json:"date_modified" format:"date-time"`
	DateReargued         time.Time `json:"date_reargued" api:"nullable" format:"date"`
	DateReargumentDenied time.Time `json:"date_reargument_denied" api:"nullable" format:"date"`
	DateTerminated       time.Time `json:"date_terminated" api:"nullable" format:"date"`
	// The docket number assigned by the court.
	DocketNumber string `json:"docket_number"`
	// Normalized core docket number.
	DocketNumberCore string `json:"docket_number_core"`
	// URL to the Internet Archive docket file.
	FilepathIa string `json:"filepath_ia" api:"nullable" format:"uri"`
	// URL to the Internet Archive JSON docket file.
	FilepathIaJson       string    `json:"filepath_ia_json" api:"nullable" format:"uri"`
	IaDateFirstChange    time.Time `json:"ia_date_first_change" api:"nullable" format:"date-time"`
	IaNeedsUpload        bool      `json:"ia_needs_upload" api:"nullable"`
	IaUploadFailureCount int64     `json:"ia_upload_failure_count" api:"nullable"`
	// Integrated database data, if available.
	IdbData          map[string]any `json:"idb_data" api:"nullable"`
	JurisdictionType string         `json:"jurisdiction_type" api:"nullable"`
	JuryDemand       string         `json:"jury_demand" api:"nullable"`
	MdlStatus        string         `json:"mdl_status"`
	NatureOfSuit     string         `json:"nature_of_suit" api:"nullable"`
	// Original court information, if available.
	OriginalCourtInfo map[string]any `json:"original_court_info" api:"nullable"`
	PacerCaseID       string         `json:"pacer_case_id" api:"nullable"`
	// API URLs of judges on the panel.
	Panel    []string `json:"panel" format:"uri"`
	PanelStr string   `json:"panel_str"`
	// API URL of the referred judge.
	ReferredTo    string `json:"referred_to" api:"nullable" format:"uri"`
	ReferredToStr string `json:"referred_to_str"`
	ResourceUri   string `json:"resource_uri" format:"uri"`
	Slug          string `json:"slug"`
	// Numeric source identifier.
	Source int64 `json:"source"`
	// API URLs of tags on this docket.
	Tags []string `json:"tags" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                           respjson.Field
		AbsoluteURL                  respjson.Field
		AppealFrom                   respjson.Field
		AppealFromStr                respjson.Field
		AppellateCaseTypeInformation respjson.Field
		AppellateFeeStatus           respjson.Field
		AssignedTo                   respjson.Field
		AssignedToStr                respjson.Field
		AudioFiles                   respjson.Field
		BankruptcyInformation        respjson.Field
		Blocked                      respjson.Field
		CaseName                     respjson.Field
		CaseNameFull                 respjson.Field
		CaseNameShort                respjson.Field
		Cause                        respjson.Field
		Clusters                     respjson.Field
		Court                        respjson.Field
		CourtID                      respjson.Field
		DateArgued                   respjson.Field
		DateBlocked                  respjson.Field
		DateCertDenied               respjson.Field
		DateCertGranted              respjson.Field
		DateCreated                  respjson.Field
		DateFiled                    respjson.Field
		DateLastFiling               respjson.Field
		DateLastIndex                respjson.Field
		DateModified                 respjson.Field
		DateReargued                 respjson.Field
		DateReargumentDenied         respjson.Field
		DateTerminated               respjson.Field
		DocketNumber                 respjson.Field
		DocketNumberCore             respjson.Field
		FilepathIa                   respjson.Field
		FilepathIaJson               respjson.Field
		IaDateFirstChange            respjson.Field
		IaNeedsUpload                respjson.Field
		IaUploadFailureCount         respjson.Field
		IdbData                      respjson.Field
		JurisdictionType             respjson.Field
		JuryDemand                   respjson.Field
		MdlStatus                    respjson.Field
		NatureOfSuit                 respjson.Field
		OriginalCourtInfo            respjson.Field
		PacerCaseID                  respjson.Field
		Panel                        respjson.Field
		PanelStr                     respjson.Field
		ReferredTo                   respjson.Field
		ReferredToStr                respjson.Field
		ResourceUri                  respjson.Field
		Slug                         respjson.Field
		Source                       respjson.Field
		Tags                         respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Docket) RawJSON() string { return r.JSON.raw }
func (r *Docket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocketGetParams struct {
	// Comma-separated list of fields to include. Supports nested fields via
	// double-underscore notation (e.g. `educations__id`).
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Comma-separated list of fields to exclude. Supports nested fields via
	// double-underscore notation.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Response serialization format. JSON is default when no `Accept` header is
	// provided.
	//
	// Any of "json", "xml", "html".
	Format DocketGetParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocketGetParams]'s query parameters as `url.Values`.
func (r DocketGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Response serialization format. JSON is default when no `Accept` header is
// provided.
type DocketGetParamsFormat string

const (
	DocketGetParamsFormatJson DocketGetParamsFormat = "json"
	DocketGetParamsFormatXml  DocketGetParamsFormat = "xml"
	DocketGetParamsFormatHTML DocketGetParamsFormat = "html"
)

type DocketListParams struct {
	// Filter by docket ID (exact).
	ID param.Opt[int64] `query:"id,omitzero" json:"-"`
	// Filter for blocked/unblocked dockets.
	Blocked param.Opt[bool] `query:"blocked,omitzero" json:"-"`
	// Filter by case name.
	CaseName param.Opt[string] `query:"case_name,omitzero" json:"-"`
	// Filter by cause.
	Cause param.Opt[string] `query:"cause,omitzero" json:"-"`
	// Filter by court identifier (e.g. `scotus`). Supports related court filters via
	// `court__` prefix.
	Court param.Opt[string] `query:"court,omitzero" json:"-"`
	// Filter by the court's jurisdiction type (e.g. `F`, `FD`, `S`).
	QueryCourtJurisdiction1 param.Opt[string] `query:"court__jurisdiction,omitzero" json:"-"`
	// Exclude dockets from this jurisdiction type.
	QueryCourtJurisdiction2 param.Opt[string] `query:"court__jurisdiction!,omitzero" json:"-"`
	// Cursor token for deep pagination. Returned in the `next` / `previous` fields of
	// paginated responses. Available when ordering by `id`, `date_modified`, or
	// `date_created`.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter by exact creation date.
	DateCreated param.Opt[time.Time] `query:"date_created,omitzero" format:"date-time" json:"-"`
	// Created on or after this date.
	DateCreatedGte param.Opt[time.Time] `query:"date_created__gte,omitzero" format:"date-time" json:"-"`
	// Created on or before this date.
	DateCreatedLte param.Opt[time.Time] `query:"date_created__lte,omitzero" format:"date-time" json:"-"`
	// Filter by filing date.
	DateFiled param.Opt[time.Time] `query:"date_filed,omitzero" format:"date" json:"-"`
	// Filed on or after this date.
	DateFiledGte param.Opt[time.Time] `query:"date_filed__gte,omitzero" format:"date" json:"-"`
	// Filed on or before this date.
	DateFiledLte param.Opt[time.Time] `query:"date_filed__lte,omitzero" format:"date" json:"-"`
	// Filter by exact modification date.
	DateModified param.Opt[time.Time] `query:"date_modified,omitzero" format:"date-time" json:"-"`
	// Modified on or after this date.
	DateModifiedGte param.Opt[time.Time] `query:"date_modified__gte,omitzero" format:"date-time" json:"-"`
	// Modified on or before this date.
	DateModifiedLte param.Opt[time.Time] `query:"date_modified__lte,omitzero" format:"date-time" json:"-"`
	// Filter by termination date.
	DateTerminated    param.Opt[time.Time] `query:"date_terminated,omitzero" format:"date" json:"-"`
	DateTerminatedGte param.Opt[time.Time] `query:"date_terminated__gte,omitzero" format:"date" json:"-"`
	DateTerminatedLte param.Opt[time.Time] `query:"date_terminated__lte,omitzero" format:"date" json:"-"`
	// Filter by exact docket number (e.g. `23A994`).
	DocketNumber param.Opt[string] `query:"docket_number,omitzero" json:"-"`
	// Comma-separated list of fields to include. Supports nested fields via
	// double-underscore notation (e.g. `educations__id`).
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Docket IDs greater than this value.
	IDGt param.Opt[int64] `query:"id__gt,omitzero" json:"-"`
	// Docket IDs greater than or equal to this value.
	IDGte param.Opt[int64] `query:"id__gte,omitzero" json:"-"`
	// Docket IDs less than this value.
	IDLt param.Opt[int64] `query:"id__lt,omitzero" json:"-"`
	// Docket IDs less than or equal to this value.
	IDLte param.Opt[int64] `query:"id__lte,omitzero" json:"-"`
	// Docket IDs within an inclusive range (e.g. `500,1000`).
	IDRange param.Opt[string] `query:"id__range,omitzero" json:"-"`
	// Filter by nature of suit.
	NatureOfSuit param.Opt[string] `query:"nature_of_suit,omitzero" json:"-"`
	// Comma-separated list of fields to exclude. Supports nested fields via
	// double-underscore notation.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Comma-separated list of fields to order by. Prefix with `-` for descending
	// order. Use a secondary field as a tie-breaker for deterministic ordering (e.g.
	// `date_filed,id`).
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	// Page number for standard pagination (limited to 100 pages).
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter by docket source.
	Source param.Opt[int64] `query:"source,omitzero" json:"-"`
	// Set to `on` to return only the total count of matching items without result
	// data. When enabled, pagination parameters are ignored.
	//
	// Any of "on".
	Count DocketListParamsCount `query:"count,omitzero" json:"-"`
	// Response serialization format. JSON is default when no `Accept` header is
	// provided.
	//
	// Any of "json", "xml", "html".
	Format DocketListParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocketListParams]'s query parameters as `url.Values`.
func (r DocketListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Set to `on` to return only the total count of matching items without result
// data. When enabled, pagination parameters are ignored.
type DocketListParamsCount string

const (
	DocketListParamsCountOn DocketListParamsCount = "on"
)

// Response serialization format. JSON is default when no `Accept` header is
// provided.
type DocketListParamsFormat string

const (
	DocketListParamsFormatJson DocketListParamsFormat = "json"
	DocketListParamsFormatXml  DocketListParamsFormat = "xml"
	DocketListParamsFormatHTML DocketListParamsFormat = "html"
)
