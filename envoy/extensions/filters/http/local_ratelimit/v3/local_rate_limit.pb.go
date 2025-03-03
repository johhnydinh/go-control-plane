// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: envoy/extensions/filters/http/local_ratelimit/v3/local_rate_limit.proto

package local_ratelimitv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v32 "github.com/envoyproxy/go-control-plane/envoy/extensions/common/ratelimit/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// [#next-free-field: 12]
type LocalRateLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The human readable prefix to use when emitting stats.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// This field allows for a custom HTTP response status code to the downstream client when
	// the request has been rate limited.
	// Defaults to 429 (TooManyRequests).
	//
	// .. note::
	//   If this is set to < 400, 429 will be used instead.
	Status *v3.HttpStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// The token bucket configuration to use for rate limiting requests that are processed by this
	// filter. Each request processed by the filter consumes a single token. If the token is available,
	// the request will be allowed. If no tokens are available, the request will receive the configured
	// rate limit status.
	//
	// .. note::
	//   It's fine for the token bucket to be unset for the global configuration since the rate limit
	//   can be applied at a the virtual host or route level. Thus, the token bucket must be set
	//   for the per route configuration otherwise the config will be rejected.
	//
	// .. note::
	//   When using per route configuration, the bucket becomes unique to that route.
	//
	// .. note::
	//   In the current implementation the token bucket's :ref:`fill_interval
	//   <envoy_v3_api_field_type.v3.TokenBucket.fill_interval>` must be >= 50ms to avoid too aggressive
	//   refills.
	TokenBucket *v3.TokenBucket `protobuf:"bytes,3,opt,name=token_bucket,json=tokenBucket,proto3" json:"token_bucket,omitempty"`
	// If set, this will enable -- but not necessarily enforce -- the rate limit for the given
	// fraction of requests.
	// Defaults to 0% of requests for safety.
	FilterEnabled *v31.RuntimeFractionalPercent `protobuf:"bytes,4,opt,name=filter_enabled,json=filterEnabled,proto3" json:"filter_enabled,omitempty"`
	// If set, this will enforce the rate limit decisions for the given fraction of requests.
	//
	// Note: this only applies to the fraction of enabled requests.
	//
	// Defaults to 0% of requests for safety.
	FilterEnforced *v31.RuntimeFractionalPercent `protobuf:"bytes,5,opt,name=filter_enforced,json=filterEnforced,proto3" json:"filter_enforced,omitempty"`
	// Specifies a list of HTTP headers that should be added to each request that
	// has been rate limited and is also forwarded upstream. This can only occur when the
	// filter is enabled but not enforced.
	RequestHeadersToAddWhenNotEnforced []*v31.HeaderValueOption `protobuf:"bytes,10,rep,name=request_headers_to_add_when_not_enforced,json=requestHeadersToAddWhenNotEnforced,proto3" json:"request_headers_to_add_when_not_enforced,omitempty"`
	// Specifies a list of HTTP headers that should be added to each response for requests that
	// have been rate limited. This occurs when the filter is either enabled or fully enforced.
	ResponseHeadersToAdd []*v31.HeaderValueOption `protobuf:"bytes,6,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	// The rate limit descriptor list to use in the local rate limit to override
	// on. The rate limit descriptor is selected by the first full match from the
	// request descriptors.
	//
	// Example on how to use ::ref:`this <config_http_filters_local_rate_limit_descriptors>`
	//
	// .. note::
	//
	//   In the current implementation the descriptor's token bucket :ref:`fill_interval
	//   <envoy_v3_api_field_type.v3.TokenBucket.fill_interval>` must be a multiple
	//   global :ref:`token bucket's<envoy_v3_api_field_extensions.filters.http.local_ratelimit.v3.LocalRateLimit.token_bucket>` fill interval.
	//
	//   The descriptors must match verbatim for rate limiting to apply. There is no partial
	//   match by a subset of descriptor entries in the current implementation.
	Descriptors []*v32.LocalRateLimitDescriptor `protobuf:"bytes,8,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	// Specifies the rate limit configurations to be applied with the same
	// stage number. If not set, the default stage number is 0.
	//
	// .. note::
	//
	//  The filter supports a range of 0 - 10 inclusively for stage numbers.
	Stage uint32 `protobuf:"varint,9,opt,name=stage,proto3" json:"stage,omitempty"`
	// Specifies the scope of the rate limiter's token bucket.
	// If set to false, the token bucket is shared across all worker threads,
	// thus the rate limits are applied per Envoy process.
	// If set to true, a token bucket is allocated for each connection.
	// Thus the rate limits are applied per connection thereby allowing
	// one to rate limit requests on a per connection basis.
	// If unspecified, the default value is false.
	LocalRateLimitPerDownstreamConnection bool `protobuf:"varint,11,opt,name=local_rate_limit_per_downstream_connection,json=localRateLimitPerDownstreamConnection,proto3" json:"local_rate_limit_per_downstream_connection,omitempty"`
}

func (x *LocalRateLimit) Reset() {
	*x = LocalRateLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalRateLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalRateLimit) ProtoMessage() {}

func (x *LocalRateLimit) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalRateLimit.ProtoReflect.Descriptor instead.
func (*LocalRateLimit) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_rawDescGZIP(), []int{0}
}

func (x *LocalRateLimit) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *LocalRateLimit) GetStatus() *v3.HttpStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *LocalRateLimit) GetTokenBucket() *v3.TokenBucket {
	if x != nil {
		return x.TokenBucket
	}
	return nil
}

func (x *LocalRateLimit) GetFilterEnabled() *v31.RuntimeFractionalPercent {
	if x != nil {
		return x.FilterEnabled
	}
	return nil
}

func (x *LocalRateLimit) GetFilterEnforced() *v31.RuntimeFractionalPercent {
	if x != nil {
		return x.FilterEnforced
	}
	return nil
}

func (x *LocalRateLimit) GetRequestHeadersToAddWhenNotEnforced() []*v31.HeaderValueOption {
	if x != nil {
		return x.RequestHeadersToAddWhenNotEnforced
	}
	return nil
}

func (x *LocalRateLimit) GetResponseHeadersToAdd() []*v31.HeaderValueOption {
	if x != nil {
		return x.ResponseHeadersToAdd
	}
	return nil
}

func (x *LocalRateLimit) GetDescriptors() []*v32.LocalRateLimitDescriptor {
	if x != nil {
		return x.Descriptors
	}
	return nil
}

func (x *LocalRateLimit) GetStage() uint32 {
	if x != nil {
		return x.Stage
	}
	return 0
}

func (x *LocalRateLimit) GetLocalRateLimitPerDownstreamConnection() bool {
	if x != nil {
		return x.LocalRateLimitPerDownstreamConnection
	}
	return false
}

var File_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_rawDesc = []byte{
	0x0a, 0x47, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f,
	0x76, 0x33, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f,
	0x76, 0x33, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76,
	0x33, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x76, 0x33, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x06,
	0x0a, 0x0e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a,
	0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x76, 0x33, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x55, 0x0a, 0x0e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x57, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x6e,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x33, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x0e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x12, 0x87, 0x01, 0x0a,
	0x28, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x77, 0x68, 0x65, 0x6e, 0x5f, 0x6e, 0x6f, 0x74,
	0x5f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02,
	0x10, 0x0a, 0x52, 0x22, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x57, 0x68, 0x65, 0x6e, 0x4e, 0x6f, 0x74, 0x45, 0x6e,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x12, 0x68, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x64,
	0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x10, 0x0a, 0x52, 0x14, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64,
	0x12, 0x60, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x73, 0x12, 0x1d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x0a, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x12, 0x59, 0x0a, 0x2a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x25, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x65, 0x72, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xca, 0x01, 0x0a,
	0x3e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x42,
	0x13, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x69, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f,
	0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x76, 0x33, 0x3b,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x76,
	0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_rawDescData = file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_rawDesc
)

func file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_rawDescData
}

var file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_goTypes = []interface{}{
	(*LocalRateLimit)(nil),               // 0: envoy.extensions.filters.http.local_ratelimit.v3.LocalRateLimit
	(*v3.HttpStatus)(nil),                // 1: envoy.type.v3.HttpStatus
	(*v3.TokenBucket)(nil),               // 2: envoy.type.v3.TokenBucket
	(*v31.RuntimeFractionalPercent)(nil), // 3: envoy.config.core.v3.RuntimeFractionalPercent
	(*v31.HeaderValueOption)(nil),        // 4: envoy.config.core.v3.HeaderValueOption
	(*v32.LocalRateLimitDescriptor)(nil), // 5: envoy.extensions.common.ratelimit.v3.LocalRateLimitDescriptor
}
var file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.http.local_ratelimit.v3.LocalRateLimit.status:type_name -> envoy.type.v3.HttpStatus
	2, // 1: envoy.extensions.filters.http.local_ratelimit.v3.LocalRateLimit.token_bucket:type_name -> envoy.type.v3.TokenBucket
	3, // 2: envoy.extensions.filters.http.local_ratelimit.v3.LocalRateLimit.filter_enabled:type_name -> envoy.config.core.v3.RuntimeFractionalPercent
	3, // 3: envoy.extensions.filters.http.local_ratelimit.v3.LocalRateLimit.filter_enforced:type_name -> envoy.config.core.v3.RuntimeFractionalPercent
	4, // 4: envoy.extensions.filters.http.local_ratelimit.v3.LocalRateLimit.request_headers_to_add_when_not_enforced:type_name -> envoy.config.core.v3.HeaderValueOption
	4, // 5: envoy.extensions.filters.http.local_ratelimit.v3.LocalRateLimit.response_headers_to_add:type_name -> envoy.config.core.v3.HeaderValueOption
	5, // 6: envoy.extensions.filters.http.local_ratelimit.v3.LocalRateLimit.descriptors:type_name -> envoy.extensions.common.ratelimit.v3.LocalRateLimitDescriptor
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_init() }
func file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_init() {
	if File_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalRateLimit); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto = out.File
	file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_rawDesc = nil
	file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_goTypes = nil
	file_envoy_extensions_filters_http_local_ratelimit_v3_local_rate_limit_proto_depIdxs = nil
}
