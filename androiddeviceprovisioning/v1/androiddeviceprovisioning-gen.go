// Package androiddeviceprovisioning provides access to the Android Device Provisioning Partner API.
//
// See https://developers.google.com/zero-touch/
//
// Usage example:
//
//   import "google.golang.org/api/androiddeviceprovisioning/v1"
//   ...
//   androiddeviceprovisioningService, err := androiddeviceprovisioning.New(oauthHttpClient)
package androiddeviceprovisioning // import "google.golang.org/api/androiddeviceprovisioning/v1"

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

const apiId = "androiddeviceprovisioning:v1"
const apiName = "androiddeviceprovisioning"
const apiVersion = "v1"
const basePath = "https://androiddeviceprovisioning.googleapis.com/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Customers = NewCustomersService(s)
	s.Operations = NewOperationsService(s)
	s.Partners = NewPartnersService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Customers *CustomersService

	Operations *OperationsService

	Partners *PartnersService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewCustomersService(s *Service) *CustomersService {
	rs := &CustomersService{s: s}
	rs.Configurations = NewCustomersConfigurationsService(s)
	rs.Devices = NewCustomersDevicesService(s)
	rs.Dpcs = NewCustomersDpcsService(s)
	return rs
}

type CustomersService struct {
	s *Service

	Configurations *CustomersConfigurationsService

	Devices *CustomersDevicesService

	Dpcs *CustomersDpcsService
}

func NewCustomersConfigurationsService(s *Service) *CustomersConfigurationsService {
	rs := &CustomersConfigurationsService{s: s}
	return rs
}

type CustomersConfigurationsService struct {
	s *Service
}

func NewCustomersDevicesService(s *Service) *CustomersDevicesService {
	rs := &CustomersDevicesService{s: s}
	return rs
}

type CustomersDevicesService struct {
	s *Service
}

func NewCustomersDpcsService(s *Service) *CustomersDpcsService {
	rs := &CustomersDpcsService{s: s}
	return rs
}

type CustomersDpcsService struct {
	s *Service
}

func NewOperationsService(s *Service) *OperationsService {
	rs := &OperationsService{s: s}
	return rs
}

type OperationsService struct {
	s *Service
}

func NewPartnersService(s *Service) *PartnersService {
	rs := &PartnersService{s: s}
	rs.Customers = NewPartnersCustomersService(s)
	rs.Devices = NewPartnersDevicesService(s)
	rs.Vendors = NewPartnersVendorsService(s)
	return rs
}

type PartnersService struct {
	s *Service

	Customers *PartnersCustomersService

	Devices *PartnersDevicesService

	Vendors *PartnersVendorsService
}

func NewPartnersCustomersService(s *Service) *PartnersCustomersService {
	rs := &PartnersCustomersService{s: s}
	return rs
}

type PartnersCustomersService struct {
	s *Service
}

func NewPartnersDevicesService(s *Service) *PartnersDevicesService {
	rs := &PartnersDevicesService{s: s}
	return rs
}

type PartnersDevicesService struct {
	s *Service
}

func NewPartnersVendorsService(s *Service) *PartnersVendorsService {
	rs := &PartnersVendorsService{s: s}
	rs.Customers = NewPartnersVendorsCustomersService(s)
	return rs
}

type PartnersVendorsService struct {
	s *Service

	Customers *PartnersVendorsCustomersService
}

func NewPartnersVendorsCustomersService(s *Service) *PartnersVendorsCustomersService {
	rs := &PartnersVendorsCustomersService{s: s}
	return rs
}

type PartnersVendorsCustomersService struct {
	s *Service
}

