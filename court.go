// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courtlistenersdk

import (
	"context"
	"errors"
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

// Metadata about courts in the CourtListener database.
//
// CourtService contains methods and other services that help with interacting with
// the court-listener API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCourtService] method instead.
type CourtService struct {
	Options []option.RequestOption
}

// NewCourtService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCourtService(opts ...option.RequestOption) (r CourtService) {
	r = CourtService{}
	r.Options = opts
	return
}

// Retrieve a single court
func (r *CourtService) Get(ctx context.Context, id string, query CourtGetParams, opts ...option.RequestOption) (res *Court, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("courts/%s/", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of courts. Results can generally be cached as court
// data changes infrequently.
func (r *CourtService) List(ctx context.Context, query CourtListParams, opts ...option.RequestOption) (res *pagination.CursorURLPage[Court], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "courts/"
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

// Returns a paginated list of courts. Results can generally be cached as court
// data changes infrequently.
func (r *CourtService) ListAutoPaging(ctx context.Context, query CourtListParams, opts ...option.RequestOption) *pagination.CursorURLPageAutoPager[Court] {
	return pagination.NewCursorURLPageAutoPager(r.List(ctx, query, opts...))
}

type Court struct {
	// Court identifier (e.g. `scotus`, `ca9`, `dcd`).
	ID string `json:"id"`
	// String used for citations from this court.
	CitationString string    `json:"citation_string"`
	DateCreated    time.Time `json:"date_created" format:"date-time"`
	DateModified   time.Time `json:"date_modified" format:"date-time"`
	// Date the court was dissolved, if applicable.
	EndDate time.Time `json:"end_date" api:"nullable" format:"date"`
	// Full name of the court.
	FullName string `json:"full_name"`
	// Whether this court is currently active.
	InUse bool `json:"in_use"`
	// Jurisdiction type code.
	Jurisdiction string `json:"jurisdiction"`
	// Sort position for display.
	Position float64 `json:"position"`
	// Canonical API URL for this court.
	ResourceUri string `json:"resource_uri" format:"uri"`
	// Short/abbreviated name.
	ShortName string `json:"short_name"`
	// Date the court was founded.
	StartDate time.Time `json:"start_date" api:"nullable" format:"date"`
	// Court's website URL.
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CitationString respjson.Field
		DateCreated    respjson.Field
		DateModified   respjson.Field
		EndDate        respjson.Field
		FullName       respjson.Field
		InUse          respjson.Field
		Jurisdiction   respjson.Field
		Position       respjson.Field
		ResourceUri    respjson.Field
		ShortName      respjson.Field
		StartDate      respjson.Field
		URL            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Court) RawJSON() string { return r.JSON.raw }
func (r *Court) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CourtGetParams struct {
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
	Format CourtGetParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CourtGetParams]'s query parameters as `url.Values`.
func (r CourtGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Response serialization format. JSON is default when no `Accept` header is
// provided.
type CourtGetParamsFormat string

const (
	CourtGetParamsFormatJson CourtGetParamsFormat = "json"
	CourtGetParamsFormatXml  CourtGetParamsFormat = "xml"
	CourtGetParamsFormatHTML CourtGetParamsFormat = "html"
)

type CourtListParams struct {
	// Filter by court identifier (e.g. `scotus`, `ca9`, `dcd`).
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Cursor token for deep pagination. Returned in the `next` / `previous` fields of
	// paginated responses. Available when ordering by `id`, `date_modified`, or
	// `date_created`.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter by exact date modified (ISO-8601).
	DateModified param.Opt[time.Time] `query:"date_modified,omitzero" format:"date-time" json:"-"`
	// Filter courts modified on or after this date.
	DateModifiedGte param.Opt[time.Time] `query:"date_modified__gte,omitzero" format:"date-time" json:"-"`
	// Filter courts modified on or before this date.
	DateModifiedLte param.Opt[time.Time] `query:"date_modified__lte,omitzero" format:"date-time" json:"-"`
	// Comma-separated list of fields to include. Supports nested fields via
	// double-underscore notation (e.g. `educations__id`).
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Filter by the full name of the court.
	FullName param.Opt[string] `query:"full_name,omitzero" json:"-"`
	// Filter courts whose full name starts with the given value.
	FullNameStartswith param.Opt[string] `query:"full_name__startswith,omitzero" json:"-"`
	// Filter by multiple court identifiers (comma-separated).
	IDIn param.Opt[string] `query:"id__in,omitzero" json:"-"`
	// Filter by jurisdiction type. Common values: `F` (Federal Appellate), `FD`
	// (Federal District), `FB` (Federal Bankruptcy), `FBP` (Federal Bankruptcy Panel),
	// `FS` (Federal Special), `S` (State Supreme), `SA` (State Appellate), `ST` (State
	// Trial), `SS` (State Special), `SAG` (State Attorney General), `T` (Tribal), `I`
	// (International), `C` (Committee), `TES` (Testing).
	Jurisdiction param.Opt[string] `query:"jurisdiction,omitzero" json:"-"`
	// Comma-separated list of fields to exclude. Supports nested fields via
	// double-underscore notation.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Comma-separated list of fields to order by. Prefix with `-` for descending
	// order. Use a secondary field as a tie-breaker for deterministic ordering (e.g.
	// `date_filed,id`).
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	// Page number for standard pagination (limited to 100 pages).
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Set to `on` to return only the total count of matching items without result
	// data. When enabled, pagination parameters are ignored.
	//
	// Any of "on".
	Count CourtListParamsCount `query:"count,omitzero" json:"-"`
	// Response serialization format. JSON is default when no `Accept` header is
	// provided.
	//
	// Any of "json", "xml", "html".
	Format CourtListParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CourtListParams]'s query parameters as `url.Values`.
func (r CourtListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Set to `on` to return only the total count of matching items without result
// data. When enabled, pagination parameters are ignored.
type CourtListParamsCount string

const (
	CourtListParamsCountOn CourtListParamsCount = "on"
)

// Response serialization format. JSON is default when no `Accept` header is
// provided.
type CourtListParamsFormat string

const (
	CourtListParamsFormatJson CourtListParamsFormat = "json"
	CourtListParamsFormatXml  CourtListParamsFormat = "xml"
	CourtListParamsFormatHTML CourtListParamsFormat = "html"
)
