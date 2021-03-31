package pinpoint

import (
	SDK "github.com/aws/aws-sdk-go-v2/service/pinpoint"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

type GPSPointDimension struct {
	CoordinatesLatitude  float64
	CoordinatesLongitude float64

	// optional
	RangeInKilometers float64
}

func newGPSPointDimension(o *SDK.GPSPointDimension) GPSPointDimension {
	result := GPSPointDimension{}
	if o == nil {
		return result
	}

	if v := o.Coordinates; v != nil {
		if v.Latitude != nil {
			result.CoordinatesLatitude = *v.Latitude
		}
		if v.Longitude != nil {
			result.CoordinatesLongitude = *v.Longitude
		}
	}

	if o.RangeInKilometers != nil {
		result.RangeInKilometers = *o.RangeInKilometers
	}
	return result
}

func (r GPSPointDimension) ToSDK() *SDK.GPSPointDimension {
	o := SDK.GPSPointDimension{}

	o.Coordinates = &SDK.GPSCoordinates{
		Latitude:  pointers.Float64(r.CoordinatesLatitude),
		Longitude: pointers.Float64(r.CoordinatesLongitude),
	}

	if r.RangeInKilometers != 0 {
		o.RangeInKilometers = pointers.Float64(r.RangeInKilometers)
	}
	return &o
}

type SegmentBehaviors struct {
	RecencyDuration Duration
	RecencyType     RecencyType
}

func newSegmentBehaviors(o *SDK.SegmentBehaviors) SegmentBehaviors {
	result := SegmentBehaviors{}
	if o == nil {
		return result
	}

	if o.Recency != nil {
		result.RecencyDuration = Duration(o.Recency.Duration)
		result.RecencyType = RecencyType(o.Recency.RecencyType)
	}
	return result
}

func (r SegmentBehaviors) ToSDK() *SDK.SegmentBehaviors {
	o := SDK.SegmentBehaviors{}

	o.Recency.Duration = SDK.Duration(r.RecencyDuration)
	o.Recency.RecencyType = SDK.RecencyType(r.RecencyType)
	return &o
}

type SegmentDemographics struct {
	AppVersion SetDimension
	Channel    SetDimension
	DeviceType SetDimension
	Make       SetDimension
	Model      SetDimension
	Platform   SetDimension
}

func newSegmentDemographics(o *SDK.SegmentDemographics) SegmentDemographics {
	result := SegmentDemographics{}
	if o == nil {
		return result
	}

	result.AppVersion = newSetDimension(o.AppVersion)
	result.Channel = newSetDimension(o.Channel)
	result.DeviceType = newSetDimension(o.DeviceType)
	result.Make = newSetDimension(o.Make)
	result.Model = newSetDimension(o.Model)
	result.Platform = newSetDimension(o.Platform)
	return result
}

func (r SegmentDemographics) ToSDK() *SDK.SegmentDemographics {
	o := SDK.SegmentDemographics{}

	o.AppVersion = r.AppVersion.ToSDK()
	o.Channel = r.Channel.ToSDK()
	o.DeviceType = r.DeviceType.ToSDK()
	o.Make = r.Make.ToSDK()
	o.Model = r.Model.ToSDK()
	o.Platform = r.Platform.ToSDK()
	return &o
}

type SegmentDimensions struct {
	Attributes     map[string]AttributeDimension
	Behavior       SegmentBehaviors
	Demographic    SegmentDemographics
	Location       SegmentLocation
	Metrics        map[string]MetricDimension
	UserAttributes map[string]AttributeDimension
}

func newSegmentDimensions(o *SDK.SegmentDimensions) SegmentDimensions {
	result := SegmentDimensions{}
	if o == nil {
		return result
	}

	result.Attributes = newAttributeDimensionMap(o.Attributes)
	result.Behavior = newSegmentBehaviors(o.Behavior)
	result.Demographic = newSegmentDemographics(o.Demographic)
	result.Location = newSegmentLocation(o.Location)
	result.Metrics = newMetricDimensionMap(o.Metrics)
	result.UserAttributes = newAttributeDimensionMap(o.UserAttributes)
	return result
}

func (r SegmentDimensions) ToSDK() SDK.SegmentDimensions {
	o := SDK.SegmentDimensions{}

	o.Attributes = attributeDimensionMapToSDK(r.Attributes)
	o.Behavior = r.Behavior.ToSDK()
	o.Demographic = r.Demographic.ToSDK()
	o.Location = r.Location.ToSDK()
	o.Metrics = metricDimensionMapToSDK(r.Metrics)
	o.UserAttributes = attributeDimensionMapToSDK(r.UserAttributes)
	return o
}

type SegmentImportResource struct {
	Format  Format
	RoleARN string
	S3URL   string
	Size    int64

	// optional
	ChannelCounts map[string]int64
}

func newSegmentImportResource(o *SDK.SegmentImportResource) SegmentImportResource {
	result := SegmentImportResource{}
	if o == nil {
		return result
	}

	result.Format = Format(o.Format)

	if o.RoleArn != nil {
		result.RoleARN = *o.RoleArn
	}
	if o.S3Url != nil {
		result.S3URL = *o.S3Url
	}
	if o.Size != nil {
		result.Size = *o.Size
	}

	result.ChannelCounts = o.ChannelCounts
	return result
}

type SegmentGroup struct {
	Dimensions     []SegmentDimensions
	SourceSegments []SegmentReference
	SourceType     SourceType
	Type           Type
}

func newSegmentGroup(o *SDK.SegmentGroup) SegmentGroup {
	result := SegmentGroup{}
	if o == nil {
		return result
	}

	if len(o.Dimensions) != 0 {
		list := make([]SegmentDimensions, len(o.Dimensions))
		for i, v := range o.Dimensions {
			v := v
			list[i] = newSegmentDimensions(&v)
		}
		result.Dimensions = list
	}

	if len(o.SourceSegments) != 0 {
		list := make([]SegmentReference, len(o.SourceSegments))
		for i, v := range o.SourceSegments {
			v := v
			list[i] = newSegmentReference(&v)
		}
		result.SourceSegments = list
	}

	result.SourceType = SourceType(o.SourceType)
	result.Type = Type(o.Type)
	return result
}

func (r SegmentGroup) ToSDK() SDK.SegmentGroup {
	o := SDK.SegmentGroup{}

	if len(r.Dimensions) != 0 {
		list := make([]SDK.SegmentDimensions, len(r.Dimensions))
		for i, v := range r.Dimensions {
			list[i] = v.ToSDK()
		}
		o.Dimensions = list
	}

	if len(r.SourceSegments) != 0 {
		list := make([]SDK.SegmentReference, len(r.SourceSegments))
		for i, v := range r.SourceSegments {
			list[i] = v.ToSDK()
		}
		o.SourceSegments = list
	}

	o.SourceType = SDK.SourceType(r.SourceType)
	o.Type = SDK.Type(r.Type)
	return o
}

type SegmentGroupList struct {
	Groups  []SegmentGroup
	Include Include
}

func newSegmentGroupList(o *SDK.SegmentGroupList) SegmentGroupList {
	result := SegmentGroupList{}
	if o == nil {
		return result
	}

	if len(o.Groups) != 0 {
		list := make([]SegmentGroup, len(o.Groups))
		for i, v := range o.Groups {
			v := v
			list[i] = newSegmentGroup(&v)
		}
		result.Groups = list
	}

	result.Include = Include(o.Include)
	return result
}

func (r SegmentGroupList) ToSDK() *SDK.SegmentGroupList {
	o := SDK.SegmentGroupList{}

	if len(r.Groups) != 0 {
		list := make([]SDK.SegmentGroup, len(r.Groups))
		for i, v := range r.Groups {
			list[i] = v.ToSDK()
		}
		o.Groups = list
	}

	o.Include = SDK.Include(r.Include)
	return &o
}

type SegmentLocation struct {
	Country  SetDimension
	GPSPoint GPSPointDimension
}

func newSegmentLocation(o *SDK.SegmentLocation) SegmentLocation {
	result := SegmentLocation{}
	if o == nil {
		return result
	}

	result.Country = newSetDimension(o.Country)
	result.GPSPoint = newGPSPointDimension(o.GPSPoint)
	return result
}

func (r SegmentLocation) ToSDK() *SDK.SegmentLocation {
	o := SDK.SegmentLocation{}

	o.Country = r.Country.ToSDK()
	o.GPSPoint = r.GPSPoint.ToSDK()
	return &o
}

type SegmentReference struct {
	ID string

	// optional
	Version int64
}

func newSegmentReference(o *SDK.SegmentReference) SegmentReference {
	result := SegmentReference{}
	if o == nil {
		return result
	}

	if o.Id != nil {
		result.ID = *o.Id
	}
	if o.Version != nil {
		result.Version = *o.Version
	}
	return result
}

func (r SegmentReference) ToSDK() SDK.SegmentReference {
	o := SDK.SegmentReference{}

	if r.ID != "" {
		o.Id = pointers.String(r.ID)
	}
	if r.Version != 0 {
		o.Version = pointers.Long64(r.Version)
	}
	return o
}

type SegmentResponse struct {
	ApplicationID string
	ARN           string
	CreationDate  string
	ID            string
	SegmentType   SegmentType

	// optional
	Dimensions       SegmentDimensions
	ImportDefinition SegmentImportResource
	LastModifiedDate string
	Name             string
	SegmentGroups    SegmentGroupList
	Tags             map[string]string
	Version          int64
}

func newSegmentResponse(o *SDK.SegmentResponse) SegmentResponse {
	result := SegmentResponse{}
	if o == nil {
		return result
	}

	if o.ApplicationId != nil {
		result.ApplicationID = *o.ApplicationId
	}
	if o.Arn != nil {
		result.ARN = *o.Arn
	}
	if o.CreationDate != nil {
		result.CreationDate = *o.CreationDate
	}
	if o.Id != nil {
		result.ID = *o.Id
	}
	result.SegmentType = SegmentType(o.SegmentType)

	result.Dimensions = newSegmentDimensions(o.Dimensions)
	result.ImportDefinition = newSegmentImportResource(o.ImportDefinition)

	if o.LastModifiedDate != nil {
		result.LastModifiedDate = *o.LastModifiedDate
	}
	if o.Name != nil {
		result.Name = *o.Name
	}

	result.SegmentGroups = newSegmentGroupList(o.SegmentGroups)
	result.Tags = o.Tags

	if o.Version != nil {
		result.Version = *o.Version
	}
	return result
}

type SegmentsResponse struct {
	Item      []SegmentResponse
	NextToken string
}

func newSegmentsResponse(o *SDK.SegmentsResponse) SegmentsResponse {
	result := SegmentsResponse{}
	if o == nil {
		return result
	}

	if len(o.Item) != 0 {
		list := make([]SegmentResponse, len(o.Item))
		for i, v := range o.Item {
			v := v
			list[i] = newSegmentResponse(&v)
		}
		result.Item = list
	}

	if o.NextToken != nil {
		result.NextToken = *o.NextToken
	}

	return result
}

type WriteSegmentRequest struct {
	Dimensions    SegmentDimensions
	Name          string
	SegmentGroups SegmentGroupList
	Tags          map[string]string
}

func (r WriteSegmentRequest) ToSDK() *SDK.WriteSegmentRequest {
	o := SDK.WriteSegmentRequest{}

	vv := r.Dimensions.ToSDK()
	o.Dimensions = &vv

	if r.Name != "" {
		o.Name = pointers.String(r.Name)
	}

	o.SegmentGroups = r.SegmentGroups.ToSDK()
	o.Tags = r.Tags
	return &o
}