// ClaimDeviceRequest: Request message to claim a device on behalf of a
// customer.
type ClaimDeviceRequest struct {
	// CustomerId: Required. The ID of the customer for whom the device is
	// being claimed.
	CustomerId int64 `json:"customerId,omitempty,string"`

	// DeviceIdentifier: Required. The device identifier of the device to
	// claim.
	DeviceIdentifier *DeviceIdentifier `json:"deviceIdentifier,omitempty"`

	// SectionType: Required. The section type of the device's provisioning
	// record.
	//
	// Possible values:
	//   "SECTION_TYPE_UNSPECIFIED" - Unspecified section type.
	//   "SECTION_TYPE_SIM_LOCK" - SIM-lock section type.
	//   "SECTION_TYPE_ZERO_TOUCH" - Zero-touch enrollment section type.
	SectionType string `json:"sectionType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CustomerId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CustomerId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClaimDeviceRequest) MarshalJSON() ([]byte, error) {
	type NoMethod ClaimDeviceRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClaimDeviceResponse: Response message containing device id of the
// claim.
type ClaimDeviceResponse struct {
	// DeviceId: The device ID of the claimed device.
	DeviceId int64 `json:"deviceId,omitempty,string"`

	// DeviceName: The resource name of the device in the
	// format
	// `partners/[PARTNER_ID]/devices/[DEVICE_ID]`.
	DeviceName string `json:"deviceName,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DeviceId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeviceId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClaimDeviceResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ClaimDeviceResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClaimDevicesRequest: Request to claim devices asynchronously in
// batch. Claiming a device adds the
// device to zero-touch enrollment and shows the device in the
// customer's view
// of the portal.
type ClaimDevicesRequest struct {
	// Claims: Required. A list of device claims.
	Claims []*PartnerClaim `json:"claims,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Claims") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Claims") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClaimDevicesRequest) MarshalJSON() ([]byte, error) {
	type NoMethod ClaimDevicesRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Company: A reseller, vendor, or customer in the zero-touch reseller
// and customer APIs.
type Company struct {
	// AdminEmails: Input only. Optional. Email address of customer's users
	// in the admin role.
	// Each email address must be associated with a Google Account.
	AdminEmails []string `json:"adminEmails,omitempty"`

	// CompanyId: Output only. The ID of the company. Assigned by the
	// server.
	CompanyId int64 `json:"companyId,omitempty,string"`

	// CompanyName: Required. The name of the company. For example _XYZ
	// Corp_. Displayed to the
	// company's employees in the zero-touch enrollment portal.
	CompanyName string `json:"companyName,omitempty"`

	// Name: Output only. The API resource name of the company. The resource
	// name is one
	// of the following formats:
	//
	// * `partners/[PARTNER_ID]/customers/[CUSTOMER_ID]`
	// * `partners/[PARTNER_ID]/vendors/[VENDOR_ID]`
	// *
	// `partners/[PARTNER_ID]/vendors/[VENDOR_ID]/customers/[CUSTOMER_ID]`
	//
	// A
	// ssigned by the server.
	Name string `json:"name,omitempty"`

	// OwnerEmails: Input only. Email address of customer's users in the
	// owner role. At least
	// one `owner_email` is required. Each email address must be associated
	// with a
	// Google Account. Owners share the same access as admins but can also
	// add,
	// delete, and edit your organization's portal users.
	OwnerEmails []string `json:"ownerEmails,omitempty"`

	// TermsStatus: Output only. Whether any user from the company has
	// accepted the latest
	// Terms of Service (ToS). See
	// TermsStatus.
	//
	// Possible values:
	//   "TERMS_STATUS_UNSPECIFIED" - Default value. This value should never
	// be set if the enum is present.
	//   "TERMS_STATUS_NOT_ACCEPTED" - None of the company's users have
	// accepted the ToS.
	//   "TERMS_STATUS_ACCEPTED" - One (or more) of the company's users has
	// accepted the ToS.
	//   "TERMS_STATUS_STALE" - None of the company's users has accepted the
	// current ToS but at least one
	// user accepted a previous ToS.
	TermsStatus string `json:"termsStatus,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AdminEmails") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AdminEmails") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Company) MarshalJSON() ([]byte, error) {
	type NoMethod Company
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Configuration: A configuration collects the provisioning options for
// Android devices. Each
// configuration combines the following:
//
// * The EMM device policy controller (DPC) installed on the devices.
// * EMM policies enforced on the devices.
// * Metadata displayed on the device to help users during
// setup.
//
// Customers can add as many configurations as they need. However,
// zero-touch
// enrollment works best when a customer sets a default configuration
// that's
// applied to any new devices the organization purchases.
type Configuration struct {
	// CompanyName: Required. The name of the organization. Zero-touch
	// enrollment shows this
	// organization name to device users during device provisioning.
	CompanyName string `json:"companyName,omitempty"`

	// ConfigurationId: Output only. The ID of the configuration. Assigned
	// by the server.
	ConfigurationId int64 `json:"configurationId,omitempty,string"`

	// ConfigurationName: Required. A short name that describes the
	// configuration's purpose. For
	// example, _Sales team_ or _Temporary employees_. The zero-touch
	// enrollment
	// portal displays this name to IT admins.
	ConfigurationName string `json:"configurationName,omitempty"`

	// ContactEmail: Required. The email address that device users can
	// contact to get help.
	// Zero-touch enrollment shows this email address to device users
	// before
	// device provisioning. The value is validated on input.
	ContactEmail string `json:"contactEmail,omitempty"`

	// ContactPhone: Required. The telephone number that device users can
	// call, using another
	// device, to get help. Zero-touch enrollment shows this number to
	// device
	// users before device provisioning. Accepts numerals, spaces, the plus
	// sign,
	// hyphens, and parentheses.
	ContactPhone string `json:"contactPhone,omitempty"`

	// CustomMessage: A message, containing one or two sentences, to help
	// device users get help
	// or give them more details about what’s happening to their
	// device.
	// Zero-touch enrollment shows this message before the device is
	// provisioned.
	CustomMessage string `json:"customMessage,omitempty"`

	// DpcExtras: The JSON-formatted EMM provisioning extras that are passed
	// to the DPC.
	DpcExtras string `json:"dpcExtras,omitempty"`

	// DpcResourcePath: Required. The resource name of the selected DPC
	// (device policy controller)
	// in the format `customers/[CUSTOMER_ID]/dpcs/*`. To list the supported
	// DPCs,
	// call
	// `customers.dpcs.list`.
	DpcResourcePath string `json:"dpcResourcePath,omitempty"`

	// IsDefault: Required. Whether this is the default configuration that
	// zero-touch
	// enrollment applies to any new devices the organization purchases in
	// the
	// future. Only one customer configuration can be the default. Setting
	// this
	// value to `true`, changes the previous default configuration's
	// `isDefault`
	// value to `false`.
	IsDefault bool `json:"isDefault,omitempty"`

	// Name: Output only. The API resource name in the
	// format
	// `customers/[CUSTOMER_ID]/configurations/[CONFIGURATION_ID]`. Assigned
	// by
	// the server.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CompanyName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CompanyName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Configuration) MarshalJSON() ([]byte, error) {
	type NoMethod Configuration
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreateCustomerRequest: Request message to create a customer.
type CreateCustomerRequest struct {
	// Customer: Required. The company data to populate the new customer.
	// Must contain a
	// value for `companyName` and at least one `owner_email` that's
	// associated
	// with a Google Account. The values for `companyId` and `name` must be
	// empty.
	Customer *Company `json:"customer,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Customer") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Customer") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CreateCustomerRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CreateCustomerRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CustomerApplyConfigurationRequest: Request message for customer to
// assign a configuration to device.
type CustomerApplyConfigurationRequest struct {
	// Configuration: Required. The configuration applied to the device in
	// the
	// format
	// `customers/[CUSTOMER_ID]/configurations/[CONFIGURATION_ID]`.
	Configuration string `json:"configuration,omitempty"`

	// Device: Required. The device the configuration is applied to.
	Device *DeviceReference `json:"device,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Configuration") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Configuration") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CustomerApplyConfigurationRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CustomerApplyConfigurationRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CustomerListConfigurationsResponse: Response message of customer's
// listing configuration.
type CustomerListConfigurationsResponse struct {
	// Configurations: The configurations.
	Configurations []*Configuration `json:"configurations,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Configurations") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Configurations") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *CustomerListConfigurationsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod CustomerListConfigurationsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CustomerListCustomersResponse: Response message for listing my
// customers.
type CustomerListCustomersResponse struct {
	// Customers: The customer accounts the calling user is a member of.
	Customers []*Company `json:"customers,omitempty"`

	// NextPageToken: A token used to access the next page of results.
	// Omitted if no further
	// results are available.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Customers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Customers") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CustomerListCustomersResponse) MarshalJSON() ([]byte, error) {
	type NoMethod CustomerListCustomersResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CustomerListDevicesResponse: Response message of customer's liting
// devices.
type CustomerListDevicesResponse struct {
	// Devices: The customer's devices.
	Devices []*Device `json:"devices,omitempty"`

	// NextPageToken: A token used to access the next page of results.
	// Omitted if no further
	// results are available.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Devices") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Devices") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CustomerListDevicesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod CustomerListDevicesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CustomerListDpcsResponse: Response message of customer's listing
// DPCs.
type CustomerListDpcsResponse struct {
	// Dpcs: The list of DPCs available to the customer that support
	// zero-touch
	// enrollment.
	Dpcs []*Dpc `json:"dpcs,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Dpcs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Dpcs") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CustomerListDpcsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod CustomerListDpcsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CustomerRemoveConfigurationRequest: Request message for customer to
// remove the configuration from device.
type CustomerRemoveConfigurationRequest struct {
	// Device: Required. The device to remove the configuration from.
	Device *DeviceReference `json:"device,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Device") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Device") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CustomerRemoveConfigurationRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CustomerRemoveConfigurationRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CustomerUnclaimDeviceRequest: Request message for customer to unclaim
// a device.
type CustomerUnclaimDeviceRequest struct {
	// Device: Required. The device to unclaim.
	Device *DeviceReference `json:"device,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Device") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Device") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CustomerUnclaimDeviceRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CustomerUnclaimDeviceRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Device: An Android device registered for zero-touch enrollment.
type Device struct {
	// Claims: Output only. The provisioning claims for a device. Devices
	// claimed for
	// zero-touch enrollment have a claim with the type
	// `SECTION_TYPE_ZERO_TOUCH`.
	// Call
	// `partners.devices.unclaim`
	// or
	// `partner
	// s.devices.unclaimAsync`
	// to remove the device from zero-touch enrollment.
	Claims []*DeviceClaim `json:"claims,omitempty"`

	// Configuration: Not available to resellers.
	Configuration string `json:"configuration,omitempty"`

	// DeviceId: Output only. The ID of the device. Assigned by the server.
	DeviceId int64 `json:"deviceId,omitempty,string"`

	// DeviceIdentifier: The hardware IDs that identify a manufactured
	// device. To learn more,
	// read
	// [Identifiers](/zero-touch/guides/identifiers).
	DeviceIdentifier *DeviceIdentifier `json:"deviceIdentifier,omitempty"`

	// DeviceMetadata: The metadata attached to the device. Structured as
	// key-value pairs. To
	// learn more, read [Device metadata](/zero-touch/guides/metadata).
	DeviceMetadata *DeviceMetadata `json:"deviceMetadata,omitempty"`

	// Name: Output only. The API resource name in the
	// format
	// `partners/[PARTNER_ID]/devices/[DEVICE_ID]`. Assigned by the server.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Claims") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Claims") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Device) MarshalJSON() ([]byte, error) {
	type NoMethod Device
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DeviceClaim: A record of a device claimed by a reseller for a
// customer. Devices claimed
// for zero-touch enrollment have a claim with the
// type
// `SECTION_TYPE_ZERO_TOUCH`. To learn more, read
// [Claim devices for customers](/zero-touch/guides/how-it-works#claim).
type DeviceClaim struct {
	// OwnerCompanyId: The ID of the Customer that purchased the device.
	OwnerCompanyId int64 `json:"ownerCompanyId,omitempty,string"`

	// ResellerId: The ID of the reseller that claimed the device.
	ResellerId int64 `json:"resellerId,omitempty,string"`

	// SectionType: Output only. The type of claim made on the device.
	//
	// Possible values:
	//   "SECTION_TYPE_UNSPECIFIED" - Unspecified section type.
	//   "SECTION_TYPE_SIM_LOCK" - SIM-lock section type.
	//   "SECTION_TYPE_ZERO_TOUCH" - Zero-touch enrollment section type.
	SectionType string `json:"sectionType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OwnerCompanyId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OwnerCompanyId") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DeviceClaim) MarshalJSON() ([]byte, error) {
	type NoMethod DeviceClaim
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DeviceIdentifier: Encapsulates hardware and product IDs to identify a
// manufactured device.
// To understand requirements on identifier sets,
// read
// [Identifiers](/zero-touch/guides/identifiers).
type DeviceIdentifier struct {
	// Imei: The device’s IMEI number. Validated on input.
	Imei string `json:"imei,omitempty"`

	// Manufacturer: The device manufacturer’s name. Matches the device's
	// built-in
	// value returned from `android.os.Build.MANUFACTURER`. Allowed values
	// are
	// listed
	// in
	// [manufacturers](/zero-touch/resources/manufacturer-names#manufactur
	// ers-names).
	Manufacturer string `json:"manufacturer,omitempty"`

	// Meid: The device’s MEID number.
	Meid string `json:"meid,omitempty"`

	// Model: The device model's name. Matches the device's built-in value
	// returned from
	// `android.os.Build.MODEL`. Allowed values are listed
	// in
	// [models](/zero-touch/resources/manufacturer-names#model-names).
	Model string `json:"model,omitempty"`

	// SerialNumber: The manufacturer's serial number for the device. This
	// value might not be
	// unique across different device models.
	SerialNumber string `json:"serialNumber,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Imei") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Imei") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DeviceIdentifier) MarshalJSON() ([]byte, error) {
	type NoMethod DeviceIdentifier
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DeviceMetadata: Metadata entries that can be attached to a `Device`.
// To learn more, read
// [Device metadata](/zero-touch/guides/metadata).
type DeviceMetadata struct {
	// Entries: Metadata entries recorded as key-value pairs.
	Entries map[string]string `json:"entries,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Entries") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Entries") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DeviceMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod DeviceMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DeviceReference: A `DeviceReference` is an API abstraction that lets
// you supply a _device_
// argument to a method using one of the following identifier types:
//
// * A numeric API resource ID.
// * Real-world hardware IDs, such as IMEI number, belonging to the
// manufactured
//   device.
//
// Methods that operate on devices take a `DeviceReference` as a
// parameter type
// because it's more flexible for the caller. To learn more about
// device
// identifiers, read [Identifiers](/zero-touch/guides/identifiers).
type DeviceReference struct {
	// DeviceId: The ID of the device.
	DeviceId int64 `json:"deviceId,omitempty,string"`

	// DeviceIdentifier: The hardware IDs of the device.
	DeviceIdentifier *DeviceIdentifier `json:"deviceIdentifier,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DeviceId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeviceId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DeviceReference) MarshalJSON() ([]byte, error) {
	type NoMethod DeviceReference
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DevicesLongRunningOperationMetadata: Tracks the status of a
// long-running operation to asynchronously update a
// batch of reseller metadata attached to devices. To learn more,
// read
// [Long‑running batch
// operations](/zero-touch/guides/how-it-works#operations).
type DevicesLongRunningOperationMetadata struct {
	// DevicesCount: The number of metadata updates in the operation. This
	// might be different
	// from the number of updates in the request if the API can't parse some
	// of
	// the updates.
	DevicesCount int64 `json:"devicesCount,omitempty"`

	// ProcessingStatus: The processing status of the operation.
	//
	// Possible values:
	//   "BATCH_PROCESS_STATUS_UNSPECIFIED" - Invalid code. Shouldn't be
	// used.
	//   "BATCH_PROCESS_PENDING" - Pending.
	//   "BATCH_PROCESS_IN_PROGRESS" - In progress.
	//   "BATCH_PROCESS_PROCESSED" - Processed.
	// This doesn't mean all items were processed successfully, you
	// should
	// check the `response` field for the result of every item.
	ProcessingStatus string `json:"processingStatus,omitempty"`

	// Progress: The processing progress of the operation. Measured as a
	// number from 0 to
	// 100. A value of 10O doesnt always mean the operation
	// completed—check for
	// the inclusion of a `done` field.
	Progress int64 `json:"progress,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DevicesCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DevicesCount") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DevicesLongRunningOperationMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod DevicesLongRunningOperationMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DevicesLongRunningOperationResponse: Tracks the status of a
// long-running operation to claim, unclaim, or attach
// metadata to devices. To learn more, read
// [Long‑running batch
// operations](/zero-touch/guides/how-it-works#operations).
type DevicesLongRunningOperationResponse struct {
	// PerDeviceStatus: The processing status for each device in the
	// operation.
	// One `PerDeviceStatus` per device. The list order matches the items in
	// the
	// original request.
	PerDeviceStatus []*OperationPerDevice `json:"perDeviceStatus,omitempty"`

	// SuccessCount: A summary of how many items in the operation the server
	// processed
	// successfully. Updated as the operation progresses.
	SuccessCount int64 `json:"successCount,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PerDeviceStatus") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PerDeviceStatus") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DevicesLongRunningOperationResponse) MarshalJSON() ([]byte, error) {
	type NoMethod DevicesLongRunningOperationResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Dpc: An EMM's DPC ([device
// policy
// controller](http://developer.android.com/work/dpc/build-dpc.htm
// l)).
// Zero-touch enrollment installs a DPC (listed in the `Configuration`)
// on a
// device to maintain the customer's mobile policies. All the DPCs
// listed by the
// API support zero-touch enrollment and are available in Google Play.
type Dpc struct {
	// DpcName: Output only. The title of the DPC app in Google Play. For
	// example, _Google
	// Apps Device Policy_. Useful in an application's user interface.
	DpcName string `json:"dpcName,omitempty"`

	// Name: Output only. The API resource name in the
	// format
	// `customers/[CUSTOMER_ID]/dpcs/[DPC_ID]`. Assigned by
	// the server. To maintain a reference to a DPC across customer
	// accounts,
	// persist and match the last path component (`DPC_ID`).
	Name string `json:"name,omitempty"`

	// PackageName: Output only. The DPC's Android application ID that looks
	// like a Java
	// package name. Zero-touch enrollment installs the DPC app onto a
	// device
	// using this identifier.
	PackageName string `json:"packageName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DpcName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DpcName") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Dpc) MarshalJSON() ([]byte, error) {
	type NoMethod Dpc
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

// FindDevicesByDeviceIdentifierRequest: Request to find devices.
type FindDevicesByDeviceIdentifierRequest struct {
	// DeviceIdentifier: Required. The device identifier to search for.
	DeviceIdentifier *DeviceIdentifier `json:"deviceIdentifier,omitempty"`

	// Limit: Required. The maximum number of devices to show in a page of
	// results. Must
	// be between 1 and 100 inclusive.
	Limit int64 `json:"limit,omitempty,string"`

	// PageToken: A token specifying which result page to return.
	PageToken string `json:"pageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DeviceIdentifier") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeviceIdentifier") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *FindDevicesByDeviceIdentifierRequest) MarshalJSON() ([]byte, error) {
	type NoMethod FindDevicesByDeviceIdentifierRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FindDevicesByDeviceIdentifierResponse: Response containing found
// devices.
type FindDevicesByDeviceIdentifierResponse struct {
	// Devices: Found devices.
	Devices []*Device `json:"devices,omitempty"`

	// NextPageToken: A token used to access the next page of results.
	// Omitted if no further
	// results are available.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalSize: The total count of items in the list irrespective of
	// pagination.
	TotalSize int64 `json:"totalSize,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Devices") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Devices") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FindDevicesByDeviceIdentifierResponse) MarshalJSON() ([]byte, error) {
	type NoMethod FindDevicesByDeviceIdentifierResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FindDevicesByOwnerRequest: Request to find devices by customers.
type FindDevicesByOwnerRequest struct {
	// CustomerId: Required. The list of customer IDs to search for.
	CustomerId googleapi.Int64s `json:"customerId,omitempty"`

	// Limit: Required. The maximum number of devices to show in a page of
	// results. Must
	// be between 1 and 100 inclusive.
	Limit int64 `json:"limit,omitempty,string"`

	// PageToken: A token specifying which result page to return.
	PageToken string `json:"pageToken,omitempty"`

	// SectionType: Required. The section type of the device's provisioning
	// record.
	//
	// Possible values:
	//   "SECTION_TYPE_UNSPECIFIED" - Unspecified section type.
	//   "SECTION_TYPE_SIM_LOCK" - SIM-lock section type.
	//   "SECTION_TYPE_ZERO_TOUCH" - Zero-touch enrollment section type.
	SectionType string `json:"sectionType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CustomerId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CustomerId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FindDevicesByOwnerRequest) MarshalJSON() ([]byte, error) {
	type NoMethod FindDevicesByOwnerRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FindDevicesByOwnerResponse: Response containing found devices.
type FindDevicesByOwnerResponse struct {
	// Devices: The customer's devices.
	Devices []*Device `json:"devices,omitempty"`

	// NextPageToken: A token used to access the next page of
	// results.
	// Omitted if no further results are available.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalSize: The total count of items in the list irrespective of
	// pagination.
	TotalSize int64 `json:"totalSize,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Devices") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Devices") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FindDevicesByOwnerResponse) MarshalJSON() ([]byte, error) {
	type NoMethod FindDevicesByOwnerResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListCustomersResponse: Response message of all customers related to
// this partner.
type ListCustomersResponse struct {
	// Customers: List of customers related to this reseller partner.
	Customers []*Company `json:"customers,omitempty"`

	// NextPageToken: A token to retrieve the next page of results. Omitted
	// if no further results
	// are available.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalSize: The total count of items in the list irrespective of
	// pagination.
	TotalSize int64 `json:"totalSize,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Customers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Customers") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListCustomersResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListCustomersResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListVendorCustomersResponse: Response message to list customers of
// the vendor.
type ListVendorCustomersResponse struct {
	// Customers: List of customers of the vendor.
	Customers []*Company `json:"customers,omitempty"`

	// NextPageToken: A token to retrieve the next page of results. Omitted
	// if no further results
	// are available.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalSize: The total count of items in the list irrespective of
	// pagination.
	TotalSize int64 `json:"totalSize,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Customers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Customers") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListVendorCustomersResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListVendorCustomersResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListVendorsResponse: Response message to list vendors of the partner.
type ListVendorsResponse struct {
	// NextPageToken: A token to retrieve the next page of results. Omitted
	// if no further results
	// are available.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalSize: The total count of items in the list irrespective of
	// pagination.
	TotalSize int64 `json:"totalSize,omitempty"`

	// Vendors: List of vendors of the reseller partner. Fields `name`,
	// `companyId` and
	// `companyName` are populated to the Company object.
	Vendors []*Company `json:"vendors,omitempty"`

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

func (s *ListVendorsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListVendorsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Operation: This resource represents a long-running operation that is
// the result of a
// network API call.
type Operation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress.
	// If `true`, the operation is completed, and either `error` or
	// `response` is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: This field will always be not set if the operation is created
	// by `claimAsync`, `unclaimAsync`, or `updateMetadataAsync`. In this
	// case, error information for each device is set in
	// `response.perDeviceStatus.result.status`.
	Error *Status `json:"error,omitempty"`

	// Metadata: This field will contain a
	// `DevicesLongRunningOperationMetadata` object if the operation is
	// created by `claimAsync`, `unclaimAsync`, or `updateMetadataAsync`.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that
	// originally returns it. If you use the default HTTP mapping,
	// the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `json:"name,omitempty"`

	// Response: This field will contain a
	// `DevicesLongRunningOperationResponse` object if the operation is
	// created by `claimAsync`, `unclaimAsync`, or `updateMetadataAsync`.
	Response googleapi.RawMessage `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Operation) MarshalJSON() ([]byte, error) {
	type NoMethod Operation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OperationPerDevice: A task for each device in the operation.
// Corresponds to each device
// change in the request.
type OperationPerDevice struct {
	// Claim: A copy of the original device-claim request received by the
	// server.
	Claim *PartnerClaim `json:"claim,omitempty"`

	// Result: The processing result for each device.
	Result *PerDeviceStatusInBatch `json:"result,omitempty"`

	// Unclaim: A copy of the original device-unclaim request received by
	// the server.
	Unclaim *PartnerUnclaim `json:"unclaim,omitempty"`

	// UpdateMetadata: A copy of the original metadata-update request
	// received by the server.
	UpdateMetadata *UpdateMetadataArguments `json:"updateMetadata,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Claim") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Claim") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OperationPerDevice) MarshalJSON() ([]byte, error) {
	type NoMethod OperationPerDevice
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PartnerClaim: Identifies one claim request.
type PartnerClaim struct {
	// CustomerId: Required. The ID of the customer for whom the device is
	// being claimed.
	CustomerId int64 `json:"customerId,omitempty,string"`

	// DeviceIdentifier: Required. Device identifier of the device.
	DeviceIdentifier *DeviceIdentifier `json:"deviceIdentifier,omitempty"`

	// DeviceMetadata: Required. The metadata to attach to the device at
	// claim.
	DeviceMetadata *DeviceMetadata `json:"deviceMetadata,omitempty"`

	// SectionType: Required. The section type of the device's provisioning
	// record.
	//
	// Possible values:
	//   "SECTION_TYPE_UNSPECIFIED" - Unspecified section type.
	//   "SECTION_TYPE_SIM_LOCK" - SIM-lock section type.
	//   "SECTION_TYPE_ZERO_TOUCH" - Zero-touch enrollment section type.
	SectionType string `json:"sectionType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CustomerId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CustomerId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PartnerClaim) MarshalJSON() ([]byte, error) {
	type NoMethod PartnerClaim
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PartnerUnclaim: Identifies one unclaim request.
type PartnerUnclaim struct {
	// DeviceId: Device ID of the device.
	DeviceId int64 `json:"deviceId,omitempty,string"`

	// DeviceIdentifier: Device identifier of the device.
	DeviceIdentifier *DeviceIdentifier `json:"deviceIdentifier,omitempty"`

	// SectionType: Required. The section type of the device's provisioning
	// record.
	//
	// Possible values:
	//   "SECTION_TYPE_UNSPECIFIED" - Unspecified section type.
	//   "SECTION_TYPE_SIM_LOCK" - SIM-lock section type.
	//   "SECTION_TYPE_ZERO_TOUCH" - Zero-touch enrollment section type.
	SectionType string `json:"sectionType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DeviceId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeviceId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PartnerUnclaim) MarshalJSON() ([]byte, error) {
	type NoMethod PartnerUnclaim
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PerDeviceStatusInBatch: Captures the processing status for each
// device in the operation.
type PerDeviceStatusInBatch struct {
	// DeviceId: If processing succeeds, the device ID of the device.
	DeviceId int64 `json:"deviceId,omitempty,string"`

	// ErrorIdentifier: If processing fails, the error type.
	ErrorIdentifier string `json:"errorIdentifier,omitempty"`

	// ErrorMessage: If processing fails, a developer message explaining
	// what went wrong.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Status: The result status of the device after processing.
	//
	// Possible values:
	//   "SINGLE_DEVICE_STATUS_UNSPECIFIED" - Invalid code. Shouldn't be
	// used.
	//   "SINGLE_DEVICE_STATUS_UNKNOWN_ERROR" - Unknown error.
	// We don't expect this error to occur here.
	//   "SINGLE_DEVICE_STATUS_OTHER_ERROR" - Other error.
	// We know/expect this error, but there's no defined error code for
	// the
	// error.
	//   "SINGLE_DEVICE_STATUS_SUCCESS" - Success.
	//   "SINGLE_DEVICE_STATUS_PERMISSION_DENIED" - Permission denied.
	//   "SINGLE_DEVICE_STATUS_INVALID_DEVICE_IDENTIFIER" - Invalid device
	// identifier.
	//   "SINGLE_DEVICE_STATUS_INVALID_SECTION_TYPE" - Invalid section type.
	//   "SINGLE_DEVICE_STATUS_SECTION_NOT_YOURS" - This section is claimed
	// by another company.
	Status string `json:"status,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DeviceId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeviceId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PerDeviceStatusInBatch) MarshalJSON() ([]byte, error) {
	type NoMethod PerDeviceStatusInBatch
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There is a
	// common set of
	// message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
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

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UnclaimDeviceRequest: Request message to unclaim a device.
type UnclaimDeviceRequest struct {
	// DeviceId: The device ID returned by `ClaimDevice`.
	DeviceId int64 `json:"deviceId,omitempty,string"`

	// DeviceIdentifier: The device identifier you used when you claimed
	// this device.
	DeviceIdentifier *DeviceIdentifier `json:"deviceIdentifier,omitempty"`

	// SectionType: Required. The section type of the device's provisioning
	// record.
	//
	// Possible values:
	//   "SECTION_TYPE_UNSPECIFIED" - Unspecified section type.
	//   "SECTION_TYPE_SIM_LOCK" - SIM-lock section type.
	//   "SECTION_TYPE_ZERO_TOUCH" - Zero-touch enrollment section type.
	SectionType string `json:"sectionType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DeviceId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeviceId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UnclaimDeviceRequest) MarshalJSON() ([]byte, error) {
	type NoMethod UnclaimDeviceRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UnclaimDevicesRequest: Request to unclaim devices asynchronously in
// batch.
type UnclaimDevicesRequest struct {
	// Unclaims: Required. The list of devices to unclaim.
	Unclaims []*PartnerUnclaim `json:"unclaims,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Unclaims") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Unclaims") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UnclaimDevicesRequest) MarshalJSON() ([]byte, error) {
	type NoMethod UnclaimDevicesRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UpdateDeviceMetadataInBatchRequest: Request to update device metadata
// in batch.
type UpdateDeviceMetadataInBatchRequest struct {
	// Updates: Required. The list of metadata updates.
	Updates []*UpdateMetadataArguments `json:"updates,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Updates") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Updates") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UpdateDeviceMetadataInBatchRequest) MarshalJSON() ([]byte, error) {
	type NoMethod UpdateDeviceMetadataInBatchRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UpdateDeviceMetadataRequest: Request to set metadata for a device.
type UpdateDeviceMetadataRequest struct {
	// DeviceMetadata: Required. The metdata to attach to the device.
	DeviceMetadata *DeviceMetadata `json:"deviceMetadata,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DeviceMetadata") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeviceMetadata") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *UpdateDeviceMetadataRequest) MarshalJSON() ([]byte, error) {
	type NoMethod UpdateDeviceMetadataRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UpdateMetadataArguments: Identifies metdata updates to one device.
type UpdateMetadataArguments struct {
	// DeviceId: Device ID of the device.
	DeviceId int64 `json:"deviceId,omitempty,string"`

	// DeviceIdentifier: Device identifier.
	DeviceIdentifier *DeviceIdentifier `json:"deviceIdentifier,omitempty"`

	// DeviceMetadata: Required. The metadata to update.
	DeviceMetadata *DeviceMetadata `json:"deviceMetadata,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DeviceId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeviceId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UpdateMetadataArguments) MarshalJSON() ([]byte, error) {
	type NoMethod UpdateMetadataArguments
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "androiddeviceprovisioning.customers.list":

type CustomersListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the user's customer accounts.
func (r *CustomersService) List() *CustomersListCall {
	c := &CustomersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of customers to show in a page of results.
// A number between 1 and 100 (inclusive).
func (c *CustomersListCall) PageSize(pageSize int64) *CustomersListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token specifying
// which result page to return.
func (c *CustomersListCall) PageToken(pageToken string) *CustomersListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersListCall) Fields(s ...googleapi.Field) *CustomersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CustomersListCall) IfNoneMatch(entityTag string) *CustomersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomersListCall) Context(ctx context.Context) *CustomersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CustomersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CustomersListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/customers")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.customers.list"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.customers.list" call.
// Exactly one of *CustomerListCustomersResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *CustomerListCustomersResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CustomersListCall) Do(opts ...googleapi.CallOption) (*CustomerListCustomersResponse, error) {
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
	ret := &CustomerListCustomersResponse{
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
	//   "description": "Lists the user's customer accounts.",
	//   "flatPath": "v1/customers",
	//   "httpMethod": "GET",
	//   "id": "androiddeviceprovisioning.customers.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of customers to show in a page of results.\nA number between 1 and 100 (inclusive).",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token specifying which result page to return.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/customers",
	//   "response": {
	//     "$ref": "CustomerListCustomersResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *CustomersListCall) Pages(ctx context.Context, f func(*CustomerListCustomersResponse) error) error {
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

// method id "androiddeviceprovisioning.customers.configurations.create":

type CustomersConfigurationsCreateCall struct {
	s             *Service
	parent        string
	configuration *Configuration
	urlParams_    gensupport.URLParams
	ctx_          context.Context
	header_       http.Header
}

// Create: Creates a new configuration. Once created, a customer can
// apply the
// configuration to devices.
func (r *CustomersConfigurationsService) Create(parent string, configuration *Configuration) *CustomersConfigurationsCreateCall {
	c := &CustomersConfigurationsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.configuration = configuration
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersConfigurationsCreateCall) Fields(s ...googleapi.Field) *CustomersConfigurationsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomersConfigurationsCreateCall) Context(ctx context.Context) *CustomersConfigurationsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CustomersConfigurationsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CustomersConfigurationsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.configuration)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/configurations")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.customers.configurations.create"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.customers.configurations.create" call.
// Exactly one of *Configuration or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Configuration.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CustomersConfigurationsCreateCall) Do(opts ...googleapi.CallOption) (*Configuration, error) {
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
	ret := &Configuration{
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
	//   "description": "Creates a new configuration. Once created, a customer can apply the\nconfiguration to devices.",
	//   "flatPath": "v1/customers/{customersId}/configurations",
	//   "httpMethod": "POST",
	//   "id": "androiddeviceprovisioning.customers.configurations.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The customer that manages the configuration. An API resource name\nin the format `customers/[CUSTOMER_ID]`.",
	//       "location": "path",
	//       "pattern": "^customers/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/configurations",
	//   "request": {
	//     "$ref": "Configuration"
	//   },
	//   "response": {
	//     "$ref": "Configuration"
	//   }
	// }

}

// method id "androiddeviceprovisioning.customers.configurations.delete":

type CustomersConfigurationsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes an unused configuration. The API call fails if the
// customer has
// devices with the configuration applied.
func (r *CustomersConfigurationsService) Delete(name string) *CustomersConfigurationsDeleteCall {
	c := &CustomersConfigurationsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersConfigurationsDeleteCall) Fields(s ...googleapi.Field) *CustomersConfigurationsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomersConfigurationsDeleteCall) Context(ctx context.Context) *CustomersConfigurationsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CustomersConfigurationsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CustomersConfigurationsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.customers.configurations.delete"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.customers.configurations.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *CustomersConfigurationsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	//   "description": "Deletes an unused configuration. The API call fails if the customer has\ndevices with the configuration applied.",
	//   "flatPath": "v1/customers/{customersId}/configurations/{configurationsId}",
	//   "httpMethod": "DELETE",
	//   "id": "androiddeviceprovisioning.customers.configurations.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The configuration to delete. An API resource name in the format\n`customers/[CUSTOMER_ID]/configurations/[CONFIGURATION_ID]`. If the\nconfiguration is applied to any devices, the API call fails.",
	//       "location": "path",
	//       "pattern": "^customers/[^/]+/configurations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   }
	// }

}

// method id "androiddeviceprovisioning.customers.configurations.get":

type CustomersConfigurationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the details of a configuration.
func (r *CustomersConfigurationsService) Get(name string) *CustomersConfigurationsGetCall {
	c := &CustomersConfigurationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersConfigurationsGetCall) Fields(s ...googleapi.Field) *CustomersConfigurationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CustomersConfigurationsGetCall) IfNoneMatch(entityTag string) *CustomersConfigurationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomersConfigurationsGetCall) Context(ctx context.Context) *CustomersConfigurationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CustomersConfigurationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CustomersConfigurationsGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.customers.configurations.get"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.customers.configurations.get" call.
// Exactly one of *Configuration or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Configuration.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CustomersConfigurationsGetCall) Do(opts ...googleapi.CallOption) (*Configuration, error) {
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
	ret := &Configuration{
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
	//   "description": "Gets the details of a configuration.",
	//   "flatPath": "v1/customers/{customersId}/configurations/{configurationsId}",
	//   "httpMethod": "GET",
	//   "id": "androiddeviceprovisioning.customers.configurations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The configuration to get. An API resource name in the format\n`customers/[CUSTOMER_ID]/configurations/[CONFIGURATION_ID]`.",
	//       "location": "path",
	//       "pattern": "^customers/[^/]+/configurations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Configuration"
	//   }
	// }

}

// method id "androiddeviceprovisioning.customers.configurations.list":

type CustomersConfigurationsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists a customer's configurations.
func (r *CustomersConfigurationsService) List(parent string) *CustomersConfigurationsListCall {
	c := &CustomersConfigurationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersConfigurationsListCall) Fields(s ...googleapi.Field) *CustomersConfigurationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CustomersConfigurationsListCall) IfNoneMatch(entityTag string) *CustomersConfigurationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomersConfigurationsListCall) Context(ctx context.Context) *CustomersConfigurationsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CustomersConfigurationsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CustomersConfigurationsListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/configurations")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.customers.configurations.list"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.customers.configurations.list" call.
// Exactly one of *CustomerListConfigurationsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *CustomerListConfigurationsResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *CustomersConfigurationsListCall) Do(opts ...googleapi.CallOption) (*CustomerListConfigurationsResponse, error) {
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
	ret := &CustomerListConfigurationsResponse{
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
	//   "description": "Lists a customer's configurations.",
	//   "flatPath": "v1/customers/{customersId}/configurations",
	//   "httpMethod": "GET",
	//   "id": "androiddeviceprovisioning.customers.configurations.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The customer that manages the listed configurations. An API\nresource name in the format `customers/[CUSTOMER_ID]`.",
	//       "location": "path",
	//       "pattern": "^customers/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/configurations",
	//   "response": {
	//     "$ref": "CustomerListConfigurationsResponse"
	//   }
	// }

}

// method id "androiddeviceprovisioning.customers.configurations.patch":

type CustomersConfigurationsPatchCall struct {
	s             *Service
	name          string
	configuration *Configuration
	urlParams_    gensupport.URLParams
	ctx_          context.Context
	header_       http.Header
}

// Patch: Updates a configuration's field values.
func (r *CustomersConfigurationsService) Patch(name string, configuration *Configuration) *CustomersConfigurationsPatchCall {
	c := &CustomersConfigurationsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.configuration = configuration
	return c
}

// UpdateMask sets the optional parameter "updateMask": Required. The
// field mask applied to the target `Configuration` before
// updating the fields. To learn more about using field masks,
// read
// [FieldMask](/protocol-buffers/docs/reference/google.protobuf#fiel
// dmask) in
// the Protocol Buffers documentation.
func (c *CustomersConfigurationsPatchCall) UpdateMask(updateMask string) *CustomersConfigurationsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersConfigurationsPatchCall) Fields(s ...googleapi.Field) *CustomersConfigurationsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomersConfigurationsPatchCall) Context(ctx context.Context) *CustomersConfigurationsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CustomersConfigurationsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CustomersConfigurationsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.configuration)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.customers.configurations.patch"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.customers.configurations.patch" call.
// Exactly one of *Configuration or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Configuration.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CustomersConfigurationsPatchCall) Do(opts ...googleapi.CallOption) (*Configuration, error) {
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
	ret := &Configuration{
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
	//   "description": "Updates a configuration's field values.",
	//   "flatPath": "v1/customers/{customersId}/configurations/{configurationsId}",
	//   "httpMethod": "PATCH",
	//   "id": "androiddeviceprovisioning.customers.configurations.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Output only. The API resource name in the format\n`customers/[CUSTOMER_ID]/configurations/[CONFIGURATION_ID]`. Assigned by\nthe server.",
	//       "location": "path",
	//       "pattern": "^customers/[^/]+/configurations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Required. The field mask applied to the target `Configuration` before\nupdating the fields. To learn more about using field masks, read\n[FieldMask](/protocol-buffers/docs/reference/google.protobuf#fieldmask) in\nthe Protocol Buffers documentation.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "request": {
	//     "$ref": "Configuration"
	//   },
	//   "response": {
	//     "$ref": "Configuration"
	//   }
	// }

}

// method id "androiddeviceprovisioning.customers.devices.applyConfiguration":

type CustomersDevicesApplyConfigurationCall struct {
	s                                 *Service
	parent                            string
	customerapplyconfigurationrequest *CustomerApplyConfigurationRequest
	urlParams_                        gensupport.URLParams
	ctx_                              context.Context
	header_                           http.Header
}

// ApplyConfiguration: Applies a Configuration to the device to register
// the device for zero-touch
// enrollment. After applying a configuration to a device, the
// device
// automatically provisions itself on first boot, or next factory reset.
func (r *CustomersDevicesService) ApplyConfiguration(parent string, customerapplyconfigurationrequest *CustomerApplyConfigurationRequest) *CustomersDevicesApplyConfigurationCall {
	c := &CustomersDevicesApplyConfigurationCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.customerapplyconfigurationrequest = customerapplyconfigurationrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersDevicesApplyConfigurationCall) Fields(s ...googleapi.Field) *CustomersDevicesApplyConfigurationCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomersDevicesApplyConfigurationCall) Context(ctx context.Context) *CustomersDevicesApplyConfigurationCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CustomersDevicesApplyConfigurationCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CustomersDevicesApplyConfigurationCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.customerapplyconfigurationrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/devices:applyConfiguration")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.customers.devices.applyConfiguration"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.customers.devices.applyConfiguration" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *CustomersDevicesApplyConfigurationCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	//   "description": "Applies a Configuration to the device to register the device for zero-touch\nenrollment. After applying a configuration to a device, the device\nautomatically provisions itself on first boot, or next factory reset.",
	//   "flatPath": "v1/customers/{customersId}/devices:applyConfiguration",
	//   "httpMethod": "POST",
	//   "id": "androiddeviceprovisioning.customers.devices.applyConfiguration",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The customer managing the device. An API resource name in the\nformat `customers/[CUSTOMER_ID]`.",
	//       "location": "path",
	//       "pattern": "^customers/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/devices:applyConfiguration",
	//   "request": {
	//     "$ref": "CustomerApplyConfigurationRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   }
	// }

}

// method id "androiddeviceprovisioning.customers.devices.get":

type CustomersDevicesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the details of a device.
func (r *CustomersDevicesService) Get(name string) *CustomersDevicesGetCall {
	c := &CustomersDevicesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersDevicesGetCall) Fields(s ...googleapi.Field) *CustomersDevicesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CustomersDevicesGetCall) IfNoneMatch(entityTag string) *CustomersDevicesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomersDevicesGetCall) Context(ctx context.Context) *CustomersDevicesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CustomersDevicesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CustomersDevicesGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.customers.devices.get"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.customers.devices.get" call.
// Exactly one of *Device or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Device.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *CustomersDevicesGetCall) Do(opts ...googleapi.CallOption) (*Device, error) {
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
	ret := &Device{
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
	//   "description": "Gets the details of a device.",
	//   "flatPath": "v1/customers/{customersId}/devices/{devicesId}",
	//   "httpMethod": "GET",
	//   "id": "androiddeviceprovisioning.customers.devices.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The device to get. An API resource name in the format\n`customers/[CUSTOMER_ID]/devices/[DEVICE_ID]`.",
	//       "location": "path",
	//       "pattern": "^customers/[^/]+/devices/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Device"
	//   }
	// }

}

// method id "androiddeviceprovisioning.customers.devices.list":

type CustomersDevicesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists a customer's devices.
func (r *CustomersDevicesService) List(parent string) *CustomersDevicesListCall {
	c := &CustomersDevicesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of devices to show in a page of results.
// Must be between 1 and 100 inclusive.
func (c *CustomersDevicesListCall) PageSize(pageSize int64) *CustomersDevicesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token specifying
// which result page to return.
func (c *CustomersDevicesListCall) PageToken(pageToken string) *CustomersDevicesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersDevicesListCall) Fields(s ...googleapi.Field) *CustomersDevicesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CustomersDevicesListCall) IfNoneMatch(entityTag string) *CustomersDevicesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomersDevicesListCall) Context(ctx context.Context) *CustomersDevicesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CustomersDevicesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CustomersDevicesListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/devices")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.customers.devices.list"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.customers.devices.list" call.
// Exactly one of *CustomerListDevicesResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *CustomerListDevicesResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CustomersDevicesListCall) Do(opts ...googleapi.CallOption) (*CustomerListDevicesResponse, error) {
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
	ret := &CustomerListDevicesResponse{
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
	//   "description": "Lists a customer's devices.",
	//   "flatPath": "v1/customers/{customersId}/devices",
	//   "httpMethod": "GET",
	//   "id": "androiddeviceprovisioning.customers.devices.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of devices to show in a page of results.\nMust be between 1 and 100 inclusive.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "A token specifying which result page to return.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The customer managing the devices. An API resource name in the\nformat `customers/[CUSTOMER_ID]`.",
	//       "location": "path",
	//       "pattern": "^customers/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/devices",
	//   "response": {
	//     "$ref": "CustomerListDevicesResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *CustomersDevicesListCall) Pages(ctx context.Context, f func(*CustomerListDevicesResponse) error) error {
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

// method id "androiddeviceprovisioning.customers.devices.removeConfiguration":

type CustomersDevicesRemoveConfigurationCall struct {
	s                                  *Service
	parent                             string
	customerremoveconfigurationrequest *CustomerRemoveConfigurationRequest
	urlParams_                         gensupport.URLParams
	ctx_                               context.Context
	header_                            http.Header
}

// RemoveConfiguration: Removes a configuration from device.
func (r *CustomersDevicesService) RemoveConfiguration(parent string, customerremoveconfigurationrequest *CustomerRemoveConfigurationRequest) *CustomersDevicesRemoveConfigurationCall {
	c := &CustomersDevicesRemoveConfigurationCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.customerremoveconfigurationrequest = customerremoveconfigurationrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersDevicesRemoveConfigurationCall) Fields(s ...googleapi.Field) *CustomersDevicesRemoveConfigurationCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomersDevicesRemoveConfigurationCall) Context(ctx context.Context) *CustomersDevicesRemoveConfigurationCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CustomersDevicesRemoveConfigurationCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CustomersDevicesRemoveConfigurationCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.customerremoveconfigurationrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/devices:removeConfiguration")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.customers.devices.removeConfiguration"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.customers.devices.removeConfiguration" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *CustomersDevicesRemoveConfigurationCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	//   "description": "Removes a configuration from device.",
	//   "flatPath": "v1/customers/{customersId}/devices:removeConfiguration",
	//   "httpMethod": "POST",
	//   "id": "androiddeviceprovisioning.customers.devices.removeConfiguration",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The customer managing the device in the format\n`customers/[CUSTOMER_ID]`.",
	//       "location": "path",
	//       "pattern": "^customers/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/devices:removeConfiguration",
	//   "request": {
	//     "$ref": "CustomerRemoveConfigurationRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   }
	// }

}

// method id "androiddeviceprovisioning.customers.devices.unclaim":

type CustomersDevicesUnclaimCall struct {
	s                            *Service
	parent                       string
	customerunclaimdevicerequest *CustomerUnclaimDeviceRequest
	urlParams_                   gensupport.URLParams
	ctx_                         context.Context
	header_                      http.Header
}

// Unclaim: Unclaims a device from a customer and removes it from
// zero-touch
// enrollment.
//
// After removing a device, a customer must contact their reseller to
// register
// the device into zero-touch enrollment again.
func (r *CustomersDevicesService) Unclaim(parent string, customerunclaimdevicerequest *CustomerUnclaimDeviceRequest) *CustomersDevicesUnclaimCall {
	c := &CustomersDevicesUnclaimCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.customerunclaimdevicerequest = customerunclaimdevicerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersDevicesUnclaimCall) Fields(s ...googleapi.Field) *CustomersDevicesUnclaimCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomersDevicesUnclaimCall) Context(ctx context.Context) *CustomersDevicesUnclaimCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CustomersDevicesUnclaimCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CustomersDevicesUnclaimCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.customerunclaimdevicerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/devices:unclaim")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.customers.devices.unclaim"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.customers.devices.unclaim" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *CustomersDevicesUnclaimCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	//   "description": "Unclaims a device from a customer and removes it from zero-touch\nenrollment.\n\nAfter removing a device, a customer must contact their reseller to register\nthe device into zero-touch enrollment again.",
	//   "flatPath": "v1/customers/{customersId}/devices:unclaim",
	//   "httpMethod": "POST",
	//   "id": "androiddeviceprovisioning.customers.devices.unclaim",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The customer managing the device. An API resource name in the\nformat `customers/[CUSTOMER_ID]`.",
	//       "location": "path",
	//       "pattern": "^customers/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/devices:unclaim",
	//   "request": {
	//     "$ref": "CustomerUnclaimDeviceRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   }
	// }

}

// method id "androiddeviceprovisioning.customers.dpcs.list":

type CustomersDpcsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the DPCs (device policy controllers) that support
// zero-touch
// enrollment.
func (r *CustomersDpcsService) List(parent string) *CustomersDpcsListCall {
	c := &CustomersDpcsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersDpcsListCall) Fields(s ...googleapi.Field) *CustomersDpcsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CustomersDpcsListCall) IfNoneMatch(entityTag string) *CustomersDpcsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomersDpcsListCall) Context(ctx context.Context) *CustomersDpcsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CustomersDpcsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CustomersDpcsListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/dpcs")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.customers.dpcs.list"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.customers.dpcs.list" call.
// Exactly one of *CustomerListDpcsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *CustomerListDpcsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CustomersDpcsListCall) Do(opts ...googleapi.CallOption) (*CustomerListDpcsResponse, error) {
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
	ret := &CustomerListDpcsResponse{
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
	//   "description": "Lists the DPCs (device policy controllers) that support zero-touch\nenrollment.",
	//   "flatPath": "v1/customers/{customersId}/dpcs",
	//   "httpMethod": "GET",
	//   "id": "androiddeviceprovisioning.customers.dpcs.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The customer that can use the DPCs in configurations. An API\nresource name in the format `customers/[CUSTOMER_ID]`.",
	//       "location": "path",
	//       "pattern": "^customers/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/dpcs",
	//   "response": {
	//     "$ref": "CustomerListDpcsResponse"
	//   }
	// }

}

// method id "androiddeviceprovisioning.operations.get":

type OperationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the latest state of a long-running operation.  Clients can
// use this
// method to poll the operation result at intervals as recommended by
// the API
// service.
func (r *OperationsService) Get(name string) *OperationsGetCall {
	c := &OperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.operations.get"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.operations.get" call.
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
	//   "description": "Gets the latest state of a long-running operation.  Clients can use this\nmethod to poll the operation result at intervals as recommended by the API\nservice.",
	//   "flatPath": "v1/operations/{operationsId}",
	//   "httpMethod": "GET",
	//   "id": "androiddeviceprovisioning.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource.",
	//       "location": "path",
	//       "pattern": "^operations/.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   }
	// }

}

// method id "androiddeviceprovisioning.partners.customers.create":

type PartnersCustomersCreateCall struct {
	s                     *Service
	parent                string
	createcustomerrequest *CreateCustomerRequest
	urlParams_            gensupport.URLParams
	ctx_                  context.Context
	header_               http.Header
}

// Create: Creates a customer for zero-touch enrollment. After the
// method returns
// successfully, admin and owner roles can manage devices and EMM
// configs
// by calling API methods or using their zero-touch enrollment
// portal.
// The customer receives an email that welcomes them to zero-touch
// enrollment
// and explains how to sign into the portal.
func (r *PartnersCustomersService) Create(parent string, createcustomerrequest *CreateCustomerRequest) *PartnersCustomersCreateCall {
	c := &PartnersCustomersCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.createcustomerrequest = createcustomerrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PartnersCustomersCreateCall) Fields(s ...googleapi.Field) *PartnersCustomersCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PartnersCustomersCreateCall) Context(ctx context.Context) *PartnersCustomersCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PartnersCustomersCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PartnersCustomersCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createcustomerrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/customers")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.partners.customers.create"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.partners.customers.create" call.
// Exactly one of *Company or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Company.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *PartnersCustomersCreateCall) Do(opts ...googleapi.CallOption) (*Company, error) {
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
	ret := &Company{
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
	//   "description": "Creates a customer for zero-touch enrollment. After the method returns\nsuccessfully, admin and owner roles can manage devices and EMM configs\nby calling API methods or using their zero-touch enrollment portal.\nThe customer receives an email that welcomes them to zero-touch enrollment\nand explains how to sign into the portal.",
	//   "flatPath": "v1/partners/{partnersId}/customers",
	//   "httpMethod": "POST",
	//   "id": "androiddeviceprovisioning.partners.customers.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The parent resource ID in the format `partners/[PARTNER_ID]` that\nidentifies the reseller.",
	//       "location": "path",
	//       "pattern": "^partners/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/customers",
	//   "request": {
	//     "$ref": "CreateCustomerRequest"
	//   },
	//   "response": {
	//     "$ref": "Company"
	//   }
	// }

}

// method id "androiddeviceprovisioning.partners.customers.list":

type PartnersCustomersListCall struct {
	s            *Service
	partnerId    int64
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the customers that are enrolled to the reseller
// identified by the
// `partnerId` argument. This list includes customers that the
// reseller
// created and customers that enrolled themselves using the portal.
func (r *PartnersCustomersService) List(partnerId int64) *PartnersCustomersListCall {
	c := &PartnersCustomersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.partnerId = partnerId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to be returned. If not specified or 0, all
// the records are returned.
func (c *PartnersCustomersListCall) PageSize(pageSize int64) *PartnersCustomersListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results returned by the server.
func (c *PartnersCustomersListCall) PageToken(pageToken string) *PartnersCustomersListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PartnersCustomersListCall) Fields(s ...googleapi.Field) *PartnersCustomersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *PartnersCustomersListCall) IfNoneMatch(entityTag string) *PartnersCustomersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PartnersCustomersListCall) Context(ctx context.Context) *PartnersCustomersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PartnersCustomersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PartnersCustomersListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/partners/{+partnerId}/customers")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"partnerId": strconv.FormatInt(c.partnerId, 10),
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.partners.customers.list"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.partners.customers.list" call.
// Exactly one of *ListCustomersResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListCustomersResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PartnersCustomersListCall) Do(opts ...googleapi.CallOption) (*ListCustomersResponse, error) {
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
	ret := &ListCustomersResponse{
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
	//   "description": "Lists the customers that are enrolled to the reseller identified by the\n`partnerId` argument. This list includes customers that the reseller\ncreated and customers that enrolled themselves using the portal.",
	//   "flatPath": "v1/partners/{partnersId}/customers",
	//   "httpMethod": "GET",
	//   "id": "androiddeviceprovisioning.partners.customers.list",
	//   "parameterOrder": [
	//     "partnerId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of results to be returned. If not specified or 0, all\nthe records are returned.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results returned by the server.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "partnerId": {
	//       "description": "Required. The ID of the reseller partner.",
	//       "format": "int64",
	//       "location": "path",
	//       "pattern": "^[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/partners/{+partnerId}/customers",
	//   "response": {
	//     "$ref": "ListCustomersResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *PartnersCustomersListCall) Pages(ctx context.Context, f func(*ListCustomersResponse) error) error {
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

// method id "androiddeviceprovisioning.partners.devices.claim":

type PartnersDevicesClaimCall struct {
	s                  *Service
	partnerId          int64
	claimdevicerequest *ClaimDeviceRequest
	urlParams_         gensupport.URLParams
	ctx_               context.Context
	header_            http.Header
}

// Claim: Claims a device for a customer and adds it to zero-touch
// enrollment. If the
// device is already claimed by another customer, the call returns an
// error.
func (r *PartnersDevicesService) Claim(partnerId int64, claimdevicerequest *ClaimDeviceRequest) *PartnersDevicesClaimCall {
	c := &PartnersDevicesClaimCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.partnerId = partnerId
	c.claimdevicerequest = claimdevicerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PartnersDevicesClaimCall) Fields(s ...googleapi.Field) *PartnersDevicesClaimCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PartnersDevicesClaimCall) Context(ctx context.Context) *PartnersDevicesClaimCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PartnersDevicesClaimCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PartnersDevicesClaimCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.claimdevicerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/partners/{+partnerId}/devices:claim")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"partnerId": strconv.FormatInt(c.partnerId, 10),
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.partners.devices.claim"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.partners.devices.claim" call.
// Exactly one of *ClaimDeviceResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ClaimDeviceResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PartnersDevicesClaimCall) Do(opts ...googleapi.CallOption) (*ClaimDeviceResponse, error) {
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
	ret := &ClaimDeviceResponse{
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
	//   "description": "Claims a device for a customer and adds it to zero-touch enrollment. If the\ndevice is already claimed by another customer, the call returns an error.",
	//   "flatPath": "v1/partners/{partnersId}/devices:claim",
	//   "httpMethod": "POST",
	//   "id": "androiddeviceprovisioning.partners.devices.claim",
	//   "parameterOrder": [
	//     "partnerId"
	//   ],
	//   "parameters": {
	//     "partnerId": {
	//       "description": "Required. The ID of the reseller partner.",
	//       "format": "int64",
	//       "location": "path",
	//       "pattern": "^[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/partners/{+partnerId}/devices:claim",
	//   "request": {
	//     "$ref": "ClaimDeviceRequest"
	//   },
	//   "response": {
	//     "$ref": "ClaimDeviceResponse"
	//   }
	// }

}

// method id "androiddeviceprovisioning.partners.devices.claimAsync":

type PartnersDevicesClaimAsyncCall struct {
	s                   *Service
	partnerId           int64
	claimdevicesrequest *ClaimDevicesRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// ClaimAsync: Claims a batch of devices for a customer asynchronously.
// Adds the devices
// to zero-touch enrollment. To learn more, read [Long‑running
// batch
// operations](/zero-touch/guides/how-it-works#operations).
func (r *PartnersDevicesService) ClaimAsync(partnerId int64, claimdevicesrequest *ClaimDevicesRequest) *PartnersDevicesClaimAsyncCall {
	c := &PartnersDevicesClaimAsyncCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.partnerId = partnerId
	c.claimdevicesrequest = claimdevicesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PartnersDevicesClaimAsyncCall) Fields(s ...googleapi.Field) *PartnersDevicesClaimAsyncCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PartnersDevicesClaimAsyncCall) Context(ctx context.Context) *PartnersDevicesClaimAsyncCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PartnersDevicesClaimAsyncCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PartnersDevicesClaimAsyncCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.claimdevicesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/partners/{+partnerId}/devices:claimAsync")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"partnerId": strconv.FormatInt(c.partnerId, 10),
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.partners.devices.claimAsync"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.partners.devices.claimAsync" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *PartnersDevicesClaimAsyncCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Claims a batch of devices for a customer asynchronously. Adds the devices\nto zero-touch enrollment. To learn more, read [Long‑running batch\noperations](/zero-touch/guides/how-it-works#operations).",
	//   "flatPath": "v1/partners/{partnersId}/devices:claimAsync",
	//   "httpMethod": "POST",
	//   "id": "androiddeviceprovisioning.partners.devices.claimAsync",
	//   "parameterOrder": [
	//     "partnerId"
	//   ],
	//   "parameters": {
	//     "partnerId": {
	//       "description": "Required. The ID of the reseller partner.",
	//       "format": "int64",
	//       "location": "path",
	//       "pattern": "^[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/partners/{+partnerId}/devices:claimAsync",
	//   "request": {
	//     "$ref": "ClaimDevicesRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   }
	// }

}

// method id "androiddeviceprovisioning.partners.devices.findByIdentifier":

type PartnersDevicesFindByIdentifierCall struct {
	s                                    *Service
	partnerId                            int64
	finddevicesbydeviceidentifierrequest *FindDevicesByDeviceIdentifierRequest
	urlParams_                           gensupport.URLParams
	ctx_                                 context.Context
	header_                              http.Header
}

// FindByIdentifier: Finds devices by hardware identifiers, such as
// IMEI.
func (r *PartnersDevicesService) FindByIdentifier(partnerId int64, finddevicesbydeviceidentifierrequest *FindDevicesByDeviceIdentifierRequest) *PartnersDevicesFindByIdentifierCall {
	c := &PartnersDevicesFindByIdentifierCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.partnerId = partnerId
	c.finddevicesbydeviceidentifierrequest = finddevicesbydeviceidentifierrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PartnersDevicesFindByIdentifierCall) Fields(s ...googleapi.Field) *PartnersDevicesFindByIdentifierCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PartnersDevicesFindByIdentifierCall) Context(ctx context.Context) *PartnersDevicesFindByIdentifierCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PartnersDevicesFindByIdentifierCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PartnersDevicesFindByIdentifierCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.finddevicesbydeviceidentifierrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/partners/{+partnerId}/devices:findByIdentifier")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"partnerId": strconv.FormatInt(c.partnerId, 10),
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.partners.devices.findByIdentifier"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.partners.devices.findByIdentifier" call.
// Exactly one of *FindDevicesByDeviceIdentifierResponse or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *FindDevicesByDeviceIdentifierResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PartnersDevicesFindByIdentifierCall) Do(opts ...googleapi.CallOption) (*FindDevicesByDeviceIdentifierResponse, error) {
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
	ret := &FindDevicesByDeviceIdentifierResponse{
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
	//   "description": "Finds devices by hardware identifiers, such as IMEI.",
	//   "flatPath": "v1/partners/{partnersId}/devices:findByIdentifier",
	//   "httpMethod": "POST",
	//   "id": "androiddeviceprovisioning.partners.devices.findByIdentifier",
	//   "parameterOrder": [
	//     "partnerId"
	//   ],
	//   "parameters": {
	//     "partnerId": {
	//       "description": "Required. The ID of the reseller partner.",
	//       "format": "int64",
	//       "location": "path",
	//       "pattern": "^[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/partners/{+partnerId}/devices:findByIdentifier",
	//   "request": {
	//     "$ref": "FindDevicesByDeviceIdentifierRequest"
	//   },
	//   "response": {
	//     "$ref": "FindDevicesByDeviceIdentifierResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *PartnersDevicesFindByIdentifierCall) Pages(ctx context.Context, f func(*FindDevicesByDeviceIdentifierResponse) error) error {
	c.ctx_ = ctx
	defer func(pt string) { c.finddevicesbydeviceidentifierrequest.PageToken = pt }(c.finddevicesbydeviceidentifierrequest.PageToken) // reset paging to original point
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
		c.finddevicesbydeviceidentifierrequest.PageToken = x.NextPageToken
	}
}

// method id "androiddeviceprovisioning.partners.devices.findByOwner":

type PartnersDevicesFindByOwnerCall struct {
	s                         *Service
	partnerId                 int64
	finddevicesbyownerrequest *FindDevicesByOwnerRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// FindByOwner: Finds devices claimed for customers. The results only
// contain devices
// registered to the reseller that's identified by the `partnerId`
// argument.
// The customer's devices purchased from other resellers don't appear in
// the
// results.
func (r *PartnersDevicesService) FindByOwner(partnerId int64, finddevicesbyownerrequest *FindDevicesByOwnerRequest) *PartnersDevicesFindByOwnerCall {
	c := &PartnersDevicesFindByOwnerCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.partnerId = partnerId
	c.finddevicesbyownerrequest = finddevicesbyownerrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PartnersDevicesFindByOwnerCall) Fields(s ...googleapi.Field) *PartnersDevicesFindByOwnerCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PartnersDevicesFindByOwnerCall) Context(ctx context.Context) *PartnersDevicesFindByOwnerCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PartnersDevicesFindByOwnerCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PartnersDevicesFindByOwnerCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.finddevicesbyownerrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/partners/{+partnerId}/devices:findByOwner")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"partnerId": strconv.FormatInt(c.partnerId, 10),
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.partners.devices.findByOwner"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.partners.devices.findByOwner" call.
// Exactly one of *FindDevicesByOwnerResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *FindDevicesByOwnerResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PartnersDevicesFindByOwnerCall) Do(opts ...googleapi.CallOption) (*FindDevicesByOwnerResponse, error) {
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
	ret := &FindDevicesByOwnerResponse{
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
	//   "description": "Finds devices claimed for customers. The results only contain devices\nregistered to the reseller that's identified by the `partnerId` argument.\nThe customer's devices purchased from other resellers don't appear in the\nresults.",
	//   "flatPath": "v1/partners/{partnersId}/devices:findByOwner",
	//   "httpMethod": "POST",
	//   "id": "androiddeviceprovisioning.partners.devices.findByOwner",
	//   "parameterOrder": [
	//     "partnerId"
	//   ],
	//   "parameters": {
	//     "partnerId": {
	//       "description": "Required. The ID of the reseller partner.",
	//       "format": "int64",
	//       "location": "path",
	//       "pattern": "^[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/partners/{+partnerId}/devices:findByOwner",
	//   "request": {
	//     "$ref": "FindDevicesByOwnerRequest"
	//   },
	//   "response": {
	//     "$ref": "FindDevicesByOwnerResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *PartnersDevicesFindByOwnerCall) Pages(ctx context.Context, f func(*FindDevicesByOwnerResponse) error) error {
	c.ctx_ = ctx
	defer func(pt string) { c.finddevicesbyownerrequest.PageToken = pt }(c.finddevicesbyownerrequest.PageToken) // reset paging to original point
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
		c.finddevicesbyownerrequest.PageToken = x.NextPageToken
	}
}

// method id "androiddeviceprovisioning.partners.devices.get":

type PartnersDevicesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a device.
func (r *PartnersDevicesService) Get(name string) *PartnersDevicesGetCall {
	c := &PartnersDevicesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PartnersDevicesGetCall) Fields(s ...googleapi.Field) *PartnersDevicesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *PartnersDevicesGetCall) IfNoneMatch(entityTag string) *PartnersDevicesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PartnersDevicesGetCall) Context(ctx context.Context) *PartnersDevicesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PartnersDevicesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PartnersDevicesGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.partners.devices.get"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.partners.devices.get" call.
// Exactly one of *Device or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Device.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *PartnersDevicesGetCall) Do(opts ...googleapi.CallOption) (*Device, error) {
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
	ret := &Device{
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
	//   "description": "Gets a device.",
	//   "flatPath": "v1/partners/{partnersId}/devices/{devicesId}",
	//   "httpMethod": "GET",
	//   "id": "androiddeviceprovisioning.partners.devices.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The device API resource name in the format\n`partners/[PARTNER_ID]/devices/[DEVICE_ID]`.",
	//       "location": "path",
	//       "pattern": "^partners/[^/]+/devices/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Device"
	//   }
	// }

}

// method id "androiddeviceprovisioning.partners.devices.metadata":

type PartnersDevicesMetadataCall struct {
	s                           *Service
	metadataOwnerId             int64
	deviceId                    int64
	updatedevicemetadatarequest *UpdateDeviceMetadataRequest
	urlParams_                  gensupport.URLParams
	ctx_                        context.Context
	header_                     http.Header
}

// Metadata: Updates reseller metadata associated with the device.
func (r *PartnersDevicesService) Metadata(metadataOwnerId int64, deviceId int64, updatedevicemetadatarequest *UpdateDeviceMetadataRequest) *PartnersDevicesMetadataCall {
	c := &PartnersDevicesMetadataCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.metadataOwnerId = metadataOwnerId
	c.deviceId = deviceId
	c.updatedevicemetadatarequest = updatedevicemetadatarequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PartnersDevicesMetadataCall) Fields(s ...googleapi.Field) *PartnersDevicesMetadataCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PartnersDevicesMetadataCall) Context(ctx context.Context) *PartnersDevicesMetadataCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PartnersDevicesMetadataCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PartnersDevicesMetadataCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.updatedevicemetadatarequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/partners/{+metadataOwnerId}/devices/{+deviceId}/metadata")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"metadataOwnerId": strconv.FormatInt(c.metadataOwnerId, 10),
		"deviceId":        strconv.FormatInt(c.deviceId, 10),
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.partners.devices.metadata"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.partners.devices.metadata" call.
// Exactly one of *DeviceMetadata or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *DeviceMetadata.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PartnersDevicesMetadataCall) Do(opts ...googleapi.CallOption) (*DeviceMetadata, error) {
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
	ret := &DeviceMetadata{
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
	//   "description": "Updates reseller metadata associated with the device.",
	//   "flatPath": "v1/partners/{partnersId}/devices/{devicesId}/metadata",
	//   "httpMethod": "POST",
	//   "id": "androiddeviceprovisioning.partners.devices.metadata",
	//   "parameterOrder": [
	//     "metadataOwnerId",
	//     "deviceId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "Required. The ID of the reseller partner.",
	//       "format": "int64",
	//       "location": "path",
	//       "pattern": "^[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "metadataOwnerId": {
	//       "description": "Required. The owner of the newly set metadata. Set this to the partner ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "pattern": "^[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/partners/{+metadataOwnerId}/devices/{+deviceId}/metadata",
	//   "request": {
	//     "$ref": "UpdateDeviceMetadataRequest"
	//   },
	//   "response": {
	//     "$ref": "DeviceMetadata"
	//   }
	// }

}

// method id "androiddeviceprovisioning.partners.devices.unclaim":

type PartnersDevicesUnclaimCall struct {
	s                    *Service
	partnerId            int64
	unclaimdevicerequest *UnclaimDeviceRequest
	urlParams_           gensupport.URLParams
	ctx_                 context.Context
	header_              http.Header
}

// Unclaim: Unclaims a device from a customer and removes it from
// zero-touch
// enrollment.
func (r *PartnersDevicesService) Unclaim(partnerId int64, unclaimdevicerequest *UnclaimDeviceRequest) *PartnersDevicesUnclaimCall {
	c := &PartnersDevicesUnclaimCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.partnerId = partnerId
	c.unclaimdevicerequest = unclaimdevicerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PartnersDevicesUnclaimCall) Fields(s ...googleapi.Field) *PartnersDevicesUnclaimCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PartnersDevicesUnclaimCall) Context(ctx context.Context) *PartnersDevicesUnclaimCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PartnersDevicesUnclaimCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PartnersDevicesUnclaimCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.unclaimdevicerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/partners/{+partnerId}/devices:unclaim")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"partnerId": strconv.FormatInt(c.partnerId, 10),
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.partners.devices.unclaim"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.partners.devices.unclaim" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *PartnersDevicesUnclaimCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	//   "description": "Unclaims a device from a customer and removes it from zero-touch\nenrollment.",
	//   "flatPath": "v1/partners/{partnersId}/devices:unclaim",
	//   "httpMethod": "POST",
	//   "id": "androiddeviceprovisioning.partners.devices.unclaim",
	//   "parameterOrder": [
	//     "partnerId"
	//   ],
	//   "parameters": {
	//     "partnerId": {
	//       "description": "Required. The ID of the reseller partner.",
	//       "format": "int64",
	//       "location": "path",
	//       "pattern": "^[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/partners/{+partnerId}/devices:unclaim",
	//   "request": {
	//     "$ref": "UnclaimDeviceRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   }
	// }

}

// method id "androiddeviceprovisioning.partners.devices.unclaimAsync":

type PartnersDevicesUnclaimAsyncCall struct {
	s                     *Service
	partnerId             int64
	unclaimdevicesrequest *UnclaimDevicesRequest
	urlParams_            gensupport.URLParams
	ctx_                  context.Context
	header_               http.Header
}

// UnclaimAsync: Unclaims a batch of devices for a customer
// asynchronously. Removes the
// devices from zero-touch enrollment. To learn more, read
// [Long‑running
// batch
// operations](/zero-touch/guides/how-it-works#operations).
func (r *PartnersDevicesService) UnclaimAsync(partnerId int64, unclaimdevicesrequest *UnclaimDevicesRequest) *PartnersDevicesUnclaimAsyncCall {
	c := &PartnersDevicesUnclaimAsyncCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.partnerId = partnerId
	c.unclaimdevicesrequest = unclaimdevicesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PartnersDevicesUnclaimAsyncCall) Fields(s ...googleapi.Field) *PartnersDevicesUnclaimAsyncCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PartnersDevicesUnclaimAsyncCall) Context(ctx context.Context) *PartnersDevicesUnclaimAsyncCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PartnersDevicesUnclaimAsyncCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PartnersDevicesUnclaimAsyncCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.unclaimdevicesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/partners/{+partnerId}/devices:unclaimAsync")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"partnerId": strconv.FormatInt(c.partnerId, 10),
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.partners.devices.unclaimAsync"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.partners.devices.unclaimAsync" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *PartnersDevicesUnclaimAsyncCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Unclaims a batch of devices for a customer asynchronously. Removes the\ndevices from zero-touch enrollment. To learn more, read [Long‑running batch\noperations](/zero-touch/guides/how-it-works#operations).",
	//   "flatPath": "v1/partners/{partnersId}/devices:unclaimAsync",
	//   "httpMethod": "POST",
	//   "id": "androiddeviceprovisioning.partners.devices.unclaimAsync",
	//   "parameterOrder": [
	//     "partnerId"
	//   ],
	//   "parameters": {
	//     "partnerId": {
	//       "description": "Required. The reseller partner ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "pattern": "^[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/partners/{+partnerId}/devices:unclaimAsync",
	//   "request": {
	//     "$ref": "UnclaimDevicesRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   }
	// }

}

// method id "androiddeviceprovisioning.partners.devices.updateMetadataAsync":

type PartnersDevicesUpdateMetadataAsyncCall struct {
	s                                  *Service
	partnerId                          int64
	updatedevicemetadatainbatchrequest *UpdateDeviceMetadataInBatchRequest
	urlParams_                         gensupport.URLParams
	ctx_                               context.Context
	header_                            http.Header
}

// UpdateMetadataAsync: Updates the reseller metadata attached to a
// batch of devices. This method
// updates devices asynchronously and returns an `Operation` that can be
// used
// to track progress. Read [Long‑running
// batch
// operations](/zero-touch/guides/how-it-works#operations).
func (r *PartnersDevicesService) UpdateMetadataAsync(partnerId int64, updatedevicemetadatainbatchrequest *UpdateDeviceMetadataInBatchRequest) *PartnersDevicesUpdateMetadataAsyncCall {
	c := &PartnersDevicesUpdateMetadataAsyncCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.partnerId = partnerId
	c.updatedevicemetadatainbatchrequest = updatedevicemetadatainbatchrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PartnersDevicesUpdateMetadataAsyncCall) Fields(s ...googleapi.Field) *PartnersDevicesUpdateMetadataAsyncCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PartnersDevicesUpdateMetadataAsyncCall) Context(ctx context.Context) *PartnersDevicesUpdateMetadataAsyncCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PartnersDevicesUpdateMetadataAsyncCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PartnersDevicesUpdateMetadataAsyncCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.updatedevicemetadatainbatchrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/partners/{+partnerId}/devices:updateMetadataAsync")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"partnerId": strconv.FormatInt(c.partnerId, 10),
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.partners.devices.updateMetadataAsync"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.partners.devices.updateMetadataAsync" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *PartnersDevicesUpdateMetadataAsyncCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Updates the reseller metadata attached to a batch of devices. This method\nupdates devices asynchronously and returns an `Operation` that can be used\nto track progress. Read [Long‑running batch\noperations](/zero-touch/guides/how-it-works#operations).",
	//   "flatPath": "v1/partners/{partnersId}/devices:updateMetadataAsync",
	//   "httpMethod": "POST",
	//   "id": "androiddeviceprovisioning.partners.devices.updateMetadataAsync",
	//   "parameterOrder": [
	//     "partnerId"
	//   ],
	//   "parameters": {
	//     "partnerId": {
	//       "description": "Required. The reseller partner ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "pattern": "^[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/partners/{+partnerId}/devices:updateMetadataAsync",
	//   "request": {
	//     "$ref": "UpdateDeviceMetadataInBatchRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   }
	// }

}

// method id "androiddeviceprovisioning.partners.vendors.list":

type PartnersVendorsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the vendors of the partner.
func (r *PartnersVendorsService) List(parent string) *PartnersVendorsListCall {
	c := &PartnersVendorsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to be returned.
func (c *PartnersVendorsListCall) PageSize(pageSize int64) *PartnersVendorsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results returned by the server.
func (c *PartnersVendorsListCall) PageToken(pageToken string) *PartnersVendorsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PartnersVendorsListCall) Fields(s ...googleapi.Field) *PartnersVendorsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *PartnersVendorsListCall) IfNoneMatch(entityTag string) *PartnersVendorsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PartnersVendorsListCall) Context(ctx context.Context) *PartnersVendorsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PartnersVendorsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PartnersVendorsListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/vendors")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.partners.vendors.list"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.partners.vendors.list" call.
// Exactly one of *ListVendorsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListVendorsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PartnersVendorsListCall) Do(opts ...googleapi.CallOption) (*ListVendorsResponse, error) {
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
	ret := &ListVendorsResponse{
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
	//   "description": "Lists the vendors of the partner.",
	//   "flatPath": "v1/partners/{partnersId}/vendors",
	//   "httpMethod": "GET",
	//   "id": "androiddeviceprovisioning.partners.vendors.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of results to be returned.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results returned by the server.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The resource name in the format `partners/[PARTNER_ID]`.",
	//       "location": "path",
	//       "pattern": "^partners/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/vendors",
	//   "response": {
	//     "$ref": "ListVendorsResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *PartnersVendorsListCall) Pages(ctx context.Context, f func(*ListVendorsResponse) error) error {
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

// method id "androiddeviceprovisioning.partners.vendors.customers.list":

type PartnersVendorsCustomersListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the customers of the vendor.
func (r *PartnersVendorsCustomersService) List(parent string) *PartnersVendorsCustomersListCall {
	c := &PartnersVendorsCustomersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to be returned.
func (c *PartnersVendorsCustomersListCall) PageSize(pageSize int64) *PartnersVendorsCustomersListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results returned by the server.
func (c *PartnersVendorsCustomersListCall) PageToken(pageToken string) *PartnersVendorsCustomersListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PartnersVendorsCustomersListCall) Fields(s ...googleapi.Field) *PartnersVendorsCustomersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *PartnersVendorsCustomersListCall) IfNoneMatch(entityTag string) *PartnersVendorsCustomersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PartnersVendorsCustomersListCall) Context(ctx context.Context) *PartnersVendorsCustomersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PartnersVendorsCustomersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PartnersVendorsCustomersListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/customers")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(googleapi.MethodIDToContext(c.ctx_, "androiddeviceprovisioning.partners.vendors.customers.list"), c.s.client, req)
}

// Do executes the "androiddeviceprovisioning.partners.vendors.customers.list" call.
// Exactly one of *ListVendorCustomersResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListVendorCustomersResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PartnersVendorsCustomersListCall) Do(opts ...googleapi.CallOption) (*ListVendorCustomersResponse, error) {
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
	ret := &ListVendorCustomersResponse{
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
	//   "description": "Lists the customers of the vendor.",
	//   "flatPath": "v1/partners/{partnersId}/vendors/{vendorsId}/customers",
	//   "httpMethod": "GET",
	//   "id": "androiddeviceprovisioning.partners.vendors.customers.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of results to be returned.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results returned by the server.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The resource name in the format\n`partners/[PARTNER_ID]/vendors/[VENDOR_ID]`.",
	//       "location": "path",
	//       "pattern": "^partners/[^/]+/vendors/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/customers",
	//   "response": {
	//     "$ref": "ListVendorCustomersResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *PartnersVendorsCustomersListCall) Pages(ctx context.Context, f func(*ListVendorCustomersResponse) error) error {
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
