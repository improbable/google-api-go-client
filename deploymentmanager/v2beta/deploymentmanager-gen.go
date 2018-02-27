// Package deploymentmanager provides access to the Google Cloud Deployment Manager API V2Beta Methods.
//
// See https://developers.google.com/deployment-manager/
//
// Usage example:
//
//   import "google.golang.org/api/deploymentmanager/v2beta"
//   ...
//   deploymentmanagerService, err := deploymentmanager.New(oauthHttpClient)
package deploymentmanager // import "google.golang.org/api/deploymentmanager/v2beta"

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

const apiId = "deploymentmanager:v2beta"
const apiName = "deploymentmanager"
const apiVersion = "v2beta"
const basePath = "https://www.googleapis.com/deploymentmanager/v2beta/projects/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View your data across Google Cloud Platform services
	CloudPlatformReadOnlyScope = "https://www.googleapis.com/auth/cloud-platform.read-only"

	// View and manage your Google Cloud Platform management resources and
	// deployment status information
	NdevCloudmanScope = "https://www.googleapis.com/auth/ndev.cloudman"

	// View your Google Cloud Platform management resources and deployment
	// status information
	NdevCloudmanReadonlyScope = "https://www.googleapis.com/auth/ndev.cloudman.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.CompositeTypes = NewCompositeTypesService(s)
	s.Deployments = NewDeploymentsService(s)
	s.Manifests = NewManifestsService(s)
	s.Operations = NewOperationsService(s)
	s.Resources = NewResourcesService(s)
	s.TypeProviders = NewTypeProvidersService(s)
	s.Types = NewTypesService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	CompositeTypes *CompositeTypesService

	Deployments *DeploymentsService

	Manifests *ManifestsService

	Operations *OperationsService

	Resources *ResourcesService

	TypeProviders *TypeProvidersService

	Types *TypesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewCompositeTypesService(s *Service) *CompositeTypesService {
	rs := &CompositeTypesService{s: s}
	return rs
}

type CompositeTypesService struct {
	s *Service
}

func NewDeploymentsService(s *Service) *DeploymentsService {
	rs := &DeploymentsService{s: s}
	return rs
}

type DeploymentsService struct {
	s *Service
}

func NewManifestsService(s *Service) *ManifestsService {
	rs := &ManifestsService{s: s}
	return rs
}

type ManifestsService struct {
	s *Service
}

func NewOperationsService(s *Service) *OperationsService {
	rs := &OperationsService{s: s}
	return rs
}

type OperationsService struct {
	s *Service
}

func NewResourcesService(s *Service) *ResourcesService {
	rs := &ResourcesService{s: s}
	return rs
}

type ResourcesService struct {
	s *Service
}

func NewTypeProvidersService(s *Service) *TypeProvidersService {
	rs := &TypeProvidersService{s: s}
	return rs
}

type TypeProvidersService struct {
	s *Service
}

func NewTypesService(s *Service) *TypesService {
	rs := &TypesService{s: s}
	return rs
}

type TypesService struct {
	s *Service
}

