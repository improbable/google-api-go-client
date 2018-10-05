// Package servicebroker provides access to the Service Broker API.
//
// See https://cloud.google.com/kubernetes-engine/docs/concepts/add-on/service-broker
//
// Usage example:
//
//   import "google.golang.org/api/servicebroker/v1beta1"
//   ...
//   servicebrokerService, err := servicebroker.New(oauthHttpClient)
package servicebroker // import "google.golang.org/api/servicebroker/v1beta1"

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

const apiId = "servicebroker:v1beta1"
const apiName = "servicebroker"
const apiVersion = "v1beta1"
const basePath = "https://servicebroker.googleapis.com/"

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
	s.V1beta1 = NewV1beta1Service(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService

	V1beta1 *V1beta1Service
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Brokers = NewProjectsBrokersService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Brokers *ProjectsBrokersService
}

func NewProjectsBrokersService(s *Service) *ProjectsBrokersService {
	rs := &ProjectsBrokersService{s: s}
	rs.Instances = NewProjectsBrokersInstancesService(s)
	rs.V2 = NewProjectsBrokersV2Service(s)
	return rs
}

type ProjectsBrokersService struct {
	s *Service

	Instances *ProjectsBrokersInstancesService

	V2 *ProjectsBrokersV2Service
}

func NewProjectsBrokersInstancesService(s *Service) *ProjectsBrokersInstancesService {
	rs := &ProjectsBrokersInstancesService{s: s}
	rs.Bindings = NewProjectsBrokersInstancesBindingsService(s)
	return rs
}

type ProjectsBrokersInstancesService struct {
	s *Service

	Bindings *ProjectsBrokersInstancesBindingsService
}

func NewProjectsBrokersInstancesBindingsService(s *Service) *ProjectsBrokersInstancesBindingsService {
	rs := &ProjectsBrokersInstancesBindingsService{s: s}
	return rs
}

type ProjectsBrokersInstancesBindingsService struct {
	s *Service
}

func NewProjectsBrokersV2Service(s *Service) *ProjectsBrokersV2Service {
	rs := &ProjectsBrokersV2Service{s: s}
	rs.Catalog = NewProjectsBrokersV2CatalogService(s)
	rs.ServiceInstances = NewProjectsBrokersV2ServiceInstancesService(s)
	return rs
}

type ProjectsBrokersV2Service struct {
	s *Service

	Catalog *ProjectsBrokersV2CatalogService

	ServiceInstances *ProjectsBrokersV2ServiceInstancesService
}

func NewProjectsBrokersV2CatalogService(s *Service) *ProjectsBrokersV2CatalogService {
	rs := &ProjectsBrokersV2CatalogService{s: s}
	return rs
}

type ProjectsBrokersV2CatalogService struct {
	s *Service
}

func NewProjectsBrokersV2ServiceInstancesService(s *Service) *ProjectsBrokersV2ServiceInstancesService {
	rs := &ProjectsBrokersV2ServiceInstancesService{s: s}
	rs.ServiceBindings = NewProjectsBrokersV2ServiceInstancesServiceBindingsService(s)
	return rs
}

type ProjectsBrokersV2ServiceInstancesService struct {
	s *Service

	ServiceBindings *ProjectsBrokersV2ServiceInstancesServiceBindingsService
}

func NewProjectsBrokersV2ServiceInstancesServiceBindingsService(s *Service) *ProjectsBrokersV2ServiceInstancesServiceBindingsService {
	rs := &ProjectsBrokersV2ServiceInstancesServiceBindingsService{s: s}
	return rs
}

type ProjectsBrokersV2ServiceInstancesServiceBindingsService struct {
	s *Service
}

func NewV1beta1Service(s *Service) *V1beta1Service {
	rs := &V1beta1Service{s: s}
	return rs
}

type V1beta1Service struct {
	s *Service
}

