// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package policyanalyzer provides access to the Policy Analyzer API.
//
// For product documentation, see: https://www.google.com
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/policyanalyzer/v1beta1"
//   ...
//   ctx := context.Background()
//   policyanalyzerService, err := policyanalyzer.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   policyanalyzerService, err := policyanalyzer.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   policyanalyzerService, err := policyanalyzer.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package policyanalyzer // import "google.golang.org/api/policyanalyzer/v1beta1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint

const apiId = "policyanalyzer:v1beta1"
const apiName = "policyanalyzer"
const apiVersion = "v1beta1"
const basePath = "https://policyanalyzer.googleapis.com/"
const mtlsBasePath = "https://policyanalyzer.mtls.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// See, edit, configure, and delete your Google Cloud data and see the
	// email address for your Google Account.
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := option.WithScopes(
		"https://www.googleapis.com/auth/cloud-platform",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Locations = NewProjectsLocationsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Locations *ProjectsLocationsService
}

func NewProjectsLocationsService(s *Service) *ProjectsLocationsService {
	rs := &ProjectsLocationsService{s: s}
	rs.ActivityTypes = NewProjectsLocationsActivityTypesService(s)
	return rs
}

type ProjectsLocationsService struct {
	s *Service

	ActivityTypes *ProjectsLocationsActivityTypesService
}

func NewProjectsLocationsActivityTypesService(s *Service) *ProjectsLocationsActivityTypesService {
	rs := &ProjectsLocationsActivityTypesService{s: s}
	rs.Activities = NewProjectsLocationsActivityTypesActivitiesService(s)
	return rs
}

type ProjectsLocationsActivityTypesService struct {
	s *Service

	Activities *ProjectsLocationsActivityTypesActivitiesService
}

func NewProjectsLocationsActivityTypesActivitiesService(s *Service) *ProjectsLocationsActivityTypesActivitiesService {
	rs := &ProjectsLocationsActivityTypesActivitiesService{s: s}
	return rs
}

type ProjectsLocationsActivityTypesActivitiesService struct {
	s *Service
}