// AsyncOptions: Async options that determine when a resource should
// finish.
type AsyncOptions struct {
	// MethodMatch: Method regex where this policy will apply.
	MethodMatch string `json:"methodMatch,omitempty"`

	// PollingOptions: Deployment manager will poll instances for this API
	// resource setting a RUNNING state, and blocking until polling
	// conditions tell whether the resource is completed or failed.
	PollingOptions *PollingOptions `json:"pollingOptions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MethodMatch") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MethodMatch") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AsyncOptions) MarshalJSON() ([]byte, error) {
	type NoMethod AsyncOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AuditConfig: Specifies the audit configuration for a service. The
// configuration determines which permission types are logged, and what
// identities, if any, are exempted from logging. An AuditConfig must
// have one or more AuditLogConfigs.
//
// If there are AuditConfigs for both `allServices` and a specific
// service, the union of the two AuditConfigs is used for that service:
// the log_types specified in each AuditConfig are enabled, and the
// exempted_members in each AuditLogConfig are exempted.
//
// Example Policy with multiple AuditConfigs:
//
// { "audit_configs": [ { "service": "allServices" "audit_log_configs":
// [ { "log_type": "DATA_READ", "exempted_members": [
// "user:foo@gmail.com" ] }, { "log_type": "DATA_WRITE", }, {
// "log_type": "ADMIN_READ", } ] }, { "service":
// "fooservice.googleapis.com" "audit_log_configs": [ { "log_type":
// "DATA_READ", }, { "log_type": "DATA_WRITE", "exempted_members": [
// "user:bar@gmail.com" ] } ] } ] }
//
// For fooservice, this policy enables DATA_READ, DATA_WRITE and
// ADMIN_READ logging. It also exempts foo@gmail.com from DATA_READ
// logging, and bar@gmail.com from DATA_WRITE logging.
type AuditConfig struct {
	// AuditLogConfigs: The configuration for logging of each type of
	// permission.
	AuditLogConfigs []*AuditLogConfig `json:"auditLogConfigs,omitempty"`

	ExemptedMembers []string `json:"exemptedMembers,omitempty"`

	// Service: Specifies a service that will be enabled for audit logging.
	// For example, `storage.googleapis.com`, `cloudsql.googleapis.com`.
	// `allServices` is a special value that covers all services.
	Service string `json:"service,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AuditLogConfigs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AuditLogConfigs") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AuditConfig) MarshalJSON() ([]byte, error) {
	type NoMethod AuditConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AuditLogConfig: Provides the configuration for logging a type of
// permissions. Example:
//
// { "audit_log_configs": [ { "log_type": "DATA_READ",
// "exempted_members": [ "user:foo@gmail.com" ] }, { "log_type":
// "DATA_WRITE", } ] }
//
// This enables 'DATA_READ' and 'DATA_WRITE' logging, while exempting
// foo@gmail.com from DATA_READ logging.
type AuditLogConfig struct {
	// ExemptedMembers: Specifies the identities that do not cause logging
	// for this type of permission. Follows the same format of
	// [Binding.members][].
	ExemptedMembers []string `json:"exemptedMembers,omitempty"`

	// LogType: The log type that this config enables.
	LogType string `json:"logType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExemptedMembers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExemptedMembers") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AuditLogConfig) MarshalJSON() ([]byte, error) {
	type NoMethod AuditLogConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AuthorizationLoggingOptions: Authorization-related information used
// by Cloud Audit Logging.
type AuthorizationLoggingOptions struct {
	// PermissionType: The type of the permission that was checked.
	PermissionType string `json:"permissionType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PermissionType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PermissionType") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AuthorizationLoggingOptions) MarshalJSON() ([]byte, error) {
	type NoMethod AuthorizationLoggingOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BaseType: BaseType that describes a service-backed Type.
type BaseType struct {
	// CollectionOverrides: Allows resource handling overrides for specific
	// collections
	CollectionOverrides []*CollectionOverride `json:"collectionOverrides,omitempty"`

	// Credential: Credential used when interacting with this type.
	Credential *Credential `json:"credential,omitempty"`

	// DescriptorUrl: Descriptor Url for the this type.
	DescriptorUrl string `json:"descriptorUrl,omitempty"`

	// Options: Options to apply when handling any resources in this
	// service.
	Options *Options `json:"options,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CollectionOverrides")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CollectionOverrides") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *BaseType) MarshalJSON() ([]byte, error) {
	type NoMethod BaseType
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BasicAuth: Basic Auth used as a credential.
type BasicAuth struct {
	Password string `json:"password,omitempty"`

	User string `json:"user,omitempty"`

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

func (s *BasicAuth) MarshalJSON() ([]byte, error) {
	type NoMethod BasicAuth
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Binding: Associates `members` with a `role`.
type Binding struct {
	// Condition: The condition that is associated with this binding. NOTE:
	// an unsatisfied condition will not allow user access via current
	// binding. Different bindings, including their conditions, are examined
	// independently. This field is only visible as GOOGLE_INTERNAL or
	// CONDITION_TRUSTED_TESTER.
	Condition *Expr `json:"condition,omitempty"`

	// Members: Specifies the identities requesting access for a Cloud
	// Platform resource. `members` can have the following values:
	//
	// * `allUsers`: A special identifier that represents anyone who is on
	// the internet; with or without a Google account.
	//
	// * `allAuthenticatedUsers`: A special identifier that represents
	// anyone who is authenticated with a Google account or a service
	// account.
	//
	// * `user:{emailid}`: An email address that represents a specific
	// Google account. For example, `alice@gmail.com` or
	// `joe@example.com`.
	//
	//
	//
	// * `serviceAccount:{emailid}`: An email address that represents a
	// service account. For example,
	// `my-other-app@appspot.gserviceaccount.com`.
	//
	// * `group:{emailid}`: An email address that represents a Google group.
	// For example, `admins@example.com`.
	//
	//
	//
	// * `domain:{domain}`: A Google Apps domain name that represents all
	// the users of that domain. For example, `google.com` or `example.com`.
	Members []string `json:"members,omitempty"`

	// Role: Role that is assigned to `members`. For example,
	// `roles/viewer`, `roles/editor`, or `roles/owner`.
	Role string `json:"role,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Condition") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Condition") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Binding) MarshalJSON() ([]byte, error) {
	type NoMethod Binding
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CollectionOverride: CollectionOverride allows resource handling
// overrides for specific resources within a BaseType
type CollectionOverride struct {
	// Collection: The collection that identifies this resource within its
	// service.
	Collection string `json:"collection,omitempty"`

	// Options: The options to apply to this resource-level override
	Options *Options `json:"options,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Collection") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Collection") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CollectionOverride) MarshalJSON() ([]byte, error) {
	type NoMethod CollectionOverride
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CompositeType: Holds the composite type.
type CompositeType struct {
	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// Id: Output only. Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// InsertTime: Output only. Timestamp when the composite type was
	// created, in RFC3339 text format.
	InsertTime string `json:"insertTime,omitempty"`

	// Labels: Map of labels; provided by the client when the resource is
	// created or updated. Specifically: Label keys must be between 1 and 63
	// characters long and must conform to the following regular expression:
	// [a-z]([-a-z0-9]*[a-z0-9])? Label values must be between 0 and 63
	// characters long and must conform to the regular expression
	// ([a-z]([-a-z0-9]*[a-z0-9])?)?
	Labels []*CompositeTypeLabelEntry `json:"labels,omitempty"`

	// Name: Name of the composite type, must follow the expression:
	// [a-z]([-a-z0-9_.]{0,61}[a-z0-9])?.
	Name string `json:"name,omitempty"`

	// Operation: Output only. The Operation that most recently ran, or is
	// currently running, on this composite type.
	Operation *Operation `json:"operation,omitempty"`

	// SelfLink: Output only. Self link for the type provider.
	SelfLink string `json:"selfLink,omitempty"`

	Status string `json:"status,omitempty"`

	// TemplateContents: Files for the template type.
	TemplateContents *TemplateContents `json:"templateContents,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CompositeType) MarshalJSON() ([]byte, error) {
	type NoMethod CompositeType
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type CompositeTypeLabelEntry struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CompositeTypeLabelEntry) MarshalJSON() ([]byte, error) {
	type NoMethod CompositeTypeLabelEntry
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CompositeTypesListResponse: A response that returns all Composite
// Types supported by Deployment Manager
type CompositeTypesListResponse struct {
	// CompositeTypes: Output only. A list of resource composite types
	// supported by Deployment Manager.
	CompositeTypes []*CompositeType `json:"compositeTypes,omitempty"`

	// NextPageToken: A token used to continue a truncated list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CompositeTypes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CompositeTypes") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *CompositeTypesListResponse) MarshalJSON() ([]byte, error) {
	type NoMethod CompositeTypesListResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Condition: A condition to be met.
type Condition struct {
	// Iam: Trusted attributes supplied by the IAM system.
	Iam string `json:"iam,omitempty"`

	// Op: An operator to apply the subject with.
	Op string `json:"op,omitempty"`

	// Svc: Trusted attributes discharged by the service.
	Svc string `json:"svc,omitempty"`

	// Sys: Trusted attributes supplied by any service that owns resources
	// and uses the IAM system for access control.
	Sys string `json:"sys,omitempty"`

	// Value: DEPRECATED. Use 'values' instead.
	Value string `json:"value,omitempty"`

	// Values: The objects of the condition. This is mutually exclusive with
	// 'value'.
	Values []string `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Iam") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Iam") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Condition) MarshalJSON() ([]byte, error) {
	type NoMethod Condition
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ConfigFile struct {
	// Content: The contents of the file.
	Content string `json:"content,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Content") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Content") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ConfigFile) MarshalJSON() ([]byte, error) {
	type NoMethod ConfigFile
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Credential: The credential used by Deployment Manager and
// TypeProvider. Only one of the options is permitted.
type Credential struct {
	// BasicAuth: Basic Auth Credential, only used by TypeProvider.
	BasicAuth *BasicAuth `json:"basicAuth,omitempty"`

	// ServiceAccount: Service Account Credential, only used by Deployment.
	ServiceAccount *ServiceAccount `json:"serviceAccount,omitempty"`

	// UseProjectDefault: Specify to use the project default credential,
	// only supported by Deployment.
	UseProjectDefault bool `json:"useProjectDefault,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BasicAuth") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BasicAuth") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Credential) MarshalJSON() ([]byte, error) {
	type NoMethod Credential
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Deployment struct {
	// Description: An optional user-provided description of the deployment.
	Description string `json:"description,omitempty"`

	// Fingerprint: Provides a fingerprint to use in requests to modify a
	// deployment, such as update(), stop(), and cancelPreview() requests. A
	// fingerprint is a randomly generated value that must be provided with
	// update(), stop(), and cancelPreview() requests to perform optimistic
	// locking. This ensures optimistic concurrency so that only one request
	// happens at a time.
	//
	// The fingerprint is initially generated by Deployment Manager and
	// changes after every request to modify data. To get the latest
	// fingerprint value, perform a get() request to a deployment.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Id: Output only. Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// InsertTime: Output only. Timestamp when the deployment was created,
	// in RFC3339 text format .
	InsertTime string `json:"insertTime,omitempty"`

	// Labels: Map of labels; provided by the client when the resource is
	// created or updated. Specifically: Label keys must be between 1 and 63
	// characters long and must conform to the following regular expression:
	// [a-z]([-a-z0-9]*[a-z0-9])? Label values must be between 0 and 63
	// characters long and must conform to the regular expression
	// ([a-z]([-a-z0-9]*[a-z0-9])?)?
	Labels []*DeploymentLabelEntry `json:"labels,omitempty"`

	// Manifest: Output only. URL of the manifest representing the last
	// manifest that was successfully deployed.
	Manifest string `json:"manifest,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name string `json:"name,omitempty"`

	// Operation: Output only. The Operation that most recently ran, or is
	// currently running, on this deployment.
	Operation *Operation `json:"operation,omitempty"`

	// SelfLink: Output only. Self link for the deployment.
	SelfLink string `json:"selfLink,omitempty"`

	// Target: [Input Only] The parameters that define your deployment,
	// including the deployment configuration and relevant templates.
	Target *TargetConfiguration `json:"target,omitempty"`

	// Update: Output only. If Deployment Manager is currently updating or
	// previewing an update to this deployment, the updated configuration
	// appears here.
	Update *DeploymentUpdate `json:"update,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Deployment) MarshalJSON() ([]byte, error) {
	type NoMethod Deployment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type DeploymentLabelEntry struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DeploymentLabelEntry) MarshalJSON() ([]byte, error) {
	type NoMethod DeploymentLabelEntry
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type DeploymentUpdate struct {
	// Description: Output only. An optional user-provided description of
	// the deployment after the current update has been applied.
	Description string `json:"description,omitempty"`

	// Labels: Output only. Map of labels; provided by the client when the
	// resource is created or updated. Specifically: Label keys must be
	// between 1 and 63 characters long and must conform to the following
	// regular expression: [a-z]([-a-z0-9]*[a-z0-9])? Label values must be
	// between 0 and 63 characters long and must conform to the regular
	// expression ([a-z]([-a-z0-9]*[a-z0-9])?)?
	Labels []*DeploymentUpdateLabelEntry `json:"labels,omitempty"`

	// Manifest: Output only. URL of the manifest representing the update
	// configuration of this deployment.
	Manifest string `json:"manifest,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DeploymentUpdate) MarshalJSON() ([]byte, error) {
	type NoMethod DeploymentUpdate
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type DeploymentUpdateLabelEntry struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DeploymentUpdateLabelEntry) MarshalJSON() ([]byte, error) {
	type NoMethod DeploymentUpdateLabelEntry
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type DeploymentsCancelPreviewRequest struct {
	// Fingerprint: Specifies a fingerprint for cancelPreview() requests. A
	// fingerprint is a randomly generated value that must be provided in
	// cancelPreview() requests to perform optimistic locking. This ensures
	// optimistic concurrency so that the deployment does not have
	// conflicting requests (e.g. if someone attempts to make a new update
	// request while another user attempts to cancel a preview, this would
	// prevent one of the requests).
	//
	// The fingerprint is initially generated by Deployment Manager and
	// changes after every request to modify a deployment. To get the latest
	// fingerprint value, perform a get() request on the deployment.
	Fingerprint string `json:"fingerprint,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Fingerprint") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Fingerprint") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DeploymentsCancelPreviewRequest) MarshalJSON() ([]byte, error) {
	type NoMethod DeploymentsCancelPreviewRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DeploymentsListResponse: A response containing a partial list of
// deployments and a page token used to build the next request if the
// request has been truncated.
type DeploymentsListResponse struct {
	// Deployments: Output only. The deployments contained in this response.
	Deployments []*Deployment `json:"deployments,omitempty"`

	// NextPageToken: Output only. A token used to continue a truncated list
	// request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Deployments") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Deployments") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DeploymentsListResponse) MarshalJSON() ([]byte, error) {
	type NoMethod DeploymentsListResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type DeploymentsStopRequest struct {
	// Fingerprint: Specifies a fingerprint for stop() requests. A
	// fingerprint is a randomly generated value that must be provided in
	// stop() requests to perform optimistic locking. This ensures
	// optimistic concurrency so that the deployment does not have
	// conflicting requests (e.g. if someone attempts to make a new update
	// request while another user attempts to stop an ongoing update
	// request, this would prevent a collision).
	//
	// The fingerprint is initially generated by Deployment Manager and
	// changes after every request to modify a deployment. To get the latest
	// fingerprint value, perform a get() request on the deployment.
	Fingerprint string `json:"fingerprint,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Fingerprint") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Fingerprint") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DeploymentsStopRequest) MarshalJSON() ([]byte, error) {
	type NoMethod DeploymentsStopRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Diagnostic struct {
	// Field: JsonPath expression on the resource that if non empty,
	// indicates that this field needs to be extracted as a diagnostic.
	Field string `json:"field,omitempty"`

	// Level: Level to record this diagnostic.
	Level string `json:"level,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Field") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Field") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Diagnostic) MarshalJSON() ([]byte, error) {
	type NoMethod Diagnostic
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Expr: Represents an expression text. Example:
//
// title: "User account presence" description: "Determines whether the
// request has a user account" expression: "size(request.user) > 0"
type Expr struct {
	// Description: An optional description of the expression. This is a
	// longer text which describes the expression, e.g. when hovered over it
	// in a UI.
	Description string `json:"description,omitempty"`

	// Expression: Textual representation of an expression in Common
	// Expression Language syntax.
	//
	// The application context of the containing message determines which
	// well-known feature set of CEL is supported.
	Expression string `json:"expression,omitempty"`

	// Location: An optional string indicating the location of the
	// expression for error reporting, e.g. a file name and a position in
	// the file.
	Location string `json:"location,omitempty"`

	// Title: An optional title for the expression, i.e. a short string
	// describing its purpose. This can be used e.g. in UIs which allow to
	// enter the expression.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Expr) MarshalJSON() ([]byte, error) {
	type NoMethod Expr
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ImportFile struct {
	// Content: The contents of the file.
	Content string `json:"content,omitempty"`

	// Name: The name of the file.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Content") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Content") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ImportFile) MarshalJSON() ([]byte, error) {
	type NoMethod ImportFile
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// InputMapping: InputMapping creates a 'virtual' property that will be
// injected into the properties before sending the request to the
// underlying API.
type InputMapping struct {
	// FieldName: The name of the field that is going to be injected.
	FieldName string `json:"fieldName,omitempty"`

	// Location: The location where this mapping applies.
	Location string `json:"location,omitempty"`

	// MethodMatch: Regex to evaluate on method to decide if input applies.
	MethodMatch string `json:"methodMatch,omitempty"`

	// Value: A jsonPath expression to select an element.
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FieldName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FieldName") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *InputMapping) MarshalJSON() ([]byte, error) {
	type NoMethod InputMapping
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LogConfig: Specifies what kind of log the caller must write
type LogConfig struct {
	// CloudAudit: Cloud audit options.
	CloudAudit *LogConfigCloudAuditOptions `json:"cloudAudit,omitempty"`

	// Counter: Counter options.
	Counter *LogConfigCounterOptions `json:"counter,omitempty"`

	// DataAccess: Data access options.
	DataAccess *LogConfigDataAccessOptions `json:"dataAccess,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CloudAudit") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CloudAudit") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LogConfig) MarshalJSON() ([]byte, error) {
	type NoMethod LogConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LogConfigCloudAuditOptions: Write a Cloud Audit log
type LogConfigCloudAuditOptions struct {
	// AuthorizationLoggingOptions: Information used by the Cloud Audit
	// Logging pipeline.
	AuthorizationLoggingOptions *AuthorizationLoggingOptions `json:"authorizationLoggingOptions,omitempty"`

	// LogName: The log_name to populate in the Cloud Audit Record.
	LogName string `json:"logName,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AuthorizationLoggingOptions") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "AuthorizationLoggingOptions") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LogConfigCloudAuditOptions) MarshalJSON() ([]byte, error) {
	type NoMethod LogConfigCloudAuditOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LogConfigCounterOptions: Increment a streamz counter with the
// specified metric and field names.
//
// Metric names should start with a '/', generally be lowercase-only,
// and end in "_count". Field names should not contain an initial slash.
// The actual exported metric names will have "/iam/policy"
// prepended.
//
// Field names correspond to IAM request parameters and field values are
// their respective values.
//
// At present the only supported field names are - "iam_principal",
// corresponding to IAMContext.principal; - "" (empty string), resulting
// in one aggretated counter with no field.
//
// Examples: counter { metric: "/debug_access_count" field:
// "iam_principal" } ==> increment counter
// /iam/policy/backend_debug_access_count {iam_principal=[value of
// IAMContext.principal]}
//
// At this time we do not support: * multiple field names (though this
// may be supported in the future) * decrementing the counter *
// incrementing it by anything other than 1
type LogConfigCounterOptions struct {
	// Field: The field value to attribute.
	Field string `json:"field,omitempty"`

	// Metric: The metric to update.
	Metric string `json:"metric,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Field") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Field") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LogConfigCounterOptions) MarshalJSON() ([]byte, error) {
	type NoMethod LogConfigCounterOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LogConfigDataAccessOptions: Write a Data Access (Gin) log
type LogConfigDataAccessOptions struct {
	// LogMode: Whether Gin logging should happen in a fail-closed manner at
	// the caller. This is relevant only in the LocalIAM implementation, for
	// now.
	LogMode string `json:"logMode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LogMode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LogMode") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LogConfigDataAccessOptions) MarshalJSON() ([]byte, error) {
	type NoMethod LogConfigDataAccessOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Manifest struct {
	// Config: Output only. The YAML configuration for this manifest.
	Config *ConfigFile `json:"config,omitempty"`

	// ExpandedConfig: Output only. The fully-expanded configuration file,
	// including any templates and references.
	ExpandedConfig string `json:"expandedConfig,omitempty"`

	// Id: Output only. Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Imports: Output only. The imported files for this manifest.
	Imports []*ImportFile `json:"imports,omitempty"`

	// InsertTime: Output only. Timestamp when the manifest was created, in
	// RFC3339 text format.
	InsertTime string `json:"insertTime,omitempty"`

	// Layout: Output only. The YAML layout for this manifest.
	Layout string `json:"layout,omitempty"`

	// Name: Output only.
	//
	// The name of the manifest.
	Name string `json:"name,omitempty"`

	// SelfLink: Output only. Self link for the manifest.
	SelfLink string `json:"selfLink,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Config") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Config") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Manifest) MarshalJSON() ([]byte, error) {
	type NoMethod Manifest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ManifestsListResponse: A response containing a partial list of
// manifests and a page token used to build the next request if the
// request has been truncated.
type ManifestsListResponse struct {
	// Manifests: Output only. Manifests contained in this list response.
	Manifests []*Manifest `json:"manifests,omitempty"`

	// NextPageToken: Output only. A token used to continue a truncated list
	// request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Manifests") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Manifests") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ManifestsListResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ManifestsListResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Operation: An Operation resource, used to manage asynchronous API
// requests. (== resource_for v1.globalOperations ==) (== resource_for
// beta.globalOperations ==) (== resource_for v1.regionOperations ==)
// (== resource_for beta.regionOperations ==) (== resource_for
// v1.zoneOperations ==) (== resource_for beta.zoneOperations ==)
type Operation struct {
	// ClientOperationId: [Output Only] Reserved for future use.
	ClientOperationId string `json:"clientOperationId,omitempty"`

	// CreationTimestamp: [Deprecated] This field is deprecated.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: [Output Only] A textual description of the operation,
	// which is set when the operation is created.
	Description string `json:"description,omitempty"`

	// EndTime: [Output Only] The time that this operation was completed.
	// This value is in RFC3339 text format.
	EndTime string `json:"endTime,omitempty"`

	// Error: [Output Only] If errors are generated during processing of the
	// operation, this field will be populated.
	Error *OperationError `json:"error,omitempty"`

	// HttpErrorMessage: [Output Only] If the operation fails, this field
	// contains the HTTP error message that was returned, such as NOT FOUND.
	HttpErrorMessage string `json:"httpErrorMessage,omitempty"`

	// HttpErrorStatusCode: [Output Only] If the operation fails, this field
	// contains the HTTP error status code that was returned. For example, a
	// 404 means the resource was not found.
	HttpErrorStatusCode int64 `json:"httpErrorStatusCode,omitempty"`

	// Id: [Output Only] The unique identifier for the resource. This
	// identifier is defined by the server.
	Id uint64 `json:"id,omitempty,string"`

	// InsertTime: [Output Only] The time that this operation was requested.
	// This value is in RFC3339 text format.
	InsertTime string `json:"insertTime,omitempty"`

	// Kind: [Output Only] Type of the resource. Always compute#operation
	// for Operation resources.
	Kind string `json:"kind,omitempty"`

	// Name: [Output Only] Name of the resource.
	Name string `json:"name,omitempty"`

	// OperationType: [Output Only] The type of operation, such as insert,
	// update, or delete, and so on.
	OperationType string `json:"operationType,omitempty"`

	// Progress: [Output Only] An optional progress indicator that ranges
	// from 0 to 100. There is no requirement that this be linear or support
	// any granularity of operations. This should not be used to guess when
	// the operation will be complete. This number should monotonically
	// increase as the operation progresses.
	Progress int64 `json:"progress,omitempty"`

	// Region: [Output Only] The URL of the region where the operation
	// resides. Only available when performing regional operations. You must
	// specify this field as part of the HTTP request URL. It is not
	// settable as a field in the request body.
	Region string `json:"region,omitempty"`

	// SelfLink: [Output Only] Server-defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// StartTime: [Output Only] The time that this operation was started by
	// the server. This value is in RFC3339 text format.
	StartTime string `json:"startTime,omitempty"`

	// Status: [Output Only] The status of the operation, which can be one
	// of the following: PENDING, RUNNING, or DONE.
	Status string `json:"status,omitempty"`

	// StatusMessage: [Output Only] An optional textual description of the
	// current status of the operation.
	StatusMessage string `json:"statusMessage,omitempty"`

	// TargetId: [Output Only] The unique target ID, which identifies a
	// specific incarnation of the target resource.
	TargetId uint64 `json:"targetId,omitempty,string"`

	// TargetLink: [Output Only] The URL of the resource that the operation
	// modifies. For operations related to creating a snapshot, this points
	// to the persistent disk that the snapshot was created from.
	TargetLink string `json:"targetLink,omitempty"`

	// User: [Output Only] User who requested the operation, for example:
	// user@example.com.
	User string `json:"user,omitempty"`

	// Warnings: [Output Only] If warning messages are generated during
	// processing of the operation, this field will be populated.
	Warnings []*OperationWarnings `json:"warnings,omitempty"`

	// Zone: [Output Only] The URL of the zone where the operation resides.
	// Only available when performing per-zone operations. You must specify
	// this field as part of the HTTP request URL. It is not settable as a
	// field in the request body.
	Zone string `json:"zone,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ClientOperationId")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClientOperationId") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Operation) MarshalJSON() ([]byte, error) {
	type NoMethod Operation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OperationError: [Output Only] If errors are generated during
// processing of the operation, this field will be populated.
type OperationError struct {
	// Errors: [Output Only] The array of errors encountered while
	// processing this operation.
	Errors []*OperationErrorErrors `json:"errors,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Errors") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Errors") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OperationError) MarshalJSON() ([]byte, error) {
	type NoMethod OperationError
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type OperationErrorErrors struct {
	// Code: [Output Only] The error type identifier for this error.
	Code string `json:"code,omitempty"`

	// Location: [Output Only] Indicates the field in the request that
	// caused the error. This property is optional.
	Location string `json:"location,omitempty"`

	// Message: [Output Only] An optional, human-readable error message.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OperationErrorErrors) MarshalJSON() ([]byte, error) {
	type NoMethod OperationErrorErrors
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type OperationWarnings struct {
	// Code: [Output Only] A warning code, if applicable. For example,
	// Compute Engine returns NO_RESULTS_ON_PAGE if there are no results in
	// the response.
	Code string `json:"code,omitempty"`

	// Data: [Output Only] Metadata about this warning in key: value format.
	// For example:
	// "data": [ { "key": "scope", "value": "zones/us-east1-d" }
	Data []*OperationWarningsData `json:"data,omitempty"`

	// Message: [Output Only] A human-readable description of the warning
	// code.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OperationWarnings) MarshalJSON() ([]byte, error) {
	type NoMethod OperationWarnings
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type OperationWarningsData struct {
	// Key: [Output Only] A key that provides more detail on the warning
	// being returned. For example, for warnings where there are no results
	// in a list request for a particular zone, this key might be scope and
	// the key value might be the zone name. Other examples might be a key
	// indicating a deprecated resource and a suggested replacement, or a
	// warning about invalid network settings (for example, if an instance
	// attempts to perform IP forwarding but is not enabled for IP
	// forwarding).
	Key string `json:"key,omitempty"`

	// Value: [Output Only] A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OperationWarningsData) MarshalJSON() ([]byte, error) {
	type NoMethod OperationWarningsData
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OperationsListResponse: A response containing a partial list of
// operations and a page token used to build the next request if the
// request has been truncated.
type OperationsListResponse struct {
	// NextPageToken: Output only. A token used to continue a truncated list
	// request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Operations: Output only. Operations contained in this list response.
	Operations []*Operation `json:"operations,omitempty"`

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

func (s *OperationsListResponse) MarshalJSON() ([]byte, error) {
	type NoMethod OperationsListResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Options: Options allows customized resource handling by Deployment
// Manager.
type Options struct {
	// AsyncOptions: Options regarding how to thread async requests.
	AsyncOptions []*AsyncOptions `json:"asyncOptions,omitempty"`

	// InputMappings: The mappings that apply for requests.
	InputMappings []*InputMapping `json:"inputMappings,omitempty"`

	// ValidationOptions: Options for how to validate and process properties
	// on a resource.
	ValidationOptions *ValidationOptions `json:"validationOptions,omitempty"`

	// VirtualProperties: Additional properties block described as a
	// jsonSchema, these properties will never be part of the json payload,
	// but they can be consumed by InputMappings, this must be a valid json
	// schema draft-04. The properties specified here will be decouple in a
	// different section. This schema will be merged to the schema
	// validation, and properties here will be extracted From the payload
	// and consumed explicitly by InputMappings. ex: field1: type: string
	// field2: type: number
	VirtualProperties string `json:"virtualProperties,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AsyncOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AsyncOptions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Options) MarshalJSON() ([]byte, error) {
	type NoMethod Options
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Policy: Defines an Identity and Access Management (IAM) policy. It is
// used to specify access control policies for Cloud Platform
// resources.
//
//
//
// A `Policy` consists of a list of `bindings`. A `Binding` binds a list
// of `members` to a `role`, where the members can be user accounts,
// Google groups, Google domains, and service accounts. A `role` is a
// named list of permissions defined by IAM.
//
// **Example**
//
// { "bindings": [ { "role": "roles/owner", "members": [
// "user:mike@example.com", "group:admins@example.com",
// "domain:google.com",
// "serviceAccount:my-other-app@appspot.gserviceaccount.com", ] }, {
// "role": "roles/viewer", "members": ["user:sean@example.com"] } ]
// }
//
// For a description of IAM and its features, see the [IAM developer's
// guide](https://cloud.google.com/iam/docs).
type Policy struct {
	// AuditConfigs: Specifies cloud audit logging configuration for this
	// policy.
	AuditConfigs []*AuditConfig `json:"auditConfigs,omitempty"`

	// Bindings: Associates a list of `members` to a `role`. `bindings` with
	// no members will result in an error.
	Bindings []*Binding `json:"bindings,omitempty"`

	// Etag: `etag` is used for optimistic concurrency control as a way to
	// help prevent simultaneous updates of a policy from overwriting each
	// other. It is strongly suggested that systems make use of the `etag`
	// in the read-modify-write cycle to perform policy updates in order to
	// avoid race conditions: An `etag` is returned in the response to
	// `getIamPolicy`, and systems are expected to put that etag in the
	// request to `setIamPolicy` to ensure that their change will be applied
	// to the same version of the policy.
	//
	// If no `etag` is provided in the call to `setIamPolicy`, then the
	// existing policy is overwritten blindly.
	Etag string `json:"etag,omitempty"`

	IamOwned bool `json:"iamOwned,omitempty"`

	// Rules: If more than one rule is specified, the rules are applied in
	// the following manner: - All matching LOG rules are always applied. -
	// If any DENY/DENY_WITH_LOG rule matches, permission is denied. Logging
	// will be applied if one or more matching rule requires logging. -
	// Otherwise, if any ALLOW/ALLOW_WITH_LOG rule matches, permission is
	// granted. Logging will be applied if one or more matching rule
	// requires logging. - Otherwise, if no rule applies, permission is
	// denied.
	Rules []*Rule `json:"rules,omitempty"`

	// Version: Deprecated.
	Version int64 `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AuditConfigs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AuditConfigs") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Policy) MarshalJSON() ([]byte, error) {
	type NoMethod Policy
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PollingOptions struct {
	// Diagnostics: An array of diagnostics to be collected by Deployment
	// Manager, these diagnostics will be displayed to the user.
	Diagnostics []*Diagnostic `json:"diagnostics,omitempty"`

	// FailCondition: JsonPath expression that determines if the request
	// failed.
	FailCondition string `json:"failCondition,omitempty"`

	// FinishCondition: JsonPath expression that determines if the request
	// is completed.
	FinishCondition string `json:"finishCondition,omitempty"`

	// PollingLink: JsonPath expression that evaluates to string, it
	// indicates where to poll.
	PollingLink string `json:"pollingLink,omitempty"`

	// TargetLink: JsonPath expression, after polling is completed,
	// indicates where to fetch the resource.
	TargetLink string `json:"targetLink,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Diagnostics") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Diagnostics") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PollingOptions) MarshalJSON() ([]byte, error) {
	type NoMethod PollingOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Resource struct {
	// AccessControl: The Access Control Policy set on this resource.
	AccessControl *ResourceAccessControl `json:"accessControl,omitempty"`

	// FinalProperties: Output only. The evaluated properties of the
	// resource with references expanded. Returned as serialized YAML.
	FinalProperties string `json:"finalProperties,omitempty"`

	// Id: Output only. Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// InsertTime: Output only. Timestamp when the resource was created or
	// acquired, in RFC3339 text format .
	InsertTime string `json:"insertTime,omitempty"`

	// Manifest: Output only. URL of the manifest representing the current
	// configuration of this resource.
	Manifest string `json:"manifest,omitempty"`

	// Name: Output only. The name of the resource as it appears in the YAML
	// config.
	Name string `json:"name,omitempty"`

	// Properties: Output only. The current properties of the resource
	// before any references have been filled in. Returned as serialized
	// YAML.
	Properties string `json:"properties,omitempty"`

	// Type: Output only. The type of the resource, for example
	// compute.v1.instance, or cloudfunctions.v1beta1.function.
	Type string `json:"type,omitempty"`

	// Update: Output only. If Deployment Manager is currently updating or
	// previewing an update to this resource, the updated configuration
	// appears here.
	Update *ResourceUpdate `json:"update,omitempty"`

	// UpdateTime: Output only. Timestamp when the resource was updated, in
	// RFC3339 text format .
	UpdateTime string `json:"updateTime,omitempty"`

	// Url: Output only. The URL of the actual resource.
	Url string `json:"url,omitempty"`

	// Warnings: Output only. If warning messages are generated during
	// processing of this resource, this field will be populated.
	Warnings []*ResourceWarnings `json:"warnings,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AccessControl") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccessControl") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Resource) MarshalJSON() ([]byte, error) {
	type NoMethod Resource
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ResourceWarnings struct {
	// Code: [Output Only] A warning code, if applicable. For example,
	// Compute Engine returns NO_RESULTS_ON_PAGE if there are no results in
	// the response.
	Code string `json:"code,omitempty"`

	// Data: [Output Only] Metadata about this warning in key: value format.
	// For example:
	// "data": [ { "key": "scope", "value": "zones/us-east1-d" }
	Data []*ResourceWarningsData `json:"data,omitempty"`

	// Message: [Output Only] A human-readable description of the warning
	// code.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResourceWarnings) MarshalJSON() ([]byte, error) {
	type NoMethod ResourceWarnings
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ResourceWarningsData struct {
	// Key: [Output Only] A key that provides more detail on the warning
	// being returned. For example, for warnings where there are no results
	// in a list request for a particular zone, this key might be scope and
	// the key value might be the zone name. Other examples might be a key
	// indicating a deprecated resource and a suggested replacement, or a
	// warning about invalid network settings (for example, if an instance
	// attempts to perform IP forwarding but is not enabled for IP
	// forwarding).
	Key string `json:"key,omitempty"`

	// Value: [Output Only] A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResourceWarningsData) MarshalJSON() ([]byte, error) {
	type NoMethod ResourceWarningsData
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ResourceAccessControl: The access controls set on the resource.
type ResourceAccessControl struct {
	// GcpIamPolicy: The GCP IAM Policy to set on the resource.
	GcpIamPolicy string `json:"gcpIamPolicy,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GcpIamPolicy") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "GcpIamPolicy") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResourceAccessControl) MarshalJSON() ([]byte, error) {
	type NoMethod ResourceAccessControl
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ResourceUpdate struct {
	// AccessControl: The Access Control Policy to set on this resource
	// after updating the resource itself.
	AccessControl *ResourceAccessControl `json:"accessControl,omitempty"`

	// Error: Output only. If errors are generated during update of the
	// resource, this field will be populated.
	Error *ResourceUpdateError `json:"error,omitempty"`

	// FinalProperties: Output only. The expanded properties of the resource
	// with reference values expanded. Returned as serialized YAML.
	FinalProperties string `json:"finalProperties,omitempty"`

	// Intent: Output only. The intent of the resource: PREVIEW, UPDATE, or
	// CANCEL.
	Intent string `json:"intent,omitempty"`

	// Manifest: Output only. URL of the manifest representing the update
	// configuration of this resource.
	Manifest string `json:"manifest,omitempty"`

	// Properties: Output only. The set of updated properties for this
	// resource, before references are expanded. Returned as serialized
	// YAML.
	Properties string `json:"properties,omitempty"`

	// State: Output only. The state of the resource.
	State string `json:"state,omitempty"`

	// Warnings: Output only. If warning messages are generated during
	// processing of this resource, this field will be populated.
	Warnings []*ResourceUpdateWarnings `json:"warnings,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AccessControl") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccessControl") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResourceUpdate) MarshalJSON() ([]byte, error) {
	type NoMethod ResourceUpdate
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ResourceUpdateError: Output only. If errors are generated during
// update of the resource, this field will be populated.
type ResourceUpdateError struct {
	// Errors: [Output Only] The array of errors encountered while
	// processing this operation.
	Errors []*ResourceUpdateErrorErrors `json:"errors,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Errors") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Errors") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResourceUpdateError) MarshalJSON() ([]byte, error) {
	type NoMethod ResourceUpdateError
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ResourceUpdateErrorErrors struct {
	// Code: [Output Only] The error type identifier for this error.
	Code string `json:"code,omitempty"`

	// Location: [Output Only] Indicates the field in the request that
	// caused the error. This property is optional.
	Location string `json:"location,omitempty"`

	// Message: [Output Only] An optional, human-readable error message.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResourceUpdateErrorErrors) MarshalJSON() ([]byte, error) {
	type NoMethod ResourceUpdateErrorErrors
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ResourceUpdateWarnings struct {
	// Code: [Output Only] A warning code, if applicable. For example,
	// Compute Engine returns NO_RESULTS_ON_PAGE if there are no results in
	// the response.
	Code string `json:"code,omitempty"`

	// Data: [Output Only] Metadata about this warning in key: value format.
	// For example:
	// "data": [ { "key": "scope", "value": "zones/us-east1-d" }
	Data []*ResourceUpdateWarningsData `json:"data,omitempty"`

	// Message: [Output Only] A human-readable description of the warning
	// code.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResourceUpdateWarnings) MarshalJSON() ([]byte, error) {
	type NoMethod ResourceUpdateWarnings
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ResourceUpdateWarningsData struct {
	// Key: [Output Only] A key that provides more detail on the warning
	// being returned. For example, for warnings where there are no results
	// in a list request for a particular zone, this key might be scope and
	// the key value might be the zone name. Other examples might be a key
	// indicating a deprecated resource and a suggested replacement, or a
	// warning about invalid network settings (for example, if an instance
	// attempts to perform IP forwarding but is not enabled for IP
	// forwarding).
	Key string `json:"key,omitempty"`

	// Value: [Output Only] A warning data value corresponding to the key.
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResourceUpdateWarningsData) MarshalJSON() ([]byte, error) {
	type NoMethod ResourceUpdateWarningsData
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ResourcesListResponse: A response containing a partial list of
// resources and a page token used to build the next request if the
// request has been truncated.
type ResourcesListResponse struct {
	// NextPageToken: A token used to continue a truncated list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Resources: Resources contained in this list response.
	Resources []*Resource `json:"resources,omitempty"`

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

func (s *ResourcesListResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ResourcesListResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Rule: A rule to be applied in a Policy.
type Rule struct {
	// Action: Required
	Action string `json:"action,omitempty"`

	// Conditions: Additional restrictions that must be met. All conditions
	// must pass for the rule to match.
	Conditions []*Condition `json:"conditions,omitempty"`

	// Description: Human-readable description of the rule.
	Description string `json:"description,omitempty"`

	// Ins: If one or more 'in' clauses are specified, the rule matches if
	// the PRINCIPAL/AUTHORITY_SELECTOR is in at least one of these entries.
	Ins []string `json:"ins,omitempty"`

	// LogConfigs: The config returned to callers of
	// tech.iam.IAM.CheckPolicy for any entries that match the LOG action.
	LogConfigs []*LogConfig `json:"logConfigs,omitempty"`

	// NotIns: If one or more 'not_in' clauses are specified, the rule
	// matches if the PRINCIPAL/AUTHORITY_SELECTOR is in none of the
	// entries.
	NotIns []string `json:"notIns,omitempty"`

	// Permissions: A permission is a string of form '..' (e.g.,
	// 'storage.buckets.list'). A value of '*' matches all permissions, and
	// a verb part of '*' (e.g., 'storage.buckets.*') matches all verbs.
	Permissions []string `json:"permissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Action") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Action") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Rule) MarshalJSON() ([]byte, error) {
	type NoMethod Rule
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ServiceAccount: Service Account used as a credential.
type ServiceAccount struct {
	// Email: The IAM service account email address like
	// test@myproject.iam.gserviceaccount.com
	Email string `json:"email,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Email") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Email") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ServiceAccount) MarshalJSON() ([]byte, error) {
	type NoMethod ServiceAccount
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type TargetConfiguration struct {
	// Config: The configuration to use for this deployment.
	Config *ConfigFile `json:"config,omitempty"`

	// Imports: Specifies any files to import for this configuration. This
	// can be used to import templates or other files. For example, you
	// might import a text file in order to use the file in a template.
	Imports []*ImportFile `json:"imports,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Config") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Config") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TargetConfiguration) MarshalJSON() ([]byte, error) {
	type NoMethod TargetConfiguration
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TemplateContents: Files that make up the template contents of a
// template type.
type TemplateContents struct {
	// Imports: Import files referenced by the main template.
	Imports []*ImportFile `json:"imports,omitempty"`

	// Interpreter: Which interpreter (python or jinja) should be used
	// during expansion.
	Interpreter string `json:"interpreter,omitempty"`

	// MainTemplate: The filename of the mainTemplate
	MainTemplate string `json:"mainTemplate,omitempty"`

	// Schema: The contents of the template schema.
	Schema string `json:"schema,omitempty"`

	// Template: The contents of the main template file.
	Template string `json:"template,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Imports") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Imports") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TemplateContents) MarshalJSON() ([]byte, error) {
	type NoMethod TemplateContents
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type TestPermissionsRequest struct {
	// Permissions: The set of permissions to check for the 'resource'.
	// Permissions with wildcards (such as '*' or 'storage.*') are not
	// allowed.
	Permissions []string `json:"permissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Permissions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TestPermissionsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod TestPermissionsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type TestPermissionsResponse struct {
	// Permissions: A subset of `TestPermissionsRequest.permissions` that
	// the caller is allowed.
	Permissions []string `json:"permissions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Permissions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TestPermissionsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod TestPermissionsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Type: A resource type supported by Deployment Manager.
type Type struct {
	// Base: Base Type (configurable service) that backs this Type.
	Base *BaseType `json:"base,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// Id: Output only. Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// InsertTime: Output only. Timestamp when the type was created, in
	// RFC3339 text format.
	InsertTime string `json:"insertTime,omitempty"`

	// Labels: Map of labels; provided by the client when the resource is
	// created or updated. Specifically: Label keys must be between 1 and 63
	// characters long and must conform to the following regular expression:
	// [a-z]([-a-z0-9]*[a-z0-9])? Label values must be between 0 and 63
	// characters long and must conform to the regular expression
	// ([a-z]([-a-z0-9]*[a-z0-9])?)?
	Labels []*TypeLabelEntry `json:"labels,omitempty"`

	// Name: Name of the type.
	Name string `json:"name,omitempty"`

	// Operation: Output only. The Operation that most recently ran, or is
	// currently running, on this type.
	Operation *Operation `json:"operation,omitempty"`

	// SelfLink: Output only. Self link for the type.
	SelfLink string `json:"selfLink,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Base") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Base") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Type) MarshalJSON() ([]byte, error) {
	type NoMethod Type
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TypeInfo: Contains detailed information about a composite type, base
// type, or base type with specific collection.
type TypeInfo struct {
	// Description: The description of the type.
	Description string `json:"description,omitempty"`

	// DocumentationLink: For swagger 2.0 externalDocs field will be used.
	// For swagger 1.2 this field will be empty.
	DocumentationLink string `json:"documentationLink,omitempty"`

	// Kind: Output only. Type of the output. Always
	// deploymentManager#TypeInfo for TypeInfo.
	Kind string `json:"kind,omitempty"`

	// Name: The base type or composite type name.
	Name string `json:"name,omitempty"`

	// Schema: For base types with a collection, we return a schema and
	// documentation link For template types, we return only a schema
	Schema *TypeInfoSchemaInfo `json:"schema,omitempty"`

	// SelfLink: Output only. Server-defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Title: The title on the API descriptor URL provided.
	Title string `json:"title,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TypeInfo) MarshalJSON() ([]byte, error) {
	type NoMethod TypeInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type TypeInfoSchemaInfo struct {
	// Input: The properties that this composite type or base type
	// collection accept as input, represented as a json blob, format is:
	// JSON Schema Draft V4
	Input string `json:"input,omitempty"`

	// Output: The properties that this composite type or base type
	// collection exposes as output, these properties can be used for
	// references, represented as json blob, format is: JSON Schema Draft V4
	Output string `json:"output,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Input") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Input") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TypeInfoSchemaInfo) MarshalJSON() ([]byte, error) {
	type NoMethod TypeInfoSchemaInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type TypeLabelEntry struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TypeLabelEntry) MarshalJSON() ([]byte, error) {
	type NoMethod TypeLabelEntry
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TypeProvider: A type provider that describes a service-backed Type.
type TypeProvider struct {
	// CollectionOverrides: Allows resource handling overrides for specific
	// collections
	CollectionOverrides []*CollectionOverride `json:"collectionOverrides,omitempty"`

	// Credential: Credential used when interacting with this type.
	Credential *Credential `json:"credential,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// DescriptorUrl: Descriptor Url for the this type provider.
	DescriptorUrl string `json:"descriptorUrl,omitempty"`

	// Id: Output only. Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// InsertTime: Output only. Timestamp when the type provider was
	// created, in RFC3339 text format.
	InsertTime string `json:"insertTime,omitempty"`

	// Labels: Map of labels; provided by the client when the resource is
	// created or updated. Specifically: Label keys must be between 1 and 63
	// characters long and must conform to the following regular expression:
	// [a-z]([-a-z0-9]*[a-z0-9])? Label values must be between 0 and 63
	// characters long and must conform to the regular expression
	// ([a-z]([-a-z0-9]*[a-z0-9])?)?
	Labels []*TypeProviderLabelEntry `json:"labels,omitempty"`

	// Name: Name of the resource; provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name string `json:"name,omitempty"`

	// Operation: Output only. The Operation that most recently ran, or is
	// currently running, on this type provider.
	Operation *Operation `json:"operation,omitempty"`

	// Options: Options to apply when handling any resources in this
	// service.
	Options *Options `json:"options,omitempty"`

	// SelfLink: Output only. Self link for the type provider.
	SelfLink string `json:"selfLink,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CollectionOverrides")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CollectionOverrides") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *TypeProvider) MarshalJSON() ([]byte, error) {
	type NoMethod TypeProvider
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type TypeProviderLabelEntry struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TypeProviderLabelEntry) MarshalJSON() ([]byte, error) {
	type NoMethod TypeProviderLabelEntry
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TypeProvidersListResponse: A response that returns all Type Providers
// supported by Deployment Manager
type TypeProvidersListResponse struct {
	// NextPageToken: A token used to continue a truncated list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TypeProviders: Output only. A list of resource type providers
	// supported by Deployment Manager.
	TypeProviders []*TypeProvider `json:"typeProviders,omitempty"`

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

func (s *TypeProvidersListResponse) MarshalJSON() ([]byte, error) {
	type NoMethod TypeProvidersListResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type TypeProvidersListTypesResponse struct {
	// NextPageToken: A token used to continue a truncated list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Types: Output only. A list of resource type info.
	Types []*TypeInfo `json:"types,omitempty"`

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

func (s *TypeProvidersListTypesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod TypeProvidersListTypesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TypesListResponse: A response that returns all Types supported by
// Deployment Manager
type TypesListResponse struct {
	// NextPageToken: A token used to continue a truncated list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Types: Output only. A list of resource types supported by Deployment
	// Manager.
	Types []*Type `json:"types,omitempty"`

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

func (s *TypesListResponse) MarshalJSON() ([]byte, error) {
	type NoMethod TypesListResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ValidationOptions: Options for how to validate and process properties
// on a resource.
type ValidationOptions struct {
	// SchemaValidation: Customize how deployment manager will validate the
	// resource against schema errors.
	SchemaValidation string `json:"schemaValidation,omitempty"`

	// UndeclaredProperties: Specify what to do with extra properties when
	// executing a request.
	UndeclaredProperties string `json:"undeclaredProperties,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SchemaValidation") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SchemaValidation") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ValidationOptions) MarshalJSON() ([]byte, error) {
	type NoMethod ValidationOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "deploymentmanager.compositeTypes.delete":

type CompositeTypesDeleteCall struct {
	s             *Service
	project       string
	compositeType string
	urlParams_    gensupport.URLParams
	ctx_          context.Context
	header_       http.Header
}

// Delete: Deletes a composite type.
func (r *CompositeTypesService) Delete(project string, compositeType string) *CompositeTypesDeleteCall {
	c := &CompositeTypesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.compositeType = compositeType
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CompositeTypesDeleteCall) Fields(s ...googleapi.Field) *CompositeTypesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CompositeTypesDeleteCall) Context(ctx context.Context) *CompositeTypesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CompositeTypesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CompositeTypesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/compositeTypes/{compositeType}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":       c.project,
		"compositeType": c.compositeType,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.compositeTypes.delete"), c.s.client, req)
}

// Do executes the "deploymentmanager.compositeTypes.delete" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *CompositeTypesDeleteCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Deletes a composite type.",
	//   "httpMethod": "DELETE",
	//   "id": "deploymentmanager.compositeTypes.delete",
	//   "parameterOrder": [
	//     "project",
	//     "compositeType"
	//   ],
	//   "parameters": {
	//     "compositeType": {
	//       "description": "The name of the type for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9_.]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/compositeTypes/{compositeType}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.compositeTypes.get":

type CompositeTypesGetCall struct {
	s             *Service
	project       string
	compositeType string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// Get: Gets information about a specific composite type.
func (r *CompositeTypesService) Get(project string, compositeType string) *CompositeTypesGetCall {
	c := &CompositeTypesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.compositeType = compositeType
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CompositeTypesGetCall) Fields(s ...googleapi.Field) *CompositeTypesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CompositeTypesGetCall) IfNoneMatch(entityTag string) *CompositeTypesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CompositeTypesGetCall) Context(ctx context.Context) *CompositeTypesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CompositeTypesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CompositeTypesGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/compositeTypes/{compositeType}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":       c.project,
		"compositeType": c.compositeType,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.compositeTypes.get"), c.s.client, req)
}

// Do executes the "deploymentmanager.compositeTypes.get" call.
// Exactly one of *CompositeType or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *CompositeType.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CompositeTypesGetCall) Do(opts ...googleapi.CallOption) (*CompositeType, error) {
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
	ret := &CompositeType{
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
	//   "description": "Gets information about a specific composite type.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.compositeTypes.get",
	//   "parameterOrder": [
	//     "project",
	//     "compositeType"
	//   ],
	//   "parameters": {
	//     "compositeType": {
	//       "description": "The name of the composite type for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9_.]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/compositeTypes/{compositeType}",
	//   "response": {
	//     "$ref": "CompositeType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.compositeTypes.insert":

type CompositeTypesInsertCall struct {
	s             *Service
	project       string
	compositetype *CompositeType
	urlParams_    gensupport.URLParams
	ctx_          context.Context
	header_       http.Header
}

// Insert: Creates a composite type.
func (r *CompositeTypesService) Insert(project string, compositetype *CompositeType) *CompositeTypesInsertCall {
	c := &CompositeTypesInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.compositetype = compositetype
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CompositeTypesInsertCall) Fields(s ...googleapi.Field) *CompositeTypesInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CompositeTypesInsertCall) Context(ctx context.Context) *CompositeTypesInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CompositeTypesInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CompositeTypesInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.compositetype)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/compositeTypes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.compositeTypes.insert"), c.s.client, req)
}

// Do executes the "deploymentmanager.compositeTypes.insert" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *CompositeTypesInsertCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Creates a composite type.",
	//   "httpMethod": "POST",
	//   "id": "deploymentmanager.compositeTypes.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/compositeTypes",
	//   "request": {
	//     "$ref": "CompositeType"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.compositeTypes.list":

type CompositeTypesListCall struct {
	s            *Service
	project      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all composite types for Deployment Manager.
func (r *CompositeTypesService) List(project string) *CompositeTypesListCall {
	c := &CompositeTypesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	return c
}

// Filter sets the optional parameter "filter": A filter expression that
// filters resources listed in the response. The expression must specify
// the field name, a comparison operator, and the value that you want to
// use for filtering. The value must be a string, a number, or a
// boolean. The comparison operator must be either =, !=, >, or <.
//
// For example, if you are filtering Compute Engine instances, you can
// exclude instances named example-instance by specifying name !=
// example-instance.
//
// You can also filter nested fields. For example, you could specify
// scheduling.automaticRestart = false to include instances only if they
// are not scheduled for automatic restarts. You can use filtering on
// nested fields to filter based on resource labels.
//
// To filter on multiple expressions, provide each separate expression
// within parentheses. For example, (scheduling.automaticRestart = true)
// (cpuPlatform = "Intel Skylake"). By default, each expression is an
// AND expression. However, you can include AND and OR expressions
// explicitly. For example, (cpuPlatform = "Intel Skylake") OR
// (cpuPlatform = "Intel Broadwell") AND (scheduling.automaticRestart =
// true).
func (c *CompositeTypesListCall) Filter(filter string) *CompositeTypesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of results per page that should be returned. If the number of
// available results is larger than maxResults, Compute Engine returns a
// nextPageToken that can be used to get the next page of results in
// subsequent list requests. Acceptable values are 0 to 500, inclusive.
// (Default: 500)
func (c *CompositeTypesListCall) MaxResults(maxResults int64) *CompositeTypesListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// OrderBy sets the optional parameter "orderBy": Sorts list results by
// a certain order. By default, results are returned in alphanumerical
// order based on the resource name.
//
// You can also sort results in descending order based on the creation
// timestamp using orderBy="creationTimestamp desc". This sorts results
// based on the creationTimestamp field in reverse chronological order
// (newest result first). Use this to sort resources like operations so
// that the newest operation is returned first.
//
// Currently, only sorting by name or creationTimestamp desc is
// supported.
func (c *CompositeTypesListCall) OrderBy(orderBy string) *CompositeTypesListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a page
// token to use. Set pageToken to the nextPageToken returned by a
// previous list request to get the next page of results.
func (c *CompositeTypesListCall) PageToken(pageToken string) *CompositeTypesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CompositeTypesListCall) Fields(s ...googleapi.Field) *CompositeTypesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CompositeTypesListCall) IfNoneMatch(entityTag string) *CompositeTypesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CompositeTypesListCall) Context(ctx context.Context) *CompositeTypesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CompositeTypesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CompositeTypesListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/compositeTypes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.compositeTypes.list"), c.s.client, req)
}

// Do executes the "deploymentmanager.compositeTypes.list" call.
// Exactly one of *CompositeTypesListResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *CompositeTypesListResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CompositeTypesListCall) Do(opts ...googleapi.CallOption) (*CompositeTypesListResponse, error) {
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
	ret := &CompositeTypesListResponse{
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
	//   "description": "Lists all composite types for Deployment Manager.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.compositeTypes.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "A filter expression that filters resources listed in the response. The expression must specify the field name, a comparison operator, and the value that you want to use for filtering. The value must be a string, a number, or a boolean. The comparison operator must be either =, !=, \u003e, or \u003c.\n\nFor example, if you are filtering Compute Engine instances, you can exclude instances named example-instance by specifying name != example-instance.\n\nYou can also filter nested fields. For example, you could specify scheduling.automaticRestart = false to include instances only if they are not scheduled for automatic restarts. You can use filtering on nested fields to filter based on resource labels.\n\nTo filter on multiple expressions, provide each separate expression within parentheses. For example, (scheduling.automaticRestart = true) (cpuPlatform = \"Intel Skylake\"). By default, each expression is an AND expression. However, you can include AND and OR expressions explicitly. For example, (cpuPlatform = \"Intel Skylake\") OR (cpuPlatform = \"Intel Broadwell\") AND (scheduling.automaticRestart = true).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "The maximum number of results per page that should be returned. If the number of available results is larger than maxResults, Compute Engine returns a nextPageToken that can be used to get the next page of results in subsequent list requests. Acceptable values are 0 to 500, inclusive. (Default: 500)",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Sorts list results by a certain order. By default, results are returned in alphanumerical order based on the resource name.\n\nYou can also sort results in descending order based on the creation timestamp using orderBy=\"creationTimestamp desc\". This sorts results based on the creationTimestamp field in reverse chronological order (newest result first). Use this to sort resources like operations so that the newest operation is returned first.\n\nCurrently, only sorting by name or creationTimestamp desc is supported.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a page token to use. Set pageToken to the nextPageToken returned by a previous list request to get the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/compositeTypes",
	//   "response": {
	//     "$ref": "CompositeTypesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *CompositeTypesListCall) Pages(ctx context.Context, f func(*CompositeTypesListResponse) error) error {
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

// method id "deploymentmanager.compositeTypes.patch":

type CompositeTypesPatchCall struct {
	s             *Service
	project       string
	compositeType string
	compositetype *CompositeType
	urlParams_    gensupport.URLParams
	ctx_          context.Context
	header_       http.Header
}

// Patch: Updates a composite type. This method supports patch
// semantics.
func (r *CompositeTypesService) Patch(project string, compositeType string, compositetype *CompositeType) *CompositeTypesPatchCall {
	c := &CompositeTypesPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.compositeType = compositeType
	c.compositetype = compositetype
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CompositeTypesPatchCall) Fields(s ...googleapi.Field) *CompositeTypesPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CompositeTypesPatchCall) Context(ctx context.Context) *CompositeTypesPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CompositeTypesPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CompositeTypesPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.compositetype)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/compositeTypes/{compositeType}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":       c.project,
		"compositeType": c.compositeType,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.compositeTypes.patch"), c.s.client, req)
}

// Do executes the "deploymentmanager.compositeTypes.patch" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *CompositeTypesPatchCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Updates a composite type. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "deploymentmanager.compositeTypes.patch",
	//   "parameterOrder": [
	//     "project",
	//     "compositeType"
	//   ],
	//   "parameters": {
	//     "compositeType": {
	//       "description": "The name of the composite type for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9_.]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/compositeTypes/{compositeType}",
	//   "request": {
	//     "$ref": "CompositeType"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.compositeTypes.update":

type CompositeTypesUpdateCall struct {
	s             *Service
	project       string
	compositeType string
	compositetype *CompositeType
	urlParams_    gensupport.URLParams
	ctx_          context.Context
	header_       http.Header
}

// Update: Updates a composite type.
func (r *CompositeTypesService) Update(project string, compositeType string, compositetype *CompositeType) *CompositeTypesUpdateCall {
	c := &CompositeTypesUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.compositeType = compositeType
	c.compositetype = compositetype
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CompositeTypesUpdateCall) Fields(s ...googleapi.Field) *CompositeTypesUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CompositeTypesUpdateCall) Context(ctx context.Context) *CompositeTypesUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CompositeTypesUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CompositeTypesUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.compositetype)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/compositeTypes/{compositeType}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":       c.project,
		"compositeType": c.compositeType,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.compositeTypes.update"), c.s.client, req)
}