// GoogleCloudServicebrokerV1beta1__Binding: Describes the binding.
type GoogleCloudServicebrokerV1beta1__Binding struct {
	// BindResource: A JSON object that contains data for platform resources
	// associated with
	// the binding to be created.
	BindResource googleapi.RawMessage `json:"bind_resource,omitempty"`

	// BindingId: The id of the binding. Must be unique within GCP
	// project.
	// Maximum length is 64, GUID recommended.
	// Required.
	BindingId string `json:"binding_id,omitempty"`

	// CreateTime: Output only.
	// Timestamp for when the binding was created.
	CreateTime string `json:"createTime,omitempty"`

	// DeploymentName: Output only.
	// String containing the Deployment Manager deployment name that was
	// created
	// for this binding,
	DeploymentName string `json:"deploymentName,omitempty"`

	// Parameters: Configuration options for the service binding.
	Parameters googleapi.RawMessage `json:"parameters,omitempty"`

	// PlanId: The ID of the plan. See `Service` and `Plan` resources for
	// details.
	// Maximum length is 64, GUID recommended.
	// Required.
	PlanId string `json:"plan_id,omitempty"`

	// ResourceName: Output only.
	// The resource name of the binding,
	// e.g.
	// projects/project_id/brokers/broker_id/service_instances/instance_
	// id/bindings/binding_id.
	ResourceName string `json:"resourceName,omitempty"`

	// ServiceId: The id of the service. Must be a valid identifier of a
	// service
	// contained in the list from a `ListServices()` call.
	// Maximum length is 64, GUID recommended.
	// Required.
	ServiceId string `json:"service_id,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BindResource") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BindResource") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudServicebrokerV1beta1__Binding) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudServicebrokerV1beta1__Binding
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudServicebrokerV1beta1__Broker: Broker represents a
// consumable collection of Service Registry catalogs
// exposed as an OSB Broker.
type GoogleCloudServicebrokerV1beta1__Broker struct {
	// CreateTime: Output only.
	// Timestamp for when the broker was created.
	CreateTime string `json:"createTime,omitempty"`

	// Name: Name of the broker in the
	// format:
	// <projects>/<project-id>/brokers/<broker>.
	// This allows for multiple brokers per project which can be used
	// to
	// enable having custom brokers per GKE cluster, for example.
	Name string `json:"name,omitempty"`

	// Title: User friendly title of the broker.
	// Limited to 1024 characters. Requests with longer titles will be
	// rejected.
	Title string `json:"title,omitempty"`

	// Url: Output only.
	// URL of the broker OSB-compliant endpoint, for
	// example:
	// https://servicebroker.googleapis.com/projects/<project>/broke
	// rs/<broker>
	Url string `json:"url,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudServicebrokerV1beta1__Broker) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudServicebrokerV1beta1__Broker
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudServicebrokerV1beta1__CreateBindingResponse: Response for
// the `CreateBinding()` method.
type GoogleCloudServicebrokerV1beta1__CreateBindingResponse struct {
	// Credentials: Credentials to use the binding.
	Credentials googleapi.RawMessage `json:"credentials,omitempty"`

	// Description: Used to communicate description of the response. Usually
	// for non-standard
	// error
	// codes.
	// https://github.com/openservicebrokerapi/servicebroker/blob/mast
	// er/spec.md#service-broker-errors
	Description string `json:"description,omitempty"`

	// Operation: If broker executes operation asynchronously, this is the
	// operation ID that
	// can be polled to check the completion status of said operation.
	// This broker always executes all create/delete operations
	// asynchronously.
	Operation string `json:"operation,omitempty"`

	// RouteServiceUrl: A URL to which the platform may proxy requests for
	// the address sent with
	// bind_resource.route
	RouteServiceUrl string `json:"route_service_url,omitempty"`

	// SyslogDrainUrl: From where to read system logs.
	SyslogDrainUrl string `json:"syslog_drain_url,omitempty"`

	// VolumeMounts: An array of configuration for mounting volumes.
	VolumeMounts []googleapi.RawMessage `json:"volume_mounts,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Credentials") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Credentials") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudServicebrokerV1beta1__CreateBindingResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudServicebrokerV1beta1__CreateBindingResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudServicebrokerV1beta1__CreateServiceInstanceResponse:
// Response for the `CreateServiceInstance()` method.
type GoogleCloudServicebrokerV1beta1__CreateServiceInstanceResponse struct {
	// Description: Used to communicate description of the response. Usually
	// for non-standard
	// error
	// codes.
	// https://github.com/openservicebrokerapi/servicebroker/blob/mast
	// er/spec.md#service-broker-errors
	Description string `json:"description,omitempty"`

	// Operation: If broker executes operation asynchronously, this is the
	// operation ID that
	// can be polled to check the completion status of said operation.
	// This broker always will return a non-empty operation on success.
	Operation string `json:"operation,omitempty"`

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

func (s *GoogleCloudServicebrokerV1beta1__CreateServiceInstanceResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudServicebrokerV1beta1__CreateServiceInstanceResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudServicebrokerV1beta1__DashboardClient: Message containing
// information required to activate Dashboard SSO feature.
type GoogleCloudServicebrokerV1beta1__DashboardClient struct {
	// Id: The id of the Oauth client that the dashboard will use.
	Id string `json:"id,omitempty"`

	// RedirectUri: A URI for the service dashboard.
	// Validated by the OAuth token server when the dashboard requests a
	// token.
	RedirectUri string `json:"redirect_uri,omitempty"`

	// Secret: A secret for the dashboard client.
	Secret string `json:"secret,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Id") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudServicebrokerV1beta1__DashboardClient) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudServicebrokerV1beta1__DashboardClient
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudServicebrokerV1beta1__DeleteBindingResponse: Response for
// the `DeleteBinding()` method.
type GoogleCloudServicebrokerV1beta1__DeleteBindingResponse struct {
	// Description: Used to communicate description of the response. Usually
	// for non-standard
	// error
	// codes.
	// https://github.com/openservicebrokerapi/servicebroker/blob/mast
	// er/spec.md#service-broker-errors
	Description string `json:"description,omitempty"`

	// Operation: If broker executes operation asynchronously, this is the
	// operation ID that
	// can be polled to check the completion status of said operation.
	Operation string `json:"operation,omitempty"`

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

func (s *GoogleCloudServicebrokerV1beta1__DeleteBindingResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudServicebrokerV1beta1__DeleteBindingResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudServicebrokerV1beta1__DeleteServiceInstanceResponse:
// Response for the `DeleteServiceInstance()` method.
type GoogleCloudServicebrokerV1beta1__DeleteServiceInstanceResponse struct {
	// Description: Used to communicate description of the response. Usually
	// for non-standard
	// error
	// codes.
	// https://github.com/openservicebrokerapi/servicebroker/blob/mast
	// er/spec.md#service-broker-errors
	Description string `json:"description,omitempty"`

	// Operation: If broker executes operation asynchronously, this is the
	// operation ID that
	// can be polled to check the completion status of said operation.
	Operation string `json:"operation,omitempty"`

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

func (s *GoogleCloudServicebrokerV1beta1__DeleteServiceInstanceResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudServicebrokerV1beta1__DeleteServiceInstanceResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudServicebrokerV1beta1__GetBindingResponse: Response for the
// `GetBinding()` method.
type GoogleCloudServicebrokerV1beta1__GetBindingResponse struct {
	// Credentials: Credentials to use the binding.
	Credentials googleapi.RawMessage `json:"credentials,omitempty"`

	// DeploymentName: String containing the Deployment Manager deployment
	// name that was created
	// for this binding,
	DeploymentName string `json:"deploymentName,omitempty"`

	// Description: Used to communicate description of the response. Usually
	// for non-standard
	// error
	// codes.
	// https://github.com/openservicebrokerapi/servicebroker/blob/mast
	// er/spec.md#service-broker-errors
	Description string `json:"description,omitempty"`

	// ResourceName: Output only.
	// The resource name of the binding,
	// e.g.
	// projects/project_id/brokers/broker_id/service_instances/instance_
	// id/bindings/binding_id.
	ResourceName string `json:"resourceName,omitempty"`

	// RouteServiceUrl: A URL to which the platform may proxy requests for
	// the address sent with
	// bind_resource.route
	RouteServiceUrl string `json:"route_service_url,omitempty"`

	// SyslogDrainUrl: From where to read system logs.
	SyslogDrainUrl string `json:"syslog_drain_url,omitempty"`

	// VolumeMounts: An array of configurations for mounting volumes.
	VolumeMounts []googleapi.RawMessage `json:"volume_mounts,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Credentials") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Credentials") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudServicebrokerV1beta1__GetBindingResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudServicebrokerV1beta1__GetBindingResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudServicebrokerV1beta1__ListBindingsResponse: The response
// for the `ListBindings()` method.
type GoogleCloudServicebrokerV1beta1__ListBindingsResponse struct {
	// Bindings: The list of bindings in the instance.
	Bindings []*GoogleCloudServicebrokerV1beta1__Binding `json:"bindings,omitempty"`

	// Description: Used to communicate description of the response. Usually
	// for non-standard
	// error
	// codes.
	// https://github.com/openservicebrokerapi/servicebroker/blob/mast
	// er/spec.md#service-broker-errors
	Description string `json:"description,omitempty"`

	// NextPageToken: This token allows you to get the next page of results
	// for list requests.
	// If the number of results is larger than `pageSize`, use the
	// `nextPageToken`
	// as a value for the query parameter `pageToken` in the next list
	// request.
	// Subsequent list requests will have their own `nextPageToken` to
	// continue
	// paging through the results
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Bindings") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bindings") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudServicebrokerV1beta1__ListBindingsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudServicebrokerV1beta1__ListBindingsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudServicebrokerV1beta1__ListBrokersResponse: The response
// for the `ListBrokers()` method.
type GoogleCloudServicebrokerV1beta1__ListBrokersResponse struct {
	// Brokers: The list of brokers in the container.
	Brokers []*GoogleCloudServicebrokerV1beta1__Broker `json:"brokers,omitempty"`

	// NextPageToken: This token allows you to get the next page of results
	// for list requests.
	// If the number of results is larger than `pageSize`, use the
	// `nextPageToken`
	// as a value for the query parameter `pageToken` in the next list
	// request.
	// Subsequent list requests will have their own `nextPageToken` to
	// continue
	// paging through the results
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Brokers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Brokers") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudServicebrokerV1beta1__ListBrokersResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudServicebrokerV1beta1__ListBrokersResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudServicebrokerV1beta1__ListCatalogResponse: Response
// message for the `ListCatalog()` method.
type GoogleCloudServicebrokerV1beta1__ListCatalogResponse struct {
	// Description: Used to communicate description of the response. Usually
	// for non-standard
	// error
	// codes.
	// https://github.com/openservicebrokerapi/servicebroker/blob/mast
	// er/spec.md#service-broker-errors
	Description string `json:"description,omitempty"`

	// NextPageToken: This token allows you to get the next page of results
	// for list requests.
	// If the number of results is larger than `pageSize`, use the
	// `nextPageToken`
	// as a value for the query parameter `pageToken` in the next list
	// request.
	// Subsequent list requests will have their own `nextPageToken` to
	// continue
	// paging through the results
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Services: The services available for the requested GCP project.
	Services []*GoogleCloudServicebrokerV1beta1__Service `json:"services,omitempty"`

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

func (s *GoogleCloudServicebrokerV1beta1__ListCatalogResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudServicebrokerV1beta1__ListCatalogResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudServicebrokerV1beta1__ListServiceInstancesResponse: The
// response for the `ListServiceInstances()` method.
type GoogleCloudServicebrokerV1beta1__ListServiceInstancesResponse struct {
	// Description: Used to communicate description of the response. Usually
	// for non-standard
	// error
	// codes.
	// https://github.com/openservicebrokerapi/servicebroker/blob/mast
	// er/spec.md#service-broker-errors
	Description string `json:"description,omitempty"`

	// Instances: The list of instances in the broker.
	Instances []*GoogleCloudServicebrokerV1beta1__ServiceInstance `json:"instances,omitempty"`

	// NextPageToken: This token allows you to get the next page of results
	// for list requests.
	// If the number of results is larger than `pageSize`, use the
	// `nextPageToken`
	// as a value for the query parameter `pageToken` in the next list
	// request.
	// Subsequent list requests will have their own `nextPageToken` to
	// continue
	// paging through the results
	NextPageToken string `json:"nextPageToken,omitempty"`

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

func (s *GoogleCloudServicebrokerV1beta1__ListServiceInstancesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudServicebrokerV1beta1__ListServiceInstancesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudServicebrokerV1beta1__Operation: Describes a long running
// operation, which conforms to OpenService API.
type GoogleCloudServicebrokerV1beta1__Operation struct {
	// Description: Optional description of the Operation state.
	Description string `json:"description,omitempty"`

	// State: The state of the operation.
	// Valid values are: "in progress", "succeeded", and "failed".
	State string `json:"state,omitempty"`

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

func (s *GoogleCloudServicebrokerV1beta1__Operation) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudServicebrokerV1beta1__Operation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudServicebrokerV1beta1__Plan: Plan message describes a
// Service Plan.
type GoogleCloudServicebrokerV1beta1__Plan struct {
	// Bindable: Specifies whether instances of the service can be bound to
	// applications.
	// If not specified, `Service.bindable` will be presumed.
	Bindable bool `json:"bindable,omitempty"`

	// Description: Textual description of the plan. Optional.
	Description string `json:"description,omitempty"`

	// Free: Whether the service is free.
	Free bool `json:"free,omitempty"`

	// Id: ID is a globally unique identifier used to uniquely identify the
	// plan.
	// User must make no presumption about the format of this field.
	Id string `json:"id,omitempty"`

	// Metadata: A list of metadata for a service offering.
	// Metadata is an arbitrary JSON object.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: User friendly name of the plan.
	// The name must be globally unique within GCP project.
	// Note, which is different from ("This must be globally unique within
	// a
	// platform marketplace").
	Name string `json:"name,omitempty"`

	// Schemas: Schema definitions for service instances and bindings for
	// the plan.
	Schemas googleapi.RawMessage `json:"schemas,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Bindable") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bindable") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudServicebrokerV1beta1__Plan) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudServicebrokerV1beta1__Plan
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudServicebrokerV1beta1__Service: The resource model mostly
// follows the Open Service Broker API, as
// described
// here:
// https://github.com/openservicebrokerapi/servicebroker/blob/maste
// r/_spec.md
// Though due to Google Specifics it has additional optional fields.
type GoogleCloudServicebrokerV1beta1__Service struct {
	// Bindable: Specifies whether instances of the service can be bound to
	// applications.
	// Required.
	Bindable bool `json:"bindable,omitempty"`

	// BindingRetrievable: Whether the service provides an endpoint to get
	// service bindings.
	BindingRetrievable bool `json:"binding_retrievable,omitempty"`

	// BindingsRetrievable: Whether the service provides an endpoint to get
	// service bindings.
	BindingsRetrievable bool `json:"bindings_retrievable,omitempty"`

	// DashboardClient: Information to activate Dashboard SSO feature.
	DashboardClient *GoogleCloudServicebrokerV1beta1__DashboardClient `json:"dashboard_client,omitempty"`

	// Description: Textual description of the service. Required.
	Description string `json:"description,omitempty"`

	// Id: ID is a globally unique identifier used to uniquely identify the
	// service.
	// ID is an opaque string.
	Id string `json:"id,omitempty"`

	// InstancesRetrievable: Whether the service provides an endpoint to get
	// service instances.
	InstancesRetrievable bool `json:"instances_retrievable,omitempty"`

	// Metadata: A list of metadata for a service offering.
	// Metadata is an arbitrary JSON object.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: User friendly service name.
	// Name must match [a-z0-9]+ regexp.
	// The name must be globally unique within GCP project.
	// Note, which is different from ("This must be globally unique within
	// a
	// platform marketplace").
	// Required.
	Name string `json:"name,omitempty"`

	// PlanUpdateable: Whether the service supports upgrade/downgrade for
	// some plans.
	PlanUpdateable bool `json:"plan_updateable,omitempty"`

	// Plans: A list of plans for this service.
	// At least one plan is required.
	Plans []*GoogleCloudServicebrokerV1beta1__Plan `json:"plans,omitempty"`

	// Tags: Tags provide a flexible mechanism to expose a classification,
	// attribute, or
	// base technology of a service.
	Tags []string `json:"tags,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Bindable") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bindable") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudServicebrokerV1beta1__Service) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudServicebrokerV1beta1__Service
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudServicebrokerV1beta1__ServiceInstance: Message describing
// inputs to Provision and Update Service instance requests.
type GoogleCloudServicebrokerV1beta1__ServiceInstance struct {
	// Context: Platform specific contextual information under which the
	// service instance
	// is to be provisioned. This replaces organization_guid and
	// space_guid.
	// But can also contain anything.
	// Currently only used for logging context information.
	Context googleapi.RawMessage `json:"context,omitempty"`

	// CreateTime: Output only.
	// Timestamp for when the instance was created.
	CreateTime string `json:"createTime,omitempty"`

	// DeploymentName: Output only.
	// String containing the Deployment Manager deployment name that was
	// created
	// for this instance,
	DeploymentName string `json:"deploymentName,omitempty"`

	// Description: To return errors when GetInstance call is done via HTTP
	// to be unified with
	// other methods.
	Description string `json:"description,omitempty"`

	// InstanceId: The id of the service instance. Must be unique within GCP
	// project.
	// Maximum length is 64, GUID recommended.
	// Required.
	InstanceId string `json:"instance_id,omitempty"`

	// OrganizationGuid: The platform GUID for the organization under which
	// the service is to be
	// provisioned.
	// Required.
	OrganizationGuid string `json:"organization_guid,omitempty"`

	// Parameters: Configuration options for the service
	// instance.
	// Parameters is JSON object serialized to string.
	Parameters googleapi.RawMessage `json:"parameters,omitempty"`

	// PlanId: The ID of the plan. See `Service` and `Plan` resources for
	// details.
	// Maximum length is 64, GUID recommended.
	// Required.
	PlanId string `json:"plan_id,omitempty"`

	// PreviousValues: Used only in UpdateServiceInstance request to
	// optionally specify previous
	// fields.
	PreviousValues googleapi.RawMessage `json:"previous_values,omitempty"`

	// ResourceName: Output only.
	// The resource name of the instance,
	// e.g.
	// projects/project_id/brokers/broker_id/service_instances/instance_
	// id
	ResourceName string `json:"resourceName,omitempty"`

	// ServiceId: The id of the service. Must be a valid identifier of a
	// service
	// contained in the list from a `ListServices()` call.
	// Maximum length is 64, GUID recommended.
	// Required.
	ServiceId string `json:"service_id,omitempty"`

	// SpaceGuid: The identifier for the project space within the platform
	// organization.
	// Required.
	SpaceGuid string `json:"space_guid,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Context") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Context") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudServicebrokerV1beta1__ServiceInstance) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudServicebrokerV1beta1__ServiceInstance
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudServicebrokerV1beta1__UpdateServiceInstanceResponse:
// Response for the `UpdateServiceInstance()` method.
type GoogleCloudServicebrokerV1beta1__UpdateServiceInstanceResponse struct {
	// Description: Used to communicate description of the response. Usually
	// for non-standard
	// error
	// codes.
	// https://github.com/openservicebrokerapi/servicebroker/blob/mast
	// er/spec.md#service-broker-errors
	Description string `json:"description,omitempty"`

	// Operation: If broker executes operation asynchronously, this is the
	// operation ID that
	// can be polled to check the completion status of said operation.
	Operation string `json:"operation,omitempty"`

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

func (s *GoogleCloudServicebrokerV1beta1__UpdateServiceInstanceResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudServicebrokerV1beta1__UpdateServiceInstanceResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIamV1__Binding: Associates `members` with a `role`.
type GoogleIamV1__Binding struct {
	// Condition: Unimplemented. The condition that is associated with this
	// binding.
	// NOTE: an unsatisfied condition will not allow user access via
	// current
	// binding. Different bindings, including their conditions, are
	// examined
	// independently.
	Condition *GoogleType__Expr `json:"condition,omitempty"`

	// Members: Specifies the identities requesting access for a Cloud
	// Platform resource.
	// `members` can have the following values:
	//
	// * `allUsers`: A special identifier that represents anyone who is
	//    on the internet; with or without a Google account.
	//
	// * `allAuthenticatedUsers`: A special identifier that represents
	// anyone
	//    who is authenticated with a Google account or a service
	// account.
	//
	// * `user:{emailid}`: An email address that represents a specific
	// Google
	//    account. For example, `alice@gmail.com` .
	//
	//
	// * `serviceAccount:{emailid}`: An email address that represents a
	// service
	//    account. For example,
	// `my-other-app@appspot.gserviceaccount.com`.
	//
	// * `group:{emailid}`: An email address that represents a Google
	// group.
	//    For example, `admins@example.com`.
	//
	//
	// * `domain:{domain}`: A Google Apps domain name that represents all
	// the
	//    users of that domain. For example, `google.com` or
	// `example.com`.
	//
	//
	Members []string `json:"members,omitempty"`

	// Role: Role that is assigned to `members`.
	// For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
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

func (s *GoogleIamV1__Binding) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIamV1__Binding
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIamV1__Policy: Defines an Identity and Access Management (IAM)
// policy. It is used to
// specify access control policies for Cloud Platform resources.
//
//
// A `Policy` consists of a list of `bindings`. A `binding` binds a list
// of
// `members` to a `role`, where the members can be user accounts, Google
// groups,
// Google domains, and service accounts. A `role` is a named list of
// permissions
// defined by IAM.
//
// **JSON Example**
//
//     {
//       "bindings": [
//         {
//           "role": "roles/owner",
//           "members": [
//             "user:mike@example.com",
//             "group:admins@example.com",
//             "domain:google.com",
//
// "serviceAccount:my-other-app@appspot.gserviceaccount.com"
//           ]
//         },
//         {
//           "role": "roles/viewer",
//           "members": ["user:sean@example.com"]
//         }
//       ]
//     }
//
// **YAML Example**
//
//     bindings:
//     - members:
//       - user:mike@example.com
//       - group:admins@example.com
//       - domain:google.com
//       - serviceAccount:my-other-app@appspot.gserviceaccount.com
//       role: roles/owner
//     - members:
//       - user:sean@example.com
//       role: roles/viewer
//
//
// For a description of IAM and its features, see the
// [IAM developer's guide](https://cloud.google.com/iam/docs).
type GoogleIamV1__Policy struct {
	// Bindings: Associates a list of `members` to a `role`.
	// `bindings` with no members will result in an error.
	Bindings []*GoogleIamV1__Binding `json:"bindings,omitempty"`

	// Etag: `etag` is used for optimistic concurrency control as a way to
	// help
	// prevent simultaneous updates of a policy from overwriting each
	// other.
	// It is strongly suggested that systems make use of the `etag` in
	// the
	// read-modify-write cycle to perform policy updates in order to avoid
	// race
	// conditions: An `etag` is returned in the response to `getIamPolicy`,
	// and
	// systems are expected to put that etag in the request to
	// `setIamPolicy` to
	// ensure that their change will be applied to the same version of the
	// policy.
	//
	// If no `etag` is provided in the call to `setIamPolicy`, then the
	// existing
	// policy is overwritten blindly.
	Etag string `json:"etag,omitempty"`

	// Version: Deprecated.
	Version int64 `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Bindings") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bindings") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIamV1__Policy) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIamV1__Policy
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIamV1__SetIamPolicyRequest: Request message for `SetIamPolicy`
// method.
type GoogleIamV1__SetIamPolicyRequest struct {
	// Policy: REQUIRED: The complete policy to be applied to the
	// `resource`. The size of
	// the policy is limited to a few 10s of KB. An empty policy is a
	// valid policy but certain Cloud Platform services (such as
	// Projects)
	// might reject them.
	Policy *GoogleIamV1__Policy `json:"policy,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Policy") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Policy") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIamV1__SetIamPolicyRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIamV1__SetIamPolicyRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIamV1__TestIamPermissionsRequest: Request message for
// `TestIamPermissions` method.
type GoogleIamV1__TestIamPermissionsRequest struct {
	// Permissions: The set of permissions to check for the `resource`.
	// Permissions with
	// wildcards (such as '*' or 'storage.*') are not allowed. For
	// more
	// information see
	// [IAM
	// Overview](https://cloud.google.com/iam/docs/overview#permissions).
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

func (s *GoogleIamV1__TestIamPermissionsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIamV1__TestIamPermissionsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIamV1__TestIamPermissionsResponse: Response message for
// `TestIamPermissions` method.
type GoogleIamV1__TestIamPermissionsResponse struct {
	// Permissions: A subset of `TestPermissionsRequest.permissions` that
	// the caller is
	// allowed.
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

func (s *GoogleIamV1__TestIamPermissionsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIamV1__TestIamPermissionsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleProtobuf__Empty: A generic empty message that you can re-use to
// avoid defining duplicated
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
type GoogleProtobuf__Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// GoogleType__Expr: Represents an expression text. Example:
//
//     title: "User account presence"
//     description: "Determines whether the request has a user account"
//     expression: "size(request.user) > 0"
type GoogleType__Expr struct {
	// Description: An optional description of the expression. This is a
	// longer text which
	// describes the expression, e.g. when hovered over it in a UI.
	Description string `json:"description,omitempty"`

	// Expression: Textual representation of an expression in
	// Common Expression Language syntax.
	//
	// The application context of the containing message determines
	// which
	// well-known feature set of CEL is supported.
	Expression string `json:"expression,omitempty"`

	// Location: An optional string indicating the location of the
	// expression for error
	// reporting, e.g. a file name and a position in the file.
	Location string `json:"location,omitempty"`

	// Title: An optional title for the expression, i.e. a short string
	// describing
	// its purpose. This can be used e.g. in UIs which allow to enter
	// the
	// expression.
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

func (s *GoogleType__Expr) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleType__Expr
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "servicebroker.projects.brokers.create":

type ProjectsBrokersCreateCall struct {
	s                                       *Service
	parent                                  string
	googlecloudservicebrokerv1beta1__broker *GoogleCloudServicebrokerV1beta1__Broker
	urlParams_                              gensupport.URLParams
	ctx_                                    context.Context
	header_                                 http.Header
}

// Create: CreateBroker creates a Broker.
func (r *ProjectsBrokersService) Create(parent string, googlecloudservicebrokerv1beta1__broker *GoogleCloudServicebrokerV1beta1__Broker) *ProjectsBrokersCreateCall {
	c := &ProjectsBrokersCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googlecloudservicebrokerv1beta1__broker = googlecloudservicebrokerv1beta1__broker
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersCreateCall) Fields(s ...googleapi.Field) *ProjectsBrokersCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersCreateCall) Context(ctx context.Context) *ProjectsBrokersCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudservicebrokerv1beta1__broker)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/brokers")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.create"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.create" call.
// Exactly one of *GoogleCloudServicebrokerV1beta1__Broker or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleCloudServicebrokerV1beta1__Broker.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsBrokersCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudServicebrokerV1beta1__Broker, error) {
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
	ret := &GoogleCloudServicebrokerV1beta1__Broker{
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
	//   "description": "CreateBroker creates a Broker.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers",
	//   "httpMethod": "POST",
	//   "id": "servicebroker.projects.brokers.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "The project in which to create broker.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/brokers",
	//   "request": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__Broker"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__Broker"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.projects.brokers.delete":

type ProjectsBrokersDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: DeleteBroker deletes a Broker.
func (r *ProjectsBrokersService) Delete(name string) *ProjectsBrokersDeleteCall {
	c := &ProjectsBrokersDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersDeleteCall) Fields(s ...googleapi.Field) *ProjectsBrokersDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersDeleteCall) Context(ctx context.Context) *ProjectsBrokersDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.delete"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.delete" call.
// Exactly one of *GoogleProtobuf__Empty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobuf__Empty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsBrokersDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleProtobuf__Empty, error) {
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
	ret := &GoogleProtobuf__Empty{
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
	//   "description": "DeleteBroker deletes a Broker.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers/{brokersId}",
	//   "httpMethod": "DELETE",
	//   "id": "servicebroker.projects.brokers.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The broker to delete.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/brokers/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleProtobuf__Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.projects.brokers.list":

type ProjectsBrokersListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: ListBrokers lists brokers.
func (r *ProjectsBrokersService) List(parent string) *ProjectsBrokersListCall {
	c := &ProjectsBrokersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": Specifies the number
// of results to return per page. If there are fewer
// elements than the specified number, returns all elements.
//  Acceptable values are 0 to 200, inclusive. (Default: 100)
func (c *ProjectsBrokersListCall) PageSize(pageSize int64) *ProjectsBrokersListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a page
// token to use. Set `pageToken` to a `nextPageToken`
// returned by a previous list request to get the next page of results.
func (c *ProjectsBrokersListCall) PageToken(pageToken string) *ProjectsBrokersListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersListCall) Fields(s ...googleapi.Field) *ProjectsBrokersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsBrokersListCall) IfNoneMatch(entityTag string) *ProjectsBrokersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersListCall) Context(ctx context.Context) *ProjectsBrokersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersListCall) doRequest(alt string) (*http.Response, error) {
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
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/brokers")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.list"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.list" call.
// Exactly one of *GoogleCloudServicebrokerV1beta1__ListBrokersResponse
// or error will be non-nil. Any non-2xx status code is an error.
// Response headers are in either
// *GoogleCloudServicebrokerV1beta1__ListBrokersResponse.ServerResponse.H
// eader or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsBrokersListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudServicebrokerV1beta1__ListBrokersResponse, error) {
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
	ret := &GoogleCloudServicebrokerV1beta1__ListBrokersResponse{
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
	//   "description": "ListBrokers lists brokers.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers",
	//   "httpMethod": "GET",
	//   "id": "servicebroker.projects.brokers.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Specifies the number of results to return per page. If there are fewer\nelements than the specified number, returns all elements.\nOptional. Acceptable values are 0 to 200, inclusive. (Default: 100)",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a page token to use. Set `pageToken` to a `nextPageToken`\nreturned by a previous list request to get the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Parent must match `projects/[PROJECT_ID]/brokers`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/brokers",
	//   "response": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__ListBrokersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsBrokersListCall) Pages(ctx context.Context, f func(*GoogleCloudServicebrokerV1beta1__ListBrokersResponse) error) error {
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

// method id "servicebroker.projects.brokers.instances.get":

type ProjectsBrokersInstancesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the given service instance from the system.
// The API call accepts both OSB style API and standard google style
// API
// resource path.
// i.e. both `projects/*/brokers/*/instances/*`
//  and `projects/*/brokers/*/v2/service_instances/*` are acceptable
// paths.
func (r *ProjectsBrokersInstancesService) Get(name string) *ProjectsBrokersInstancesGetCall {
	c := &ProjectsBrokersInstancesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersInstancesGetCall) Fields(s ...googleapi.Field) *ProjectsBrokersInstancesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsBrokersInstancesGetCall) IfNoneMatch(entityTag string) *ProjectsBrokersInstancesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersInstancesGetCall) Context(ctx context.Context) *ProjectsBrokersInstancesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersInstancesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersInstancesGetCall) doRequest(alt string) (*http.Response, error) {
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
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.instances.get"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.instances.get" call.
// Exactly one of *GoogleCloudServicebrokerV1beta1__ServiceInstance or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudServicebrokerV1beta1__ServiceInstance.ServerResponse.Heade
// r or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsBrokersInstancesGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudServicebrokerV1beta1__ServiceInstance, error) {
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
	ret := &GoogleCloudServicebrokerV1beta1__ServiceInstance{
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
	//   "description": "Gets the given service instance from the system.\nThe API call accepts both OSB style API and standard google style API\nresource path.\ni.e. both `projects/*/brokers/*/instances/*`\n and `projects/*/brokers/*/v2/service_instances/*` are acceptable paths.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers/{brokersId}/instances/{instancesId}",
	//   "httpMethod": "GET",
	//   "id": "servicebroker.projects.brokers.instances.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The resource name of the instance to return.\nName must match\n`projects/[PROJECT_ID]/brokers/[BROKER_ID]/` +\n`v2/service_instances/[INSTANCE_ID]`\nor\n`projects/[PROJECT_ID]/brokers/[BROKER_ID]/instances/[INSTANCE_ID]`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/brokers/[^/]+/instances/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__ServiceInstance"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.projects.brokers.instances.getLast_operation":

type ProjectsBrokersInstancesGetLastOperationCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetLastOperation: Returns the state of the last operation for the
// service instance.
// Only last (or current) operation can be polled.
func (r *ProjectsBrokersInstancesService) GetLastOperation(name string) *ProjectsBrokersInstancesGetLastOperationCall {
	c := &ProjectsBrokersInstancesGetLastOperationCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Operation sets the optional parameter "operation": If `operation` was
// returned during mutation operation, this field must be
// populated with the provided value.
func (c *ProjectsBrokersInstancesGetLastOperationCall) Operation(operation string) *ProjectsBrokersInstancesGetLastOperationCall {
	c.urlParams_.Set("operation", operation)
	return c
}

// PlanId sets the optional parameter "planId": Plan id.
func (c *ProjectsBrokersInstancesGetLastOperationCall) PlanId(planId string) *ProjectsBrokersInstancesGetLastOperationCall {
	c.urlParams_.Set("planId", planId)
	return c
}

// ServiceId sets the optional parameter "serviceId": Service id.
func (c *ProjectsBrokersInstancesGetLastOperationCall) ServiceId(serviceId string) *ProjectsBrokersInstancesGetLastOperationCall {
	c.urlParams_.Set("serviceId", serviceId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersInstancesGetLastOperationCall) Fields(s ...googleapi.Field) *ProjectsBrokersInstancesGetLastOperationCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsBrokersInstancesGetLastOperationCall) IfNoneMatch(entityTag string) *ProjectsBrokersInstancesGetLastOperationCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersInstancesGetLastOperationCall) Context(ctx context.Context) *ProjectsBrokersInstancesGetLastOperationCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersInstancesGetLastOperationCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersInstancesGetLastOperationCall) doRequest(alt string) (*http.Response, error) {
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
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}/last_operation")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.instances.getLast_operation"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.instances.getLast_operation" call.
// Exactly one of *GoogleCloudServicebrokerV1beta1__Operation or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudServicebrokerV1beta1__Operation.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsBrokersInstancesGetLastOperationCall) Do(opts ...googleapi.CallOption) (*GoogleCloudServicebrokerV1beta1__Operation, error) {
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
	ret := &GoogleCloudServicebrokerV1beta1__Operation{
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
	//   "description": "Returns the state of the last operation for the service instance.\nOnly last (or current) operation can be polled.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers/{brokersId}/instances/{instancesId}/last_operation",
	//   "httpMethod": "GET",
	//   "id": "servicebroker.projects.brokers.instances.getLast_operation",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name must match\n`projects/[PROJECT_ID]/brokers/[BROKER_ID]/v2/`+\n   `service_instances/[INSTANCE_ID]`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/brokers/[^/]+/instances/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "operation": {
	//       "description": "If `operation` was returned during mutation operation, this field must be\npopulated with the provided value.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "planId": {
	//       "description": "Plan id.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "serviceId": {
	//       "description": "Service id.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}/last_operation",
	//   "response": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.projects.brokers.instances.list":

type ProjectsBrokersInstancesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all the instances in the brokers
// This API is an extension and not part of the OSB spec.
// Hence the path is a standard Google API URL.
func (r *ProjectsBrokersInstancesService) List(parent string) *ProjectsBrokersInstancesListCall {
	c := &ProjectsBrokersInstancesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": Specifies the number
// of results to return per page. If there are fewer
// elements than the specified number, returns all elements.
//  Acceptable values are 0 to 200, inclusive. (Default: 100)
func (c *ProjectsBrokersInstancesListCall) PageSize(pageSize int64) *ProjectsBrokersInstancesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a page
// token to use. Set `pageToken` to a `nextPageToken`
// returned by a previous list request to get the next page of results.
func (c *ProjectsBrokersInstancesListCall) PageToken(pageToken string) *ProjectsBrokersInstancesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersInstancesListCall) Fields(s ...googleapi.Field) *ProjectsBrokersInstancesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsBrokersInstancesListCall) IfNoneMatch(entityTag string) *ProjectsBrokersInstancesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersInstancesListCall) Context(ctx context.Context) *ProjectsBrokersInstancesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersInstancesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersInstancesListCall) doRequest(alt string) (*http.Response, error) {
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
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/instances")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.instances.list"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.instances.list" call.
// Exactly one of
// *GoogleCloudServicebrokerV1beta1__ListServiceInstancesResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudServicebrokerV1beta1__ListServiceInstancesResponse.ServerR
// esponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsBrokersInstancesListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudServicebrokerV1beta1__ListServiceInstancesResponse, error) {
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
	ret := &GoogleCloudServicebrokerV1beta1__ListServiceInstancesResponse{
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
	//   "description": "Lists all the instances in the brokers\nThis API is an extension and not part of the OSB spec.\nHence the path is a standard Google API URL.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers/{brokersId}/instances",
	//   "httpMethod": "GET",
	//   "id": "servicebroker.projects.brokers.instances.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Specifies the number of results to return per page. If there are fewer\nelements than the specified number, returns all elements.\nOptional. Acceptable values are 0 to 200, inclusive. (Default: 100)",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a page token to use. Set `pageToken` to a `nextPageToken`\nreturned by a previous list request to get the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Parent must match `projects/[PROJECT_ID]/brokers/[BROKER_ID]`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/brokers/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/instances",
	//   "response": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__ListServiceInstancesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsBrokersInstancesListCall) Pages(ctx context.Context, f func(*GoogleCloudServicebrokerV1beta1__ListServiceInstancesResponse) error) error {
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

// method id "servicebroker.projects.brokers.instances.bindings.getLast_operation":

type ProjectsBrokersInstancesBindingsGetLastOperationCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetLastOperation: Returns the state of the last operation for the
// binding.
// Only last (or current) operation can be polled.
func (r *ProjectsBrokersInstancesBindingsService) GetLastOperation(name string) *ProjectsBrokersInstancesBindingsGetLastOperationCall {
	c := &ProjectsBrokersInstancesBindingsGetLastOperationCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Operation sets the optional parameter "operation": If `operation` was
// returned during mutation operation, this field must be
// populated with the provided value.
func (c *ProjectsBrokersInstancesBindingsGetLastOperationCall) Operation(operation string) *ProjectsBrokersInstancesBindingsGetLastOperationCall {
	c.urlParams_.Set("operation", operation)
	return c
}

// PlanId sets the optional parameter "planId": Plan id.
func (c *ProjectsBrokersInstancesBindingsGetLastOperationCall) PlanId(planId string) *ProjectsBrokersInstancesBindingsGetLastOperationCall {
	c.urlParams_.Set("planId", planId)
	return c
}

// ServiceId sets the optional parameter "serviceId": Service id.
func (c *ProjectsBrokersInstancesBindingsGetLastOperationCall) ServiceId(serviceId string) *ProjectsBrokersInstancesBindingsGetLastOperationCall {
	c.urlParams_.Set("serviceId", serviceId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersInstancesBindingsGetLastOperationCall) Fields(s ...googleapi.Field) *ProjectsBrokersInstancesBindingsGetLastOperationCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsBrokersInstancesBindingsGetLastOperationCall) IfNoneMatch(entityTag string) *ProjectsBrokersInstancesBindingsGetLastOperationCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersInstancesBindingsGetLastOperationCall) Context(ctx context.Context) *ProjectsBrokersInstancesBindingsGetLastOperationCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersInstancesBindingsGetLastOperationCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersInstancesBindingsGetLastOperationCall) doRequest(alt string) (*http.Response, error) {
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
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}/last_operation")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.instances.bindings.getLast_operation"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.instances.bindings.getLast_operation" call.
// Exactly one of *GoogleCloudServicebrokerV1beta1__Operation or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudServicebrokerV1beta1__Operation.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsBrokersInstancesBindingsGetLastOperationCall) Do(opts ...googleapi.CallOption) (*GoogleCloudServicebrokerV1beta1__Operation, error) {
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
	ret := &GoogleCloudServicebrokerV1beta1__Operation{
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
	//   "description": "Returns the state of the last operation for the binding.\nOnly last (or current) operation can be polled.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers/{brokersId}/instances/{instancesId}/bindings/{bindingsId}/last_operation",
	//   "httpMethod": "GET",
	//   "id": "servicebroker.projects.brokers.instances.bindings.getLast_operation",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name must match\n`projects/[PROJECT_ID]/brokers/[BROKER_ID]/v2/service_instances/[INSTANCE_ID]/service_binding/[BINDING_ID]`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/brokers/[^/]+/instances/[^/]+/bindings/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "operation": {
	//       "description": "If `operation` was returned during mutation operation, this field must be\npopulated with the provided value.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "planId": {
	//       "description": "Plan id.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "serviceId": {
	//       "description": "Service id.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}/last_operation",
	//   "response": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.projects.brokers.instances.bindings.list":

type ProjectsBrokersInstancesBindingsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all the bindings in the instance.
func (r *ProjectsBrokersInstancesBindingsService) List(parent string) *ProjectsBrokersInstancesBindingsListCall {
	c := &ProjectsBrokersInstancesBindingsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": Specifies the number
// of results to return per page. If there are fewer
// elements than the specified number, returns all elements.
//  Acceptable values are 0 to 200, inclusive. (Default: 100)
func (c *ProjectsBrokersInstancesBindingsListCall) PageSize(pageSize int64) *ProjectsBrokersInstancesBindingsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a page
// token to use. Set `pageToken` to a `nextPageToken`
// returned by a previous list request to get the next page of results.
func (c *ProjectsBrokersInstancesBindingsListCall) PageToken(pageToken string) *ProjectsBrokersInstancesBindingsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersInstancesBindingsListCall) Fields(s ...googleapi.Field) *ProjectsBrokersInstancesBindingsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsBrokersInstancesBindingsListCall) IfNoneMatch(entityTag string) *ProjectsBrokersInstancesBindingsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersInstancesBindingsListCall) Context(ctx context.Context) *ProjectsBrokersInstancesBindingsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersInstancesBindingsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersInstancesBindingsListCall) doRequest(alt string) (*http.Response, error) {
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
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/bindings")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.instances.bindings.list"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.instances.bindings.list" call.
// Exactly one of *GoogleCloudServicebrokerV1beta1__ListBindingsResponse
// or error will be non-nil. Any non-2xx status code is an error.
// Response headers are in either
// *GoogleCloudServicebrokerV1beta1__ListBindingsResponse.ServerResponse.
// Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsBrokersInstancesBindingsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudServicebrokerV1beta1__ListBindingsResponse, error) {
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
	ret := &GoogleCloudServicebrokerV1beta1__ListBindingsResponse{
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
	//   "description": "Lists all the bindings in the instance.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers/{brokersId}/instances/{instancesId}/bindings",
	//   "httpMethod": "GET",
	//   "id": "servicebroker.projects.brokers.instances.bindings.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Specifies the number of results to return per page. If there are fewer\nelements than the specified number, returns all elements.\nOptional. Acceptable values are 0 to 200, inclusive. (Default: 100)",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a page token to use. Set `pageToken` to a `nextPageToken`\nreturned by a previous list request to get the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Parent must match\n`projects/[PROJECT_ID]/brokers/[BROKER_ID]/` +\n`v2/service_instances/[INSTANCE_ID]`\nor\n`projects/[PROJECT_ID]/brokers/[BROKER_ID]/instances/[INSTANCE_ID]`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/brokers/[^/]+/instances/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/bindings",
	//   "response": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__ListBindingsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsBrokersInstancesBindingsListCall) Pages(ctx context.Context, f func(*GoogleCloudServicebrokerV1beta1__ListBindingsResponse) error) error {
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

// method id "servicebroker.projects.brokers.v2.catalog.list":

type ProjectsBrokersV2CatalogListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all the Services registered with this broker for
// consumption for
// given service registry broker, which contains an set of
// services.
// Note, that Service producer API is separate from Broker API.
func (r *ProjectsBrokersV2CatalogService) List(parent string) *ProjectsBrokersV2CatalogListCall {
	c := &ProjectsBrokersV2CatalogListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": Specifies the number
// of results to return per page. If there are fewer
// elements than the specified number, returns all elements.
//  If unset or 0, all the results will be returned.
func (c *ProjectsBrokersV2CatalogListCall) PageSize(pageSize int64) *ProjectsBrokersV2CatalogListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a page
// token to use. Set `pageToken` to a `nextPageToken`
// returned by a previous list request to get the next page of results.
func (c *ProjectsBrokersV2CatalogListCall) PageToken(pageToken string) *ProjectsBrokersV2CatalogListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersV2CatalogListCall) Fields(s ...googleapi.Field) *ProjectsBrokersV2CatalogListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsBrokersV2CatalogListCall) IfNoneMatch(entityTag string) *ProjectsBrokersV2CatalogListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersV2CatalogListCall) Context(ctx context.Context) *ProjectsBrokersV2CatalogListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersV2CatalogListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersV2CatalogListCall) doRequest(alt string) (*http.Response, error) {
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
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/v2/catalog")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.v2.catalog.list"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.v2.catalog.list" call.
// Exactly one of *GoogleCloudServicebrokerV1beta1__ListCatalogResponse
// or error will be non-nil. Any non-2xx status code is an error.
// Response headers are in either
// *GoogleCloudServicebrokerV1beta1__ListCatalogResponse.ServerResponse.H
// eader or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsBrokersV2CatalogListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudServicebrokerV1beta1__ListCatalogResponse, error) {
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
	ret := &GoogleCloudServicebrokerV1beta1__ListCatalogResponse{
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
	//   "description": "Lists all the Services registered with this broker for consumption for\ngiven service registry broker, which contains an set of services.\nNote, that Service producer API is separate from Broker API.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers/{brokersId}/v2/catalog",
	//   "httpMethod": "GET",
	//   "id": "servicebroker.projects.brokers.v2.catalog.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Specifies the number of results to return per page. If there are fewer\nelements than the specified number, returns all elements.\nOptional. If unset or 0, all the results will be returned.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a page token to use. Set `pageToken` to a `nextPageToken`\nreturned by a previous list request to get the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Parent must match `projects/[PROJECT_ID]/brokers/[BROKER_ID]`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/brokers/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/v2/catalog",
	//   "response": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__ListCatalogResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsBrokersV2CatalogListCall) Pages(ctx context.Context, f func(*GoogleCloudServicebrokerV1beta1__ListCatalogResponse) error) error {
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

// method id "servicebroker.projects.brokers.v2.service_instances.create":

type ProjectsBrokersV2ServiceInstancesCreateCall struct {
	s                                                *Service
	parent                                           string
	instanceId                                       string
	googlecloudservicebrokerv1beta1__serviceinstance *GoogleCloudServicebrokerV1beta1__ServiceInstance
	urlParams_                                       gensupport.URLParams
	ctx_                                             context.Context
	header_                                          http.Header
}

// Create: Provisions a service instance.
// If `request.accepts_incomplete` is false and Broker cannot execute
// request
// synchronously HTTP 422 error will be returned along
// with
// FAILED_PRECONDITION status.
// If `request.accepts_incomplete` is true and the Broker decides to
// execute
// resource asynchronously then HTTP 202 response code will be returned
// and a
// valid polling operation in the response will be included.
// If Broker executes the request synchronously and it succeeds HTTP
// 201
// response will be furnished.
// If identical instance exists, then HTTP 200 response will be
// returned.
// If an instance with identical ID but mismatching parameters exists,
// then
// HTTP 409 status code will be returned.
func (r *ProjectsBrokersV2ServiceInstancesService) Create(parent string, instanceId string, googlecloudservicebrokerv1beta1__serviceinstance *GoogleCloudServicebrokerV1beta1__ServiceInstance) *ProjectsBrokersV2ServiceInstancesCreateCall {
	c := &ProjectsBrokersV2ServiceInstancesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.instanceId = instanceId
	c.googlecloudservicebrokerv1beta1__serviceinstance = googlecloudservicebrokerv1beta1__serviceinstance
	return c
}

// AcceptsIncomplete sets the optional parameter "acceptsIncomplete":
// Value indicating that API client supports asynchronous operations.
// If
// Broker cannot execute the request synchronously HTTP 422 code will
// be
// returned to HTTP clients along with FAILED_PRECONDITION error.
// If true and broker will execute request asynchronously 202 HTTP code
// will
// be returned.
// This broker always requires this to be true as all mutator operations
// are
// asynchronous.
func (c *ProjectsBrokersV2ServiceInstancesCreateCall) AcceptsIncomplete(acceptsIncomplete bool) *ProjectsBrokersV2ServiceInstancesCreateCall {
	c.urlParams_.Set("acceptsIncomplete", fmt.Sprint(acceptsIncomplete))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersV2ServiceInstancesCreateCall) Fields(s ...googleapi.Field) *ProjectsBrokersV2ServiceInstancesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersV2ServiceInstancesCreateCall) Context(ctx context.Context) *ProjectsBrokersV2ServiceInstancesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersV2ServiceInstancesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersV2ServiceInstancesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudservicebrokerv1beta1__serviceinstance)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/v2/service_instances/{+instance_id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent":      c.parent,
		"instance_id": c.instanceId,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.v2.service_instances.create"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.v2.service_instances.create" call.
// Exactly one of
// *GoogleCloudServicebrokerV1beta1__CreateServiceInstanceResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudServicebrokerV1beta1__CreateServiceInstanceResponse.Server
// Response.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsBrokersV2ServiceInstancesCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudServicebrokerV1beta1__CreateServiceInstanceResponse, error) {
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
	ret := &GoogleCloudServicebrokerV1beta1__CreateServiceInstanceResponse{
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
	//   "description": "Provisions a service instance.\nIf `request.accepts_incomplete` is false and Broker cannot execute request\nsynchronously HTTP 422 error will be returned along with\nFAILED_PRECONDITION status.\nIf `request.accepts_incomplete` is true and the Broker decides to execute\nresource asynchronously then HTTP 202 response code will be returned and a\nvalid polling operation in the response will be included.\nIf Broker executes the request synchronously and it succeeds HTTP 201\nresponse will be furnished.\nIf identical instance exists, then HTTP 200 response will be returned.\nIf an instance with identical ID but mismatching parameters exists, then\nHTTP 409 status code will be returned.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers/{brokersId}/v2/service_instances/{service_instancesId}",
	//   "httpMethod": "PUT",
	//   "id": "servicebroker.projects.brokers.v2.service_instances.create",
	//   "parameterOrder": [
	//     "parent",
	//     "instance_id"
	//   ],
	//   "parameters": {
	//     "acceptsIncomplete": {
	//       "description": "Value indicating that API client supports asynchronous operations. If\nBroker cannot execute the request synchronously HTTP 422 code will be\nreturned to HTTP clients along with FAILED_PRECONDITION error.\nIf true and broker will execute request asynchronously 202 HTTP code will\nbe returned.\nThis broker always requires this to be true as all mutator operations are\nasynchronous.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "instance_id": {
	//       "description": "The id of the service instance. Must be unique within GCP project.\nMaximum length is 64, GUID recommended.\nRequired.",
	//       "location": "path",
	//       "pattern": "^[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Parent must match `projects/[PROJECT_ID]/brokers/[BROKER_ID]`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/brokers/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/v2/service_instances/{+instance_id}",
	//   "request": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__ServiceInstance"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__CreateServiceInstanceResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.projects.brokers.v2.service_instances.delete":

type ProjectsBrokersV2ServiceInstancesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deprovisions a service instance.
// For synchronous/asynchronous request details see
// CreateServiceInstance
// method.
// If service instance does not exist HTTP 410 status will be returned.
func (r *ProjectsBrokersV2ServiceInstancesService) Delete(name string) *ProjectsBrokersV2ServiceInstancesDeleteCall {
	c := &ProjectsBrokersV2ServiceInstancesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// AcceptsIncomplete sets the optional parameter "acceptsIncomplete":
// See CreateServiceInstanceRequest for details.
func (c *ProjectsBrokersV2ServiceInstancesDeleteCall) AcceptsIncomplete(acceptsIncomplete bool) *ProjectsBrokersV2ServiceInstancesDeleteCall {
	c.urlParams_.Set("acceptsIncomplete", fmt.Sprint(acceptsIncomplete))
	return c
}

// PlanId sets the optional parameter "planId": The plan id of the
// service instance.
func (c *ProjectsBrokersV2ServiceInstancesDeleteCall) PlanId(planId string) *ProjectsBrokersV2ServiceInstancesDeleteCall {
	c.urlParams_.Set("planId", planId)
	return c
}

// ServiceId sets the optional parameter "serviceId": The service id of
// the service instance.
func (c *ProjectsBrokersV2ServiceInstancesDeleteCall) ServiceId(serviceId string) *ProjectsBrokersV2ServiceInstancesDeleteCall {
	c.urlParams_.Set("serviceId", serviceId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersV2ServiceInstancesDeleteCall) Fields(s ...googleapi.Field) *ProjectsBrokersV2ServiceInstancesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersV2ServiceInstancesDeleteCall) Context(ctx context.Context) *ProjectsBrokersV2ServiceInstancesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersV2ServiceInstancesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersV2ServiceInstancesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.v2.service_instances.delete"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.v2.service_instances.delete" call.
// Exactly one of
// *GoogleCloudServicebrokerV1beta1__DeleteServiceInstanceResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudServicebrokerV1beta1__DeleteServiceInstanceResponse.Server
// Response.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsBrokersV2ServiceInstancesDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleCloudServicebrokerV1beta1__DeleteServiceInstanceResponse, error) {
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
	ret := &GoogleCloudServicebrokerV1beta1__DeleteServiceInstanceResponse{
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
	//   "description": "Deprovisions a service instance.\nFor synchronous/asynchronous request details see CreateServiceInstance\nmethod.\nIf service instance does not exist HTTP 410 status will be returned.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers/{brokersId}/v2/service_instances/{service_instancesId}",
	//   "httpMethod": "DELETE",
	//   "id": "servicebroker.projects.brokers.v2.service_instances.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "acceptsIncomplete": {
	//       "description": "See CreateServiceInstanceRequest for details.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "name": {
	//       "description": "Name must match\n`projects/[PROJECT_ID]/brokers/[BROKER_ID]/` +\n`v2/service_instances/[INSTANCE_ID]`\nor\n`projects/[PROJECT_ID]/brokers/[BROKER_ID]/instances/[INSTANCE_ID]`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/brokers/[^/]+/v2/service_instances/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "planId": {
	//       "description": "The plan id of the service instance.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "serviceId": {
	//       "description": "The service id of the service instance.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__DeleteServiceInstanceResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.projects.brokers.v2.service_instances.get":

type ProjectsBrokersV2ServiceInstancesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the given service instance from the system.
// The API call accepts both OSB style API and standard google style
// API
// resource path.
// i.e. both `projects/*/brokers/*/instances/*`
//  and `projects/*/brokers/*/v2/service_instances/*` are acceptable
// paths.
func (r *ProjectsBrokersV2ServiceInstancesService) Get(name string) *ProjectsBrokersV2ServiceInstancesGetCall {
	c := &ProjectsBrokersV2ServiceInstancesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersV2ServiceInstancesGetCall) Fields(s ...googleapi.Field) *ProjectsBrokersV2ServiceInstancesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsBrokersV2ServiceInstancesGetCall) IfNoneMatch(entityTag string) *ProjectsBrokersV2ServiceInstancesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersV2ServiceInstancesGetCall) Context(ctx context.Context) *ProjectsBrokersV2ServiceInstancesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersV2ServiceInstancesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersV2ServiceInstancesGetCall) doRequest(alt string) (*http.Response, error) {
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
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.v2.service_instances.get"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.v2.service_instances.get" call.
// Exactly one of *GoogleCloudServicebrokerV1beta1__ServiceInstance or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudServicebrokerV1beta1__ServiceInstance.ServerResponse.Heade
// r or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsBrokersV2ServiceInstancesGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudServicebrokerV1beta1__ServiceInstance, error) {
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
	ret := &GoogleCloudServicebrokerV1beta1__ServiceInstance{
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
	//   "description": "Gets the given service instance from the system.\nThe API call accepts both OSB style API and standard google style API\nresource path.\ni.e. both `projects/*/brokers/*/instances/*`\n and `projects/*/brokers/*/v2/service_instances/*` are acceptable paths.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers/{brokersId}/v2/service_instances/{service_instancesId}",
	//   "httpMethod": "GET",
	//   "id": "servicebroker.projects.brokers.v2.service_instances.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The resource name of the instance to return.\nName must match\n`projects/[PROJECT_ID]/brokers/[BROKER_ID]/` +\n`v2/service_instances/[INSTANCE_ID]`\nor\n`projects/[PROJECT_ID]/brokers/[BROKER_ID]/instances/[INSTANCE_ID]`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/brokers/[^/]+/v2/service_instances/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__ServiceInstance"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.projects.brokers.v2.service_instances.getLast_operation":

type ProjectsBrokersV2ServiceInstancesGetLastOperationCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetLastOperation: Returns the state of the last operation for the
// service instance.
// Only last (or current) operation can be polled.
func (r *ProjectsBrokersV2ServiceInstancesService) GetLastOperation(name string) *ProjectsBrokersV2ServiceInstancesGetLastOperationCall {
	c := &ProjectsBrokersV2ServiceInstancesGetLastOperationCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Operation sets the optional parameter "operation": If `operation` was
// returned during mutation operation, this field must be
// populated with the provided value.
func (c *ProjectsBrokersV2ServiceInstancesGetLastOperationCall) Operation(operation string) *ProjectsBrokersV2ServiceInstancesGetLastOperationCall {
	c.urlParams_.Set("operation", operation)
	return c
}

// PlanId sets the optional parameter "planId": Plan id.
func (c *ProjectsBrokersV2ServiceInstancesGetLastOperationCall) PlanId(planId string) *ProjectsBrokersV2ServiceInstancesGetLastOperationCall {
	c.urlParams_.Set("planId", planId)
	return c
}

// ServiceId sets the optional parameter "serviceId": Service id.
func (c *ProjectsBrokersV2ServiceInstancesGetLastOperationCall) ServiceId(serviceId string) *ProjectsBrokersV2ServiceInstancesGetLastOperationCall {
	c.urlParams_.Set("serviceId", serviceId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersV2ServiceInstancesGetLastOperationCall) Fields(s ...googleapi.Field) *ProjectsBrokersV2ServiceInstancesGetLastOperationCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsBrokersV2ServiceInstancesGetLastOperationCall) IfNoneMatch(entityTag string) *ProjectsBrokersV2ServiceInstancesGetLastOperationCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersV2ServiceInstancesGetLastOperationCall) Context(ctx context.Context) *ProjectsBrokersV2ServiceInstancesGetLastOperationCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersV2ServiceInstancesGetLastOperationCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersV2ServiceInstancesGetLastOperationCall) doRequest(alt string) (*http.Response, error) {
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
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}/last_operation")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.v2.service_instances.getLast_operation"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.v2.service_instances.getLast_operation" call.
// Exactly one of *GoogleCloudServicebrokerV1beta1__Operation or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudServicebrokerV1beta1__Operation.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsBrokersV2ServiceInstancesGetLastOperationCall) Do(opts ...googleapi.CallOption) (*GoogleCloudServicebrokerV1beta1__Operation, error) {
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
	ret := &GoogleCloudServicebrokerV1beta1__Operation{
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
	//   "description": "Returns the state of the last operation for the service instance.\nOnly last (or current) operation can be polled.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers/{brokersId}/v2/service_instances/{service_instancesId}/last_operation",
	//   "httpMethod": "GET",
	//   "id": "servicebroker.projects.brokers.v2.service_instances.getLast_operation",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name must match\n`projects/[PROJECT_ID]/brokers/[BROKER_ID]/v2/`+\n   `service_instances/[INSTANCE_ID]`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/brokers/[^/]+/v2/service_instances/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "operation": {
	//       "description": "If `operation` was returned during mutation operation, this field must be\npopulated with the provided value.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "planId": {
	//       "description": "Plan id.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "serviceId": {
	//       "description": "Service id.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}/last_operation",
	//   "response": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.projects.brokers.v2.service_instances.patch":

type ProjectsBrokersV2ServiceInstancesPatchCall struct {
	s                                                *Service
	name                                             string
	googlecloudservicebrokerv1beta1__serviceinstance *GoogleCloudServicebrokerV1beta1__ServiceInstance
	urlParams_                                       gensupport.URLParams
	ctx_                                             context.Context
	header_                                          http.Header
}

// Patch: Updates an existing service instance.
// See CreateServiceInstance for possible response codes.
func (r *ProjectsBrokersV2ServiceInstancesService) Patch(name string, googlecloudservicebrokerv1beta1__serviceinstance *GoogleCloudServicebrokerV1beta1__ServiceInstance) *ProjectsBrokersV2ServiceInstancesPatchCall {
	c := &ProjectsBrokersV2ServiceInstancesPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.googlecloudservicebrokerv1beta1__serviceinstance = googlecloudservicebrokerv1beta1__serviceinstance
	return c
}

// AcceptsIncomplete sets the optional parameter "acceptsIncomplete":
// See CreateServiceInstanceRequest for details.
func (c *ProjectsBrokersV2ServiceInstancesPatchCall) AcceptsIncomplete(acceptsIncomplete bool) *ProjectsBrokersV2ServiceInstancesPatchCall {
	c.urlParams_.Set("acceptsIncomplete", fmt.Sprint(acceptsIncomplete))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersV2ServiceInstancesPatchCall) Fields(s ...googleapi.Field) *ProjectsBrokersV2ServiceInstancesPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersV2ServiceInstancesPatchCall) Context(ctx context.Context) *ProjectsBrokersV2ServiceInstancesPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersV2ServiceInstancesPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersV2ServiceInstancesPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudservicebrokerv1beta1__serviceinstance)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.v2.service_instances.patch"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.v2.service_instances.patch" call.
// Exactly one of
// *GoogleCloudServicebrokerV1beta1__UpdateServiceInstanceResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudServicebrokerV1beta1__UpdateServiceInstanceResponse.Server
// Response.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsBrokersV2ServiceInstancesPatchCall) Do(opts ...googleapi.CallOption) (*GoogleCloudServicebrokerV1beta1__UpdateServiceInstanceResponse, error) {
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
	ret := &GoogleCloudServicebrokerV1beta1__UpdateServiceInstanceResponse{
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
	//   "description": "Updates an existing service instance.\nSee CreateServiceInstance for possible response codes.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers/{brokersId}/v2/service_instances/{service_instancesId}",
	//   "httpMethod": "PATCH",
	//   "id": "servicebroker.projects.brokers.v2.service_instances.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "acceptsIncomplete": {
	//       "description": "See CreateServiceInstanceRequest for details.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "name": {
	//       "description": "Name must match\n`projects/[PROJECT_ID]/brokers/[BROKER_ID]/v2/service_instances/[INSTANCE_ID]`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/brokers/[^/]+/v2/service_instances/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "request": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__ServiceInstance"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__UpdateServiceInstanceResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.projects.brokers.v2.service_instances.service_bindings.create":

type ProjectsBrokersV2ServiceInstancesServiceBindingsCreateCall struct {
	s                                        *Service
	parent                                   string
	bindingId                                string
	googlecloudservicebrokerv1beta1__binding *GoogleCloudServicebrokerV1beta1__Binding
	urlParams_                               gensupport.URLParams
	ctx_                                     context.Context
	header_                                  http.Header
}

// Create: CreateBinding generates a service binding to an existing
// service instance.
// See ProviServiceInstance for async operation details.
func (r *ProjectsBrokersV2ServiceInstancesServiceBindingsService) Create(parent string, bindingId string, googlecloudservicebrokerv1beta1__binding *GoogleCloudServicebrokerV1beta1__Binding) *ProjectsBrokersV2ServiceInstancesServiceBindingsCreateCall {
	c := &ProjectsBrokersV2ServiceInstancesServiceBindingsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.bindingId = bindingId
	c.googlecloudservicebrokerv1beta1__binding = googlecloudservicebrokerv1beta1__binding
	return c
}

// AcceptsIncomplete sets the optional parameter "acceptsIncomplete":
// See CreateServiceInstanceRequest for details.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsCreateCall) AcceptsIncomplete(acceptsIncomplete bool) *ProjectsBrokersV2ServiceInstancesServiceBindingsCreateCall {
	c.urlParams_.Set("acceptsIncomplete", fmt.Sprint(acceptsIncomplete))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsCreateCall) Fields(s ...googleapi.Field) *ProjectsBrokersV2ServiceInstancesServiceBindingsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsCreateCall) Context(ctx context.Context) *ProjectsBrokersV2ServiceInstancesServiceBindingsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudservicebrokerv1beta1__binding)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/service_bindings/{+binding_id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent":     c.parent,
		"binding_id": c.bindingId,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.v2.service_instances.service_bindings.create"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.v2.service_instances.service_bindings.create" call.
// Exactly one of
// *GoogleCloudServicebrokerV1beta1__CreateBindingResponse or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleCloudServicebrokerV1beta1__CreateBindingResponse.ServerResponse
// .Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudServicebrokerV1beta1__CreateBindingResponse, error) {
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
	ret := &GoogleCloudServicebrokerV1beta1__CreateBindingResponse{
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
	//   "description": "CreateBinding generates a service binding to an existing service instance.\nSee ProviServiceInstance for async operation details.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers/{brokersId}/v2/service_instances/{service_instancesId}/service_bindings/{service_bindingsId}",
	//   "httpMethod": "PUT",
	//   "id": "servicebroker.projects.brokers.v2.service_instances.service_bindings.create",
	//   "parameterOrder": [
	//     "parent",
	//     "binding_id"
	//   ],
	//   "parameters": {
	//     "acceptsIncomplete": {
	//       "description": "See CreateServiceInstanceRequest for details.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "binding_id": {
	//       "description": "The id of the binding. Must be unique within GCP project.\nMaximum length is 64, GUID recommended.\nRequired.",
	//       "location": "path",
	//       "pattern": "^[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The GCP container.\nMust match\n`projects/[PROJECT_ID]/brokers/[BROKER_ID]/v2/service_instances/[INSTANCE_ID]`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/brokers/[^/]+/v2/service_instances/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/service_bindings/{+binding_id}",
	//   "request": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__Binding"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__CreateBindingResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.projects.brokers.v2.service_instances.service_bindings.delete":

type ProjectsBrokersV2ServiceInstancesServiceBindingsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Unbinds from a service instance.
// For synchronous/asynchronous request details see
// CreateServiceInstance
// method.
// If binding does not exist HTTP 410 status will be returned.
func (r *ProjectsBrokersV2ServiceInstancesServiceBindingsService) Delete(name string) *ProjectsBrokersV2ServiceInstancesServiceBindingsDeleteCall {
	c := &ProjectsBrokersV2ServiceInstancesServiceBindingsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// AcceptsIncomplete sets the optional parameter "acceptsIncomplete":
// See CreateServiceInstanceRequest for details.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsDeleteCall) AcceptsIncomplete(acceptsIncomplete bool) *ProjectsBrokersV2ServiceInstancesServiceBindingsDeleteCall {
	c.urlParams_.Set("acceptsIncomplete", fmt.Sprint(acceptsIncomplete))
	return c
}

// PlanId sets the optional parameter "planId": The plan id of the
// service instance.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsDeleteCall) PlanId(planId string) *ProjectsBrokersV2ServiceInstancesServiceBindingsDeleteCall {
	c.urlParams_.Set("planId", planId)
	return c
}

// ServiceId sets the optional parameter "serviceId": Additional query
// parameter hints.
// The service id of the service instance.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsDeleteCall) ServiceId(serviceId string) *ProjectsBrokersV2ServiceInstancesServiceBindingsDeleteCall {
	c.urlParams_.Set("serviceId", serviceId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsDeleteCall) Fields(s ...googleapi.Field) *ProjectsBrokersV2ServiceInstancesServiceBindingsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsDeleteCall) Context(ctx context.Context) *ProjectsBrokersV2ServiceInstancesServiceBindingsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.v2.service_instances.service_bindings.delete"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.v2.service_instances.service_bindings.delete" call.
// Exactly one of
// *GoogleCloudServicebrokerV1beta1__DeleteBindingResponse or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleCloudServicebrokerV1beta1__DeleteBindingResponse.ServerResponse
// .Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleCloudServicebrokerV1beta1__DeleteBindingResponse, error) {
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
	ret := &GoogleCloudServicebrokerV1beta1__DeleteBindingResponse{
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
	//   "description": "Unbinds from a service instance.\nFor synchronous/asynchronous request details see CreateServiceInstance\nmethod.\nIf binding does not exist HTTP 410 status will be returned.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers/{brokersId}/v2/service_instances/{service_instancesId}/service_bindings/{service_bindingsId}",
	//   "httpMethod": "DELETE",
	//   "id": "servicebroker.projects.brokers.v2.service_instances.service_bindings.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "acceptsIncomplete": {
	//       "description": "See CreateServiceInstanceRequest for details.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "name": {
	//       "description": "Name must match\n`projects/[PROJECT_ID]/brokers/[BROKER_ID]/`\n`v2/service_instances/[INSTANCE_ID]/service_bindings/[BINDING_ID]`\nor\n`projects/[PROJECT_ID]/brokers/[BROKER_ID]/`\n`/instances/[INSTANCE_ID]/bindings/[BINDING_ID]`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/brokers/[^/]+/v2/service_instances/[^/]+/service_bindings/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "planId": {
	//       "description": "The plan id of the service instance.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "serviceId": {
	//       "description": "Additional query parameter hints.\nThe service id of the service instance.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__DeleteBindingResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.projects.brokers.v2.service_instances.service_bindings.get":

type ProjectsBrokersV2ServiceInstancesServiceBindingsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: GetBinding returns the binding information.
func (r *ProjectsBrokersV2ServiceInstancesServiceBindingsService) Get(name string) *ProjectsBrokersV2ServiceInstancesServiceBindingsGetCall {
	c := &ProjectsBrokersV2ServiceInstancesServiceBindingsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// PlanId sets the optional parameter "planId": Plan id.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsGetCall) PlanId(planId string) *ProjectsBrokersV2ServiceInstancesServiceBindingsGetCall {
	c.urlParams_.Set("planId", planId)
	return c
}

// ServiceId sets the optional parameter "serviceId": Service id.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsGetCall) ServiceId(serviceId string) *ProjectsBrokersV2ServiceInstancesServiceBindingsGetCall {
	c.urlParams_.Set("serviceId", serviceId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsGetCall) Fields(s ...googleapi.Field) *ProjectsBrokersV2ServiceInstancesServiceBindingsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsGetCall) IfNoneMatch(entityTag string) *ProjectsBrokersV2ServiceInstancesServiceBindingsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsGetCall) Context(ctx context.Context) *ProjectsBrokersV2ServiceInstancesServiceBindingsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsGetCall) doRequest(alt string) (*http.Response, error) {
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
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.v2.service_instances.service_bindings.get"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.v2.service_instances.service_bindings.get" call.
// Exactly one of *GoogleCloudServicebrokerV1beta1__GetBindingResponse
// or error will be non-nil. Any non-2xx status code is an error.
// Response headers are in either
// *GoogleCloudServicebrokerV1beta1__GetBindingResponse.ServerResponse.He
// ader or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudServicebrokerV1beta1__GetBindingResponse, error) {
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
	ret := &GoogleCloudServicebrokerV1beta1__GetBindingResponse{
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
	//   "description": "GetBinding returns the binding information.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers/{brokersId}/v2/service_instances/{service_instancesId}/service_bindings/{service_bindingsId}",
	//   "httpMethod": "GET",
	//   "id": "servicebroker.projects.brokers.v2.service_instances.service_bindings.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name must match\n`projects/[PROJECT_ID]/brokers/[BROKER_ID]/v2/service_instances/[INSTANCE_ID]/service_bindings`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/brokers/[^/]+/v2/service_instances/[^/]+/service_bindings/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "planId": {
	//       "description": "Plan id.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "serviceId": {
	//       "description": "Service id.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__GetBindingResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.projects.brokers.v2.service_instances.service_bindings.getLast_operation":

type ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetLastOperation: Returns the state of the last operation for the
// binding.
// Only last (or current) operation can be polled.
func (r *ProjectsBrokersV2ServiceInstancesServiceBindingsService) GetLastOperation(name string) *ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall {
	c := &ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Operation sets the optional parameter "operation": If `operation` was
// returned during mutation operation, this field must be
// populated with the provided value.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall) Operation(operation string) *ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall {
	c.urlParams_.Set("operation", operation)
	return c
}

// PlanId sets the optional parameter "planId": Plan id.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall) PlanId(planId string) *ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall {
	c.urlParams_.Set("planId", planId)
	return c
}

// ServiceId sets the optional parameter "serviceId": Service id.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall) ServiceId(serviceId string) *ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall {
	c.urlParams_.Set("serviceId", serviceId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall) Fields(s ...googleapi.Field) *ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall) IfNoneMatch(entityTag string) *ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall) Context(ctx context.Context) *ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall) doRequest(alt string) (*http.Response, error) {
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
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}/last_operation")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.projects.brokers.v2.service_instances.service_bindings.getLast_operation"), c.s.client, req)
}

