package s3

import (
	"time"

	SDK "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/private/pointers"
)

type AccessControlPolicy struct {
	Grants []Grant
	Owner  Owner
}

func (r AccessControlPolicy) ToSDK() *SDK.AccessControlPolicy {
	if len(r.Grants) == 0 {
		return nil
	}

	o := &SDK.AccessControlPolicy{}

	v := r.Owner.ToSDK()
	o.Owner = &v

	if len(r.Grants) != 0 {
		list := make([]SDK.Grant, len(r.Grants))
		for i, v := range r.Grants {
			list[i] = v.ToSDK()
		}
		o.Grants = list
	}

	return o
}

type Bucket struct {
	CreationDate time.Time
	Name         string
}

func newBuckets(list []SDK.Bucket) []Bucket {
	if len(list) == 0 {
		return nil
	}

	results := make([]Bucket, len(list))
	for i, v := range list {
		results[i] = newBucket(v)
	}
	return results
}

func newBucket(b SDK.Bucket) Bucket {
	result := Bucket{}

	if b.CreationDate != nil {
		result.CreationDate = *b.CreationDate
	}
	if b.Name != nil {
		result.Name = *b.Name
	}
	return result
}

type Delete struct {
	Objects []ObjectIdentifier
	Quiet   bool
}

// func newDelete(o *SDK.Delete) Delete {
// 	result := Delete{}
// 	if o == nil {
// 		return result
// 	}

// 	if o.Quiet != nil {
// 		result.Quiet = *o.Quiet
// 	}
// 	if len(o.Objects) != 0 {
// 		list := make([]ObjectIdentifier, len(o.Objects))
// 		for i, v := range o.Objects {
// 			list[i] = newObjectIdentifier(v)
// 		}
// 	}
// 	return result
// }

func (r Delete) ToSDK() SDK.Delete {
	o := SDK.Delete{}

	if r.Quiet {
		o.Quiet = pointers.Bool(r.Quiet)
	}
	if len(r.Objects) != 0 {
		list := make([]SDK.ObjectIdentifier, len(r.Objects))
		for i, v := range r.Objects {
			list[i] = v.ToSDK()
		}
		o.Objects = list
	}

	return o
}

type DeletedObject struct {
	DeleteMarker          bool
	DeleteMarkerVersionID string
	Key                   string
	VersionID             string
}

func newDeletedObjects(list []SDK.DeletedObject) []DeletedObject {
	if len(list) == 0 {
		return nil
	}

	result := make([]DeletedObject, len(list))
	for i, v := range list {
		result[i] = newDeletedObject(v)
	}
	return result
}

func newDeletedObject(o SDK.DeletedObject) DeletedObject {
	result := DeletedObject{}

	if o.DeleteMarker != nil {
		result.DeleteMarker = *o.DeleteMarker
	}
	if o.DeleteMarkerVersionId != nil {
		result.DeleteMarkerVersionID = *o.DeleteMarkerVersionId
	}
	if o.Key != nil {
		result.Key = *o.Key
	}
	if o.VersionId != nil {
		result.VersionID = *o.VersionId
	}
	return result
}

type DeleteMarkerEntry struct {
	IsLatest     bool
	Key          string
	LastModified time.Time
	Owner        Owner
	VersionID    string
}

func newDeleteMarkerEntries(list []SDK.DeleteMarkerEntry) []DeleteMarkerEntry {
	if len(list) == 0 {
		return nil
	}

	result := make([]DeleteMarkerEntry, len(list))
	for i, v := range list {
		result[i] = newDeleteMarkerEntry(v)
	}
	return result
}

func newDeleteMarkerEntry(o SDK.DeleteMarkerEntry) DeleteMarkerEntry {
	result := DeleteMarkerEntry{}

	if o.IsLatest != nil {
		result.IsLatest = *o.IsLatest
	}
	if o.Key != nil {
		result.Key = *o.Key
	}
	if o.LastModified != nil {
		result.LastModified = *o.LastModified
	}
	if o.VersionId != nil {
		result.VersionID = *o.VersionId
	}

	result.Owner = newOwner(o.Owner)
	return result
}