// Do executes the "deploymentmanager.compositeTypes.update" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *CompositeTypesUpdateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Updates a composite type.",
	//   "httpMethod": "PUT",
	//   "id": "deploymentmanager.compositeTypes.update",
	//   "parameterOrder": [
	//     "project",
	//     "compositeType"
	//   ],
	//   "parameters": {
	//     "compositeType": {
	//       "description": "The name of the composite type for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9_.]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/compositeTypes/{compositeType}",
	//   "request": {
	//     "$ref": "CompositeType"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.deployments.cancelPreview":

type DeploymentsCancelPreviewCall struct {
	s                               *Service
	project                         string
	deployment                      string
	deploymentscancelpreviewrequest *DeploymentsCancelPreviewRequest
	urlParams_                      gensupport.URLParams
	ctx_                            context.Context
	header_                         http.Header
}

// CancelPreview: Cancels and removes the preview currently associated
// with the deployment.
func (r *DeploymentsService) CancelPreview(project string, deployment string, deploymentscancelpreviewrequest *DeploymentsCancelPreviewRequest) *DeploymentsCancelPreviewCall {
	c := &DeploymentsCancelPreviewCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.deployment = deployment
	c.deploymentscancelpreviewrequest = deploymentscancelpreviewrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsCancelPreviewCall) Fields(s ...googleapi.Field) *DeploymentsCancelPreviewCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DeploymentsCancelPreviewCall) Context(ctx context.Context) *DeploymentsCancelPreviewCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DeploymentsCancelPreviewCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DeploymentsCancelPreviewCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.deploymentscancelpreviewrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{deployment}/cancelPreview")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.deployments.cancelPreview"), c.s.client, req)
}