// Do executes the "servicebroker.projects.brokers.v2.service_instances.service_bindings.getLast_operation" call.
// Exactly one of *GoogleCloudServicebrokerV1beta1__Operation or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudServicebrokerV1beta1__Operation.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsBrokersV2ServiceInstancesServiceBindingsGetLastOperationCall) Do(opts ...googleapi.CallOption) (*GoogleCloudServicebrokerV1beta1__Operation, error) {
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
	ret := &GoogleCloudServicebrokerV1beta1__Operation{
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
	//   "description": "Returns the state of the last operation for the binding.\nOnly last (or current) operation can be polled.",
	//   "flatPath": "v1beta1/projects/{projectsId}/brokers/{brokersId}/v2/service_instances/{service_instancesId}/service_bindings/{service_bindingsId}/last_operation",
	//   "httpMethod": "GET",
	//   "id": "servicebroker.projects.brokers.v2.service_instances.service_bindings.getLast_operation",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Name must match\n`projects/[PROJECT_ID]/brokers/[BROKER_ID]/v2/service_instances/[INSTANCE_ID]/service_binding/[BINDING_ID]`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/brokers/[^/]+/v2/service_instances/[^/]+/service_bindings/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "operation": {
	//       "description": "If `operation` was returned during mutation operation, this field must be\npopulated with the provided value.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "planId": {
	//       "description": "Plan id.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "serviceId": {
	//       "description": "Service id.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}/last_operation",
	//   "response": {
	//     "$ref": "GoogleCloudServicebrokerV1beta1__Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.getIamPolicy":

type V1beta1GetIamPolicyCall struct {
	s            *Service
	resource     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetIamPolicy: Gets the access control policy for a resource.
// Returns an empty policy if the resource exists and does not have a
// policy
// set.
func (r *V1beta1Service) GetIamPolicy(resource string) *V1beta1GetIamPolicyCall {
	c := &V1beta1GetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1beta1GetIamPolicyCall) Fields(s ...googleapi.Field) *V1beta1GetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *V1beta1GetIamPolicyCall) IfNoneMatch(entityTag string) *V1beta1GetIamPolicyCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V1beta1GetIamPolicyCall) Context(ctx context.Context) *V1beta1GetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *V1beta1GetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *V1beta1GetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
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
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+resource}:getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.getIamPolicy"), c.s.client, req)
}

