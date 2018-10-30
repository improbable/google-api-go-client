// Package websecurityscanner provides access to the Web Security Scanner API.
//
// See https://cloud.google.com/security-scanner/
//
// Usage example:
//
//   import "google.golang.org/api/websecurityscanner/v1alpha"
//   ...
//   websecurityscannerService, err := websecurityscanner.New(oauthHttpClient)
package websecurityscanner // import "google.golang.org/api/websecurityscanner/v1alpha"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
var _ = ctxhttp.Do

const apiId = "websecurityscanner:v1alpha"
const apiName = "websecurityscanner"
const apiVersion = "v1alpha"
const basePath = "https://websecurityscanner.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

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
	rs.ScanConfigs = NewProjectsScanConfigsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	ScanConfigs *ProjectsScanConfigsService
}

func NewProjectsScanConfigsService(s *Service) *ProjectsScanConfigsService {
	rs := &ProjectsScanConfigsService{s: s}
	rs.ScanRuns = NewProjectsScanConfigsScanRunsService(s)
	return rs
}

type ProjectsScanConfigsService struct {
	s *Service

	ScanRuns *ProjectsScanConfigsScanRunsService
}

func NewProjectsScanConfigsScanRunsService(s *Service) *ProjectsScanConfigsScanRunsService {
	rs := &ProjectsScanConfigsScanRunsService{s: s}
	rs.CrawledUrls = NewProjectsScanConfigsScanRunsCrawledUrlsService(s)
	rs.FindingTypeStats = NewProjectsScanConfigsScanRunsFindingTypeStatsService(s)
	rs.Findings = NewProjectsScanConfigsScanRunsFindingsService(s)
	return rs
}

type ProjectsScanConfigsScanRunsService struct {
	s *Service

	CrawledUrls *ProjectsScanConfigsScanRunsCrawledUrlsService

	FindingTypeStats *ProjectsScanConfigsScanRunsFindingTypeStatsService

	Findings *ProjectsScanConfigsScanRunsFindingsService
}

func NewProjectsScanConfigsScanRunsCrawledUrlsService(s *Service) *ProjectsScanConfigsScanRunsCrawledUrlsService {
	rs := &ProjectsScanConfigsScanRunsCrawledUrlsService{s: s}
	return rs
}

type ProjectsScanConfigsScanRunsCrawledUrlsService struct {
	s *Service
}

func NewProjectsScanConfigsScanRunsFindingTypeStatsService(s *Service) *ProjectsScanConfigsScanRunsFindingTypeStatsService {
	rs := &ProjectsScanConfigsScanRunsFindingTypeStatsService{s: s}
	return rs
}

type ProjectsScanConfigsScanRunsFindingTypeStatsService struct {
	s *Service
}

func NewProjectsScanConfigsScanRunsFindingsService(s *Service) *ProjectsScanConfigsScanRunsFindingsService {
	rs := &ProjectsScanConfigsScanRunsFindingsService{s: s}
	return rs
}

type ProjectsScanConfigsScanRunsFindingsService struct {
	s *Service
}