// Do executes the "deploymentmanager.deployments.cancelPreview" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *DeploymentsCancelPreviewCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Cancels and removes the preview currently associated with the deployment.",
	//   "httpMethod": "POST",
	//   "id": "deploymentmanager.deployments.cancelPreview",
	//   "parameterOrder": [
	//     "project",
	//     "deployment"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}/cancelPreview",
	//   "request": {
	//     "$ref": "DeploymentsCancelPreviewRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.deployments.delete":

type DeploymentsDeleteCall struct {
	s          *Service
	project    string
	deployment string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a deployment and all of the resources in the
// deployment.
func (r *DeploymentsService) Delete(project string, deployment string) *DeploymentsDeleteCall {
	c := &DeploymentsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.deployment = deployment
	return c
}

// DeletePolicy sets the optional parameter "deletePolicy": Sets the
// policy to use for deleting resources.
//
// Possible values:
//   "ABANDON"
//   "DELETE" (default)
func (c *DeploymentsDeleteCall) DeletePolicy(deletePolicy string) *DeploymentsDeleteCall {
	c.urlParams_.Set("deletePolicy", deletePolicy)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsDeleteCall) Fields(s ...googleapi.Field) *DeploymentsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DeploymentsDeleteCall) Context(ctx context.Context) *DeploymentsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DeploymentsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DeploymentsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{deployment}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.deployments.delete"), c.s.client, req)
}