// Do executes the "servicebroker.getIamPolicy" call.
// Exactly one of *GoogleIamV1__Policy or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleIamV1__Policy.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *V1beta1GetIamPolicyCall) Do(opts ...googleapi.CallOption) (*GoogleIamV1__Policy, error) {
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
	ret := &GoogleIamV1__Policy{
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
	//   "description": "Gets the access control policy for a resource.\nReturns an empty policy if the resource exists and does not have a policy\nset.",
	//   "flatPath": "v1beta1/{v1beta1Id}:getIamPolicy",
	//   "httpMethod": "GET",
	//   "id": "servicebroker.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+resource}:getIamPolicy",
	//   "response": {
	//     "$ref": "GoogleIamV1__Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.setIamPolicy":

type V1beta1SetIamPolicyCall struct {
	s                                *Service
	resource                         string
	googleiamv1__setiampolicyrequest *GoogleIamV1__SetIamPolicyRequest
	urlParams_                       gensupport.URLParams
	ctx_                             context.Context
	header_                          http.Header
}

// SetIamPolicy: Sets the access control policy on the specified
// resource. Replaces any
// existing policy.
func (r *V1beta1Service) SetIamPolicy(resource string, googleiamv1__setiampolicyrequest *GoogleIamV1__SetIamPolicyRequest) *V1beta1SetIamPolicyCall {
	c := &V1beta1SetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.googleiamv1__setiampolicyrequest = googleiamv1__setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1beta1SetIamPolicyCall) Fields(s ...googleapi.Field) *V1beta1SetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V1beta1SetIamPolicyCall) Context(ctx context.Context) *V1beta1SetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *V1beta1SetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *V1beta1SetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleiamv1__setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+resource}:setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.setIamPolicy"), c.s.client, req)
}

