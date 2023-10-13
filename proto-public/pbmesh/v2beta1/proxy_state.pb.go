// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: pbmesh/v2beta1/proxy_state.proto

package meshv2beta1

import (
	pbproxystate "github.com/hashicorp/consul/proto-public/pbmesh/v2beta1/pbproxystate"
	pbresource "github.com/hashicorp/consul/proto-public/pbresource"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProxyStateTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ProxyState is the partially filled out ProxyState resource. The Endpoints, LeafCertificates and TrustBundles fields will need filling in after the resource is stored.
	ProxyState *ProxyState `protobuf:"bytes,1,opt,name=proxy_state,json=proxyState,proto3" json:"proxy_state,omitempty"`
	// RequiredEndpoints is a map of arbitrary string names to endpoint refs that need fetching by the proxy state controller.
	RequiredEndpoints map[string]*pbproxystate.EndpointRef `protobuf:"bytes,2,rep,name=required_endpoints,json=requiredEndpoints,proto3" json:"required_endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// RequiredLeafCertificates is a map of arbitrary string names to leaf certificates that need fetching/generation by the proxy state controller.
	RequiredLeafCertificates map[string]*pbproxystate.LeafCertificateRef `protobuf:"bytes,3,rep,name=required_leaf_certificates,json=requiredLeafCertificates,proto3" json:"required_leaf_certificates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// RequiredTrustBundles is a map of arbitrary string names to trust bundle refs that need fetching by the proxy state controller.
	RequiredTrustBundles map[string]*pbproxystate.TrustBundleRef `protobuf:"bytes,4,rep,name=required_trust_bundles,json=requiredTrustBundles,proto3" json:"required_trust_bundles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ProxyStateTemplate) Reset() {
	*x = ProxyStateTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_proxy_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyStateTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyStateTemplate) ProtoMessage() {}

func (x *ProxyStateTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_proxy_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyStateTemplate.ProtoReflect.Descriptor instead.
func (*ProxyStateTemplate) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_proxy_state_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyStateTemplate) GetProxyState() *ProxyState {
	if x != nil {
		return x.ProxyState
	}
	return nil
}

func (x *ProxyStateTemplate) GetRequiredEndpoints() map[string]*pbproxystate.EndpointRef {
	if x != nil {
		return x.RequiredEndpoints
	}
	return nil
}

func (x *ProxyStateTemplate) GetRequiredLeafCertificates() map[string]*pbproxystate.LeafCertificateRef {
	if x != nil {
		return x.RequiredLeafCertificates
	}
	return nil
}

func (x *ProxyStateTemplate) GetRequiredTrustBundles() map[string]*pbproxystate.TrustBundleRef {
	if x != nil {
		return x.RequiredTrustBundles
	}
	return nil
}

type ProxyState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identity is a reference to the identity of the workload this proxy is for.
	Identity *pbresource.Reference `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	// Listeners is a list of listeners for this proxy.
	Listeners []*pbproxystate.Listener `protobuf:"bytes,2,rep,name=listeners,proto3" json:"listeners,omitempty"`
	// Clusters is a map from cluster name to clusters. The keys are referenced from listeners or routes.
	Clusters map[string]*pbproxystate.Cluster `protobuf:"bytes,3,rep,name=clusters,proto3" json:"clusters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Routes is a map from route name to routes. The keys are referenced from listeners.
	Routes map[string]*pbproxystate.Route `protobuf:"bytes,4,rep,name=routes,proto3" json:"routes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Endpoints is a map from cluster name to endpoints.
	Endpoints map[string]*pbproxystate.Endpoints `protobuf:"bytes,5,rep,name=endpoints,proto3" json:"endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// LeafCertificates is a map from UUID to leaf certificates.
	LeafCertificates map[string]*pbproxystate.LeafCertificate `protobuf:"bytes,6,rep,name=leaf_certificates,json=leafCertificates,proto3" json:"leaf_certificates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// TrustBundles is a map from peer name to trust bundles.
	TrustBundles map[string]*pbproxystate.TrustBundle `protobuf:"bytes,7,rep,name=trust_bundles,json=trustBundles,proto3" json:"trust_bundles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// TLS has TLS configuration for this proxy.
	Tls *pbproxystate.TLS `protobuf:"bytes,8,opt,name=tls,proto3" json:"tls,omitempty"`
	// Escape defines top level escape hatches. These are user configured json strings that configure an entire piece of listener or cluster Envoy configuration.
	Escape *pbproxystate.EscapeHatches `protobuf:"bytes,9,opt,name=escape,proto3" json:"escape,omitempty"`
	// AccessLogs configures access logging for this proxy.
	AccessLogs *pbproxystate.AccessLogs `protobuf:"bytes,10,opt,name=access_logs,json=accessLogs,proto3" json:"access_logs,omitempty"`
}

func (x *ProxyState) Reset() {
	*x = ProxyState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_proxy_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyState) ProtoMessage() {}

func (x *ProxyState) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_proxy_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyState.ProtoReflect.Descriptor instead.
func (*ProxyState) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_proxy_state_proto_rawDescGZIP(), []int{1}
}

func (x *ProxyState) GetIdentity() *pbresource.Reference {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *ProxyState) GetListeners() []*pbproxystate.Listener {
	if x != nil {
		return x.Listeners
	}
	return nil
}

func (x *ProxyState) GetClusters() map[string]*pbproxystate.Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *ProxyState) GetRoutes() map[string]*pbproxystate.Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *ProxyState) GetEndpoints() map[string]*pbproxystate.Endpoints {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *ProxyState) GetLeafCertificates() map[string]*pbproxystate.LeafCertificate {
	if x != nil {
		return x.LeafCertificates
	}
	return nil
}

func (x *ProxyState) GetTrustBundles() map[string]*pbproxystate.TrustBundle {
	if x != nil {
		return x.TrustBundles
	}
	return nil
}

func (x *ProxyState) GetTls() *pbproxystate.TLS {
	if x != nil {
		return x.Tls
	}
	return nil
}

func (x *ProxyState) GetEscape() *pbproxystate.EscapeHatches {
	if x != nil {
		return x.Escape
	}
	return nil
}

func (x *ProxyState) GetAccessLogs() *pbproxystate.AccessLogs {
	if x != nil {
		return x.AccessLogs
	}
	return nil
}

var File_pbmesh_v2beta1_proxy_state_proto protoreflect.FileDescriptor

var file_pbmesh_v2beta1_proxy_state_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1d, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x2d, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x29, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x62, 0x6d,
	0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x62, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x73, 0x63, 0x61, 0x70, 0x65, 0x5f, 0x68, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x70, 0x62, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x62, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x70,
	0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x62,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x07, 0x0a, 0x12, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x77, 0x0a,
	0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x8d, 0x01, 0x0a, 0x1a, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x5f, 0x6c, 0x65, 0x61, 0x66, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x66, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x18, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x66, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x12, 0x81, 0x01, 0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x54, 0x72, 0x75, 0x73, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x14, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x54, 0x72,
	0x75, 0x73, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x1a, 0x7d, 0x0a, 0x16, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x4d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x8b, 0x01, 0x0a, 0x1d, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x66, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x54, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x65, 0x61, 0x66, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x66, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x83, 0x01, 0x0a, 0x19, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x54, 0x72, 0x75, 0x73, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x50, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52,
	0x65, 0x66, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x06, 0xa2,
	0x93, 0x04, 0x02, 0x08, 0x03, 0x22, 0xaf, 0x0b, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x52, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x52,
	0x09, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x53, 0x0a, 0x08, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x4d, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x56,
	0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x38, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x6c, 0x0a, 0x11, 0x6c, 0x65, 0x61, 0x66, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3f, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x65, 0x61,
	0x66, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x10, 0x6c, 0x65, 0x61, 0x66, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x60, 0x0a, 0x0d, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x62, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x74, 0x72, 0x75, 0x73, 0x74, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x54, 0x4c, 0x53, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x12, 0x51, 0x0a, 0x06, 0x65, 0x73, 0x63,
	0x61, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x73, 0x63, 0x61, 0x70, 0x65, 0x48, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x73, 0x52, 0x06, 0x65, 0x73, 0x63, 0x61, 0x70, 0x65, 0x12, 0x57, 0x0a, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4c, 0x6f, 0x67, 0x73, 0x1a, 0x70, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x49, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x6c, 0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x47, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x73, 0x0a, 0x0e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x4b, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x80, 0x01, 0x0a, 0x15, 0x4c,
	0x65, 0x61, 0x66, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x51, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x4c, 0x65, 0x61, 0x66, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x78, 0x0a,
	0x11, 0x54, 0x72, 0x75, 0x73, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x4d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x90, 0x02, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0f, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x6d, 0x65, 0x73,
	0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x6d, 0x65, 0x73, 0x68, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x4d, 0xaa, 0x02, 0x1d, 0x48, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x4d,
	0x65, 0x73, 0x68, 0x2e, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1d, 0x48, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x4d,
	0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x29, 0x48, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x4d,
	0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x4d, 0x65, 0x73,
	0x68, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pbmesh_v2beta1_proxy_state_proto_rawDescOnce sync.Once
	file_pbmesh_v2beta1_proxy_state_proto_rawDescData = file_pbmesh_v2beta1_proxy_state_proto_rawDesc
)

func file_pbmesh_v2beta1_proxy_state_proto_rawDescGZIP() []byte {
	file_pbmesh_v2beta1_proxy_state_proto_rawDescOnce.Do(func() {
		file_pbmesh_v2beta1_proxy_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbmesh_v2beta1_proxy_state_proto_rawDescData)
	})
	return file_pbmesh_v2beta1_proxy_state_proto_rawDescData
}

var file_pbmesh_v2beta1_proxy_state_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pbmesh_v2beta1_proxy_state_proto_goTypes = []interface{}{
	(*ProxyStateTemplate)(nil),              // 0: hashicorp.consul.mesh.v2beta1.ProxyStateTemplate
	(*ProxyState)(nil),                      // 1: hashicorp.consul.mesh.v2beta1.ProxyState
	nil,                                     // 2: hashicorp.consul.mesh.v2beta1.ProxyStateTemplate.RequiredEndpointsEntry
	nil,                                     // 3: hashicorp.consul.mesh.v2beta1.ProxyStateTemplate.RequiredLeafCertificatesEntry
	nil,                                     // 4: hashicorp.consul.mesh.v2beta1.ProxyStateTemplate.RequiredTrustBundlesEntry
	nil,                                     // 5: hashicorp.consul.mesh.v2beta1.ProxyState.ClustersEntry
	nil,                                     // 6: hashicorp.consul.mesh.v2beta1.ProxyState.RoutesEntry
	nil,                                     // 7: hashicorp.consul.mesh.v2beta1.ProxyState.EndpointsEntry
	nil,                                     // 8: hashicorp.consul.mesh.v2beta1.ProxyState.LeafCertificatesEntry
	nil,                                     // 9: hashicorp.consul.mesh.v2beta1.ProxyState.TrustBundlesEntry
	(*pbresource.Reference)(nil),            // 10: hashicorp.consul.resource.Reference
	(*pbproxystate.Listener)(nil),           // 11: hashicorp.consul.mesh.v2beta1.pbproxystate.Listener
	(*pbproxystate.TLS)(nil),                // 12: hashicorp.consul.mesh.v2beta1.pbproxystate.TLS
	(*pbproxystate.EscapeHatches)(nil),      // 13: hashicorp.consul.mesh.v2beta1.pbproxystate.EscapeHatches
	(*pbproxystate.AccessLogs)(nil),         // 14: hashicorp.consul.mesh.v2beta1.pbproxystate.AccessLogs
	(*pbproxystate.EndpointRef)(nil),        // 15: hashicorp.consul.mesh.v2beta1.pbproxystate.EndpointRef
	(*pbproxystate.LeafCertificateRef)(nil), // 16: hashicorp.consul.mesh.v2beta1.pbproxystate.LeafCertificateRef
	(*pbproxystate.TrustBundleRef)(nil),     // 17: hashicorp.consul.mesh.v2beta1.pbproxystate.TrustBundleRef
	(*pbproxystate.Cluster)(nil),            // 18: hashicorp.consul.mesh.v2beta1.pbproxystate.Cluster
	(*pbproxystate.Route)(nil),              // 19: hashicorp.consul.mesh.v2beta1.pbproxystate.Route
	(*pbproxystate.Endpoints)(nil),          // 20: hashicorp.consul.mesh.v2beta1.pbproxystate.Endpoints
	(*pbproxystate.LeafCertificate)(nil),    // 21: hashicorp.consul.mesh.v2beta1.pbproxystate.LeafCertificate
	(*pbproxystate.TrustBundle)(nil),        // 22: hashicorp.consul.mesh.v2beta1.pbproxystate.TrustBundle
}
var file_pbmesh_v2beta1_proxy_state_proto_depIdxs = []int32{
	1,  // 0: hashicorp.consul.mesh.v2beta1.ProxyStateTemplate.proxy_state:type_name -> hashicorp.consul.mesh.v2beta1.ProxyState
	2,  // 1: hashicorp.consul.mesh.v2beta1.ProxyStateTemplate.required_endpoints:type_name -> hashicorp.consul.mesh.v2beta1.ProxyStateTemplate.RequiredEndpointsEntry
	3,  // 2: hashicorp.consul.mesh.v2beta1.ProxyStateTemplate.required_leaf_certificates:type_name -> hashicorp.consul.mesh.v2beta1.ProxyStateTemplate.RequiredLeafCertificatesEntry
	4,  // 3: hashicorp.consul.mesh.v2beta1.ProxyStateTemplate.required_trust_bundles:type_name -> hashicorp.consul.mesh.v2beta1.ProxyStateTemplate.RequiredTrustBundlesEntry
	10, // 4: hashicorp.consul.mesh.v2beta1.ProxyState.identity:type_name -> hashicorp.consul.resource.Reference
	11, // 5: hashicorp.consul.mesh.v2beta1.ProxyState.listeners:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.Listener
	5,  // 6: hashicorp.consul.mesh.v2beta1.ProxyState.clusters:type_name -> hashicorp.consul.mesh.v2beta1.ProxyState.ClustersEntry
	6,  // 7: hashicorp.consul.mesh.v2beta1.ProxyState.routes:type_name -> hashicorp.consul.mesh.v2beta1.ProxyState.RoutesEntry
	7,  // 8: hashicorp.consul.mesh.v2beta1.ProxyState.endpoints:type_name -> hashicorp.consul.mesh.v2beta1.ProxyState.EndpointsEntry
	8,  // 9: hashicorp.consul.mesh.v2beta1.ProxyState.leaf_certificates:type_name -> hashicorp.consul.mesh.v2beta1.ProxyState.LeafCertificatesEntry
	9,  // 10: hashicorp.consul.mesh.v2beta1.ProxyState.trust_bundles:type_name -> hashicorp.consul.mesh.v2beta1.ProxyState.TrustBundlesEntry
	12, // 11: hashicorp.consul.mesh.v2beta1.ProxyState.tls:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.TLS
	13, // 12: hashicorp.consul.mesh.v2beta1.ProxyState.escape:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.EscapeHatches
	14, // 13: hashicorp.consul.mesh.v2beta1.ProxyState.access_logs:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.AccessLogs
	15, // 14: hashicorp.consul.mesh.v2beta1.ProxyStateTemplate.RequiredEndpointsEntry.value:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.EndpointRef
	16, // 15: hashicorp.consul.mesh.v2beta1.ProxyStateTemplate.RequiredLeafCertificatesEntry.value:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.LeafCertificateRef
	17, // 16: hashicorp.consul.mesh.v2beta1.ProxyStateTemplate.RequiredTrustBundlesEntry.value:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.TrustBundleRef
	18, // 17: hashicorp.consul.mesh.v2beta1.ProxyState.ClustersEntry.value:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.Cluster
	19, // 18: hashicorp.consul.mesh.v2beta1.ProxyState.RoutesEntry.value:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.Route
	20, // 19: hashicorp.consul.mesh.v2beta1.ProxyState.EndpointsEntry.value:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.Endpoints
	21, // 20: hashicorp.consul.mesh.v2beta1.ProxyState.LeafCertificatesEntry.value:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.LeafCertificate
	22, // 21: hashicorp.consul.mesh.v2beta1.ProxyState.TrustBundlesEntry.value:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.TrustBundle
	22, // [22:22] is the sub-list for method output_type
	22, // [22:22] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_pbmesh_v2beta1_proxy_state_proto_init() }
func file_pbmesh_v2beta1_proxy_state_proto_init() {
	if File_pbmesh_v2beta1_proxy_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbmesh_v2beta1_proxy_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyStateTemplate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbmesh_v2beta1_proxy_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbmesh_v2beta1_proxy_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbmesh_v2beta1_proxy_state_proto_goTypes,
		DependencyIndexes: file_pbmesh_v2beta1_proxy_state_proto_depIdxs,
		MessageInfos:      file_pbmesh_v2beta1_proxy_state_proto_msgTypes,
	}.Build()
	File_pbmesh_v2beta1_proxy_state_proto = out.File
	file_pbmesh_v2beta1_proxy_state_proto_rawDesc = nil
	file_pbmesh_v2beta1_proxy_state_proto_goTypes = nil
	file_pbmesh_v2beta1_proxy_state_proto_depIdxs = nil
}