// Authentication: Scan authentication configuration.
type Authentication struct {
	// CustomAccount: Authentication using a custom account.
	CustomAccount *CustomAccount `json:"customAccount,omitempty"`

	// GoogleAccount: Authentication using a Google account.
	GoogleAccount *GoogleAccount `json:"googleAccount,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CustomAccount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CustomAccount") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Authentication) MarshalJSON() ([]byte, error) {
	type NoMethod Authentication
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CrawledUrl: A CrawledUrl resource represents a URL that was crawled
// during a ScanRun. Web
// Security Scanner Service crawls the web applications, following all
// links
// within the scope of sites, to find the URLs to test against.
type CrawledUrl struct {
	// Body: Output only.
	// The body of the request that was used to visit the URL.
	Body string `json:"body,omitempty"`

	// HttpMethod: Output only.
	// The http method of the request that was used to visit the URL,
	// in
	// uppercase.
	HttpMethod string `json:"httpMethod,omitempty"`

	// Url: Output only.
	// The URL that was crawled.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Body") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Body") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CrawledUrl) MarshalJSON() ([]byte, error) {
	type NoMethod CrawledUrl
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CustomAccount: Describes authentication configuration that uses a
// custom account.
type CustomAccount struct {
	// LoginUrl: Required.
	// The login form URL of the website.
	LoginUrl string `json:"loginUrl,omitempty"`

	// Password: Input only.
	// Required.
	// The password of the custom account. The credential is stored
	// encrypted
	// and not returned in any response nor included in audit logs.
	Password string `json:"password,omitempty"`

	// Username: Required.
	// The user name of the custom account.
	Username string `json:"username,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LoginUrl") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LoginUrl") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CustomAccount) MarshalJSON() ([]byte, error) {
	type NoMethod CustomAccount
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// Finding: A Finding resource represents a vulnerability instance
// identified during a
// ScanRun.
type Finding struct {
	// Body: Output only.
	// The body of the request that triggered the vulnerability.
	Body string `json:"body,omitempty"`

	// Description: Output only.
	// The description of the vulnerability.
	Description string `json:"description,omitempty"`

	// FinalUrl: Output only.
	// The URL where the browser lands when the vulnerability is detected.
	FinalUrl string `json:"finalUrl,omitempty"`

	// FindingType: Output only.
	// The type of the Finding.
	//
	// Possible values:
	//   "FINDING_TYPE_UNSPECIFIED" - The invalid finding type.
	//   "MIXED_CONTENT" - A page that was served over HTTPS also resources
	// over HTTP. A
	// man-in-the-middle attacker could tamper with the HTTP resource and
	// gain
	// full access to the website that loads the resource or to monitor
	// the
	// actions taken by the user.
	//   "OUTDATED_LIBRARY" - The version of an included library is known to
	// contain a security issue.
	// The scanner checks the version of library in use against a known list
	// of
	// vulnerable libraries. False positives are possible if the
	// version
	// detection fails or if the library has been manually patched.
	//   "ROSETTA_FLASH" - This type of vulnerability occurs when the value
	// of a request parameter
	// is reflected at the beginning of the response, for example, in
	// requests
	// using JSONP. Under certain circumstances, an attacker may be able
	// to
	// supply an alphanumeric-only Flash file in the vulnerable
	// parameter
	// causing the browser to execute the Flash file as if it originated on
	// the
	// vulnerable server.
	//   "XSS_CALLBACK" - A cross-site scripting (XSS) bug is found via
	// JavaScript callback. For
	// detailed explanations on XSS,
	// see
	// https://www.google.com/about/appsecurity/learning/xss/.
	//   "XSS_ERROR" - A potential cross-site scripting (XSS) bug due to
	// JavaScript breakage.
	// In some circumstances, the application under test might modify the
	// test
	// string before it is parsed by the browser. When the browser attempts
	// to
	// runs this modified test string, it will likely break and throw
	// a
	// JavaScript execution error, thus an injection issue is
	// occurring.
	// However, it may not be exploitable. Manual verification is needed to
	// see
	// if the test string modifications can be evaded and confirm that the
	// issue
	// is in fact an XSS vulnerability. For detailed explanations on XSS,
	// see
	// https://www.google.com/about/appsecurity/learning/xss/.
	//   "CLEAR_TEXT_PASSWORD" - An application appears to be transmitting a
	// password field in clear text.
	// An attacker can eavesdrop network traffic and sniff the password
	// field.
	//   "INVALID_CONTENT_TYPE" - An application returns sensitive content
	// with an invalid content type,
	// or without an 'X-Content-Type-Options: nosniff' header.
	//   "XSS_ANGULAR_CALLBACK" - A cross-site scripting (XSS) vulnerability
	// in AngularJS module that
	// occurs when a user-provided string is interpolated by Angular.
	//   "INVALID_HEADER" - A malformed or invalid valued header.
	//   "MISSPELLED_SECURITY_HEADER_NAME" - Misspelled security header
	// name.
	//   "MISMATCHING_SECURITY_HEADER_VALUES" - Mismatching values in a
	// duplicate security header.
	FindingType string `json:"findingType,omitempty"`

	// FrameUrl: Output only.
	// If the vulnerability was originated from nested IFrame, the
	// immediate
	// parent IFrame is reported.
	FrameUrl string `json:"frameUrl,omitempty"`

	// FuzzedUrl: Output only.
	// The URL produced by the server-side fuzzer and used in the request
	// that
	// triggered the vulnerability.
	FuzzedUrl string `json:"fuzzedUrl,omitempty"`

	// HttpMethod: Output only.
	// The http method of the request that triggered the vulnerability,
	// in
	// uppercase.
	HttpMethod string `json:"httpMethod,omitempty"`

	// Name: Output only.
	// The resource name of the Finding. The name follows the format
	// of
	// 'projects/{projectId}/scanConfigs/{scanConfigId}/scanruns/{scanRunI
	// d}/findings/{findingId}'.
	// The finding IDs are generated by the system.
	Name string `json:"name,omitempty"`

	// OutdatedLibrary: Output only.
	// An addon containing information about outdated libraries.
	OutdatedLibrary *OutdatedLibrary `json:"outdatedLibrary,omitempty"`

	// ReproductionUrl: Output only.
	// The URL containing human-readable payload that user can leverage
	// to
	// reproduce the vulnerability.
	ReproductionUrl string `json:"reproductionUrl,omitempty"`

	// TrackingId: Output only.
	// The tracking ID uniquely identifies a vulnerability instance
	// across
	// multiple ScanRuns.
	TrackingId string `json:"trackingId,omitempty"`

	// ViolatingResource: Output only.
	// An addon containing detailed information regarding any resource
	// causing the
	// vulnerability such as JavaScript sources, image, audio files, etc.
	ViolatingResource *ViolatingResource `json:"violatingResource,omitempty"`

	// VulnerableHeaders: Output only.
	// An addon containing information about vulnerable or missing HTTP
	// headers.
	VulnerableHeaders *VulnerableHeaders `json:"vulnerableHeaders,omitempty"`

	// VulnerableParameters: Output only.
	// An addon containing information about request parameters which were
	// found
	// to be vulnerable.
	VulnerableParameters *VulnerableParameters `json:"vulnerableParameters,omitempty"`

	// Xss: Output only.
	// An addon containing information reported for an XSS, if any.
	Xss *Xss `json:"xss,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Body") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Body") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Finding) MarshalJSON() ([]byte, error) {
	type NoMethod Finding
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FindingTypeStats: A FindingTypeStats resource represents stats
// regarding a specific FindingType
// of Findings under a given ScanRun.
type FindingTypeStats struct {
	// FindingCount: Output only.
	// The count of findings belonging to this finding type.
	FindingCount int64 `json:"findingCount,omitempty"`

	// FindingType: Output only.
	// The finding type associated with the stats.
	//
	// Possible values:
	//   "FINDING_TYPE_UNSPECIFIED" - The invalid finding type.
	//   "MIXED_CONTENT" - A page that was served over HTTPS also resources
	// over HTTP. A
	// man-in-the-middle attacker could tamper with the HTTP resource and
	// gain
	// full access to the website that loads the resource or to monitor
	// the
	// actions taken by the user.
	//   "OUTDATED_LIBRARY" - The version of an included library is known to
	// contain a security issue.
	// The scanner checks the version of library in use against a known list
	// of
	// vulnerable libraries. False positives are possible if the
	// version
	// detection fails or if the library has been manually patched.
	//   "ROSETTA_FLASH" - This type of vulnerability occurs when the value
	// of a request parameter
	// is reflected at the beginning of the response, for example, in
	// requests
	// using JSONP. Under certain circumstances, an attacker may be able
	// to
	// supply an alphanumeric-only Flash file in the vulnerable
	// parameter
	// causing the browser to execute the Flash file as if it originated on
	// the
	// vulnerable server.
	//   "XSS_CALLBACK" - A cross-site scripting (XSS) bug is found via
	// JavaScript callback. For
	// detailed explanations on XSS,
	// see
	// https://www.google.com/about/appsecurity/learning/xss/.
	//   "XSS_ERROR" - A potential cross-site scripting (XSS) bug due to
	// JavaScript breakage.
	// In some circumstances, the application under test might modify the
	// test
	// string before it is parsed by the browser. When the browser attempts
	// to
	// runs this modified test string, it will likely break and throw
	// a
	// JavaScript execution error, thus an injection issue is
	// occurring.
	// However, it may not be exploitable. Manual verification is needed to
	// see
	// if the test string modifications can be evaded and confirm that the
	// issue
	// is in fact an XSS vulnerability. For detailed explanations on XSS,
	// see
	// https://www.google.com/about/appsecurity/learning/xss/.
	//   "CLEAR_TEXT_PASSWORD" - An application appears to be transmitting a
	// password field in clear text.
	// An attacker can eavesdrop network traffic and sniff the password
	// field.
	//   "INVALID_CONTENT_TYPE" - An application returns sensitive content
	// with an invalid content type,
	// or without an 'X-Content-Type-Options: nosniff' header.
	//   "XSS_ANGULAR_CALLBACK" - A cross-site scripting (XSS) vulnerability
	// in AngularJS module that
	// occurs when a user-provided string is interpolated by Angular.
	//   "INVALID_HEADER" - A malformed or invalid valued header.
	//   "MISSPELLED_SECURITY_HEADER_NAME" - Misspelled security header
	// name.
	//   "MISMATCHING_SECURITY_HEADER_VALUES" - Mismatching values in a
	// duplicate security header.
	FindingType string `json:"findingType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FindingCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FindingCount") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FindingTypeStats) MarshalJSON() ([]byte, error) {
	type NoMethod FindingTypeStats
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleAccount: Describes authentication configuration that uses a
// Google account.
type GoogleAccount struct {
	// Password: Input only.
	// Required.
	// The password of the Google account. The credential is stored
	// encrypted
	// and not returned in any response nor included in audit logs.
	Password string `json:"password,omitempty"`

	// Username: Required.
	// The user name of the Google account.
	Username string `json:"username,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Password") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Password") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleAccount) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleAccount
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Header: Describes a HTTP Header.
type Header struct {
	// Name: Header name.
	Name string `json:"name,omitempty"`

	// Value: Header value.
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Header) MarshalJSON() ([]byte, error) {
	type NoMethod Header
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListCrawledUrlsResponse: Response for the `ListCrawledUrls` method.
type ListCrawledUrlsResponse struct {
	// CrawledUrls: The list of CrawledUrls returned.
	CrawledUrls []*CrawledUrl `json:"crawledUrls,omitempty"`

	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CrawledUrls") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CrawledUrls") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListCrawledUrlsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListCrawledUrlsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListFindingTypeStatsResponse: Response for the `ListFindingTypeStats`
// method.
type ListFindingTypeStatsResponse struct {
	// FindingTypeStats: The list of FindingTypeStats returned.
	FindingTypeStats []*FindingTypeStats `json:"findingTypeStats,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "FindingTypeStats") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FindingTypeStats") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListFindingTypeStatsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListFindingTypeStatsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListFindingsResponse: Response for the `ListFindings` method.
type ListFindingsResponse struct {
	// Findings: The list of Findings returned.
	Findings []*Finding `json:"findings,omitempty"`

	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Findings") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Findings") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListFindingsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListFindingsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListScanConfigsResponse: Response for the `ListScanConfigs` method.
type ListScanConfigsResponse struct {
	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ScanConfigs: The list of ScanConfigs returned.
	ScanConfigs []*ScanConfig `json:"scanConfigs,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListScanConfigsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListScanConfigsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListScanRunsResponse: Response for the `ListScanRuns` method.
type ListScanRunsResponse struct {
	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ScanRuns: The list of ScanRuns returned.
	ScanRuns []*ScanRun `json:"scanRuns,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListScanRunsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListScanRunsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OutdatedLibrary: Information reported for an outdated library.
type OutdatedLibrary struct {
	// LearnMoreUrls: URLs to learn more information about the
	// vulnerabilities in the library.
	LearnMoreUrls []string `json:"learnMoreUrls,omitempty"`

	// LibraryName: The name of the outdated library.
	LibraryName string `json:"libraryName,omitempty"`

	// Version: The version number.
	Version string `json:"version,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LearnMoreUrls") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LearnMoreUrls") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OutdatedLibrary) MarshalJSON() ([]byte, error) {
	type NoMethod OutdatedLibrary
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ScanConfig: A ScanConfig resource contains the configurations to
// launch a scan.
type ScanConfig struct {
	// Authentication: The authentication configuration. If specified,
	// service will use the
	// authentication configuration during scanning.
	Authentication *Authentication `json:"authentication,omitempty"`

	// BlacklistPatterns: The blacklist URL patterns as described
	// in
	// https://cloud.google.com/security-scanner/docs/excluded-urls
	BlacklistPatterns []string `json:"blacklistPatterns,omitempty"`

	// DisplayName: Required.
	// The user provided display name of the ScanConfig.
	DisplayName string `json:"displayName,omitempty"`

	// MaxQps: The maximum QPS during scanning. A valid value ranges from 5
	// to 20
	// inclusively. If the field is unspecified or its value is set 0,
	// server will
	// default to 15. Other values outside of [5, 20] range will be rejected
	// with
	// INVALID_ARGUMENT error.
	MaxQps int64 `json:"maxQps,omitempty"`

	// Name: The resource name of the ScanConfig. The name follows the
	// format of
	// 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs
	// are
	// generated by the system.
	Name string `json:"name,omitempty"`

	// Schedule: The schedule of the ScanConfig.
	Schedule *Schedule `json:"schedule,omitempty"`

	// StartingUrls: Required.
	// The starting URLs from which the scanner finds site pages.
	StartingUrls []string `json:"startingUrls,omitempty"`

	// TargetPlatforms: Set of Cloud Platforms targeted by the scan. If
	// empty, APP_ENGINE will be
	// used as a default.
	//
	// Possible values:
	//   "TARGET_PLATFORM_UNSPECIFIED" - The target platform is unknown.
	// Requests with this enum value will be
	// rejected with INVALID_ARGUMENT error.
	//   "APP_ENGINE" - Google App Engine service.
	//   "COMPUTE" - Google Compute Engine service.
	TargetPlatforms []string `json:"targetPlatforms,omitempty"`

	// UserAgent: The user agent used during scanning.
	//
	// Possible values:
	//   "USER_AGENT_UNSPECIFIED" - The user agent is unknown. Service will
	// default to CHROME_LINUX.
	//   "CHROME_LINUX" - Chrome on Linux. This is the service default if
	// unspecified.
	//   "CHROME_ANDROID" - Chrome on Android.
	//   "SAFARI_IPHONE" - Safari on IPhone.
	UserAgent string `json:"userAgent,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Authentication") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Authentication") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ScanConfig) MarshalJSON() ([]byte, error) {
	type NoMethod ScanConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ScanRun: A ScanRun is a output-only resource representing an actual
// run of the scan.
type ScanRun struct {
	// EndTime: Output only.
	// The time at which the ScanRun reached termination state - that the
	// ScanRun
	// is either finished or stopped by user.
	EndTime string `json:"endTime,omitempty"`

	// ExecutionState: Output only.
	// The execution state of the ScanRun.
	//
	// Possible values:
	//   "EXECUTION_STATE_UNSPECIFIED" - Represents an invalid state caused
	// by internal server error. This value
	// should never be returned.
	//   "QUEUED" - The scan is waiting in the queue.
	//   "SCANNING" - The scan is in progress.
	//   "FINISHED" - The scan is either finished or stopped by user.
	ExecutionState string `json:"executionState,omitempty"`

	// HasVulnerabilities: Output only.
	// Whether the scan run has found any vulnerabilities.
	HasVulnerabilities bool `json:"hasVulnerabilities,omitempty"`

	// Name: Output only.
	// The resource name of the ScanRun. The name follows the format
	// of
	// 'projects/{projectId}/scanConfigs/{scanConfigId}/scanRuns/{scanRunI
	// d}'.
	// The ScanRun IDs are generated by the system.
	Name string `json:"name,omitempty"`

	// ProgressPercent: Output only.
	// The percentage of total completion ranging from 0 to 100.
	// If the scan is in queue, the value is 0.
	// If the scan is running, the value ranges from 0 to 100.
	// If the scan is finished, the value is 100.
	ProgressPercent int64 `json:"progressPercent,omitempty"`

	// ResultState: Output only.
	// The result state of the ScanRun. This field is only available after
	// the
	// execution state reaches "FINISHED".
	//
	// Possible values:
	//   "RESULT_STATE_UNSPECIFIED" - Default value. This value is returned
	// when the ScanRun is not yet
	// finished.
	//   "SUCCESS" - The scan finished without errors.
	//   "ERROR" - The scan finished with errors.
	//   "KILLED" - The scan was terminated by user.
	ResultState string `json:"resultState,omitempty"`

	// StartTime: Output only.
	// The time at which the ScanRun started.
	StartTime string `json:"startTime,omitempty"`

	// UrlsCrawledCount: Output only.
	// The number of URLs crawled during this ScanRun. If the scan is in
	// progress,
	// the value represents the number of URLs crawled up to now.
	UrlsCrawledCount int64 `json:"urlsCrawledCount,omitempty,string"`

	// UrlsTestedCount: Output only.
	// The number of URLs tested during this ScanRun. If the scan is in
	// progress,
	// the value represents the number of URLs tested up to now. The number
	// of
	// URLs tested is usually larger than the number URLS crawled
	// because
	// typically a crawled URL is tested with multiple test payloads.
	UrlsTestedCount int64 `json:"urlsTestedCount,omitempty,string"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "EndTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ScanRun) MarshalJSON() ([]byte, error) {
	type NoMethod ScanRun
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Schedule: Scan schedule configuration.
type Schedule struct {
	// IntervalDurationDays: Required.
	// The duration of time between executions in days.
	IntervalDurationDays int64 `json:"intervalDurationDays,omitempty"`

	// ScheduleTime: A timestamp indicates when the next run will be
	// scheduled. The value is
	// refreshed by the server after each run. If unspecified, it will
	// default
	// to current server time, which means the scan will be scheduled to
	// start
	// immediately.
	ScheduleTime string `json:"scheduleTime,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "IntervalDurationDays") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IntervalDurationDays") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Schedule) MarshalJSON() ([]byte, error) {
	type NoMethod Schedule
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StartScanRunRequest: Request for the `StartScanRun` method.
type StartScanRunRequest struct {
}

// StopScanRunRequest: Request for the `StopScanRun` method.
type StopScanRunRequest struct {
}

// ViolatingResource: Information regarding any resource causing the
// vulnerability such
// as JavaScript sources, image, audio files, etc.
type ViolatingResource struct {
	// ContentType: The MIME type of this resource.
	ContentType string `json:"contentType,omitempty"`

	// ResourceUrl: URL of this violating resource.
	ResourceUrl string `json:"resourceUrl,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ContentType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ContentType") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ViolatingResource) MarshalJSON() ([]byte, error) {
	type NoMethod ViolatingResource
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// VulnerableHeaders: Information about vulnerable or missing HTTP
// Headers.
type VulnerableHeaders struct {
	// Headers: List of vulnerable headers.
	Headers []*Header `json:"headers,omitempty"`

	// MissingHeaders: List of missing headers.
	MissingHeaders []*Header `json:"missingHeaders,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Headers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Headers") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *VulnerableHeaders) MarshalJSON() ([]byte, error) {
	type NoMethod VulnerableHeaders
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// VulnerableParameters: Information about vulnerable request
// parameters.
type VulnerableParameters struct {
	// ParameterNames: The vulnerable parameter names.
	ParameterNames []string `json:"parameterNames,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ParameterNames") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ParameterNames") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *VulnerableParameters) MarshalJSON() ([]byte, error) {
	type NoMethod VulnerableParameters
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Xss: Information reported for an XSS.
type Xss struct {
	// ErrorMessage: An error message generated by a javascript breakage.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// StackTraces: Stack traces leading to the point where the XSS
	// occurred.
	StackTraces []string `json:"stackTraces,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ErrorMessage") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ErrorMessage") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Xss) MarshalJSON() ([]byte, error) {
	type NoMethod Xss
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "websecurityscanner.projects.scanConfigs.create":

type ProjectsScanConfigsCreateCall struct {
	s          *Service
	parent     string
	scanconfig *ScanConfig
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a new ScanConfig.
func (r *ProjectsScanConfigsService) Create(parent string, scanconfig *ScanConfig) *ProjectsScanConfigsCreateCall {
	c := &ProjectsScanConfigsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.scanconfig = scanconfig
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsScanConfigsCreateCall) Fields(s ...googleapi.Field) *ProjectsScanConfigsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsScanConfigsCreateCall) Context(ctx context.Context) *ProjectsScanConfigsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsScanConfigsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsScanConfigsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.scanconfig)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+parent}/scanConfigs")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "websecurityscanner.projects.scanConfigs.create"), c.s.client, req)
}