// Do executes the "deploymentmanager.deployments.delete" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *DeploymentsDeleteCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Deletes a deployment and all of the resources in the deployment.",
	//   "httpMethod": "DELETE",
	//   "id": "deploymentmanager.deployments.delete",
	//   "parameterOrder": [
	//     "project",
	//     "deployment"
	//   ],
	//   "parameters": {
	//     "deletePolicy": {
	//       "default": "DELETE",
	//       "description": "Sets the policy to use for deleting resources.",
	//       "enum": [
	//         "ABANDON",
	//         "DELETE"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "deployment": {
	//       "description": "The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.deployments.get":

type DeploymentsGetCall struct {
	s            *Service
	project      string
	deployment   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets information about a specific deployment.
func (r *DeploymentsService) Get(project string, deployment string) *DeploymentsGetCall {
	c := &DeploymentsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.deployment = deployment
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsGetCall) Fields(s ...googleapi.Field) *DeploymentsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *DeploymentsGetCall) IfNoneMatch(entityTag string) *DeploymentsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DeploymentsGetCall) Context(ctx context.Context) *DeploymentsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DeploymentsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DeploymentsGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{deployment}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.deployments.get"), c.s.client, req)
}

// Do executes the "deploymentmanager.deployments.get" call.
// Exactly one of *Deployment or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Deployment.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *DeploymentsGetCall) Do(opts ...googleapi.CallOption) (*Deployment, error) {
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
	ret := &Deployment{
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
	//   "description": "Gets information about a specific deployment.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.deployments.get",
	//   "parameterOrder": [
	//     "project",
	//     "deployment"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}",
	//   "response": {
	//     "$ref": "Deployment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.deployments.getIamPolicy":

type DeploymentsGetIamPolicyCall struct {
	s            *Service
	project      string
	resource     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetIamPolicy: Gets the access control policy for a resource. May be
// empty if no such policy or resource exists.
func (r *DeploymentsService) GetIamPolicy(project string, resource string) *DeploymentsGetIamPolicyCall {
	c := &DeploymentsGetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.resource = resource
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsGetIamPolicyCall) Fields(s ...googleapi.Field) *DeploymentsGetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *DeploymentsGetIamPolicyCall) IfNoneMatch(entityTag string) *DeploymentsGetIamPolicyCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DeploymentsGetIamPolicyCall) Context(ctx context.Context) *DeploymentsGetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DeploymentsGetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DeploymentsGetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{resource}/getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"resource": c.resource,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.deployments.getIamPolicy"), c.s.client, req)
}

// Do executes the "deploymentmanager.deployments.getIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *DeploymentsGetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
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
	ret := &Policy{
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
	//   "description": "Gets the access control policy for a resource. May be empty if no such policy or resource exists.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.deployments.getIamPolicy",
	//   "parameterOrder": [
	//     "project",
	//     "resource"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z0-9](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resource": {
	//       "description": "Name of the resource for this request.",
	//       "location": "path",
	//       "pattern": "[a-z0-9](?:[-a-z0-9_]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{resource}/getIamPolicy",
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.deployments.insert":

type DeploymentsInsertCall struct {
	s          *Service
	project    string
	deployment *Deployment
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Insert: Creates a deployment and all of the resources described by
// the deployment manifest.
func (r *DeploymentsService) Insert(project string, deployment *Deployment) *DeploymentsInsertCall {
	c := &DeploymentsInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.deployment = deployment
	return c
}

// CreatePolicy sets the optional parameter "createPolicy":
//
// Possible values:
//   "ACQUIRE"
//   "CREATE"
//   "CREATE_OR_ACQUIRE" (default)
func (c *DeploymentsInsertCall) CreatePolicy(createPolicy string) *DeploymentsInsertCall {
	c.urlParams_.Set("createPolicy", createPolicy)
	return c
}

// Preview sets the optional parameter "preview": If set to true,
// creates a deployment and creates "shell" resources but does not
// actually instantiate these resources. This allows you to preview what
// your deployment looks like. After previewing a deployment, you can
// deploy your resources by making a request with the update() method or
// you can use the cancelPreview() method to cancel the preview
// altogether. Note that the deployment will still exist after you
// cancel the preview and you must separately delete this deployment if
// you want to remove it.
func (c *DeploymentsInsertCall) Preview(preview bool) *DeploymentsInsertCall {
	c.urlParams_.Set("preview", fmt.Sprint(preview))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsInsertCall) Fields(s ...googleapi.Field) *DeploymentsInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DeploymentsInsertCall) Context(ctx context.Context) *DeploymentsInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DeploymentsInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DeploymentsInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.deployment)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.deployments.insert"), c.s.client, req)
}

// Do executes the "deploymentmanager.deployments.insert" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *DeploymentsInsertCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Creates a deployment and all of the resources described by the deployment manifest.",
	//   "httpMethod": "POST",
	//   "id": "deploymentmanager.deployments.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "createPolicy": {
	//       "default": "CREATE_OR_ACQUIRE",
	//       "description": "",
	//       "enum": [
	//         "ACQUIRE",
	//         "CREATE",
	//         "CREATE_OR_ACQUIRE"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "preview": {
	//       "description": "If set to true, creates a deployment and creates \"shell\" resources but does not actually instantiate these resources. This allows you to preview what your deployment looks like. After previewing a deployment, you can deploy your resources by making a request with the update() method or you can use the cancelPreview() method to cancel the preview altogether. Note that the deployment will still exist after you cancel the preview and you must separately delete this deployment if you want to remove it.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments",
	//   "request": {
	//     "$ref": "Deployment"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.deployments.list":

type DeploymentsListCall struct {
	s            *Service
	project      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all deployments for a given project.
func (r *DeploymentsService) List(project string) *DeploymentsListCall {
	c := &DeploymentsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	return c
}

// Filter sets the optional parameter "filter": A filter expression that
// filters resources listed in the response. The expression must specify
// the field name, a comparison operator, and the value that you want to
// use for filtering. The value must be a string, a number, or a
// boolean. The comparison operator must be either =, !=, >, or <.
//
// For example, if you are filtering Compute Engine instances, you can
// exclude instances named example-instance by specifying name !=
// example-instance.
//
// You can also filter nested fields. For example, you could specify
// scheduling.automaticRestart = false to include instances only if they
// are not scheduled for automatic restarts. You can use filtering on
// nested fields to filter based on resource labels.
//
// To filter on multiple expressions, provide each separate expression
// within parentheses. For example, (scheduling.automaticRestart = true)
// (cpuPlatform = "Intel Skylake"). By default, each expression is an
// AND expression. However, you can include AND and OR expressions
// explicitly. For example, (cpuPlatform = "Intel Skylake") OR
// (cpuPlatform = "Intel Broadwell") AND (scheduling.automaticRestart =
// true).
func (c *DeploymentsListCall) Filter(filter string) *DeploymentsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of results per page that should be returned. If the number of
// available results is larger than maxResults, Compute Engine returns a
// nextPageToken that can be used to get the next page of results in
// subsequent list requests. Acceptable values are 0 to 500, inclusive.
// (Default: 500)
func (c *DeploymentsListCall) MaxResults(maxResults int64) *DeploymentsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// OrderBy sets the optional parameter "orderBy": Sorts list results by
// a certain order. By default, results are returned in alphanumerical
// order based on the resource name.
//
// You can also sort results in descending order based on the creation
// timestamp using orderBy="creationTimestamp desc". This sorts results
// based on the creationTimestamp field in reverse chronological order
// (newest result first). Use this to sort resources like operations so
// that the newest operation is returned first.
//
// Currently, only sorting by name or creationTimestamp desc is
// supported.
func (c *DeploymentsListCall) OrderBy(orderBy string) *DeploymentsListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a page
// token to use. Set pageToken to the nextPageToken returned by a
// previous list request to get the next page of results.
func (c *DeploymentsListCall) PageToken(pageToken string) *DeploymentsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsListCall) Fields(s ...googleapi.Field) *DeploymentsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *DeploymentsListCall) IfNoneMatch(entityTag string) *DeploymentsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DeploymentsListCall) Context(ctx context.Context) *DeploymentsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DeploymentsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DeploymentsListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.deployments.list"), c.s.client, req)
}

// Do executes the "deploymentmanager.deployments.list" call.
// Exactly one of *DeploymentsListResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *DeploymentsListResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *DeploymentsListCall) Do(opts ...googleapi.CallOption) (*DeploymentsListResponse, error) {
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
	ret := &DeploymentsListResponse{
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
	//   "description": "Lists all deployments for a given project.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.deployments.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "A filter expression that filters resources listed in the response. The expression must specify the field name, a comparison operator, and the value that you want to use for filtering. The value must be a string, a number, or a boolean. The comparison operator must be either =, !=, \u003e, or \u003c.\n\nFor example, if you are filtering Compute Engine instances, you can exclude instances named example-instance by specifying name != example-instance.\n\nYou can also filter nested fields. For example, you could specify scheduling.automaticRestart = false to include instances only if they are not scheduled for automatic restarts. You can use filtering on nested fields to filter based on resource labels.\n\nTo filter on multiple expressions, provide each separate expression within parentheses. For example, (scheduling.automaticRestart = true) (cpuPlatform = \"Intel Skylake\"). By default, each expression is an AND expression. However, you can include AND and OR expressions explicitly. For example, (cpuPlatform = \"Intel Skylake\") OR (cpuPlatform = \"Intel Broadwell\") AND (scheduling.automaticRestart = true).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "The maximum number of results per page that should be returned. If the number of available results is larger than maxResults, Compute Engine returns a nextPageToken that can be used to get the next page of results in subsequent list requests. Acceptable values are 0 to 500, inclusive. (Default: 500)",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Sorts list results by a certain order. By default, results are returned in alphanumerical order based on the resource name.\n\nYou can also sort results in descending order based on the creation timestamp using orderBy=\"creationTimestamp desc\". This sorts results based on the creationTimestamp field in reverse chronological order (newest result first). Use this to sort resources like operations so that the newest operation is returned first.\n\nCurrently, only sorting by name or creationTimestamp desc is supported.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a page token to use. Set pageToken to the nextPageToken returned by a previous list request to get the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments",
	//   "response": {
	//     "$ref": "DeploymentsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *DeploymentsListCall) Pages(ctx context.Context, f func(*DeploymentsListResponse) error) error {
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

// method id "deploymentmanager.deployments.patch":

type DeploymentsPatchCall struct {
	s           *Service
	project     string
	deployment  string
	deployment2 *Deployment
	urlParams_  gensupport.URLParams
	ctx_        context.Context
	header_     http.Header
}

// Patch: Updates a deployment and all of the resources described by the
// deployment manifest. This method supports patch semantics.
func (r *DeploymentsService) Patch(project string, deployment string, deployment2 *Deployment) *DeploymentsPatchCall {
	c := &DeploymentsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.deployment = deployment
	c.deployment2 = deployment2
	return c
}

// CreatePolicy sets the optional parameter "createPolicy": Sets the
// policy to use for creating new resources.
//
// Possible values:
//   "ACQUIRE"
//   "CREATE"
//   "CREATE_OR_ACQUIRE" (default)
func (c *DeploymentsPatchCall) CreatePolicy(createPolicy string) *DeploymentsPatchCall {
	c.urlParams_.Set("createPolicy", createPolicy)
	return c
}

// DeletePolicy sets the optional parameter "deletePolicy": Sets the
// policy to use for deleting resources.
//
// Possible values:
//   "ABANDON"
//   "DELETE" (default)
func (c *DeploymentsPatchCall) DeletePolicy(deletePolicy string) *DeploymentsPatchCall {
	c.urlParams_.Set("deletePolicy", deletePolicy)
	return c
}

// Preview sets the optional parameter "preview": If set to true,
// updates the deployment and creates and updates the "shell" resources
// but does not actually alter or instantiate these resources. This
// allows you to preview what your deployment will look like. You can
// use this intent to preview how an update would affect your
// deployment. You must provide a target.config with a configuration if
// this is set to true. After previewing a deployment, you can deploy
// your resources by making a request with the update() or you can
// cancelPreview() to remove the preview altogether. Note that the
// deployment will still exist after you cancel the preview and you must
// separately delete this deployment if you want to remove it.
func (c *DeploymentsPatchCall) Preview(preview bool) *DeploymentsPatchCall {
	c.urlParams_.Set("preview", fmt.Sprint(preview))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsPatchCall) Fields(s ...googleapi.Field) *DeploymentsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DeploymentsPatchCall) Context(ctx context.Context) *DeploymentsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DeploymentsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DeploymentsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.deployment2)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{deployment}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.deployments.patch"), c.s.client, req)
}

// Do executes the "deploymentmanager.deployments.patch" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *DeploymentsPatchCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Updates a deployment and all of the resources described by the deployment manifest. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "deploymentmanager.deployments.patch",
	//   "parameterOrder": [
	//     "project",
	//     "deployment"
	//   ],
	//   "parameters": {
	//     "createPolicy": {
	//       "default": "CREATE_OR_ACQUIRE",
	//       "description": "Sets the policy to use for creating new resources.",
	//       "enum": [
	//         "ACQUIRE",
	//         "CREATE",
	//         "CREATE_OR_ACQUIRE"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "deletePolicy": {
	//       "default": "DELETE",
	//       "description": "Sets the policy to use for deleting resources.",
	//       "enum": [
	//         "ABANDON",
	//         "DELETE"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "deployment": {
	//       "description": "The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "preview": {
	//       "default": "false",
	//       "description": "If set to true, updates the deployment and creates and updates the \"shell\" resources but does not actually alter or instantiate these resources. This allows you to preview what your deployment will look like. You can use this intent to preview how an update would affect your deployment. You must provide a target.config with a configuration if this is set to true. After previewing a deployment, you can deploy your resources by making a request with the update() or you can cancelPreview() to remove the preview altogether. Note that the deployment will still exist after you cancel the preview and you must separately delete this deployment if you want to remove it.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}",
	//   "request": {
	//     "$ref": "Deployment"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.deployments.setIamPolicy":