type Error struct {
	Code      string
	Key       string
	Message   string
	VersionID string
}

func newErrors(list []SDK.Error) []Error {
	if len(list) == 0 {
		return nil
	}

	result := make([]Error, len(list))
	for i, v := range list {
		result[i] = newError(v)
	}
	return result
}

func newError(o SDK.Error) Error {
	result := Error{}

	if o.Code != nil {
		result.Code = *o.Code
	}
	if o.Key != nil {
		result.Key = *o.Key
	}
	if o.Message != nil {
		result.Message = *o.Message
	}
	if o.VersionId != nil {
		result.VersionID = *o.VersionId
	}
	return result
}

type Grant struct {
	GranteeDisplayName  string
	GranteeEmailAddress string
	GranteeID           string
	GranteeType         Type
	GranteeURI          string
	Permission          Permission
}

func newGrants(list []SDK.Grant) []Grant {
	if len(list) == 0 {
		return nil
	}

	result := make([]Grant, len(list))
	for i, v := range list {
		result[i] = newGrant(v)
	}
	return result
}

func newGrant(o SDK.Grant) Grant {
	result := Grant{}

	g := o.Grantee
	if g.DisplayName != nil {
		result.GranteeDisplayName = *g.DisplayName
	}
	if g.EmailAddress != nil {
		result.GranteeEmailAddress = *g.EmailAddress
	}
	if g.ID != nil {
		result.GranteeID = *g.ID
	}
	if g.Type != "" {
		result.GranteeType = Type(g.Type)
	}
	if g.URI != nil {
		result.GranteeURI = *g.URI
	}
	if o.Permission != "" {
		result.Permission = Permission(o.Permission)
	}
	return result
}

func (r Grant) ToSDK() SDK.Grant {
	o := SDK.Grant{}

	g := SDK.Grantee{}
	if r.GranteeDisplayName != "" {
		g.DisplayName = pointers.String(r.GranteeDisplayName)
	}
	if r.GranteeEmailAddress != "" {
		g.EmailAddress = pointers.String(r.GranteeEmailAddress)
	}
	if r.GranteeID != "" {
		g.ID = pointers.String(r.GranteeID)
	}
	if r.GranteeType != "" {
		g.Type = SDK.Type(r.GranteeType)
	}
	if r.GranteeURI != "" {
		g.URI = pointers.String(r.GranteeURI)
	}

	o.Permission = SDK.Permission(r.Permission)
	return o
}

type LifecycleRule struct {
	Expiration                   LifecycleExpiration
	Filter                       LifecycleRuleFilter
	ID                           string
	NoncurrentVersionTransitions []NoncurrentVersionTransition
	Prefix                       string
	Status                       ExpirationStatus
	Transitions                  []Transition

	AbortIncompleteMultipartUploadDaysAfterInitiation int64
	NoncurrentVersionExpirationDays                   int64
}

func newLifecycleRules(list []SDK.LifecycleRule) []LifecycleRule {
	if len(list) == 0 {
		return nil
	}

	results := make([]LifecycleRule, len(list))
	for i, v := range list {
		results[i] = newLifecycleRule(v)
	}
	return results
}

func newLifecycleRule(o SDK.LifecycleRule) LifecycleRule {
	result := LifecycleRule{}

	if o.Filter != nil {
		result.Filter = newLifecycleRuleFilter(o.Filter)
	}
	if o.ID != nil {
		result.ID = *o.ID
	}
	if o.Prefix != nil {
		result.Prefix = *o.Prefix
	}

	if v := o.AbortIncompleteMultipartUpload; v != nil {
		if v.DaysAfterInitiation != nil {
			result.AbortIncompleteMultipartUploadDaysAfterInitiation = *v.DaysAfterInitiation
		}
	}
	if v := o.NoncurrentVersionExpiration; v != nil {
		if v.NoncurrentDays != nil {
			result.NoncurrentVersionExpirationDays = *v.NoncurrentDays
		}
	}

	result.Expiration = newLifecycleExpiration(o.Expiration)
	result.NoncurrentVersionTransitions = newNoncurrentVersionTransitions(o.NoncurrentVersionTransitions)
	result.Status = ExpirationStatus(o.Status)
	result.Transitions = newTransitions(o.Transitions)
	return result
}