// Do executes the "websecurityscanner.projects.scanConfigs.create" call.
// Exactly one of *ScanConfig or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ScanConfig.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsScanConfigsCreateCall) Do(opts ...googleapi.CallOption) (*ScanConfig, error) {
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
	ret := &ScanConfig{
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
	//   "description": "Creates a new ScanConfig.",
	//   "flatPath": "v1alpha/projects/{projectsId}/scanConfigs",
	//   "httpMethod": "POST",
	//   "id": "websecurityscanner.projects.scanConfigs.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required.\nThe parent resource name where the scan is created, which should be a\nproject resource name in the format 'projects/{projectId}'.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha/{+parent}/scanConfigs",
	//   "request": {
	//     "$ref": "ScanConfig"
	//   },
	//   "response": {
	//     "$ref": "ScanConfig"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "websecurityscanner.projects.scanConfigs.delete":

type ProjectsScanConfigsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes an existing ScanConfig and its child resources.
func (r *ProjectsScanConfigsService) Delete(name string) *ProjectsScanConfigsDeleteCall {
	c := &ProjectsScanConfigsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsScanConfigsDeleteCall) Fields(s ...googleapi.Field) *ProjectsScanConfigsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsScanConfigsDeleteCall) Context(ctx context.Context) *ProjectsScanConfigsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsScanConfigsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsScanConfigsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "websecurityscanner.projects.scanConfigs.delete"), c.s.client, req)
}

