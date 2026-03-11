// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courtlistenersdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/court-listener-sdk-go/internal/apijson"
	"github.com/stainless-sdks/court-listener-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/court-listener-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/court-listener-sdk-go/option"
	"github.com/stainless-sdks/court-listener-sdk-go/packages/pagination"
	"github.com/stainless-sdks/court-listener-sdk-go/packages/param"
	"github.com/stainless-sdks/court-listener-sdk-go/packages/respjson"
)

// Opinion clusters grouping related decisions from a single hearing.
//
// ClusterService contains methods and other services that help with interacting
// with the court-listener API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterService] method instead.
type ClusterService struct {
	Options []option.RequestOption
}

// NewClusterService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewClusterService(opts ...option.RequestOption) (r ClusterService) {
	r = ClusterService{}
	r.Options = opts
	return
}

// Look up a cluster by its ID. The cluster ID matches the ID used in CourtListener
// case law URLs (e.g. `/opinion/2812209/obergefell-v-hodges/` corresponds to
// cluster ID `2812209`).
func (r *ClusterService) Get(ctx context.Context, id int64, query ClusterGetParams, opts ...option.RequestOption) (res *Cluster, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("clusters/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of opinion clusters. Each cluster groups together
// opinions from the same panel hearing (e.g. majority, dissent, concurrence). The
// cluster `id` is used in CourtListener case law URLs.
func (r *ClusterService) List(ctx context.Context, query ClusterListParams, opts ...option.RequestOption) (res *pagination.CursorURLPage[Cluster], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "clusters/"
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

// Returns a paginated list of opinion clusters. Each cluster groups together
// opinions from the same panel hearing (e.g. majority, dissent, concurrence). The
// cluster `id` is used in CourtListener case law URLs.
func (r *ClusterService) ListAutoPaging(ctx context.Context, query ClusterListParams, opts ...option.RequestOption) *pagination.CursorURLPageAutoPager[Cluster] {
	return pagination.NewCursorURLPageAutoPager(r.List(ctx, query, opts...))
}

type Cluster struct {
	// Cluster ID — used in CourtListener case law URLs.
	ID int64 `json:"id"`
	// Relative URL path on CourtListener.
	AbsoluteURL string `json:"absolute_url"`
	Blocked     bool   `json:"blocked"`
	// Case name frozen at decision time (does not change).
	CaseName      string `json:"case_name"`
	CaseNameFull  string `json:"case_name_full"`
	CaseNameShort string `json:"case_name_short"`
	// Number of times this cluster has been cited.
	CitationCount int64 `json:"citation_count"`
	// List of parallel citation objects for this cluster.
	Citations              []ClusterCitation `json:"citations"`
	Correction             string            `json:"correction" api:"nullable"`
	CrossReference         string            `json:"cross_reference" api:"nullable"`
	DateBlocked            time.Time         `json:"date_blocked" api:"nullable" format:"date"`
	DateCreated            time.Time         `json:"date_created" format:"date-time"`
	DateFiled              time.Time         `json:"date_filed" api:"nullable" format:"date"`
	DateFiledIsApproximate bool              `json:"date_filed_is_approximate"`
	DateModified           time.Time         `json:"date_modified" format:"date-time"`
	Disposition            string            `json:"disposition" api:"nullable"`
	// API URL of the parent docket.
	Docket    string `json:"docket" format:"uri"`
	Headnotes string `json:"headnotes" api:"nullable"`
	History   string `json:"history" api:"nullable"`
	// Judge name(s) as a string (not yet normalized).
	Judges string `json:"judges" api:"nullable"`
	// API URLs of non-participating judges.
	NonParticipatingJudges []string `json:"non_participating_judges" format:"uri"`
	OtherDates             string   `json:"other_dates" api:"nullable"`
	// API URLs of judges on the panel (normalized).
	Panel []string `json:"panel" format:"uri"`
	// Whether this cluster is published, unpublished, etc.
	PrecedentialStatus string `json:"precedential_status"`
	ResourceUri        string `json:"resource_uri" format:"uri"`
	Slug               string `json:"slug"`
	// Source of this cluster data.
	Source string `json:"source"`
	// API URLs of opinions in this cluster.
	SubOpinions []string `json:"sub_opinions" format:"uri"`
	Summary     string   `json:"summary" api:"nullable"`
	Syllabus    string   `json:"syllabus" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AbsoluteURL            respjson.Field
		Blocked                respjson.Field
		CaseName               respjson.Field
		CaseNameFull           respjson.Field
		CaseNameShort          respjson.Field
		CitationCount          respjson.Field
		Citations              respjson.Field
		Correction             respjson.Field
		CrossReference         respjson.Field
		DateBlocked            respjson.Field
		DateCreated            respjson.Field
		DateFiled              respjson.Field
		DateFiledIsApproximate respjson.Field
		DateModified           respjson.Field
		Disposition            respjson.Field
		Docket                 respjson.Field
		Headnotes              respjson.Field
		History                respjson.Field
		Judges                 respjson.Field
		NonParticipatingJudges respjson.Field
		OtherDates             respjson.Field
		Panel                  respjson.Field
		PrecedentialStatus     respjson.Field
		ResourceUri            respjson.Field
		Slug                   respjson.Field
		Source                 respjson.Field
		SubOpinions            respjson.Field
		Summary                respjson.Field
		Syllabus               respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Cluster) RawJSON() string { return r.JSON.raw }
func (r *Cluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A parallel citation for an opinion cluster.
type ClusterCitation struct {
	// The starting page number in the reporter.
	Page string `json:"page" api:"nullable"`
	// The reporter abbreviation (e.g. "U.S.", "S. Ct.").
	Reporter string `json:"reporter"`
	// The citation type identifier.
	Type int64 `json:"type"`
	// The volume number of the reporter.
	Volume int64 `json:"volume" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		Reporter    respjson.Field
		Type        respjson.Field
		Volume      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterCitation) RawJSON() string { return r.JSON.raw }
func (r *ClusterCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterGetParams struct {
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
	Format ClusterGetParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ClusterGetParams]'s query parameters as `url.Values`.
func (r ClusterGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Response serialization format. JSON is default when no `Accept` header is
// provided.
type ClusterGetParamsFormat string

const (
	ClusterGetParamsFormatJson ClusterGetParamsFormat = "json"
	ClusterGetParamsFormatXml  ClusterGetParamsFormat = "xml"
	ClusterGetParamsFormatHTML ClusterGetParamsFormat = "html"
)

type ClusterListParams struct {
	// Filter by cluster ID.
	ID param.Opt[int64] `query:"id,omitzero" json:"-"`
	// Filter by citation.
	Citation param.Opt[string] `query:"citation,omitzero" json:"-"`
	// Cursor token for deep pagination. Returned in the `next` / `previous` fields of
	// paginated responses. Available when ordering by `id`, `date_modified`, or
	// `date_created`.
	Cursor         param.Opt[string]    `query:"cursor,omitzero" json:"-"`
	DateCreated    param.Opt[time.Time] `query:"date_created,omitzero" format:"date-time" json:"-"`
	DateCreatedGte param.Opt[time.Time] `query:"date_created__gte,omitzero" format:"date-time" json:"-"`
	DateCreatedLte param.Opt[time.Time] `query:"date_created__lte,omitzero" format:"date-time" json:"-"`
	// Filter by the date the cluster was filed.
	DateFiled       param.Opt[time.Time] `query:"date_filed,omitzero" format:"date" json:"-"`
	DateFiledGte    param.Opt[time.Time] `query:"date_filed__gte,omitzero" format:"date" json:"-"`
	DateFiledLte    param.Opt[time.Time] `query:"date_filed__lte,omitzero" format:"date" json:"-"`
	DateModified    param.Opt[time.Time] `query:"date_modified,omitzero" format:"date-time" json:"-"`
	DateModifiedGte param.Opt[time.Time] `query:"date_modified__gte,omitzero" format:"date-time" json:"-"`
	DateModifiedLte param.Opt[time.Time] `query:"date_modified__lte,omitzero" format:"date-time" json:"-"`
	// Filter by parent docket ID.
	Docket param.Opt[int64] `query:"docket,omitzero" json:"-"`
	// Filter by the court of the parent docket (e.g. `scotus`).
	DocketCourt param.Opt[string] `query:"docket__court,omitzero" json:"-"`
	// Filter by the docket number of the parent docket.
	DocketDocketNumber param.Opt[string] `query:"docket__docket_number,omitzero" json:"-"`
	// Comma-separated list of fields to include. Supports nested fields via
	// double-underscore notation (e.g. `educations__id`).
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	IDGt   param.Opt[int64]  `query:"id__gt,omitzero" json:"-"`
	IDGte  param.Opt[int64]  `query:"id__gte,omitzero" json:"-"`
	IDLt   param.Opt[int64]  `query:"id__lt,omitzero" json:"-"`
	IDLte  param.Opt[int64]  `query:"id__lte,omitzero" json:"-"`
	// Inclusive range (e.g. `100,500`).
	IDRange param.Opt[string] `query:"id__range,omitzero" json:"-"`
	// Filter by judge name string.
	Judges param.Opt[string] `query:"judges,omitzero" json:"-"`
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
	Count ClusterListParamsCount `query:"count,omitzero" json:"-"`
	// Response serialization format. JSON is default when no `Accept` header is
	// provided.
	//
	// Any of "json", "xml", "html".
	Format ClusterListParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ClusterListParams]'s query parameters as `url.Values`.
func (r ClusterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Set to `on` to return only the total count of matching items without result
// data. When enabled, pagination parameters are ignored.
type ClusterListParamsCount string

const (
	ClusterListParamsCountOn ClusterListParamsCount = "on"
)

// Response serialization format. JSON is default when no `Accept` header is
// provided.
type ClusterListParamsFormat string

const (
	ClusterListParamsFormatJson ClusterListParamsFormat = "json"
	ClusterListParamsFormatXml  ClusterListParamsFormat = "xml"
	ClusterListParamsFormatHTML ClusterListParamsFormat = "html"
)
