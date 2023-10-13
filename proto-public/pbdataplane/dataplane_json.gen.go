// Code generated by protoc-json-shim. DO NOT EDIT.
package pbdataplane

import (
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON is a custom marshaler for GetSupportedDataplaneFeaturesRequest
func (this *GetSupportedDataplaneFeaturesRequest) MarshalJSON() ([]byte, error) {
	str, err := DataplaneMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GetSupportedDataplaneFeaturesRequest
func (this *GetSupportedDataplaneFeaturesRequest) UnmarshalJSON(b []byte) error {
	return DataplaneUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for DataplaneFeatureSupport
func (this *DataplaneFeatureSupport) MarshalJSON() ([]byte, error) {
	str, err := DataplaneMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DataplaneFeatureSupport
func (this *DataplaneFeatureSupport) UnmarshalJSON(b []byte) error {
	return DataplaneUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for GetSupportedDataplaneFeaturesResponse
func (this *GetSupportedDataplaneFeaturesResponse) MarshalJSON() ([]byte, error) {
	str, err := DataplaneMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GetSupportedDataplaneFeaturesResponse
func (this *GetSupportedDataplaneFeaturesResponse) UnmarshalJSON(b []byte) error {
	return DataplaneUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for GetEnvoyBootstrapParamsRequest
func (this *GetEnvoyBootstrapParamsRequest) MarshalJSON() ([]byte, error) {
	str, err := DataplaneMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GetEnvoyBootstrapParamsRequest
func (this *GetEnvoyBootstrapParamsRequest) UnmarshalJSON(b []byte) error {
	return DataplaneUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for GetEnvoyBootstrapParamsResponse
func (this *GetEnvoyBootstrapParamsResponse) MarshalJSON() ([]byte, error) {
	str, err := DataplaneMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GetEnvoyBootstrapParamsResponse
func (this *GetEnvoyBootstrapParamsResponse) UnmarshalJSON(b []byte) error {
	return DataplaneUnmarshaler.Unmarshal(b, this)
}

var (
	DataplaneMarshaler   = &protojson.MarshalOptions{}
	DataplaneUnmarshaler = &protojson.UnmarshalOptions{DiscardUnknown: false}
)