// Do executes the "websecurityscanner.projects.scanConfigs.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsScanConfigsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	ret := &Empty{
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
	//   "description": "Deletes an existing ScanConfig and its child resources.",
	//   "flatPath": "v1alpha/projects/{projectsId}/scanConfigs/{scanConfigsId}",
	//   "httpMethod": "DELETE",
	//   "id": "websecurityscanner.projects.scanConfigs.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\nThe resource name of the ScanConfig to be deleted. The name follows the\nformat of 'projects/{projectId}/scanConfigs/{scanConfigId}'.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/scanConfigs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "websecurityscanner.projects.scanConfigs.get":

type ProjectsScanConfigsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a ScanConfig.
func (r *ProjectsScanConfigsService) Get(name string) *ProjectsScanConfigsGetCall {
	c := &ProjectsScanConfigsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsScanConfigsGetCall) Fields(s ...googleapi.Field) *ProjectsScanConfigsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsScanConfigsGetCall) IfNoneMatch(entityTag string) *ProjectsScanConfigsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsScanConfigsGetCall) Context(ctx context.Context) *ProjectsScanConfigsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsScanConfigsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsScanConfigsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "websecurityscanner.projects.scanConfigs.get"), c.s.client, req)
}

// Do executes the "websecurityscanner.projects.scanConfigs.get" call.
// Exactly one of *ScanConfig or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ScanConfig.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsScanConfigsGetCall) Do(opts ...googleapi.CallOption) (*ScanConfig, error) {
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
	ret := &ScanConfig{
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
	//   "description": "Gets a ScanConfig.",
	//   "flatPath": "v1alpha/projects/{projectsId}/scanConfigs/{scanConfigsId}",
	//   "httpMethod": "GET",
	//   "id": "websecurityscanner.projects.scanConfigs.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\nThe resource name of the ScanConfig to be returned. The name follows the\nformat of 'projects/{projectId}/scanConfigs/{scanConfigId}'.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/scanConfigs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha/{+name}",
	//   "response": {
	//     "$ref": "ScanConfig"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "websecurityscanner.projects.scanConfigs.list":

type ProjectsScanConfigsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists ScanConfigs under a given project.
func (r *ProjectsScanConfigsService) List(parent string) *ProjectsScanConfigsListCall {
	c := &ProjectsScanConfigsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of ScanConfigs to return, can be limited by server.
// If not specified or not positive, the implementation will select
// a
// reasonable value.
func (c *ProjectsScanConfigsListCall) PageSize(pageSize int64) *ProjectsScanConfigsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results to be returned. This should be
// a
// `next_page_token` value returned from a previous List request.
// If unspecified, the first page of results is returned.
func (c *ProjectsScanConfigsListCall) PageToken(pageToken string) *ProjectsScanConfigsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsScanConfigsListCall) Fields(s ...googleapi.Field) *ProjectsScanConfigsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsScanConfigsListCall) IfNoneMatch(entityTag string) *ProjectsScanConfigsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsScanConfigsListCall) Context(ctx context.Context) *ProjectsScanConfigsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsScanConfigsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsScanConfigsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+parent}/scanConfigs")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "websecurityscanner.projects.scanConfigs.list"), c.s.client, req)
}

