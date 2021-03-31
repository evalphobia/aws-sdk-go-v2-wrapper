package pinpoint

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

type EndpointRequest struct {
	Address        string // Email | Push Token | Tel No.
	Attributes     map[string][]string
	ChannelType    ChannelType
	Demographic    EndpointDemographic
	EffectiveDate  string
	EndpointStatus string
	Location       EndpointLocation
	Metrics        map[string]float64
	OptOut         string
	RequestID      string
	User           EndpointUser
}

func (r EndpointRequest) ToSDK() *SDK.EndpointRequest {
	o := SDK.EndpointRequest{}

	if r.Address != "" {
		o.Address = pointers.String(r.Address)
	}

	o.Attributes = r.Attributes
	o.ChannelType = SDK.ChannelType(r.ChannelType)
	o.Demographic = r.Demographic.ToSDK()

	if r.EffectiveDate != "" {
		o.EffectiveDate = pointers.String(r.EffectiveDate)
	}
	if r.EndpointStatus != "" {
		o.EndpointStatus = pointers.String(r.EndpointStatus)
	}

	o.Location = r.Location.ToSDK()
	o.Metrics = r.Metrics

	if r.OptOut != "" {
		o.OptOut = pointers.String(r.OptOut)
	}
	if r.RequestID != "" {
		o.RequestId = pointers.String(r.RequestID)
	}

	o.User = r.User.ToSDK()
	return &o
}

type EndpointBatchRequest struct {
	Item []EndpointBatchItem // max: 100
}

func (r EndpointBatchRequest) ToSDK() *SDK.EndpointBatchRequest {
	o := SDK.EndpointBatchRequest{}

	if len(r.Item) != 0 {
		list := make([]SDK.EndpointBatchItem, len(r.Item))
		for i, v := range r.Item {
			list[i] = v.ToSDK()
		}
		o.Item = list
	}

	return &o
}

type EndpointBatchItem struct {
	Address        string // Email | Push Token | Tel No.
	Attributes     map[string][]string
	ChannelType    ChannelType
	Demographic    EndpointDemographic
	EffectiveDate  string
	EndpointStatus string
	ID             string
	Location       EndpointLocation
	Metrics        map[string]float64
	OptOut         string
	RequestID      string
	User           EndpointUser
}

func (r EndpointBatchItem) ToSDK() SDK.EndpointBatchItem {
	o := SDK.EndpointBatchItem{}

	if r.Address != "" {
		o.Address = pointers.String(r.Address)
	}

	o.Attributes = r.Attributes
	o.ChannelType = SDK.ChannelType(r.ChannelType)
	o.Demographic = r.Demographic.ToSDK()

	if r.EffectiveDate != "" {
		o.EffectiveDate = pointers.String(r.EffectiveDate)
	}
	if r.EndpointStatus != "" {
		o.EndpointStatus = pointers.String(r.EndpointStatus)
	}
	if r.ID != "" {
		o.Id = pointers.String(r.ID)
	}

	o.Location = r.Location.ToSDK()
	o.Metrics = r.Metrics

	if r.OptOut != "" {
		o.OptOut = pointers.String(r.OptOut)
	}
	if r.RequestID != "" {
		o.RequestId = pointers.String(r.RequestID)
	}

	o.User = r.User.ToSDK()
	return o
}

type EndpointsResponse struct {
	Item []EndpointResponse
}

func newEndpointsResponse(o *SDK.EndpointsResponse) EndpointsResponse {
	result := EndpointsResponse{}
	if o == nil {
		return result
	}

	if len(o.Item) != 0 {
		list := make([]EndpointResponse, len(o.Item))
		for i, v := range o.Item {
			v := v
			list[i] = newEndpointResponse(&v)
		}
		result.Item = list
	}

	return result
}

type EndpointResponse struct {
	Address        string // Email | Push Token | Tel No.
	ApplicationID  string
	Attributes     map[string][]string
	ChannelType    ChannelType
	CohortID       string
	CreationDate   string
	Demographic    EndpointDemographic
	EffectiveDate  string
	EndpointStatus string
	ID             string
	Location       EndpointLocation
	Metrics        map[string]float64
	OptOut         string
	RequestID      string
	User           EndpointUser
}