type DeploymentsSetIamPolicyCall struct {
	s          *Service
	project    string
	resource   string
	policy     *Policy
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// SetIamPolicy: Sets the access control policy on the specified
// resource. Replaces any existing policy.
func (r *DeploymentsService) SetIamPolicy(project string, resource string, policy *Policy) *DeploymentsSetIamPolicyCall {
	c := &DeploymentsSetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.resource = resource
	c.policy = policy
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsSetIamPolicyCall) Fields(s ...googleapi.Field) *DeploymentsSetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DeploymentsSetIamPolicyCall) Context(ctx context.Context) *DeploymentsSetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DeploymentsSetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DeploymentsSetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.policy)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{resource}/setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"resource": c.resource,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.deployments.setIamPolicy"), c.s.client, req)
}

// Do executes the "deploymentmanager.deployments.setIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *DeploymentsSetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
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
	ret := &Policy{
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
	//   "description": "Sets the access control policy on the specified resource. Replaces any existing policy.",
	//   "httpMethod": "POST",
	//   "id": "deploymentmanager.deployments.setIamPolicy",
	//   "parameterOrder": [
	//     "project",
	//     "resource"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z0-9](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resource": {
	//       "description": "Name of the resource for this request.",
	//       "location": "path",
	//       "pattern": "[a-z0-9](?:[-a-z0-9_]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{resource}/setIamPolicy",
	//   "request": {
	//     "$ref": "Policy"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.deployments.stop":

type DeploymentsStopCall struct {
	s                      *Service
	project                string
	deployment             string
	deploymentsstoprequest *DeploymentsStopRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// Stop: Stops an ongoing operation. This does not roll back any work
// that has already been completed, but prevents any new work from being
// started.
func (r *DeploymentsService) Stop(project string, deployment string, deploymentsstoprequest *DeploymentsStopRequest) *DeploymentsStopCall {
	c := &DeploymentsStopCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.deployment = deployment
	c.deploymentsstoprequest = deploymentsstoprequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsStopCall) Fields(s ...googleapi.Field) *DeploymentsStopCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DeploymentsStopCall) Context(ctx context.Context) *DeploymentsStopCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DeploymentsStopCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DeploymentsStopCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.deploymentsstoprequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{deployment}/stop")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.deployments.stop"), c.s.client, req)
}

// Do executes the "deploymentmanager.deployments.stop" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *DeploymentsStopCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Stops an ongoing operation. This does not roll back any work that has already been completed, but prevents any new work from being started.",
	//   "httpMethod": "POST",
	//   "id": "deploymentmanager.deployments.stop",
	//   "parameterOrder": [
	//     "project",
	//     "deployment"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}/stop",
	//   "request": {
	//     "$ref": "DeploymentsStopRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.deployments.testIamPermissions":

type DeploymentsTestIamPermissionsCall struct {
	s                      *Service
	project                string
	resource               string
	testpermissionsrequest *TestPermissionsRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified resource.
func (r *DeploymentsService) TestIamPermissions(project string, resource string, testpermissionsrequest *TestPermissionsRequest) *DeploymentsTestIamPermissionsCall {
	c := &DeploymentsTestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.resource = resource
	c.testpermissionsrequest = testpermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsTestIamPermissionsCall) Fields(s ...googleapi.Field) *DeploymentsTestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DeploymentsTestIamPermissionsCall) Context(ctx context.Context) *DeploymentsTestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DeploymentsTestIamPermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DeploymentsTestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testpermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{resource}/testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"resource": c.resource,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.deployments.testIamPermissions"), c.s.client, req)
}

// Do executes the "deploymentmanager.deployments.testIamPermissions" call.
// Exactly one of *TestPermissionsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *TestPermissionsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *DeploymentsTestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*TestPermissionsResponse, error) {
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
	ret := &TestPermissionsResponse{
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
	//   "description": "Returns permissions that a caller has on the specified resource.",
	//   "httpMethod": "POST",
	//   "id": "deploymentmanager.deployments.testIamPermissions",
	//   "parameterOrder": [
	//     "project",
	//     "resource"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z0-9](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resource": {
	//       "description": "Name of the resource for this request.",
	//       "location": "path",
	//       "pattern": "(?:[-a-z0-9_]{0,62}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{resource}/testIamPermissions",
	//   "request": {
	//     "$ref": "TestPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "TestPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.deployments.update":

type DeploymentsUpdateCall struct {
	s           *Service
	project     string
	deployment  string
	deployment2 *Deployment
	urlParams_  gensupport.URLParams
	ctx_        context.Context
	header_     http.Header
}

// Update: Updates a deployment and all of the resources described by
// the deployment manifest.
func (r *DeploymentsService) Update(project string, deployment string, deployment2 *Deployment) *DeploymentsUpdateCall {
	c := &DeploymentsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.deployment = deployment
	c.deployment2 = deployment2
	return c
}

// CreatePolicy sets the optional parameter "createPolicy": Sets the
// policy to use for creating new resources.
//
// Possible values:
//   "ACQUIRE"
//   "CREATE"
//   "CREATE_OR_ACQUIRE" (default)
func (c *DeploymentsUpdateCall) CreatePolicy(createPolicy string) *DeploymentsUpdateCall {
	c.urlParams_.Set("createPolicy", createPolicy)
	return c
}

// DeletePolicy sets the optional parameter "deletePolicy": Sets the
// policy to use for deleting resources.
//
// Possible values:
//   "ABANDON"
//   "DELETE" (default)
func (c *DeploymentsUpdateCall) DeletePolicy(deletePolicy string) *DeploymentsUpdateCall {
	c.urlParams_.Set("deletePolicy", deletePolicy)
	return c
}

// Preview sets the optional parameter "preview": If set to true,
// updates the deployment and creates and updates the "shell" resources
// but does not actually alter or instantiate these resources. This
// allows you to preview what your deployment will look like. You can
// use this intent to preview how an update would affect your
// deployment. You must provide a target.config with a configuration if
// this is set to true. After previewing a deployment, you can deploy
// your resources by making a request with the update() or you can
// cancelPreview() to remove the preview altogether. Note that the
// deployment will still exist after you cancel the preview and you must
// separately delete this deployment if you want to remove it.
func (c *DeploymentsUpdateCall) Preview(preview bool) *DeploymentsUpdateCall {
	c.urlParams_.Set("preview", fmt.Sprint(preview))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeploymentsUpdateCall) Fields(s ...googleapi.Field) *DeploymentsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DeploymentsUpdateCall) Context(ctx context.Context) *DeploymentsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DeploymentsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DeploymentsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.deployment2)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{deployment}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.deployments.update"), c.s.client, req)
}

// Do executes the "deploymentmanager.deployments.update" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *DeploymentsUpdateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Updates a deployment and all of the resources described by the deployment manifest.",
	//   "httpMethod": "PUT",
	//   "id": "deploymentmanager.deployments.update",
	//   "parameterOrder": [
	//     "project",
	//     "deployment"
	//   ],
	//   "parameters": {
	//     "createPolicy": {
	//       "default": "CREATE_OR_ACQUIRE",
	//       "description": "Sets the policy to use for creating new resources.",
	//       "enum": [
	//         "ACQUIRE",
	//         "CREATE",
	//         "CREATE_OR_ACQUIRE"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "deletePolicy": {
	//       "default": "DELETE",
	//       "description": "Sets the policy to use for deleting resources.",
	//       "enum": [
	//         "ABANDON",
	//         "DELETE"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "deployment": {
	//       "description": "The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "preview": {
	//       "default": "false",
	//       "description": "If set to true, updates the deployment and creates and updates the \"shell\" resources but does not actually alter or instantiate these resources. This allows you to preview what your deployment will look like. You can use this intent to preview how an update would affect your deployment. You must provide a target.config with a configuration if this is set to true. After previewing a deployment, you can deploy your resources by making a request with the update() or you can cancelPreview() to remove the preview altogether. Note that the deployment will still exist after you cancel the preview and you must separately delete this deployment if you want to remove it.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}",
	//   "request": {
	//     "$ref": "Deployment"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.manifests.get":

type ManifestsGetCall struct {
	s            *Service
	project      string
	deployment   string
	manifest     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets information about a specific manifest.
func (r *ManifestsService) Get(project string, deployment string, manifest string) *ManifestsGetCall {
	c := &ManifestsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.deployment = deployment
	c.manifest = manifest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ManifestsGetCall) Fields(s ...googleapi.Field) *ManifestsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ManifestsGetCall) IfNoneMatch(entityTag string) *ManifestsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ManifestsGetCall) Context(ctx context.Context) *ManifestsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ManifestsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ManifestsGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{deployment}/manifests/{manifest}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
		"manifest":   c.manifest,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.manifests.get"), c.s.client, req)
}

// Do executes the "deploymentmanager.manifests.get" call.
// Exactly one of *Manifest or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Manifest.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ManifestsGetCall) Do(opts ...googleapi.CallOption) (*Manifest, error) {
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
	ret := &Manifest{
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
	//   "description": "Gets information about a specific manifest.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.manifests.get",
	//   "parameterOrder": [
	//     "project",
	//     "deployment",
	//     "manifest"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "manifest": {
	//       "description": "The name of the manifest for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}/manifests/{manifest}",
	//   "response": {
	//     "$ref": "Manifest"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.manifests.list":

type ManifestsListCall struct {
	s            *Service
	project      string
	deployment   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all manifests for a given deployment.
func (r *ManifestsService) List(project string, deployment string) *ManifestsListCall {
	c := &ManifestsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.deployment = deployment
	return c
}

// Filter sets the optional parameter "filter": A filter expression that
// filters resources listed in the response. The expression must specify
// the field name, a comparison operator, and the value that you want to
// use for filtering. The value must be a string, a number, or a
// boolean. The comparison operator must be either =, !=, >, or <.
//
// For example, if you are filtering Compute Engine instances, you can
// exclude instances named example-instance by specifying name !=
// example-instance.
//
// You can also filter nested fields. For example, you could specify
// scheduling.automaticRestart = false to include instances only if they
// are not scheduled for automatic restarts. You can use filtering on
// nested fields to filter based on resource labels.
//
// To filter on multiple expressions, provide each separate expression
// within parentheses. For example, (scheduling.automaticRestart = true)
// (cpuPlatform = "Intel Skylake"). By default, each expression is an
// AND expression. However, you can include AND and OR expressions
// explicitly. For example, (cpuPlatform = "Intel Skylake") OR
// (cpuPlatform = "Intel Broadwell") AND (scheduling.automaticRestart =
// true).
func (c *ManifestsListCall) Filter(filter string) *ManifestsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of results per page that should be returned. If the number of
// available results is larger than maxResults, Compute Engine returns a
// nextPageToken that can be used to get the next page of results in
// subsequent list requests. Acceptable values are 0 to 500, inclusive.
// (Default: 500)
func (c *ManifestsListCall) MaxResults(maxResults int64) *ManifestsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// OrderBy sets the optional parameter "orderBy": Sorts list results by
// a certain order. By default, results are returned in alphanumerical
// order based on the resource name.
//
// You can also sort results in descending order based on the creation
// timestamp using orderBy="creationTimestamp desc". This sorts results
// based on the creationTimestamp field in reverse chronological order
// (newest result first). Use this to sort resources like operations so
// that the newest operation is returned first.
//
// Currently, only sorting by name or creationTimestamp desc is
// supported.
func (c *ManifestsListCall) OrderBy(orderBy string) *ManifestsListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a page
// token to use. Set pageToken to the nextPageToken returned by a
// previous list request to get the next page of results.
func (c *ManifestsListCall) PageToken(pageToken string) *ManifestsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ManifestsListCall) Fields(s ...googleapi.Field) *ManifestsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ManifestsListCall) IfNoneMatch(entityTag string) *ManifestsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ManifestsListCall) Context(ctx context.Context) *ManifestsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ManifestsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ManifestsListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{deployment}/manifests")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.manifests.list"), c.s.client, req)
}

// Do executes the "deploymentmanager.manifests.list" call.
// Exactly one of *ManifestsListResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ManifestsListResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ManifestsListCall) Do(opts ...googleapi.CallOption) (*ManifestsListResponse, error) {
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
	ret := &ManifestsListResponse{
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
	//   "description": "Lists all manifests for a given deployment.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.manifests.list",
	//   "parameterOrder": [
	//     "project",
	//     "deployment"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "A filter expression that filters resources listed in the response. The expression must specify the field name, a comparison operator, and the value that you want to use for filtering. The value must be a string, a number, or a boolean. The comparison operator must be either =, !=, \u003e, or \u003c.\n\nFor example, if you are filtering Compute Engine instances, you can exclude instances named example-instance by specifying name != example-instance.\n\nYou can also filter nested fields. For example, you could specify scheduling.automaticRestart = false to include instances only if they are not scheduled for automatic restarts. You can use filtering on nested fields to filter based on resource labels.\n\nTo filter on multiple expressions, provide each separate expression within parentheses. For example, (scheduling.automaticRestart = true) (cpuPlatform = \"Intel Skylake\"). By default, each expression is an AND expression. However, you can include AND and OR expressions explicitly. For example, (cpuPlatform = \"Intel Skylake\") OR (cpuPlatform = \"Intel Broadwell\") AND (scheduling.automaticRestart = true).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "The maximum number of results per page that should be returned. If the number of available results is larger than maxResults, Compute Engine returns a nextPageToken that can be used to get the next page of results in subsequent list requests. Acceptable values are 0 to 500, inclusive. (Default: 500)",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Sorts list results by a certain order. By default, results are returned in alphanumerical order based on the resource name.\n\nYou can also sort results in descending order based on the creation timestamp using orderBy=\"creationTimestamp desc\". This sorts results based on the creationTimestamp field in reverse chronological order (newest result first). Use this to sort resources like operations so that the newest operation is returned first.\n\nCurrently, only sorting by name or creationTimestamp desc is supported.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a page token to use. Set pageToken to the nextPageToken returned by a previous list request to get the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}/manifests",
	//   "response": {
	//     "$ref": "ManifestsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ManifestsListCall) Pages(ctx context.Context, f func(*ManifestsListResponse) error) error {
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

// method id "deploymentmanager.operations.get":

type OperationsGetCall struct {
	s            *Service
	project      string
	operation    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets information about a specific operation.
func (r *OperationsService) Get(project string, operation string) *OperationsGetCall {
	c := &OperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.operation = operation
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsGetCall) Fields(s ...googleapi.Field) *OperationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OperationsGetCall) IfNoneMatch(entityTag string) *OperationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OperationsGetCall) Context(ctx context.Context) *OperationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OperationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OperationsGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/operations/{operation}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":   c.project,
		"operation": c.operation,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.operations.get"), c.s.client, req)
}

// Do executes the "deploymentmanager.operations.get" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *OperationsGetCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Gets information about a specific operation.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.operations.get",
	//   "parameterOrder": [
	//     "project",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "description": "The name of the operation for this request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/operations/{operation}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.operations.list":

type OperationsListCall struct {
	s            *Service
	project      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all operations for a project.
func (r *OperationsService) List(project string) *OperationsListCall {
	c := &OperationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	return c
}

// Filter sets the optional parameter "filter": A filter expression that
// filters resources listed in the response. The expression must specify
// the field name, a comparison operator, and the value that you want to
// use for filtering. The value must be a string, a number, or a
// boolean. The comparison operator must be either =, !=, >, or <.
//
// For example, if you are filtering Compute Engine instances, you can
// exclude instances named example-instance by specifying name !=
// example-instance.
//
// You can also filter nested fields. For example, you could specify
// scheduling.automaticRestart = false to include instances only if they
// are not scheduled for automatic restarts. You can use filtering on
// nested fields to filter based on resource labels.
//
// To filter on multiple expressions, provide each separate expression
// within parentheses. For example, (scheduling.automaticRestart = true)
// (cpuPlatform = "Intel Skylake"). By default, each expression is an
// AND expression. However, you can include AND and OR expressions
// explicitly. For example, (cpuPlatform = "Intel Skylake") OR
// (cpuPlatform = "Intel Broadwell") AND (scheduling.automaticRestart =
// true).
func (c *OperationsListCall) Filter(filter string) *OperationsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of results per page that should be returned. If the number of
// available results is larger than maxResults, Compute Engine returns a
// nextPageToken that can be used to get the next page of results in
// subsequent list requests. Acceptable values are 0 to 500, inclusive.
// (Default: 500)
func (c *OperationsListCall) MaxResults(maxResults int64) *OperationsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// OrderBy sets the optional parameter "orderBy": Sorts list results by
// a certain order. By default, results are returned in alphanumerical
// order based on the resource name.
//
// You can also sort results in descending order based on the creation
// timestamp using orderBy="creationTimestamp desc". This sorts results
// based on the creationTimestamp field in reverse chronological order
// (newest result first). Use this to sort resources like operations so
// that the newest operation is returned first.
//
// Currently, only sorting by name or creationTimestamp desc is
// supported.
func (c *OperationsListCall) OrderBy(orderBy string) *OperationsListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a page
// token to use. Set pageToken to the nextPageToken returned by a
// previous list request to get the next page of results.
func (c *OperationsListCall) PageToken(pageToken string) *OperationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsListCall) Fields(s ...googleapi.Field) *OperationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OperationsListCall) IfNoneMatch(entityTag string) *OperationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OperationsListCall) Context(ctx context.Context) *OperationsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OperationsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OperationsListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/operations")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.operations.list"), c.s.client, req)
}