// Do executes the "websecurityscanner.projects.scanConfigs.list" call.
// Exactly one of *ListScanConfigsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListScanConfigsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsScanConfigsListCall) Do(opts ...googleapi.CallOption) (*ListScanConfigsResponse, error) {
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
	ret := &ListScanConfigsResponse{
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
	//   "description": "Lists ScanConfigs under a given project.",
	//   "flatPath": "v1alpha/projects/{projectsId}/scanConfigs",
	//   "httpMethod": "GET",
	//   "id": "websecurityscanner.projects.scanConfigs.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of ScanConfigs to return, can be limited by server.\nIf not specified or not positive, the implementation will select a\nreasonable value.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results to be returned. This should be a\n`next_page_token` value returned from a previous List request.\nIf unspecified, the first page of results is returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required.\nThe parent resource name, which should be a project resource name in the\nformat 'projects/{projectId}'.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha/{+parent}/scanConfigs",
	//   "response": {
	//     "$ref": "ListScanConfigsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsScanConfigsListCall) Pages(ctx context.Context, f func(*ListScanConfigsResponse) error) error {
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

// method id "websecurityscanner.projects.scanConfigs.patch":

type ProjectsScanConfigsPatchCall struct {
	s          *Service
	name       string
	scanconfig *ScanConfig
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Updates a ScanConfig. This method support partial update of a
// ScanConfig.
func (r *ProjectsScanConfigsService) Patch(name string, scanconfig *ScanConfig) *ProjectsScanConfigsPatchCall {
	c := &ProjectsScanConfigsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.scanconfig = scanconfig
	return c
}

// UpdateMask sets the optional parameter "updateMask": Required.
// The update mask applies to the resource. For the `FieldMask`
// definition,
// see
// https://developers.google.com/protocol-buffers/docs/re
// ference/google.protobuf#fieldmask
func (c *ProjectsScanConfigsPatchCall) UpdateMask(updateMask string) *ProjectsScanConfigsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsScanConfigsPatchCall) Fields(s ...googleapi.Field) *ProjectsScanConfigsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsScanConfigsPatchCall) Context(ctx context.Context) *ProjectsScanConfigsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsScanConfigsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsScanConfigsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.scanconfig)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "websecurityscanner.projects.scanConfigs.patch"), c.s.client, req)
}

