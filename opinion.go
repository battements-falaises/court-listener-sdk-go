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

// Individual judicial opinions with full text and metadata.
//
// OpinionService contains methods and other services that help with interacting
// with the court-listener API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOpinionService] method instead.
type OpinionService struct {
	Options []option.RequestOption
}

// NewOpinionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOpinionService(opts ...option.RequestOption) (r OpinionService) {
	r = OpinionService{}
	r.Options = opts
	return
}

// Look up an opinion by its ID. Note that opinion IDs do **not** reliably match
// cluster IDs. If you have a CourtListener case URL, use the cluster API to look
// it up.
func (r *OpinionService) Get(ctx context.Context, id int64, query OpinionGetParams, opts ...option.RequestOption) (res *Opinion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("opinions/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a paginated list of opinions. Each opinion contains the text of a
// judicial decision and metadata about the authoring judge.
//
// **Tip**: Prefer the `html_with_citations` field for opinion text — it contains
// the raw text with identified and linked citations, and is the field used on the
// CourtListener website.
//
// Use `fields` / `omit` parameters to exclude large text fields you don't need.
func (r *OpinionService) List(ctx context.Context, query OpinionListParams, opts ...option.RequestOption) (res *pagination.CursorURLPage[Opinion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "opinions/"
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

// Returns a paginated list of opinions. Each opinion contains the text of a
// judicial decision and metadata about the authoring judge.
//
// **Tip**: Prefer the `html_with_citations` field for opinion text — it contains
// the raw text with identified and linked citations, and is the field used on the
// CourtListener website.
//
// Use `fields` / `omit` parameters to exclude large text fields you don't need.
func (r *OpinionService) ListAutoPaging(ctx context.Context, query OpinionListParams, opts ...option.RequestOption) *pagination.CursorURLPageAutoPager[Opinion] {
	return pagination.NewCursorURLPageAutoPager(r.List(ctx, query, opts...))
}

type Opinion struct {
	ID int64 `json:"id"`
	// API URL of the authoring judge (normalized).
	Author string `json:"author" api:"nullable" format:"uri"`
	// Author name as a string (not normalized).
	AuthorStr string `json:"author_str" api:"nullable"`
	// API URL of the parent cluster.
	Cluster      string    `json:"cluster" format:"uri"`
	DateCreated  time.Time `json:"date_created" format:"date-time"`
	DateModified time.Time `json:"date_modified" format:"date-time"`
	// Original URL where the opinion was scraped from. Often unreliable as many courts
	// do not maintain stable URIs.
	DownloadURL string `json:"download_url" api:"nullable" format:"uri"`
	// Whether the text was extracted via OCR.
	ExtractedByOcr bool `json:"extracted_by_ocr"`
	// HTML content from court websites (Word Perfect or HTML documents) or
	// Resource.org.
	HTML string `json:"html" api:"nullable"`
	// HTML content from the anonymous 2020 source.
	HTMLAnon2020 string `json:"html_anon_2020" api:"nullable"`
	// HTML content from the Columbia collaboration.
	HTMLColumbia string `json:"html_columbia" api:"nullable"`
	// HTML content from the Lawbox donation.
	HTMLLawbox string `json:"html_lawbox" api:"nullable"`
	// **Recommended field.** HTML with citations identified and linked. This is the
	// field used on the CourtListener website.
	HTMLWithCitations string `json:"html_with_citations" api:"nullable"`
	// API URLs of judges who joined this opinion.
	JoinedBy []string `json:"joined_by" format:"uri"`
	// Path to the binary file for the decision, if available.
	LocalPath string `json:"local_path" api:"nullable"`
	// API URLs of other opinions cited by this one.
	OpinionsCited []string `json:"opinions_cited" format:"uri"`
	// Sort order within the cluster. Only populated for opinions ingested from Harvard
	// or Columbia sources.
	OrderingKey float64 `json:"ordering_key" api:"nullable"`
	// Whether this is a per curiam opinion.
	PerCuriam bool `json:"per_curiam"`
	// Plain text of the opinion. Populated when sourced from a court website as PDF or
	// Microsoft Word document.
	PlainText   string `json:"plain_text" api:"nullable"`
	ResourceUri string `json:"resource_uri" format:"uri"`
	// SHA-1 hash of the opinion content.
	Sha1 string `json:"sha1" api:"nullable"`
	// Opinion type (e.g. combined opinion, lead opinion, concurrence, dissent). Values
	// are number-prefixed for sort priority.
	Type string `json:"type"`
	// XML content from Harvard's Caselaw Access Project. Contains rich data but may
	// have OCR artifacts.
	XmlHarvard string `json:"xml_harvard" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Author            respjson.Field
		AuthorStr         respjson.Field
		Cluster           respjson.Field
		DateCreated       respjson.Field
		DateModified      respjson.Field
		DownloadURL       respjson.Field
		ExtractedByOcr    respjson.Field
		HTML              respjson.Field
		HTMLAnon2020      respjson.Field
		HTMLColumbia      respjson.Field
		HTMLLawbox        respjson.Field
		HTMLWithCitations respjson.Field
		JoinedBy          respjson.Field
		LocalPath         respjson.Field
		OpinionsCited     respjson.Field
		OrderingKey       respjson.Field
		PerCuriam         respjson.Field
		PlainText         respjson.Field
		ResourceUri       respjson.Field
		Sha1              respjson.Field
		Type              respjson.Field
		XmlHarvard        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Opinion) RawJSON() string { return r.JSON.raw }
func (r *Opinion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpinionGetParams struct {
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
	Format OpinionGetParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OpinionGetParams]'s query parameters as `url.Values`.
func (r OpinionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Response serialization format. JSON is default when no `Accept` header is
// provided.
type OpinionGetParamsFormat string

const (
	OpinionGetParamsFormatJson OpinionGetParamsFormat = "json"
	OpinionGetParamsFormatXml  OpinionGetParamsFormat = "xml"
	OpinionGetParamsFormatHTML OpinionGetParamsFormat = "html"
)

type OpinionListParams struct {
	// Filter by opinion ID.
	ID param.Opt[int64] `query:"id,omitzero" json:"-"`
	// Filter opinions that cite this opinion ID.
	CitedOpinion param.Opt[int64] `query:"cited_opinion,omitzero" json:"-"`
	// Filter by parent cluster ID.
	Cluster param.Opt[int64] `query:"cluster,omitzero" json:"-"`
	// Filter by court via the cluster's docket (e.g. `scotus`).
	ClusterDocketCourt param.Opt[string] `query:"cluster__docket__court,omitzero" json:"-"`
	// Filter by docket number via the cluster's docket.
	ClusterDocketDocketNumber param.Opt[string] `query:"cluster__docket__docket_number,omitzero" json:"-"`
	// Cursor token for deep pagination. Returned in the `next` / `previous` fields of
	// paginated responses. Available when ordering by `id`, `date_modified`, or
	// `date_created`.
	Cursor          param.Opt[string]    `query:"cursor,omitzero" json:"-"`
	DateCreated     param.Opt[time.Time] `query:"date_created,omitzero" format:"date-time" json:"-"`
	DateCreatedGte  param.Opt[time.Time] `query:"date_created__gte,omitzero" format:"date-time" json:"-"`
	DateCreatedLte  param.Opt[time.Time] `query:"date_created__lte,omitzero" format:"date-time" json:"-"`
	DateModified    param.Opt[time.Time] `query:"date_modified,omitzero" format:"date-time" json:"-"`
	DateModifiedGte param.Opt[time.Time] `query:"date_modified__gte,omitzero" format:"date-time" json:"-"`
	DateModifiedLte param.Opt[time.Time] `query:"date_modified__lte,omitzero" format:"date-time" json:"-"`
	// Comma-separated list of fields to include. Supports nested fields via
	// double-underscore notation (e.g. `educations__id`).
	Fields  param.Opt[string] `query:"fields,omitzero" json:"-"`
	IDGt    param.Opt[int64]  `query:"id__gt,omitzero" json:"-"`
	IDGte   param.Opt[int64]  `query:"id__gte,omitzero" json:"-"`
	IDLt    param.Opt[int64]  `query:"id__lt,omitzero" json:"-"`
	IDLte   param.Opt[int64]  `query:"id__lte,omitzero" json:"-"`
	IDRange param.Opt[string] `query:"id__range,omitzero" json:"-"`
	// Comma-separated list of fields to exclude. Supports nested fields via
	// double-underscore notation.
	Omit param.Opt[string] `query:"omit,omitzero" json:"-"`
	// Comma-separated list of fields to order by. Prefix with `-` for descending
	// order. Use a secondary field as a tie-breaker for deterministic ordering (e.g.
	// `date_filed,id`).
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	// Page number for standard pagination (limited to 100 pages).
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter by opinion type. Values are prefixed with numbers for sort order. Common
	// types include combined opinion, lead opinion, concurrence, dissent, etc.
	Type param.Opt[string] `query:"type,omitzero" json:"-"`
	// Set to `on` to return only the total count of matching items without result
	// data. When enabled, pagination parameters are ignored.
	//
	// Any of "on".
	Count OpinionListParamsCount `query:"count,omitzero" json:"-"`
	// Response serialization format. JSON is default when no `Accept` header is
	// provided.
	//
	// Any of "json", "xml", "html".
	Format OpinionListParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OpinionListParams]'s query parameters as `url.Values`.
func (r OpinionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Set to `on` to return only the total count of matching items without result
// data. When enabled, pagination parameters are ignored.
type OpinionListParamsCount string

const (
	OpinionListParamsCountOn OpinionListParamsCount = "on"
)

// Response serialization format. JSON is default when no `Accept` header is
// provided.
type OpinionListParamsFormat string

const (
	OpinionListParamsFormatJson OpinionListParamsFormat = "json"
	OpinionListParamsFormatXml  OpinionListParamsFormat = "xml"
	OpinionListParamsFormatHTML OpinionListParamsFormat = "html"
)