func (r LifecycleRule) ToSDK() SDK.LifecycleRule {
	o := SDK.LifecycleRule{}

	if r.ID != "" {
		o.ID = pointers.String(r.ID)
	}
	if r.Prefix != "" {
		o.Prefix = pointers.String(r.Prefix)
	}
	if r.Status != "" {
		o.Status = SDK.ExpirationStatus(r.Status)
	}
	if r.AbortIncompleteMultipartUploadDaysAfterInitiation != 0 {
		o.AbortIncompleteMultipartUpload = &SDK.AbortIncompleteMultipartUpload{
			DaysAfterInitiation: pointers.Long64(r.AbortIncompleteMultipartUploadDaysAfterInitiation),
		}
	}
	if r.NoncurrentVersionExpirationDays != 0 {
		o.NoncurrentVersionExpiration = &SDK.NoncurrentVersionExpiration{
			NoncurrentDays: pointers.Long64(r.NoncurrentVersionExpirationDays),
		}
	}

	if len(r.NoncurrentVersionTransitions) != 0 {
		list := make([]SDK.NoncurrentVersionTransition, len(r.NoncurrentVersionTransitions))
		for i, v := range r.NoncurrentVersionTransitions {
			list[i] = v.ToSDK()
		}
		o.NoncurrentVersionTransitions = list
	}
	if len(r.Transitions) != 0 {
		list := make([]SDK.Transition, len(r.Transitions))
		for i, v := range r.Transitions {
			list[i] = v.ToSDK()
		}
		o.Transitions = list
	}

	if !r.Expiration.IsEmpty() {
		v := r.Expiration.ToSDK()
		o.Expiration = &v
	}
	if !r.Filter.IsEmpty() {
		v := r.Filter.ToSDK()
		o.Filter = &v
	}
	return o
}

type LifecycleRuleFilter struct {
	AndOperatorPrefix string
	AndOperatorTags   []Tag
	Prefix            string
	Tag               Tag
}

func newLifecycleRuleFilter(o *SDK.LifecycleRuleFilter) LifecycleRuleFilter {
	result := LifecycleRuleFilter{}
	if o == nil {
		return result
	}

	if o.And != nil {
		a := o.And
		if a.Prefix != nil {
			result.AndOperatorPrefix = *a.Prefix
		}
		if len(a.Tags) != 0 {
			result.AndOperatorTags = newTags(a.Tags)
		}

	}
	if o.Prefix != nil {
		result.Prefix = *o.Prefix
	}
	if o.Tag != nil {
		result.Tag = newTag(*o.Tag)
	}
	return result
}

func (r LifecycleRuleFilter) IsEmpty() bool {
	switch {
	case r.Prefix != "",
		r.AndOperatorPrefix != "",
		len(r.AndOperatorTags) != 0,
		!r.Tag.IsEmpty():
		return false
	}
	return true
}

func (r LifecycleRuleFilter) ToSDK() SDK.LifecycleRuleFilter {
	o := SDK.LifecycleRuleFilter{}
	if r.Prefix != "" {
		o.Prefix = pointers.String(r.Prefix)
	}
	if !r.Tag.IsEmpty() {
		v := r.Tag.ToSDK()
		o.Tag = &v
	}
	if r.AndOperatorPrefix != "" || len(r.AndOperatorTags) != 0 {
		and := &SDK.LifecycleRuleAndOperator{}
		if r.AndOperatorPrefix != "" {
			and.Prefix = pointers.String(r.AndOperatorPrefix)
		}
		if len(r.AndOperatorTags) != 0 {
			list := make([]SDK.Tag, len(r.AndOperatorTags))
			for i, v := range r.AndOperatorTags {
				list[i] = v.ToSDK()
			}
			and.Tags = list
		}
	}
	return o
}