// Do executes the "deploymentmanager.operations.list" call.
// Exactly one of *OperationsListResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *OperationsListResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *OperationsListCall) Do(opts ...googleapi.CallOption) (*OperationsListResponse, error) {
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
	ret := &OperationsListResponse{
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
	//   "description": "Lists all operations for a project.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.operations.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "A filter expression that filters resources listed in the response. The expression must specify the field name, a comparison operator, and the value that you want to use for filtering. The value must be a string, a number, or a boolean. The comparison operator must be either =, !=, \u003e, or \u003c.\n\nFor example, if you are filtering Compute Engine instances, you can exclude instances named example-instance by specifying name != example-instance.\n\nYou can also filter nested fields. For example, you could specify scheduling.automaticRestart = false to include instances only if they are not scheduled for automatic restarts. You can use filtering on nested fields to filter based on resource labels.\n\nTo filter on multiple expressions, provide each separate expression within parentheses. For example, (scheduling.automaticRestart = true) (cpuPlatform = \"Intel Skylake\"). By default, each expression is an AND expression. However, you can include AND and OR expressions explicitly. For example, (cpuPlatform = \"Intel Skylake\") OR (cpuPlatform = \"Intel Broadwell\") AND (scheduling.automaticRestart = true).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "The maximum number of results per page that should be returned. If the number of available results is larger than maxResults, Compute Engine returns a nextPageToken that can be used to get the next page of results in subsequent list requests. Acceptable values are 0 to 500, inclusive. (Default: 500)",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Sorts list results by a certain order. By default, results are returned in alphanumerical order based on the resource name.\n\nYou can also sort results in descending order based on the creation timestamp using orderBy=\"creationTimestamp desc\". This sorts results based on the creationTimestamp field in reverse chronological order (newest result first). Use this to sort resources like operations so that the newest operation is returned first.\n\nCurrently, only sorting by name or creationTimestamp desc is supported.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a page token to use. Set pageToken to the nextPageToken returned by a previous list request to get the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/operations",
	//   "response": {
	//     "$ref": "OperationsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *OperationsListCall) Pages(ctx context.Context, f func(*OperationsListResponse) error) error {
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

// method id "deploymentmanager.resources.get":

type ResourcesGetCall struct {
	s            *Service
	project      string
	deployment   string
	resource     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets information about a single resource.
func (r *ResourcesService) Get(project string, deployment string, resource string) *ResourcesGetCall {
	c := &ResourcesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.deployment = deployment
	c.resource = resource
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ResourcesGetCall) Fields(s ...googleapi.Field) *ResourcesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ResourcesGetCall) IfNoneMatch(entityTag string) *ResourcesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ResourcesGetCall) Context(ctx context.Context) *ResourcesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ResourcesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ResourcesGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{deployment}/resources/{resource}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
		"resource":   c.resource,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.resources.get"), c.s.client, req)
}

// Do executes the "deploymentmanager.resources.get" call.
// Exactly one of *Resource or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Resource.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ResourcesGetCall) Do(opts ...googleapi.CallOption) (*Resource, error) {
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
	ret := &Resource{
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
	//   "description": "Gets information about a single resource.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.resources.get",
	//   "parameterOrder": [
	//     "project",
	//     "deployment",
	//     "resource"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resource": {
	//       "description": "The name of the resource for this request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}/resources/{resource}",
	//   "response": {
	//     "$ref": "Resource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.resources.list":

type ResourcesListCall struct {
	s            *Service
	project      string
	deployment   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all resources in a given deployment.
func (r *ResourcesService) List(project string, deployment string) *ResourcesListCall {
	c := &ResourcesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.deployment = deployment
	return c
}

// Filter sets the optional parameter "filter": A filter expression that
// filters resources listed in the response. The expression must specify
// the field name, a comparison operator, and the value that you want to
// use for filtering. The value must be a string, a number, or a
// boolean. The comparison operator must be either =, !=, >, or <.
//
// For example, if you are filtering Compute Engine instances, you can
// exclude instances named example-instance by specifying name !=
// example-instance.
//
// You can also filter nested fields. For example, you could specify
// scheduling.automaticRestart = false to include instances only if they
// are not scheduled for automatic restarts. You can use filtering on
// nested fields to filter based on resource labels.
//
// To filter on multiple expressions, provide each separate expression
// within parentheses. For example, (scheduling.automaticRestart = true)
// (cpuPlatform = "Intel Skylake"). By default, each expression is an
// AND expression. However, you can include AND and OR expressions
// explicitly. For example, (cpuPlatform = "Intel Skylake") OR
// (cpuPlatform = "Intel Broadwell") AND (scheduling.automaticRestart =
// true).
func (c *ResourcesListCall) Filter(filter string) *ResourcesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of results per page that should be returned. If the number of
// available results is larger than maxResults, Compute Engine returns a
// nextPageToken that can be used to get the next page of results in
// subsequent list requests. Acceptable values are 0 to 500, inclusive.
// (Default: 500)
func (c *ResourcesListCall) MaxResults(maxResults int64) *ResourcesListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// OrderBy sets the optional parameter "orderBy": Sorts list results by
// a certain order. By default, results are returned in alphanumerical
// order based on the resource name.
//
// You can also sort results in descending order based on the creation
// timestamp using orderBy="creationTimestamp desc". This sorts results
// based on the creationTimestamp field in reverse chronological order
// (newest result first). Use this to sort resources like operations so
// that the newest operation is returned first.
//
// Currently, only sorting by name or creationTimestamp desc is
// supported.
func (c *ResourcesListCall) OrderBy(orderBy string) *ResourcesListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a page
// token to use. Set pageToken to the nextPageToken returned by a
// previous list request to get the next page of results.
func (c *ResourcesListCall) PageToken(pageToken string) *ResourcesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ResourcesListCall) Fields(s ...googleapi.Field) *ResourcesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ResourcesListCall) IfNoneMatch(entityTag string) *ResourcesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ResourcesListCall) Context(ctx context.Context) *ResourcesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ResourcesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ResourcesListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/deployments/{deployment}/resources")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"deployment": c.deployment,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.resources.list"), c.s.client, req)
}

// Do executes the "deploymentmanager.resources.list" call.
// Exactly one of *ResourcesListResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ResourcesListResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ResourcesListCall) Do(opts ...googleapi.CallOption) (*ResourcesListResponse, error) {
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
	ret := &ResourcesListResponse{
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
	//   "description": "Lists all resources in a given deployment.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.resources.list",
	//   "parameterOrder": [
	//     "project",
	//     "deployment"
	//   ],
	//   "parameters": {
	//     "deployment": {
	//       "description": "The name of the deployment for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "A filter expression that filters resources listed in the response. The expression must specify the field name, a comparison operator, and the value that you want to use for filtering. The value must be a string, a number, or a boolean. The comparison operator must be either =, !=, \u003e, or \u003c.\n\nFor example, if you are filtering Compute Engine instances, you can exclude instances named example-instance by specifying name != example-instance.\n\nYou can also filter nested fields. For example, you could specify scheduling.automaticRestart = false to include instances only if they are not scheduled for automatic restarts. You can use filtering on nested fields to filter based on resource labels.\n\nTo filter on multiple expressions, provide each separate expression within parentheses. For example, (scheduling.automaticRestart = true) (cpuPlatform = \"Intel Skylake\"). By default, each expression is an AND expression. However, you can include AND and OR expressions explicitly. For example, (cpuPlatform = \"Intel Skylake\") OR (cpuPlatform = \"Intel Broadwell\") AND (scheduling.automaticRestart = true).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "The maximum number of results per page that should be returned. If the number of available results is larger than maxResults, Compute Engine returns a nextPageToken that can be used to get the next page of results in subsequent list requests. Acceptable values are 0 to 500, inclusive. (Default: 500)",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Sorts list results by a certain order. By default, results are returned in alphanumerical order based on the resource name.\n\nYou can also sort results in descending order based on the creation timestamp using orderBy=\"creationTimestamp desc\". This sorts results based on the creationTimestamp field in reverse chronological order (newest result first). Use this to sort resources like operations so that the newest operation is returned first.\n\nCurrently, only sorting by name or creationTimestamp desc is supported.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a page token to use. Set pageToken to the nextPageToken returned by a previous list request to get the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/deployments/{deployment}/resources",
	//   "response": {
	//     "$ref": "ResourcesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ResourcesListCall) Pages(ctx context.Context, f func(*ResourcesListResponse) error) error {
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

// method id "deploymentmanager.typeProviders.delete":

type TypeProvidersDeleteCall struct {
	s            *Service
	project      string
	typeProvider string
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Delete: Deletes a type provider.
func (r *TypeProvidersService) Delete(project string, typeProvider string) *TypeProvidersDeleteCall {
	c := &TypeProvidersDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.typeProvider = typeProvider
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TypeProvidersDeleteCall) Fields(s ...googleapi.Field) *TypeProvidersDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TypeProvidersDeleteCall) Context(ctx context.Context) *TypeProvidersDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *TypeProvidersDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *TypeProvidersDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/typeProviders/{typeProvider}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":      c.project,
		"typeProvider": c.typeProvider,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.typeProviders.delete"), c.s.client, req)
}

// Do executes the "deploymentmanager.typeProviders.delete" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *TypeProvidersDeleteCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Deletes a type provider.",
	//   "httpMethod": "DELETE",
	//   "id": "deploymentmanager.typeProviders.delete",
	//   "parameterOrder": [
	//     "project",
	//     "typeProvider"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "typeProvider": {
	//       "description": "The name of the type provider for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/typeProviders/{typeProvider}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.typeProviders.get":

type TypeProvidersGetCall struct {
	s            *Service
	project      string
	typeProvider string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets information about a specific type provider.
func (r *TypeProvidersService) Get(project string, typeProvider string) *TypeProvidersGetCall {
	c := &TypeProvidersGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.typeProvider = typeProvider
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TypeProvidersGetCall) Fields(s ...googleapi.Field) *TypeProvidersGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *TypeProvidersGetCall) IfNoneMatch(entityTag string) *TypeProvidersGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TypeProvidersGetCall) Context(ctx context.Context) *TypeProvidersGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *TypeProvidersGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *TypeProvidersGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/typeProviders/{typeProvider}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":      c.project,
		"typeProvider": c.typeProvider,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.typeProviders.get"), c.s.client, req)
}

// Do executes the "deploymentmanager.typeProviders.get" call.
// Exactly one of *TypeProvider or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *TypeProvider.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *TypeProvidersGetCall) Do(opts ...googleapi.CallOption) (*TypeProvider, error) {
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
	ret := &TypeProvider{
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
	//   "description": "Gets information about a specific type provider.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.typeProviders.get",
	//   "parameterOrder": [
	//     "project",
	//     "typeProvider"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "typeProvider": {
	//       "description": "The name of the type provider for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/typeProviders/{typeProvider}",
	//   "response": {
	//     "$ref": "TypeProvider"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.typeProviders.getType":

type TypeProvidersGetTypeCall struct {
	s            *Service
	project      string
	typeProvider string
	type_        string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetType: Gets a type info for a type provided by a TypeProvider.
func (r *TypeProvidersService) GetType(project string, typeProvider string, type_ string) *TypeProvidersGetTypeCall {
	c := &TypeProvidersGetTypeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.typeProvider = typeProvider
	c.type_ = type_
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TypeProvidersGetTypeCall) Fields(s ...googleapi.Field) *TypeProvidersGetTypeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *TypeProvidersGetTypeCall) IfNoneMatch(entityTag string) *TypeProvidersGetTypeCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TypeProvidersGetTypeCall) Context(ctx context.Context) *TypeProvidersGetTypeCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *TypeProvidersGetTypeCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *TypeProvidersGetTypeCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/typeProviders/{typeProvider}/types/{type}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":      c.project,
		"typeProvider": c.typeProvider,
		"type":         c.type_,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.typeProviders.getType"), c.s.client, req)
}

// Do executes the "deploymentmanager.typeProviders.getType" call.
// Exactly one of *TypeInfo or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *TypeInfo.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *TypeProvidersGetTypeCall) Do(opts ...googleapi.CallOption) (*TypeInfo, error) {
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
	ret := &TypeInfo{
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
	//   "description": "Gets a type info for a type provided by a TypeProvider.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.typeProviders.getType",
	//   "parameterOrder": [
	//     "project",
	//     "typeProvider",
	//     "type"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "type": {
	//       "description": "The name of the type provider for this request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "typeProvider": {
	//       "description": "The name of the type provider for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/typeProviders/{typeProvider}/types/{type}",
	//   "response": {
	//     "$ref": "TypeInfo"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// method id "deploymentmanager.typeProviders.insert":

type TypeProvidersInsertCall struct {
	s            *Service
	project      string
	typeprovider *TypeProvider
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Insert: Creates a type provider.
func (r *TypeProvidersService) Insert(project string, typeprovider *TypeProvider) *TypeProvidersInsertCall {
	c := &TypeProvidersInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.typeprovider = typeprovider
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TypeProvidersInsertCall) Fields(s ...googleapi.Field) *TypeProvidersInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TypeProvidersInsertCall) Context(ctx context.Context) *TypeProvidersInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *TypeProvidersInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *TypeProvidersInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.typeprovider)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/typeProviders")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.typeProviders.insert"), c.s.client, req)
}

// Do executes the "deploymentmanager.typeProviders.insert" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *TypeProvidersInsertCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Creates a type provider.",
	//   "httpMethod": "POST",
	//   "id": "deploymentmanager.typeProviders.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/typeProviders",
	//   "request": {
	//     "$ref": "TypeProvider"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.typeProviders.list":

type TypeProvidersListCall struct {
	s            *Service
	project      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all resource type providers for Deployment Manager.
func (r *TypeProvidersService) List(project string) *TypeProvidersListCall {
	c := &TypeProvidersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	return c
}

// Filter sets the optional parameter "filter": A filter expression that
// filters resources listed in the response. The expression must specify
// the field name, a comparison operator, and the value that you want to
// use for filtering. The value must be a string, a number, or a
// boolean. The comparison operator must be either =, !=, >, or <.
//
// For example, if you are filtering Compute Engine instances, you can
// exclude instances named example-instance by specifying name !=
// example-instance.
//
// You can also filter nested fields. For example, you could specify
// scheduling.automaticRestart = false to include instances only if they
// are not scheduled for automatic restarts. You can use filtering on
// nested fields to filter based on resource labels.
//
// To filter on multiple expressions, provide each separate expression
// within parentheses. For example, (scheduling.automaticRestart = true)
// (cpuPlatform = "Intel Skylake"). By default, each expression is an
// AND expression. However, you can include AND and OR expressions
// explicitly. For example, (cpuPlatform = "Intel Skylake") OR
// (cpuPlatform = "Intel Broadwell") AND (scheduling.automaticRestart =
// true).
func (c *TypeProvidersListCall) Filter(filter string) *TypeProvidersListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of results per page that should be returned. If the number of
// available results is larger than maxResults, Compute Engine returns a
// nextPageToken that can be used to get the next page of results in
// subsequent list requests. Acceptable values are 0 to 500, inclusive.
// (Default: 500)
func (c *TypeProvidersListCall) MaxResults(maxResults int64) *TypeProvidersListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// OrderBy sets the optional parameter "orderBy": Sorts list results by
// a certain order. By default, results are returned in alphanumerical
// order based on the resource name.
//
// You can also sort results in descending order based on the creation
// timestamp using orderBy="creationTimestamp desc". This sorts results
// based on the creationTimestamp field in reverse chronological order
// (newest result first). Use this to sort resources like operations so
// that the newest operation is returned first.
//
// Currently, only sorting by name or creationTimestamp desc is
// supported.
func (c *TypeProvidersListCall) OrderBy(orderBy string) *TypeProvidersListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a page
// token to use. Set pageToken to the nextPageToken returned by a
// previous list request to get the next page of results.
func (c *TypeProvidersListCall) PageToken(pageToken string) *TypeProvidersListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TypeProvidersListCall) Fields(s ...googleapi.Field) *TypeProvidersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *TypeProvidersListCall) IfNoneMatch(entityTag string) *TypeProvidersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TypeProvidersListCall) Context(ctx context.Context) *TypeProvidersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *TypeProvidersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *TypeProvidersListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/typeProviders")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.typeProviders.list"), c.s.client, req)
}