// Do executes the "websecurityscanner.projects.scanConfigs.patch" call.
// Exactly one of *ScanConfig or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ScanConfig.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsScanConfigsPatchCall) Do(opts ...googleapi.CallOption) (*ScanConfig, error) {
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
	ret := &ScanConfig{
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
	//   "description": "Updates a ScanConfig. This method support partial update of a ScanConfig.",
	//   "flatPath": "v1alpha/projects/{projectsId}/scanConfigs/{scanConfigsId}",
	//   "httpMethod": "PATCH",
	//   "id": "websecurityscanner.projects.scanConfigs.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The resource name of the ScanConfig. The name follows the format of\n'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are\ngenerated by the system.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/scanConfigs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Required.\nThe update mask applies to the resource. For the `FieldMask` definition,\nsee\nhttps://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha/{+name}",
	//   "request": {
	//     "$ref": "ScanConfig"
	//   },
	//   "response": {
	//     "$ref": "ScanConfig"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "websecurityscanner.projects.scanConfigs.start":

type ProjectsScanConfigsStartCall struct {
	s                   *Service
	name                string
	startscanrunrequest *StartScanRunRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// Start: Start a ScanRun according to the given ScanConfig.
func (r *ProjectsScanConfigsService) Start(name string, startscanrunrequest *StartScanRunRequest) *ProjectsScanConfigsStartCall {
	c := &ProjectsScanConfigsStartCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.startscanrunrequest = startscanrunrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsScanConfigsStartCall) Fields(s ...googleapi.Field) *ProjectsScanConfigsStartCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsScanConfigsStartCall) Context(ctx context.Context) *ProjectsScanConfigsStartCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsScanConfigsStartCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsScanConfigsStartCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.startscanrunrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+name}:start")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "websecurityscanner.projects.scanConfigs.start"), c.s.client, req)
}

// Do executes the "websecurityscanner.projects.scanConfigs.start" call.
// Exactly one of *ScanRun or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *ScanRun.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsScanConfigsStartCall) Do(opts ...googleapi.CallOption) (*ScanRun, error) {
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
	ret := &ScanRun{
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
	//   "description": "Start a ScanRun according to the given ScanConfig.",
	//   "flatPath": "v1alpha/projects/{projectsId}/scanConfigs/{scanConfigsId}:start",
	//   "httpMethod": "POST",
	//   "id": "websecurityscanner.projects.scanConfigs.start",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\nThe resource name of the ScanConfig to be used. The name follows the\nformat of 'projects/{projectId}/scanConfigs/{scanConfigId}'.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/scanConfigs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha/{+name}:start",
	//   "request": {
	//     "$ref": "StartScanRunRequest"
	//   },
	//   "response": {
	//     "$ref": "ScanRun"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "websecurityscanner.projects.scanConfigs.scanRuns.get":

type ProjectsScanConfigsScanRunsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a ScanRun.
func (r *ProjectsScanConfigsScanRunsService) Get(name string) *ProjectsScanConfigsScanRunsGetCall {
	c := &ProjectsScanConfigsScanRunsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsScanConfigsScanRunsGetCall) Fields(s ...googleapi.Field) *ProjectsScanConfigsScanRunsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsScanConfigsScanRunsGetCall) IfNoneMatch(entityTag string) *ProjectsScanConfigsScanRunsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsScanConfigsScanRunsGetCall) Context(ctx context.Context) *ProjectsScanConfigsScanRunsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsScanConfigsScanRunsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsScanConfigsScanRunsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "websecurityscanner.projects.scanConfigs.scanRuns.get"), c.s.client, req)
}

// Do executes the "websecurityscanner.projects.scanConfigs.scanRuns.get" call.
// Exactly one of *ScanRun or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *ScanRun.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsScanConfigsScanRunsGetCall) Do(opts ...googleapi.CallOption) (*ScanRun, error) {
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
	ret := &ScanRun{
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
	//   "description": "Gets a ScanRun.",
	//   "flatPath": "v1alpha/projects/{projectsId}/scanConfigs/{scanConfigsId}/scanRuns/{scanRunsId}",
	//   "httpMethod": "GET",
	//   "id": "websecurityscanner.projects.scanConfigs.scanRuns.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\nThe resource name of the ScanRun to be returned. The name follows the\nformat of\n'projects/{projectId}/scanConfigs/{scanConfigId}/scanRuns/{scanRunId}'.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/scanConfigs/[^/]+/scanRuns/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha/{+name}",
	//   "response": {
	//     "$ref": "ScanRun"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "websecurityscanner.projects.scanConfigs.scanRuns.list":

type ProjectsScanConfigsScanRunsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists ScanRuns under a given ScanConfig, in descending order of
// ScanRun
// stop time.
func (r *ProjectsScanConfigsScanRunsService) List(parent string) *ProjectsScanConfigsScanRunsListCall {
	c := &ProjectsScanConfigsScanRunsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of ScanRuns to return, can be limited by server.
// If not specified or not positive, the implementation will select
// a
// reasonable value.
func (c *ProjectsScanConfigsScanRunsListCall) PageSize(pageSize int64) *ProjectsScanConfigsScanRunsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results to be returned. This should be
// a
// `next_page_token` value returned from a previous List request.
// If unspecified, the first page of results is returned.
func (c *ProjectsScanConfigsScanRunsListCall) PageToken(pageToken string) *ProjectsScanConfigsScanRunsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsScanConfigsScanRunsListCall) Fields(s ...googleapi.Field) *ProjectsScanConfigsScanRunsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsScanConfigsScanRunsListCall) IfNoneMatch(entityTag string) *ProjectsScanConfigsScanRunsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsScanConfigsScanRunsListCall) Context(ctx context.Context) *ProjectsScanConfigsScanRunsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsScanConfigsScanRunsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsScanConfigsScanRunsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+parent}/scanRuns")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "websecurityscanner.projects.scanConfigs.scanRuns.list"), c.s.client, req)
}

// Do executes the "websecurityscanner.projects.scanConfigs.scanRuns.list" call.
// Exactly one of *ListScanRunsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListScanRunsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsScanConfigsScanRunsListCall) Do(opts ...googleapi.CallOption) (*ListScanRunsResponse, error) {
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
	ret := &ListScanRunsResponse{
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
	//   "description": "Lists ScanRuns under a given ScanConfig, in descending order of ScanRun\nstop time.",
	//   "flatPath": "v1alpha/projects/{projectsId}/scanConfigs/{scanConfigsId}/scanRuns",
	//   "httpMethod": "GET",
	//   "id": "websecurityscanner.projects.scanConfigs.scanRuns.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of ScanRuns to return, can be limited by server.\nIf not specified or not positive, the implementation will select a\nreasonable value.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results to be returned. This should be a\n`next_page_token` value returned from a previous List request.\nIf unspecified, the first page of results is returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required.\nThe parent resource name, which should be a scan resource name in the\nformat 'projects/{projectId}/scanConfigs/{scanConfigId}'.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/scanConfigs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha/{+parent}/scanRuns",
	//   "response": {
	//     "$ref": "ListScanRunsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsScanConfigsScanRunsListCall) Pages(ctx context.Context, f func(*ListScanRunsResponse) error) error {
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

// method id "websecurityscanner.projects.scanConfigs.scanRuns.stop":

type ProjectsScanConfigsScanRunsStopCall struct {
	s                  *Service
	name               string
	stopscanrunrequest *StopScanRunRequest
	urlParams_         gensupport.URLParams
	ctx_               context.Context
	header_            http.Header
}

// Stop: Stops a ScanRun. The stopped ScanRun is returned.
func (r *ProjectsScanConfigsScanRunsService) Stop(name string, stopscanrunrequest *StopScanRunRequest) *ProjectsScanConfigsScanRunsStopCall {
	c := &ProjectsScanConfigsScanRunsStopCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.stopscanrunrequest = stopscanrunrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsScanConfigsScanRunsStopCall) Fields(s ...googleapi.Field) *ProjectsScanConfigsScanRunsStopCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsScanConfigsScanRunsStopCall) Context(ctx context.Context) *ProjectsScanConfigsScanRunsStopCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsScanConfigsScanRunsStopCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsScanConfigsScanRunsStopCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.stopscanrunrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+name}:stop")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "websecurityscanner.projects.scanConfigs.scanRuns.stop"), c.s.client, req)
}