func newEndpointResponse(o *SDK.EndpointResponse) EndpointResponse {
	result := EndpointResponse{}
	if o == nil {
		return result
	}

	if o.Address != nil {
		result.Address = *o.Address
	}
	if o.ApplicationId != nil {
		result.ApplicationID = *o.ApplicationId
	}

	result.Attributes = o.Attributes
	result.ChannelType = ChannelType(o.ChannelType)

	if o.CohortId != nil {
		result.CohortID = *o.CohortId
	}
	if o.CreationDate != nil {
		result.CreationDate = *o.CreationDate
	}

	result.Demographic = newEndpointDemographic(o.Demographic)

	if o.EffectiveDate != nil {
		result.EffectiveDate = *o.EffectiveDate
	}
	if o.EndpointStatus != nil {
		result.EndpointStatus = *o.EndpointStatus
	}
	if o.Id != nil {
		result.ID = *o.Id
	}

	result.Location = newEndpointLocation(o.Location)
	result.Metrics = o.Metrics

	if o.OptOut != nil {
		result.OptOut = *o.OptOut
	}
	if o.RequestId != nil {
		result.RequestID = *o.RequestId
	}

	result.User = newEndpointUser(o.User)

	return result
}

type EndpointDemographic struct {
	AppVersion      string
	Locale          string
	Make            string
	Model           string
	ModelVersion    string
	Platform        string
	PlatformVersion string
	Timezone        string
}

func newEndpointDemographic(o *SDK.EndpointDemographic) EndpointDemographic {
	result := EndpointDemographic{}
	if o.AppVersion != nil {
		result.AppVersion = *o.AppVersion
	}
	if o.Locale != nil {
		result.Locale = *o.Locale
	}
	if o.Make != nil {
		result.Make = *o.Make
	}
	if o.Model != nil {
		result.Model = *o.Model
	}
	if o.ModelVersion != nil {
		result.ModelVersion = *o.ModelVersion
	}
	if o.Platform != nil {
		result.Platform = *o.Platform
	}
	if o.PlatformVersion != nil {
		result.PlatformVersion = *o.PlatformVersion
	}
	if o.Timezone != nil {
		result.Timezone = *o.Timezone
	}
	return result
}

func (r EndpointDemographic) ToSDK() *SDK.EndpointDemographic {
	o := SDK.EndpointDemographic{}

	if r.AppVersion != "" {
		o.AppVersion = pointers.String(r.AppVersion)
	}
	if r.Locale != "" {
		o.Locale = pointers.String(r.Locale)
	}
	if r.Make != "" {
		o.Make = pointers.String(r.Make)
	}
	if r.Model != "" {
		o.Model = pointers.String(r.Model)
	}
	if r.ModelVersion != "" {
		o.ModelVersion = pointers.String(r.ModelVersion)
	}
	if r.Platform != "" {
		o.Platform = pointers.String(r.Platform)
	}
	if r.PlatformVersion != "" {
		o.PlatformVersion = pointers.String(r.PlatformVersion)
	}
	if r.Timezone != "" {
		o.Timezone = pointers.String(r.Timezone)
	}
	return &o
}

type EndpointLocation struct {
	City       string
	Country    string
	Latitude   float64
	Longitude  float64
	PostalCode string
	Region     string

	HasLatitude  bool
	HasLongitude bool
}

func newEndpointLocation(o *SDK.EndpointLocation) EndpointLocation {
	result := EndpointLocation{}
	if o.City != nil {
		result.City = *o.City
	}
	if o.Country != nil {
		result.Country = *o.Country
	}
	if o.Latitude != nil {
		result.Latitude = *o.Latitude
	}
	if o.Longitude != nil {
		result.Longitude = *o.Longitude
	}
	if o.PostalCode != nil {
		result.PostalCode = *o.PostalCode
	}
	if o.Region != nil {
		result.Region = *o.Region
	}
	return result
}

func (r EndpointLocation) ToSDK() *SDK.EndpointLocation {
	o := SDK.EndpointLocation{}

	if r.City != "" {
		o.City = pointers.String(r.City)
	}
	if r.Country != "" {
		o.Country = pointers.String(r.Country)
	}
	if r.PostalCode != "" {
		o.PostalCode = pointers.String(r.PostalCode)
	}
	if r.Region != "" {
		o.Region = pointers.String(r.Region)
	}

	if r.Latitude != 0 || r.HasLatitude {
		o.Latitude = pointers.Float64(r.Latitude)
	}
	if r.Longitude != 0 || r.HasLongitude {
		o.Longitude = pointers.Float64(r.Longitude)
	}
	return &o
}

type EndpointUser struct {
	UserAttributes map[string][]string
	UserID         string
}

func newEndpointUser(o *SDK.EndpointUser) EndpointUser {
	result := EndpointUser{}

	result.UserAttributes = o.UserAttributes

	if o.UserId != nil {
		result.UserID = *o.UserId
	}
	return result
}

func (r EndpointUser) ToSDK() *SDK.EndpointUser {
	o := SDK.EndpointUser{}

	o.UserAttributes = r.UserAttributes

	if r.UserID != "" {
		o.UserId = pointers.String(r.UserID)
	}
	return &o
}