type GoogleCloudPolicyanalyzerV1beta1Activity struct {
	// Activity: A struct of custom fields to explain the activity.
	Activity googleapi.RawMessage `json:"activity,omitempty"`

	// ActivityType: The type of the activity.
	ActivityType string `json:"activityType,omitempty"`

	// FullResourceName: The full resource name that identifies the
	// resource. For examples of full resource names for Google Cloud
	// services, see
	// https://cloud.google.com/iam/help/troubleshooter/full-resource-names.
	FullResourceName string `json:"fullResourceName,omitempty"`

	// ObservationPeriod: The data observation period to build the activity.
	ObservationPeriod *GoogleCloudPolicyanalyzerV1beta1ObservationPeriod `json:"observationPeriod,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Activity") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Activity") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudPolicyanalyzerV1beta1Activity) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudPolicyanalyzerV1beta1Activity
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudPolicyanalyzerV1beta1ObservationPeriod: Represents data
// observation period.
type GoogleCloudPolicyanalyzerV1beta1ObservationPeriod struct {
	// EndTime: The observation end time.
	EndTime string `json:"endTime,omitempty"`

	// StartTime: The observation start time.
	StartTime string `json:"startTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndTime") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudPolicyanalyzerV1beta1ObservationPeriod) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudPolicyanalyzerV1beta1ObservationPeriod
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudPolicyanalyzerV1beta1QueryActivityResponse: Response to
// the `QueryActivity` method.
type GoogleCloudPolicyanalyzerV1beta1QueryActivityResponse struct {
	// Activities: The set of activities that match the filter included in
	// the request.
	Activities []*GoogleCloudPolicyanalyzerV1beta1Activity `json:"activities,omitempty"`

	// NextPageToken: If there might be more results than those appearing in
	// this response, then `nextPageToken` is included. To get the next set
	// of results, call this method again using the value of `nextPageToken`
	// as `pageToken`.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Activities") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Activities") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudPolicyanalyzerV1beta1QueryActivityResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudPolicyanalyzerV1beta1QueryActivityResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "policyanalyzer.projects.locations.activityTypes.activities.query":

type ProjectsLocationsActivityTypesActivitiesQueryCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Query: Queries policy activities on GCP resources.
//
// - parent: The container resource on which to execute the request.
//   Acceptable formats:
//   `projects/[PROJECT_ID|PROJECT_NUMBER]/locations/[LOCATION]/activityT
//   ypes/[ACTIVITY_TYPE]` LOCATION here refers to GCP Locations:
//   https://cloud.google.com/about/locations/.
func (r *ProjectsLocationsActivityTypesActivitiesService) Query(parent string) *ProjectsLocationsActivityTypesActivitiesQueryCall {
	c := &ProjectsLocationsActivityTypesActivitiesQueryCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Filter sets the optional parameter "filter": Optional filter
// expression to restrict the activities returned. Supported filters
// are: - service_account_last_authn.full_resource_name {=} -
// service_account_key_last_authn.full_resource_name {=}
func (c *ProjectsLocationsActivityTypesActivitiesQueryCall) Filter(filter string) *ProjectsLocationsActivityTypesActivitiesQueryCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return from this request. Max limit is 1000.
// Non-positive values are ignored. The presence of `nextPageToken` in
// the response indicates that more results might be available.
func (c *ProjectsLocationsActivityTypesActivitiesQueryCall) PageSize(pageSize int64) *ProjectsLocationsActivityTypesActivitiesQueryCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If present, then
// retrieve the next batch of results from the preceding call to this
// method. `pageToken` must be the value of `nextPageToken` from the
// previous response. The values of other method parameters should be
// identical to those in the previous call.
func (c *ProjectsLocationsActivityTypesActivitiesQueryCall) PageToken(pageToken string) *ProjectsLocationsActivityTypesActivitiesQueryCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsActivityTypesActivitiesQueryCall) Fields(s ...googleapi.Field) *ProjectsLocationsActivityTypesActivitiesQueryCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsActivityTypesActivitiesQueryCall) IfNoneMatch(entityTag string) *ProjectsLocationsActivityTypesActivitiesQueryCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsActivityTypesActivitiesQueryCall) Context(ctx context.Context) *ProjectsLocationsActivityTypesActivitiesQueryCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsActivityTypesActivitiesQueryCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsActivityTypesActivitiesQueryCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20210810")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/activities:query")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "policyanalyzer.projects.locations.activityTypes.activities.query" call.
// Exactly one of *GoogleCloudPolicyanalyzerV1beta1QueryActivityResponse
// or error will be non-nil. Any non-2xx status code is an error.
// Response headers are in either
// *GoogleCloudPolicyanalyzerV1beta1QueryActivityResponse.ServerResponse.
// Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsLocationsActivityTypesActivitiesQueryCall) Do(opts ...googleapi.CallOption) (*GoogleCloudPolicyanalyzerV1beta1QueryActivityResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleCloudPolicyanalyzerV1beta1QueryActivityResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Queries policy activities on GCP resources.",
	//   "flatPath": "v1beta1/projects/{projectsId}/locations/{locationsId}/activityTypes/{activityTypesId}/activities:query",
	//   "httpMethod": "GET",
	//   "id": "policyanalyzer.projects.locations.activityTypes.activities.query",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Optional filter expression to restrict the activities returned. Supported filters are: - service_account_last_authn.full_resource_name {=} - service_account_key_last_authn.full_resource_name {=} ",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The maximum number of results to return from this request. Max limit is 1000. Non-positive values are ignored. The presence of `nextPageToken` in the response indicates that more results might be available.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. If present, then retrieve the next batch of results from the preceding call to this method. `pageToken` must be the value of `nextPageToken` from the previous response. The values of other method parameters should be identical to those in the previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The container resource on which to execute the request. Acceptable formats: `projects/[PROJECT_ID|PROJECT_NUMBER]/locations/[LOCATION]/activityTypes/[ACTIVITY_TYPE]` LOCATION here refers to GCP Locations: https://cloud.google.com/about/locations/",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/activityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/activities:query",
	//   "response": {
	//     "$ref": "GoogleCloudPolicyanalyzerV1beta1QueryActivityResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsLocationsActivityTypesActivitiesQueryCall) Pages(ctx context.Context, f func(*GoogleCloudPolicyanalyzerV1beta1QueryActivityResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}