// Do executes the "websecurityscanner.projects.scanConfigs.scanRuns.stop" call.
// Exactly one of *ScanRun or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *ScanRun.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsScanConfigsScanRunsStopCall) Do(opts ...googleapi.CallOption) (*ScanRun, error) {
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
	ret := &ScanRun{
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
	//   "description": "Stops a ScanRun. The stopped ScanRun is returned.",
	//   "flatPath": "v1alpha/projects/{projectsId}/scanConfigs/{scanConfigsId}/scanRuns/{scanRunsId}:stop",
	//   "httpMethod": "POST",
	//   "id": "websecurityscanner.projects.scanConfigs.scanRuns.stop",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\nThe resource name of the ScanRun to be stopped. The name follows the\nformat of\n'projects/{projectId}/scanConfigs/{scanConfigId}/scanRuns/{scanRunId}'.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/scanConfigs/[^/]+/scanRuns/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha/{+name}:stop",
	//   "request": {
	//     "$ref": "StopScanRunRequest"
	//   },
	//   "response": {
	//     "$ref": "ScanRun"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "websecurityscanner.projects.scanConfigs.scanRuns.crawledUrls.list":

type ProjectsScanConfigsScanRunsCrawledUrlsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List CrawledUrls under a given ScanRun.
func (r *ProjectsScanConfigsScanRunsCrawledUrlsService) List(parent string) *ProjectsScanConfigsScanRunsCrawledUrlsListCall {
	c := &ProjectsScanConfigsScanRunsCrawledUrlsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of CrawledUrls to return, can be limited by server.
// If not specified or not positive, the implementation will select
// a
// reasonable value.
func (c *ProjectsScanConfigsScanRunsCrawledUrlsListCall) PageSize(pageSize int64) *ProjectsScanConfigsScanRunsCrawledUrlsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results to be returned. This should be
// a
// `next_page_token` value returned from a previous List request.
// If unspecified, the first page of results is returned.
func (c *ProjectsScanConfigsScanRunsCrawledUrlsListCall) PageToken(pageToken string) *ProjectsScanConfigsScanRunsCrawledUrlsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsScanConfigsScanRunsCrawledUrlsListCall) Fields(s ...googleapi.Field) *ProjectsScanConfigsScanRunsCrawledUrlsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsScanConfigsScanRunsCrawledUrlsListCall) IfNoneMatch(entityTag string) *ProjectsScanConfigsScanRunsCrawledUrlsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsScanConfigsScanRunsCrawledUrlsListCall) Context(ctx context.Context) *ProjectsScanConfigsScanRunsCrawledUrlsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsScanConfigsScanRunsCrawledUrlsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsScanConfigsScanRunsCrawledUrlsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+parent}/crawledUrls")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "websecurityscanner.projects.scanConfigs.scanRuns.crawledUrls.list"), c.s.client, req)
}

// Do executes the "websecurityscanner.projects.scanConfigs.scanRuns.crawledUrls.list" call.
// Exactly one of *ListCrawledUrlsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListCrawledUrlsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsScanConfigsScanRunsCrawledUrlsListCall) Do(opts ...googleapi.CallOption) (*ListCrawledUrlsResponse, error) {
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
	ret := &ListCrawledUrlsResponse{
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
	//   "description": "List CrawledUrls under a given ScanRun.",
	//   "flatPath": "v1alpha/projects/{projectsId}/scanConfigs/{scanConfigsId}/scanRuns/{scanRunsId}/crawledUrls",
	//   "httpMethod": "GET",
	//   "id": "websecurityscanner.projects.scanConfigs.scanRuns.crawledUrls.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of CrawledUrls to return, can be limited by server.\nIf not specified or not positive, the implementation will select a\nreasonable value.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results to be returned. This should be a\n`next_page_token` value returned from a previous List request.\nIf unspecified, the first page of results is returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required.\nThe parent resource name, which should be a scan run resource name in the\nformat\n'projects/{projectId}/scanConfigs/{scanConfigId}/scanRuns/{scanRunId}'.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/scanConfigs/[^/]+/scanRuns/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha/{+parent}/crawledUrls",
	//   "response": {
	//     "$ref": "ListCrawledUrlsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsScanConfigsScanRunsCrawledUrlsListCall) Pages(ctx context.Context, f func(*ListCrawledUrlsResponse) error) error {
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

// method id "websecurityscanner.projects.scanConfigs.scanRuns.findingTypeStats.list":

type ProjectsScanConfigsScanRunsFindingTypeStatsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List all FindingTypeStats under a given ScanRun.
func (r *ProjectsScanConfigsScanRunsFindingTypeStatsService) List(parent string) *ProjectsScanConfigsScanRunsFindingTypeStatsListCall {
	c := &ProjectsScanConfigsScanRunsFindingTypeStatsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsScanConfigsScanRunsFindingTypeStatsListCall) Fields(s ...googleapi.Field) *ProjectsScanConfigsScanRunsFindingTypeStatsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsScanConfigsScanRunsFindingTypeStatsListCall) IfNoneMatch(entityTag string) *ProjectsScanConfigsScanRunsFindingTypeStatsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsScanConfigsScanRunsFindingTypeStatsListCall) Context(ctx context.Context) *ProjectsScanConfigsScanRunsFindingTypeStatsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsScanConfigsScanRunsFindingTypeStatsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsScanConfigsScanRunsFindingTypeStatsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+parent}/findingTypeStats")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "websecurityscanner.projects.scanConfigs.scanRuns.findingTypeStats.list"), c.s.client, req)
}