type LifecycleExpiration struct {
	Date                      time.Time
	Days                      int64
	ExpiredObjectDeleteMarker bool
}

func newLifecycleExpiration(o *SDK.LifecycleExpiration) LifecycleExpiration {
	result := LifecycleExpiration{}
	if o == nil {
		return result
	}

	if o.Date != nil {
		result.Date = *o.Date
	}
	if o.Days != nil {
		result.Days = *o.Days
	}
	if o.ExpiredObjectDeleteMarker != nil {
		result.ExpiredObjectDeleteMarker = *o.ExpiredObjectDeleteMarker
	}
	return result
}

func (r LifecycleExpiration) IsEmpty() bool {
	switch {
	case r.Days != 0,
		r.ExpiredObjectDeleteMarker,
		!r.Date.IsZero():
		return false
	}
	return true
}

func (r LifecycleExpiration) ToSDK() SDK.LifecycleExpiration {
	o := SDK.LifecycleExpiration{}
	if r.Days != 0 {
		o.Days = pointers.Long64(r.Days)
	}
	if !r.Date.IsZero() {
		o.Date = &r.Date
	}
	if r.ExpiredObjectDeleteMarker {
		o.ExpiredObjectDeleteMarker = pointers.Bool(r.ExpiredObjectDeleteMarker)
	}
	return o
}

type NoncurrentVersionTransition struct {
	NoncurrentDays int64
	StorageClass   TransitionStorageClass
}

func newNoncurrentVersionTransitions(list []SDK.NoncurrentVersionTransition) []NoncurrentVersionTransition {
	if len(list) == 0 {
		return nil
	}

	results := make([]NoncurrentVersionTransition, len(list))
	for i, v := range list {
		results[i] = newNoncurrentVersionTransition(v)
	}
	return results
}

func newNoncurrentVersionTransition(o SDK.NoncurrentVersionTransition) NoncurrentVersionTransition {
	result := NoncurrentVersionTransition{}

	if o.NoncurrentDays != nil {
		result.NoncurrentDays = *o.NoncurrentDays
	}
	result.StorageClass = TransitionStorageClass(o.StorageClass)
	return result
}

func (r NoncurrentVersionTransition) ToSDK() SDK.NoncurrentVersionTransition {
	o := SDK.NoncurrentVersionTransition{}
	if r.NoncurrentDays != 0 {
		o.NoncurrentDays = pointers.Long64(r.NoncurrentDays)
	}
	if r.StorageClass != "" {
		o.StorageClass = SDK.TransitionStorageClass(r.StorageClass)
	}
	return o
}

type Object struct {
	ETag             string
	Key              string
	LastModified     time.Time
	Size             int64
	StorageClass     string
	OwnerID          string
	OwnerDisplayName string
}

func NewObject(o SDK.Object) Object {
	result := Object{}

	if o.ETag != nil {
		result.ETag = *o.ETag
	}
	if o.Key != nil {
		result.Key = *o.Key
	}
	if o.Size != nil {
		result.Size = *o.Size
	}
	if o.LastModified != nil {
		result.LastModified = *o.LastModified
	}
	if o.Owner != nil {
		owner := o.Owner
		if owner.ID != nil {
			result.OwnerID = *owner.ID
		}
		if owner.DisplayName != nil {
			result.OwnerDisplayName = *owner.DisplayName
		}
	}
	return result
}

type ObjectIdentifier struct {
	Key       string
	VersionID string
}

// func newObjectIdentifier(o SDK.ObjectIdentifier) ObjectIdentifier {
// 	result := ObjectIdentifier{}

// 	if o.Key != nil {
// 		result.Key = *o.Key
// 	}
// 	if o.VersionId != nil {
// 		result.VersionID = *o.VersionId
// 	}
// 	return result
// }