// Do executes the "deploymentmanager.typeProviders.list" call.
// Exactly one of *TypeProvidersListResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *TypeProvidersListResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *TypeProvidersListCall) Do(opts ...googleapi.CallOption) (*TypeProvidersListResponse, error) {
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
	ret := &TypeProvidersListResponse{
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
	//   "description": "Lists all resource type providers for Deployment Manager.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.typeProviders.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "A filter expression that filters resources listed in the response. The expression must specify the field name, a comparison operator, and the value that you want to use for filtering. The value must be a string, a number, or a boolean. The comparison operator must be either =, !=, \u003e, or \u003c.\n\nFor example, if you are filtering Compute Engine instances, you can exclude instances named example-instance by specifying name != example-instance.\n\nYou can also filter nested fields. For example, you could specify scheduling.automaticRestart = false to include instances only if they are not scheduled for automatic restarts. You can use filtering on nested fields to filter based on resource labels.\n\nTo filter on multiple expressions, provide each separate expression within parentheses. For example, (scheduling.automaticRestart = true) (cpuPlatform = \"Intel Skylake\"). By default, each expression is an AND expression. However, you can include AND and OR expressions explicitly. For example, (cpuPlatform = \"Intel Skylake\") OR (cpuPlatform = \"Intel Broadwell\") AND (scheduling.automaticRestart = true).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "The maximum number of results per page that should be returned. If the number of available results is larger than maxResults, Compute Engine returns a nextPageToken that can be used to get the next page of results in subsequent list requests. Acceptable values are 0 to 500, inclusive. (Default: 500)",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Sorts list results by a certain order. By default, results are returned in alphanumerical order based on the resource name.\n\nYou can also sort results in descending order based on the creation timestamp using orderBy=\"creationTimestamp desc\". This sorts results based on the creationTimestamp field in reverse chronological order (newest result first). Use this to sort resources like operations so that the newest operation is returned first.\n\nCurrently, only sorting by name or creationTimestamp desc is supported.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a page token to use. Set pageToken to the nextPageToken returned by a previous list request to get the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/typeProviders",
	//   "response": {
	//     "$ref": "TypeProvidersListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *TypeProvidersListCall) Pages(ctx context.Context, f func(*TypeProvidersListResponse) error) error {
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

// method id "deploymentmanager.typeProviders.listTypes":

type TypeProvidersListTypesCall struct {
	s            *Service
	project      string
	typeProvider string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// ListTypes: Lists all the type info for a TypeProvider.
func (r *TypeProvidersService) ListTypes(project string, typeProvider string) *TypeProvidersListTypesCall {
	c := &TypeProvidersListTypesCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.typeProvider = typeProvider
	return c
}

// Filter sets the optional parameter "filter": A filter expression that
// filters resources listed in the response. The expression must specify
// the field name, a comparison operator, and the value that you want to
// use for filtering. The value must be a string, a number, or a
// boolean. The comparison operator must be either =, !=, >, or <.
//
// For example, if you are filtering Compute Engine instances, you can
// exclude instances named example-instance by specifying name !=
// example-instance.
//
// You can also filter nested fields. For example, you could specify
// scheduling.automaticRestart = false to include instances only if they
// are not scheduled for automatic restarts. You can use filtering on
// nested fields to filter based on resource labels.
//
// To filter on multiple expressions, provide each separate expression
// within parentheses. For example, (scheduling.automaticRestart = true)
// (cpuPlatform = "Intel Skylake"). By default, each expression is an
// AND expression. However, you can include AND and OR expressions
// explicitly. For example, (cpuPlatform = "Intel Skylake") OR
// (cpuPlatform = "Intel Broadwell") AND (scheduling.automaticRestart =
// true).
func (c *TypeProvidersListTypesCall) Filter(filter string) *TypeProvidersListTypesCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of results per page that should be returned. If the number of
// available results is larger than maxResults, Compute Engine returns a
// nextPageToken that can be used to get the next page of results in
// subsequent list requests. Acceptable values are 0 to 500, inclusive.
// (Default: 500)
func (c *TypeProvidersListTypesCall) MaxResults(maxResults int64) *TypeProvidersListTypesCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// OrderBy sets the optional parameter "orderBy": Sorts list results by
// a certain order. By default, results are returned in alphanumerical
// order based on the resource name.
//
// You can also sort results in descending order based on the creation
// timestamp using orderBy="creationTimestamp desc". This sorts results
// based on the creationTimestamp field in reverse chronological order
// (newest result first). Use this to sort resources like operations so
// that the newest operation is returned first.
//
// Currently, only sorting by name or creationTimestamp desc is
// supported.
func (c *TypeProvidersListTypesCall) OrderBy(orderBy string) *TypeProvidersListTypesCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a page
// token to use. Set pageToken to the nextPageToken returned by a
// previous list request to get the next page of results.
func (c *TypeProvidersListTypesCall) PageToken(pageToken string) *TypeProvidersListTypesCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TypeProvidersListTypesCall) Fields(s ...googleapi.Field) *TypeProvidersListTypesCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *TypeProvidersListTypesCall) IfNoneMatch(entityTag string) *TypeProvidersListTypesCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TypeProvidersListTypesCall) Context(ctx context.Context) *TypeProvidersListTypesCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *TypeProvidersListTypesCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *TypeProvidersListTypesCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/typeProviders/{typeProvider}/types")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":      c.project,
		"typeProvider": c.typeProvider,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.typeProviders.listTypes"), c.s.client, req)
}

// Do executes the "deploymentmanager.typeProviders.listTypes" call.
// Exactly one of *TypeProvidersListTypesResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *TypeProvidersListTypesResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *TypeProvidersListTypesCall) Do(opts ...googleapi.CallOption) (*TypeProvidersListTypesResponse, error) {
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
	ret := &TypeProvidersListTypesResponse{
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
	//   "description": "Lists all the type info for a TypeProvider.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.typeProviders.listTypes",
	//   "parameterOrder": [
	//     "project",
	//     "typeProvider"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "A filter expression that filters resources listed in the response. The expression must specify the field name, a comparison operator, and the value that you want to use for filtering. The value must be a string, a number, or a boolean. The comparison operator must be either =, !=, \u003e, or \u003c.\n\nFor example, if you are filtering Compute Engine instances, you can exclude instances named example-instance by specifying name != example-instance.\n\nYou can also filter nested fields. For example, you could specify scheduling.automaticRestart = false to include instances only if they are not scheduled for automatic restarts. You can use filtering on nested fields to filter based on resource labels.\n\nTo filter on multiple expressions, provide each separate expression within parentheses. For example, (scheduling.automaticRestart = true) (cpuPlatform = \"Intel Skylake\"). By default, each expression is an AND expression. However, you can include AND and OR expressions explicitly. For example, (cpuPlatform = \"Intel Skylake\") OR (cpuPlatform = \"Intel Broadwell\") AND (scheduling.automaticRestart = true).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "The maximum number of results per page that should be returned. If the number of available results is larger than maxResults, Compute Engine returns a nextPageToken that can be used to get the next page of results in subsequent list requests. Acceptable values are 0 to 500, inclusive. (Default: 500)",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Sorts list results by a certain order. By default, results are returned in alphanumerical order based on the resource name.\n\nYou can also sort results in descending order based on the creation timestamp using orderBy=\"creationTimestamp desc\". This sorts results based on the creationTimestamp field in reverse chronological order (newest result first). Use this to sort resources like operations so that the newest operation is returned first.\n\nCurrently, only sorting by name or creationTimestamp desc is supported.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a page token to use. Set pageToken to the nextPageToken returned by a previous list request to get the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "typeProvider": {
	//       "description": "The name of the type provider for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/typeProviders/{typeProvider}/types",
	//   "response": {
	//     "$ref": "TypeProvidersListTypesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *TypeProvidersListTypesCall) Pages(ctx context.Context, f func(*TypeProvidersListTypesResponse) error) error {
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

// method id "deploymentmanager.typeProviders.patch":

type TypeProvidersPatchCall struct {
	s            *Service
	project      string
	typeProvider string
	typeprovider *TypeProvider
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Patch: Updates a type provider. This method supports patch semantics.
func (r *TypeProvidersService) Patch(project string, typeProvider string, typeprovider *TypeProvider) *TypeProvidersPatchCall {
	c := &TypeProvidersPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.typeProvider = typeProvider
	c.typeprovider = typeprovider
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TypeProvidersPatchCall) Fields(s ...googleapi.Field) *TypeProvidersPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TypeProvidersPatchCall) Context(ctx context.Context) *TypeProvidersPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *TypeProvidersPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *TypeProvidersPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.typeprovider)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/typeProviders/{typeProvider}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":      c.project,
		"typeProvider": c.typeProvider,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.typeProviders.patch"), c.s.client, req)
}

// Do executes the "deploymentmanager.typeProviders.patch" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *TypeProvidersPatchCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Updates a type provider. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "deploymentmanager.typeProviders.patch",
	//   "parameterOrder": [
	//     "project",
	//     "typeProvider"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "typeProvider": {
	//       "description": "The name of the type provider for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/typeProviders/{typeProvider}",
	//   "request": {
	//     "$ref": "TypeProvider"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.typeProviders.update":

type TypeProvidersUpdateCall struct {
	s            *Service
	project      string
	typeProvider string
	typeprovider *TypeProvider
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Update: Updates a type provider.
func (r *TypeProvidersService) Update(project string, typeProvider string, typeprovider *TypeProvider) *TypeProvidersUpdateCall {
	c := &TypeProvidersUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.typeProvider = typeProvider
	c.typeprovider = typeprovider
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TypeProvidersUpdateCall) Fields(s ...googleapi.Field) *TypeProvidersUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TypeProvidersUpdateCall) Context(ctx context.Context) *TypeProvidersUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *TypeProvidersUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *TypeProvidersUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.typeprovider)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/typeProviders/{typeProvider}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":      c.project,
		"typeProvider": c.typeProvider,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.typeProviders.update"), c.s.client, req)
}

// Do executes the "deploymentmanager.typeProviders.update" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *TypeProvidersUpdateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Updates a type provider.",
	//   "httpMethod": "PUT",
	//   "id": "deploymentmanager.typeProviders.update",
	//   "parameterOrder": [
	//     "project",
	//     "typeProvider"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "typeProvider": {
	//       "description": "The name of the type provider for this request.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/typeProviders/{typeProvider}",
	//   "request": {
	//     "$ref": "TypeProvider"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.cloudman"
	//   ]
	// }

}

// method id "deploymentmanager.types.list":

type TypesListCall struct {
	s            *Service
	project      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all resource types for Deployment Manager.
func (r *TypesService) List(project string) *TypesListCall {
	c := &TypesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	return c
}

// Filter sets the optional parameter "filter": A filter expression that
// filters resources listed in the response. The expression must specify
// the field name, a comparison operator, and the value that you want to
// use for filtering. The value must be a string, a number, or a
// boolean. The comparison operator must be either =, !=, >, or <.
//
// For example, if you are filtering Compute Engine instances, you can
// exclude instances named example-instance by specifying name !=
// example-instance.
//
// You can also filter nested fields. For example, you could specify
// scheduling.automaticRestart = false to include instances only if they
// are not scheduled for automatic restarts. You can use filtering on
// nested fields to filter based on resource labels.
//
// To filter on multiple expressions, provide each separate expression
// within parentheses. For example, (scheduling.automaticRestart = true)
// (cpuPlatform = "Intel Skylake"). By default, each expression is an
// AND expression. However, you can include AND and OR expressions
// explicitly. For example, (cpuPlatform = "Intel Skylake") OR
// (cpuPlatform = "Intel Broadwell") AND (scheduling.automaticRestart =
// true).
func (c *TypesListCall) Filter(filter string) *TypesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of results per page that should be returned. If the number of
// available results is larger than maxResults, Compute Engine returns a
// nextPageToken that can be used to get the next page of results in
// subsequent list requests. Acceptable values are 0 to 500, inclusive.
// (Default: 500)
func (c *TypesListCall) MaxResults(maxResults int64) *TypesListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// OrderBy sets the optional parameter "orderBy": Sorts list results by
// a certain order. By default, results are returned in alphanumerical
// order based on the resource name.
//
// You can also sort results in descending order based on the creation
// timestamp using orderBy="creationTimestamp desc". This sorts results
// based on the creationTimestamp field in reverse chronological order
// (newest result first). Use this to sort resources like operations so
// that the newest operation is returned first.
//
// Currently, only sorting by name or creationTimestamp desc is
// supported.
func (c *TypesListCall) OrderBy(orderBy string) *TypesListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a page
// token to use. Set pageToken to the nextPageToken returned by a
// previous list request to get the next page of results.
func (c *TypesListCall) PageToken(pageToken string) *TypesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TypesListCall) Fields(s ...googleapi.Field) *TypesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *TypesListCall) IfNoneMatch(entityTag string) *TypesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TypesListCall) Context(ctx context.Context) *TypesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *TypesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *TypesListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/global/types")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "deploymentmanager.types.list"), c.s.client, req)
}

// Do executes the "deploymentmanager.types.list" call.
// Exactly one of *TypesListResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *TypesListResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *TypesListCall) Do(opts ...googleapi.CallOption) (*TypesListResponse, error) {
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
	ret := &TypesListResponse{
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
	//   "description": "Lists all resource types for Deployment Manager.",
	//   "httpMethod": "GET",
	//   "id": "deploymentmanager.types.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "A filter expression that filters resources listed in the response. The expression must specify the field name, a comparison operator, and the value that you want to use for filtering. The value must be a string, a number, or a boolean. The comparison operator must be either =, !=, \u003e, or \u003c.\n\nFor example, if you are filtering Compute Engine instances, you can exclude instances named example-instance by specifying name != example-instance.\n\nYou can also filter nested fields. For example, you could specify scheduling.automaticRestart = false to include instances only if they are not scheduled for automatic restarts. You can use filtering on nested fields to filter based on resource labels.\n\nTo filter on multiple expressions, provide each separate expression within parentheses. For example, (scheduling.automaticRestart = true) (cpuPlatform = \"Intel Skylake\"). By default, each expression is an AND expression. However, you can include AND and OR expressions explicitly. For example, (cpuPlatform = \"Intel Skylake\") OR (cpuPlatform = \"Intel Broadwell\") AND (scheduling.automaticRestart = true).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "The maximum number of results per page that should be returned. If the number of available results is larger than maxResults, Compute Engine returns a nextPageToken that can be used to get the next page of results in subsequent list requests. Acceptable values are 0 to 500, inclusive. (Default: 500)",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Sorts list results by a certain order. By default, results are returned in alphanumerical order based on the resource name.\n\nYou can also sort results in descending order based on the creation timestamp using orderBy=\"creationTimestamp desc\". This sorts results based on the creationTimestamp field in reverse chronological order (newest result first). Use this to sort resources like operations so that the newest operation is returned first.\n\nCurrently, only sorting by name or creationTimestamp desc is supported.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a page token to use. Set pageToken to the nextPageToken returned by a previous list request to get the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/global/types",
	//   "response": {
	//     "$ref": "TypesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/ndev.cloudman",
	//     "https://www.googleapis.com/auth/ndev.cloudman.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *TypesListCall) Pages(ctx context.Context, f func(*TypesListResponse) error) error {
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
