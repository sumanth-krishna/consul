// Code generated by protoc-json-shim. DO NOT EDIT.
package tenancyv1alpha1

import (
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON is a custom marshaler for Namespace
func (this *Namespace) MarshalJSON() ([]byte, error) {
	str, err := NamespaceMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Namespace
func (this *Namespace) UnmarshalJSON(b []byte) error {
	return NamespaceUnmarshaler.Unmarshal(b, this)
}

var (
	NamespaceMarshaler   = &protojson.MarshalOptions{}
	NamespaceUnmarshaler = &protojson.UnmarshalOptions{DiscardUnknown: false}
)