// Do executes the "websecurityscanner.projects.scanConfigs.scanRuns.findingTypeStats.list" call.
// Exactly one of *ListFindingTypeStatsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListFindingTypeStatsResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsScanConfigsScanRunsFindingTypeStatsListCall) Do(opts ...googleapi.CallOption) (*ListFindingTypeStatsResponse, error) {
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
	ret := &ListFindingTypeStatsResponse{
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
	//   "description": "List all FindingTypeStats under a given ScanRun.",
	//   "flatPath": "v1alpha/projects/{projectsId}/scanConfigs/{scanConfigsId}/scanRuns/{scanRunsId}/findingTypeStats",
	//   "httpMethod": "GET",
	//   "id": "websecurityscanner.projects.scanConfigs.scanRuns.findingTypeStats.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required.\nThe parent resource name, which should be a scan run resource name in the\nformat\n'projects/{projectId}/scanConfigs/{scanConfigId}/scanRuns/{scanRunId}'.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/scanConfigs/[^/]+/scanRuns/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha/{+parent}/findingTypeStats",
	//   "response": {
	//     "$ref": "ListFindingTypeStatsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "websecurityscanner.projects.scanConfigs.scanRuns.findings.get":

type ProjectsScanConfigsScanRunsFindingsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a Finding.
func (r *ProjectsScanConfigsScanRunsFindingsService) Get(name string) *ProjectsScanConfigsScanRunsFindingsGetCall {
	c := &ProjectsScanConfigsScanRunsFindingsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsScanConfigsScanRunsFindingsGetCall) Fields(s ...googleapi.Field) *ProjectsScanConfigsScanRunsFindingsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsScanConfigsScanRunsFindingsGetCall) IfNoneMatch(entityTag string) *ProjectsScanConfigsScanRunsFindingsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsScanConfigsScanRunsFindingsGetCall) Context(ctx context.Context) *ProjectsScanConfigsScanRunsFindingsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsScanConfigsScanRunsFindingsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsScanConfigsScanRunsFindingsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "websecurityscanner.projects.scanConfigs.scanRuns.findings.get"), c.s.client, req)
}

// Do executes the "websecurityscanner.projects.scanConfigs.scanRuns.findings.get" call.
// Exactly one of *Finding or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Finding.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsScanConfigsScanRunsFindingsGetCall) Do(opts ...googleapi.CallOption) (*Finding, error) {
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
	ret := &Finding{
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
	//   "description": "Gets a Finding.",
	//   "flatPath": "v1alpha/projects/{projectsId}/scanConfigs/{scanConfigsId}/scanRuns/{scanRunsId}/findings/{findingsId}",
	//   "httpMethod": "GET",
	//   "id": "websecurityscanner.projects.scanConfigs.scanRuns.findings.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\nThe resource name of the Finding to be returned. The name follows the\nformat of\n'projects/{projectId}/scanConfigs/{scanConfigId}/scanRuns/{scanRunId}/findings/{findingId}'.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/scanConfigs/[^/]+/scanRuns/[^/]+/findings/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha/{+name}",
	//   "response": {
	//     "$ref": "Finding"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "websecurityscanner.projects.scanConfigs.scanRuns.findings.list":

type ProjectsScanConfigsScanRunsFindingsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List Findings under a given ScanRun.
func (r *ProjectsScanConfigsScanRunsFindingsService) List(parent string) *ProjectsScanConfigsScanRunsFindingsListCall {
	c := &ProjectsScanConfigsScanRunsFindingsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Filter sets the optional parameter "filter": The filter expression.
// The expression must be in the format: <field>
// <operator> <value>.
// Supported field: 'finding_type'.
// Supported operator: '='.
func (c *ProjectsScanConfigsScanRunsFindingsListCall) Filter(filter string) *ProjectsScanConfigsScanRunsFindingsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of Findings to return, can be limited by server.
// If not specified or not positive, the implementation will select
// a
// reasonable value.
func (c *ProjectsScanConfigsScanRunsFindingsListCall) PageSize(pageSize int64) *ProjectsScanConfigsScanRunsFindingsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results to be returned. This should be
// a
// `next_page_token` value returned from a previous List request.
// If unspecified, the first page of results is returned.
func (c *ProjectsScanConfigsScanRunsFindingsListCall) PageToken(pageToken string) *ProjectsScanConfigsScanRunsFindingsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsScanConfigsScanRunsFindingsListCall) Fields(s ...googleapi.Field) *ProjectsScanConfigsScanRunsFindingsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsScanConfigsScanRunsFindingsListCall) IfNoneMatch(entityTag string) *ProjectsScanConfigsScanRunsFindingsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsScanConfigsScanRunsFindingsListCall) Context(ctx context.Context) *ProjectsScanConfigsScanRunsFindingsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsScanConfigsScanRunsFindingsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsScanConfigsScanRunsFindingsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha/{+parent}/findings")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "websecurityscanner.projects.scanConfigs.scanRuns.findings.list"), c.s.client, req)
}

// Do executes the "websecurityscanner.projects.scanConfigs.scanRuns.findings.list" call.
// Exactly one of *ListFindingsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListFindingsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsScanConfigsScanRunsFindingsListCall) Do(opts ...googleapi.CallOption) (*ListFindingsResponse, error) {
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
	ret := &ListFindingsResponse{
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
	//   "description": "List Findings under a given ScanRun.",
	//   "flatPath": "v1alpha/projects/{projectsId}/scanConfigs/{scanConfigsId}/scanRuns/{scanRunsId}/findings",
	//   "httpMethod": "GET",
	//   "id": "websecurityscanner.projects.scanConfigs.scanRuns.findings.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "The filter expression. The expression must be in the format: \u003cfield\u003e\n\u003coperator\u003e \u003cvalue\u003e.\nSupported field: 'finding_type'.\nSupported operator: '='.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of Findings to return, can be limited by server.\nIf not specified or not positive, the implementation will select a\nreasonable value.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results to be returned. This should be a\n`next_page_token` value returned from a previous List request.\nIf unspecified, the first page of results is returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required.\nThe parent resource name, which should be a scan run resource name in the\nformat\n'projects/{projectId}/scanConfigs/{scanConfigId}/scanRuns/{scanRunId}'.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/scanConfigs/[^/]+/scanRuns/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha/{+parent}/findings",
	//   "response": {
	//     "$ref": "ListFindingsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsScanConfigsScanRunsFindingsListCall) Pages(ctx context.Context, f func(*ListFindingsResponse) error) error {
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