func (r ObjectIdentifier) ToSDK() SDK.ObjectIdentifier {
	o := SDK.ObjectIdentifier{}

	if r.Key != "" {
		o.Key = pointers.String(r.Key)
	}
	if r.VersionID != "" {
		o.VersionId = pointers.String(r.VersionID)
	}
	return o
}

type ObjectVersion struct {
	ETag         string
	IsLatest     bool
	Key          string
	LastModified time.Time
	Owner        Owner
	Size         int64
	StorageClass ObjectVersionStorageClass
	VersionID    string
}

func newObjectVersions(list []SDK.ObjectVersion) []ObjectVersion {
	if len(list) == 0 {
		return nil
	}

	result := make([]ObjectVersion, len(list))
	for i, v := range list {
		result[i] = newObjectVersion(v)
	}
	return result
}

func newObjectVersion(o SDK.ObjectVersion) ObjectVersion {
	result := ObjectVersion{}

	if o.ETag != nil {
		result.ETag = *o.ETag
	}
	if o.IsLatest != nil {
		result.IsLatest = *o.IsLatest
	}
	if o.Key != nil {
		result.Key = *o.Key
	}
	if o.LastModified != nil {
		result.LastModified = *o.LastModified
	}
	if o.Size != nil {
		result.Size = *o.Size
	}
	if o.VersionId != nil {
		result.VersionID = *o.VersionId
	}

	if o.StorageClass != "" {
		result.StorageClass = ObjectVersionStorageClass(o.StorageClass)
	}

	result.Owner = newOwner(o.Owner)
	return result
}

type Owner struct {
	DisplayName string
	ID          string
}

func newOwner(o *SDK.Owner) Owner {
	result := Owner{}
	if o == nil {
		return result
	}

	if o.DisplayName != nil {
		result.DisplayName = *o.DisplayName
	}
	if o.ID != nil {
		result.ID = *o.ID
	}
	return result
}

func (r Owner) ToSDK() SDK.Owner {
	o := SDK.Owner{}

	if r.DisplayName != "" {
		o.DisplayName = pointers.String(r.DisplayName)
	}
	if r.ID != "" {
		o.ID = pointers.String(r.ID)
	}
	return o
}

type Tag struct {
	Key   string
	Value string
}

func newTags(list []SDK.Tag) []Tag {
	if len(list) == 0 {
		return nil
	}

	results := make([]Tag, len(list))
	for i, v := range list {
		results[i] = newTag(v)
	}
	return results
}

func newTag(o SDK.Tag) Tag {
	result := Tag{}

	if o.Key != nil {
		result.Key = *o.Key
	}
	if o.Value != nil {
		result.Value = *o.Value
	}
	return result
}

func (r Tag) IsEmpty() bool {
	return r.Key == ""
}

func (r Tag) ToSDK() SDK.Tag {
	o := SDK.Tag{}

	if r.Key != "" {
		o.Key = pointers.String(r.Key)
	}
	if r.Value != "" {
		o.Value = pointers.String(r.Value)
	}
	return o
}

type Transition struct {
	Date         time.Time
	Days         int64
	StorageClass TransitionStorageClass
}

func newTransitions(list []SDK.Transition) []Transition {
	if len(list) == 0 {
		return nil
	}

	results := make([]Transition, len(list))
	for i, v := range list {
		results[i] = newTransition(v)
	}
	return results
}

func newTransition(o SDK.Transition) Transition {
	result := Transition{}

	if o.Date != nil {
		result.Date = *o.Date
	}
	if o.Days != nil {
		result.Days = *o.Days
	}
	result.StorageClass = TransitionStorageClass(o.StorageClass)
	return result
}

func (r Transition) ToSDK() SDK.Transition {
	o := SDK.Transition{}
	if !r.Date.IsZero() {
		o.Date = &r.Date
	}
	if r.Days != 0 {
		o.Days = pointers.Long64(r.Days)
	}
	if r.StorageClass != "" {
		o.StorageClass = SDK.TransitionStorageClass(r.StorageClass)
	}
	return o
}
