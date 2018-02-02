package webservices

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/xml"
	"reflect"
	"time"
	"unsafe"
)

const (
	rfc3339Format = "2006-01-02T15:04:05.0000000Z07:00"
)

// used to convert times from UTC to GMT before sending across the wire
var gmt = time.FixedZone("GMT", 0)

// internal type used for marshalling time in RFC1123 format
type timeRFC1123 struct {
	time.Time
}

// MarshalText implements the encoding.TextMarshaler interface for timeRFC1123.
func (t timeRFC1123) MarshalText() ([]byte, error) {
	return []byte(t.Format(time.RFC1123)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for timeRFC1123.
func (t *timeRFC1123) UnmarshalText(data []byte) (err error) {
	t.Time, err = time.Parse(time.RFC1123, string(data))
	return
}

// internal type used for marshalling time in RFC3339 format
type timeRFC3339 struct {
	time.Time
}

// MarshalText implements the encoding.TextMarshaler interface for timeRFC3339.
func (t timeRFC3339) MarshalText() ([]byte, error) {
	return []byte(t.Format(rfc3339Format)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for timeRFC3339.
func (t *timeRFC3339) UnmarshalText(data []byte) (err error) {
	t.Time, err = time.Parse(rfc3339Format, string(data))
	return
}

// internal type used for marshalling
type diagnosticsConfiguration struct {
	Level  DiagnosticsLevelType `json:"level,omitempty"`
	Expiry *timeRFC3339         `json:"expiry,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for DiagnosticsConfiguration.
func (dc DiagnosticsConfiguration) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if reflect.TypeOf((*DiagnosticsConfiguration)(nil)).Elem().Size() != reflect.TypeOf((*diagnosticsConfiguration)(nil)).Elem().Size() {
		panic("size mismatch between DiagnosticsConfiguration and diagnosticsConfiguration")
	}
	dc2 := (*diagnosticsConfiguration)(unsafe.Pointer(&dc))
	return e.EncodeElement(*dc2, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for DiagnosticsConfiguration.
func (dc *DiagnosticsConfiguration) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if reflect.TypeOf((*DiagnosticsConfiguration)(nil)).Elem().Size() != reflect.TypeOf((*diagnosticsConfiguration)(nil)).Elem().Size() {
		panic("size mismatch between DiagnosticsConfiguration and diagnosticsConfiguration")
	}
	dc2 := (*diagnosticsConfiguration)(unsafe.Pointer(dc))
	return d.DecodeElement(dc2, &start)
}

// internal type used for marshalling
type properties struct {
	Title                    *string                          `json:"title,omitempty"`
	Description              *string                          `json:"description,omitempty"`
	CreatedOn                *timeRFC3339                     `json:"createdOn,omitempty"`
	ModifiedOn               *timeRFC3339                     `json:"modifiedOn,omitempty"`
	ProvisioningState        ProvisioningStateType            `json:"provisioningState,omitempty"`
	Keys                     *Keys                            `json:"keys,omitempty"`
	ReadOnly                 *bool                            `json:"readOnly,omitempty"`
	SwaggerLocation          *string                          `json:"swaggerLocation,omitempty"`
	ExposeSampleData         *bool                            `json:"exposeSampleData,omitempty"`
	RealtimeConfiguration    *RealtimeConfiguration           `json:"realtimeConfiguration,omitempty"`
	Diagnostics              *DiagnosticsConfiguration        `json:"diagnostics,omitempty"`
	StorageAccount           *StorageAccount                  `json:"storageAccount,omitempty"`
	MachineLearningWorkspace *MachineLearningWorkspace        `json:"machineLearningWorkspace,omitempty"`
	CommitmentPlan           *CommitmentPlan                  `json:"commitmentPlan,omitempty"`
	Input                    *ServiceInputOutputSpecification `json:"input,omitempty"`
	Output                   *ServiceInputOutputSpecification `json:"output,omitempty"`
	ExampleRequest           *ExampleRequest                  `json:"exampleRequest,omitempty"`
	Assets                   map[string]AssetItem             `json:"assets,omitempty"`
	Parameters               map[string]Parameter             `json:"parameters,omitempty"`
	PayloadsInBlobStorage    *bool                            `json:"payloadsInBlobStorage,omitempty"`
	PayloadsLocation         *BlobLocation                    `json:"payloadsLocation,omitempty"`
	PackageType              PackageType                      `json:"packageType,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for Properties.
func (p Properties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if reflect.TypeOf((*Properties)(nil)).Elem().Size() != reflect.TypeOf((*properties)(nil)).Elem().Size() {
		panic("size mismatch between Properties and properties")
	}
	p2 := (*properties)(unsafe.Pointer(&p))
	return e.EncodeElement(*p2, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for Properties.
func (p *Properties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if reflect.TypeOf((*Properties)(nil)).Elem().Size() != reflect.TypeOf((*properties)(nil)).Elem().Size() {
		panic("size mismatch between Properties and properties")
	}
	p2 := (*properties)(unsafe.Pointer(p))
	return d.DecodeElement(p2, &start)
}

// internal type used for marshalling
type asyncOperationStatus struct {
	ID                *string                  `json:"id,omitempty"`
	Name              *string                  `json:"name,omitempty"`
	ProvisioningState ProvisioningStateType    `json:"provisioningState,omitempty"`
	StartTime         *timeRFC3339             `json:"startTime,omitempty"`
	EndTime           *timeRFC3339             `json:"endTime,omitempty"`
	PercentComplete   *float64                 `json:"percentComplete,omitempty"`
	ErrorInfo         *AsyncOperationErrorInfo `json:"errorInfo,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for AsyncOperationStatus.
func (aos AsyncOperationStatus) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if reflect.TypeOf((*AsyncOperationStatus)(nil)).Elem().Size() != reflect.TypeOf((*asyncOperationStatus)(nil)).Elem().Size() {
		panic("size mismatch between AsyncOperationStatus and asyncOperationStatus")
	}
	aos2 := (*asyncOperationStatus)(unsafe.Pointer(&aos))
	return e.EncodeElement(*aos2, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for AsyncOperationStatus.
func (aos *AsyncOperationStatus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if reflect.TypeOf((*AsyncOperationStatus)(nil)).Elem().Size() != reflect.TypeOf((*asyncOperationStatus)(nil)).Elem().Size() {
		panic("size mismatch between AsyncOperationStatus and asyncOperationStatus")
	}
	aos2 := (*asyncOperationStatus)(unsafe.Pointer(aos))
	return d.DecodeElement(aos2, &start)
}