// Do executes the "servicebroker.setIamPolicy" call.
// Exactly one of *GoogleIamV1__Policy or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleIamV1__Policy.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *V1beta1SetIamPolicyCall) Do(opts ...googleapi.CallOption) (*GoogleIamV1__Policy, error) {
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
	ret := &GoogleIamV1__Policy{
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
	//   "description": "Sets the access control policy on the specified resource. Replaces any\nexisting policy.",
	//   "flatPath": "v1beta1/{v1beta1Id}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "servicebroker.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being specified.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+resource}:setIamPolicy",
	//   "request": {
	//     "$ref": "GoogleIamV1__SetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleIamV1__Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.testIamPermissions":

type V1beta1TestIamPermissionsCall struct {
	s                                      *Service
	resource                               string
	googleiamv1__testiampermissionsrequest *GoogleIamV1__TestIamPermissionsRequest
	urlParams_                             gensupport.URLParams
	ctx_                                   context.Context
	header_                                http.Header
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified resource.
// If the resource does not exist, this will return an empty set
// of
// permissions, not a NOT_FOUND error.
//
// Note: This operation is designed to be used for building
// permission-aware
// UIs and command-line tools, not for authorization checking. This
// operation
// may "fail open" without warning.
func (r *V1beta1Service) TestIamPermissions(resource string, googleiamv1__testiampermissionsrequest *GoogleIamV1__TestIamPermissionsRequest) *V1beta1TestIamPermissionsCall {
	c := &V1beta1TestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.googleiamv1__testiampermissionsrequest = googleiamv1__testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1beta1TestIamPermissionsCall) Fields(s ...googleapi.Field) *V1beta1TestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V1beta1TestIamPermissionsCall) Context(ctx context.Context) *V1beta1TestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *V1beta1TestIamPermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *V1beta1TestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleiamv1__testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+resource}:testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "servicebroker.testIamPermissions"), c.s.client, req)
}

// Do executes the "servicebroker.testIamPermissions" call.
// Exactly one of *GoogleIamV1__TestIamPermissionsResponse or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleIamV1__TestIamPermissionsResponse.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *V1beta1TestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*GoogleIamV1__TestIamPermissionsResponse, error) {
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
	ret := &GoogleIamV1__TestIamPermissionsResponse{
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
	//   "description": "Returns permissions that a caller has on the specified resource.\nIf the resource does not exist, this will return an empty set of\npermissions, not a NOT_FOUND error.\n\nNote: This operation is designed to be used for building permission-aware\nUIs and command-line tools, not for authorization checking. This operation\nmay \"fail open\" without warning.",
	//   "flatPath": "v1beta1/{v1beta1Id}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "servicebroker.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy detail is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+resource}:testIamPermissions",
	//   "request": {
	//     "$ref": "GoogleIamV1__TestIamPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleIamV1__TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
