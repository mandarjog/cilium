// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/rds.proto

package v2

import (
	context "context"
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	route "github.com/cilium/proxy/go/envoy/api/v2/route"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// [#comment:next free field: 10]
type RouteConfiguration struct {
	// The name of the route configuration. For example, it might match
	// :ref:`route_config_name
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.Rds.route_config_name>` in
	// :ref:`envoy_api_msg_config.filter.network.http_connection_manager.v2.Rds`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// An array of virtual hosts that make up the route table.
	VirtualHosts []*route.VirtualHost `protobuf:"bytes,2,rep,name=virtual_hosts,json=virtualHosts,proto3" json:"virtual_hosts,omitempty"`
	// An array of virtual hosts will be dynamically loaded via the VHDS API.
	// Both *virtual_hosts* and *vhds* fields will be used when present. *virtual_hosts* can be used
	// for a base routing table or for infrequently changing virtual hosts. *vhds* is used for
	// on-demand discovery of virtual hosts. The contents of these two fields will be merged to
	// generate a routing table for a given RouteConfiguration, with *vhds* derived configuration
	// taking precedence.
	// [#not-implemented-hide:]
	Vhds *Vhds `protobuf:"bytes,9,opt,name=vhds,proto3" json:"vhds,omitempty"`
	// Optionally specifies a list of HTTP headers that the connection manager
	// will consider to be internal only. If they are found on external requests they will be cleaned
	// prior to filter invocation. See :ref:`config_http_conn_man_headers_x-envoy-internal` for more
	// information.
	InternalOnlyHeaders []string `protobuf:"bytes,3,rep,name=internal_only_headers,json=internalOnlyHeaders,proto3" json:"internal_only_headers,omitempty"`
	// Specifies a list of HTTP headers that should be added to each response that
	// the connection manager encodes. Headers specified at this level are applied
	// after headers from any enclosed :ref:`envoy_api_msg_route.VirtualHost` or
	// :ref:`envoy_api_msg_route.RouteAction`. For more information, including details on
	// header value syntax, see the documentation on :ref:`custom request headers
	// <config_http_conn_man_headers_custom_request_headers>`.
	ResponseHeadersToAdd []*core.HeaderValueOption `protobuf:"bytes,4,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each response
	// that the connection manager encodes.
	ResponseHeadersToRemove []string `protobuf:"bytes,5,rep,name=response_headers_to_remove,json=responseHeadersToRemove,proto3" json:"response_headers_to_remove,omitempty"`
	// Specifies a list of HTTP headers that should be added to each request
	// routed by the HTTP connection manager. Headers specified at this level are
	// applied after headers from any enclosed :ref:`envoy_api_msg_route.VirtualHost` or
	// :ref:`envoy_api_msg_route.RouteAction`. For more information, including details on
	// header value syntax, see the documentation on :ref:`custom request headers
	// <config_http_conn_man_headers_custom_request_headers>`.
	RequestHeadersToAdd []*core.HeaderValueOption `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each request
	// routed by the HTTP connection manager.
	RequestHeadersToRemove []string `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	// An optional boolean that specifies whether the clusters that the route
	// table refers to will be validated by the cluster manager. If set to true
	// and a route refers to a non-existent cluster, the route table will not
	// load. If set to false and a route refers to a non-existent cluster, the
	// route table will load and the router filter will return a 404 if the route
	// is selected at runtime. This setting defaults to true if the route table
	// is statically defined via the :ref:`route_config
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.HttpConnectionManager.route_config>`
	// option. This setting default to false if the route table is loaded dynamically via the
	// :ref:`rds
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.HttpConnectionManager.rds>`
	// option. Users may wish to override the default behavior in certain cases (for example when
	// using CDS with a static route table).
	ValidateClusters     *wrappers.BoolValue `protobuf:"bytes,7,opt,name=validate_clusters,json=validateClusters,proto3" json:"validate_clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_78812f46dcff924a, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetVirtualHosts() []*route.VirtualHost {
	if m != nil {
		return m.VirtualHosts
	}
	return nil
}

func (m *RouteConfiguration) GetVhds() *Vhds {
	if m != nil {
		return m.Vhds
	}
	return nil
}

func (m *RouteConfiguration) GetInternalOnlyHeaders() []string {
	if m != nil {
		return m.InternalOnlyHeaders
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToAdd() []*core.HeaderValueOption {
	if m != nil {
		return m.ResponseHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToRemove() []string {
	if m != nil {
		return m.ResponseHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToAdd() []*core.HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToRemove() []string {
	if m != nil {
		return m.RequestHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetValidateClusters() *wrappers.BoolValue {
	if m != nil {
		return m.ValidateClusters
	}
	return nil
}

// [#not-implemented-hide:]
type Vhds struct {
	// Configuration source specifier for VHDS.
	ConfigSource         *core.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Vhds) Reset()         { *m = Vhds{} }
func (m *Vhds) String() string { return proto.CompactTextString(m) }
func (*Vhds) ProtoMessage()    {}
func (*Vhds) Descriptor() ([]byte, []int) {
	return fileDescriptor_78812f46dcff924a, []int{1}
}

func (m *Vhds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vhds.Unmarshal(m, b)
}
func (m *Vhds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vhds.Marshal(b, m, deterministic)
}
func (m *Vhds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vhds.Merge(m, src)
}
func (m *Vhds) XXX_Size() int {
	return xxx_messageInfo_Vhds.Size(m)
}
func (m *Vhds) XXX_DiscardUnknown() {
	xxx_messageInfo_Vhds.DiscardUnknown(m)
}

var xxx_messageInfo_Vhds proto.InternalMessageInfo

func (m *Vhds) GetConfigSource() *core.ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.api.v2.RouteConfiguration")
	proto.RegisterType((*Vhds)(nil), "envoy.api.v2.Vhds")
}

func init() { proto.RegisterFile("envoy/api/v2/rds.proto", fileDescriptor_78812f46dcff924a) }

var fileDescriptor_78812f46dcff924a = []byte{
	// 661 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x3f, 0x6f, 0xd3, 0x4e,
	0x18, 0xee, 0x25, 0xe9, 0xbf, 0x4b, 0xfa, 0x53, 0x7b, 0xfd, 0x97, 0xe6, 0x87, 0x9a, 0x28, 0x14,
	0x14, 0x31, 0xd8, 0xc8, 0x4c, 0x94, 0x89, 0xb4, 0x82, 0x8a, 0xa5, 0x95, 0x5b, 0x32, 0x21, 0xac,
	0xab, 0x7d, 0x8d, 0x2d, 0xb9, 0x7e, 0xcd, 0xdd, 0xd9, 0x90, 0x15, 0x16, 0xc4, 0x08, 0x5f, 0x82,
	0xcf, 0xc0, 0x02, 0x23, 0x3b, 0x3b, 0x03, 0x30, 0xc0, 0xb7, 0x40, 0x3e, 0xfb, 0x20, 0x4e, 0x0a,
	0x03, 0xea, 0x12, 0x5d, 0xee, 0x79, 0x9e, 0xf7, 0x7d, 0xde, 0xbb, 0xe7, 0x8c, 0x37, 0x58, 0x94,
	0xc2, 0xc8, 0xa4, 0x71, 0x60, 0xa6, 0x96, 0xc9, 0x3d, 0x61, 0xc4, 0x1c, 0x24, 0x90, 0x86, 0xda,
	0x37, 0x68, 0x1c, 0x18, 0xa9, 0xd5, 0xba, 0x52, 0x62, 0xb9, 0xc0, 0x99, 0x79, 0x4a, 0x05, 0xcb,
	0xb9, 0xad, 0x6b, 0xd3, 0xa8, 0x0b, 0xd1, 0x59, 0x30, 0x74, 0x04, 0x24, 0xdc, 0xd5, 0xb4, 0x72,
	0x11, 0x2f, 0x10, 0x2e, 0xa4, 0x8c, 0x8f, 0x0a, 0x74, 0xbb, 0x6c, 0x04, 0x12, 0xc9, 0xf2, 0x5f,
	0xad, 0x1e, 0x02, 0x0c, 0x43, 0xa6, 0x08, 0x34, 0x8a, 0x40, 0x52, 0x19, 0x40, 0x24, 0xb4, 0xba,
	0x40, 0xd5, 0xbf, 0xd3, 0xe4, 0xcc, 0x7c, 0xca, 0x69, 0x1c, 0x33, 0xae, 0xf1, 0xcd, 0x94, 0x86,
	0x81, 0x47, 0x25, 0x33, 0xf5, 0xa2, 0x00, 0xd6, 0x86, 0x30, 0x04, 0xb5, 0x34, 0xb3, 0x55, 0xbe,
	0xdb, 0xfd, 0x5a, 0xc3, 0xc4, 0xce, 0x9a, 0xef, 0xa9, 0x39, 0x12, 0xae, 0x9a, 0x11, 0x82, 0x6b,
	0x11, 0x3d, 0x67, 0x4d, 0xd4, 0x41, 0xbd, 0x45, 0x5b, 0xad, 0xc9, 0x03, 0xbc, 0x94, 0x06, 0x5c,
	0x26, 0x34, 0x74, 0x7c, 0x10, 0x52, 0x34, 0x2b, 0x9d, 0x6a, 0xaf, 0x6e, 0xb5, 0x8d, 0xf1, 0x03,
	0x34, 0xf2, 0x49, 0x06, 0x39, 0xf1, 0x00, 0x84, 0xec, 0xd7, 0x3e, 0x7e, 0x6e, 0xcf, 0xd8, 0x8d,
	0xf4, 0xf7, 0x96, 0x20, 0xd7, 0x71, 0x2d, 0xf5, 0x3d, 0xd1, 0x5c, 0xec, 0xa0, 0x5e, 0xdd, 0x22,
	0xe5, 0x12, 0x03, 0xdf, 0x13, 0xb6, 0xc2, 0x89, 0x85, 0xd7, 0x83, 0x48, 0x32, 0x1e, 0xd1, 0xd0,
	0x81, 0x28, 0x1c, 0x39, 0x3e, 0xa3, 0x1e, 0xe3, 0xa2, 0x59, 0xed, 0x54, 0x7b, 0x8b, 0xf6, 0xaa,
	0x06, 0x0f, 0xa3, 0x70, 0x74, 0x90, 0x43, 0xc4, 0xc7, 0x9b, 0x9c, 0x89, 0x18, 0x22, 0xc1, 0x34,
	0xdd, 0x91, 0xe0, 0x50, 0xcf, 0x6b, 0xd6, 0x94, 0xe3, 0x9d, 0x72, 0xbb, 0xec, 0x1a, 0x8d, 0x5c,
	0x3c, 0xa0, 0x61, 0xc2, 0x0e, 0xe3, 0xec, 0x08, 0xfa, 0xf5, 0x77, 0x3f, 0x3e, 0x54, 0xe7, 0x5e,
	0xa3, 0xea, 0xf2, 0xf7, 0x79, 0x7b, 0x4d, 0x57, 0x2c, 0x9a, 0x9c, 0xc0, 0x5d, 0xcf, 0x23, 0x77,
	0x70, 0xeb, 0xa2, 0x4e, 0x9c, 0x9d, 0x43, 0xca, 0x9a, 0xb3, 0xca, 0xe2, 0xe6, 0x94, 0xd2, 0x56,
	0x30, 0x39, 0xc3, 0x1b, 0x9c, 0x3d, 0x49, 0x98, 0x90, 0x93, 0x2e, 0xe7, 0xfe, 0xd5, 0xe5, 0x6a,
	0x51, 0xb0, 0x64, 0xf2, 0x36, 0xde, 0xba, 0xa0, 0x4f, 0xe1, 0x71, 0x41, 0x79, 0xdc, 0x98, 0xd4,
	0x15, 0x16, 0xef, 0xe3, 0x15, 0x1d, 0x22, 0xc7, 0x0d, 0x13, 0x21, 0xb3, 0x93, 0x9f, 0x57, 0x57,
	0xd6, 0x32, 0xf2, 0x1c, 0x1a, 0x3a, 0x87, 0x46, 0x1f, 0x20, 0x54, 0xce, 0xec, 0x65, 0x2d, 0xda,
	0x2b, 0x34, 0xdd, 0x47, 0xb8, 0x96, 0x5d, 0x2a, 0x39, 0xc1, 0x4b, 0xa5, 0xf7, 0xa2, 0xf2, 0x35,
	0x15, 0x21, 0x35, 0x6a, 0x9e, 0xc7, 0x63, 0x45, 0xeb, 0xff, 0x97, 0x45, 0x28, 0x9b, 0x74, 0xf6,
	0x15, 0xaa, 0x2c, 0x23, 0xbb, 0xe1, 0x8e, 0xa1, 0xd6, 0xfb, 0x0a, 0x5e, 0x57, 0x19, 0xde, 0xd7,
	0x2f, 0xed, 0x98, 0xf1, 0x34, 0x70, 0x19, 0x79, 0x88, 0x1b, 0xc7, 0x92, 0x33, 0x7a, 0xae, 0x60,
	0x41, 0xb6, 0xcb, 0x8d, 0x7e, 0xf1, 0xed, 0x7c, 0xfe, 0x56, 0xfb, 0x8f, 0x78, 0x7e, 0x87, 0xdd,
	0x99, 0x1e, 0xba, 0x89, 0xc8, 0x63, 0x5c, 0xdf, 0x67, 0xa1, 0xa4, 0x45, 0xd5, 0xab, 0x13, 0xaa,
	0x0c, 0x9a, 0x2a, 0xbd, 0xf3, 0x77, 0x52, 0xa9, 0x3e, 0xe0, 0xfa, 0x3d, 0x26, 0x5d, 0xff, 0xb2,
	0x5c, 0xb7, 0x9f, 0x7f, 0xfa, 0xf6, 0xa6, 0xb2, 0xd5, 0x5d, 0x2b, 0x7d, 0x8e, 0x76, 0xd5, 0x53,
	0x15, 0xbb, 0xe8, 0x86, 0xf5, 0x02, 0xe1, 0xff, 0xc7, 0x9e, 0xec, 0xd4, 0x39, 0x7a, 0x78, 0x45,
	0x19, 0x1e, 0x8c, 0xbf, 0xe1, 0xcb, 0x1e, 0xbb, 0x6f, 0xbe, 0xfd, 0xb2, 0x8d, 0x70, 0x2b, 0x80,
	0x5c, 0x13, 0x73, 0x78, 0x36, 0x2a, 0xc9, 0xfb, 0x0b, 0xb6, 0x27, 0x8e, 0xb2, 0xc0, 0x1d, 0xa1,
	0x97, 0x08, 0x9d, 0xce, 0xa9, 0xf0, 0xdd, 0xfa, 0x19, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x62, 0xd0,
	0xe3, 0xdb, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouteDiscoveryServiceClient is the client API for RouteDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouteDiscoveryServiceClient interface {
	StreamRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_StreamRoutesClient, error)
	DeltaRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_DeltaRoutesClient, error)
	FetchRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type routeDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRouteDiscoveryServiceClient(cc *grpc.ClientConn) RouteDiscoveryServiceClient {
	return &routeDiscoveryServiceClient{cc}
}

func (c *routeDiscoveryServiceClient) StreamRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_StreamRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.RouteDiscoveryService/StreamRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeDiscoveryServiceStreamRoutesClient{stream}
	return x, nil
}

type RouteDiscoveryService_StreamRoutesClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type routeDiscoveryServiceStreamRoutesClient struct {
	grpc.ClientStream
}

func (x *routeDiscoveryServiceStreamRoutesClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeDiscoveryServiceStreamRoutesClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeDiscoveryServiceClient) DeltaRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_DeltaRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v2.RouteDiscoveryService/DeltaRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeDiscoveryServiceDeltaRoutesClient{stream}
	return x, nil
}

type RouteDiscoveryService_DeltaRoutesClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type routeDiscoveryServiceDeltaRoutesClient struct {
	grpc.ClientStream
}

func (x *routeDiscoveryServiceDeltaRoutesClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeDiscoveryServiceDeltaRoutesClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeDiscoveryServiceClient) FetchRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.RouteDiscoveryService/FetchRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteDiscoveryServiceServer is the server API for RouteDiscoveryService service.
type RouteDiscoveryServiceServer interface {
	StreamRoutes(RouteDiscoveryService_StreamRoutesServer) error
	DeltaRoutes(RouteDiscoveryService_DeltaRoutesServer) error
	FetchRoutes(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

func RegisterRouteDiscoveryServiceServer(s *grpc.Server, srv RouteDiscoveryServiceServer) {
	s.RegisterService(&_RouteDiscoveryService_serviceDesc, srv)
}

func _RouteDiscoveryService_StreamRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteDiscoveryServiceServer).StreamRoutes(&routeDiscoveryServiceStreamRoutesServer{stream})
}

type RouteDiscoveryService_StreamRoutesServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type routeDiscoveryServiceStreamRoutesServer struct {
	grpc.ServerStream
}

func (x *routeDiscoveryServiceStreamRoutesServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeDiscoveryServiceStreamRoutesServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteDiscoveryService_DeltaRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteDiscoveryServiceServer).DeltaRoutes(&routeDiscoveryServiceDeltaRoutesServer{stream})
}

type RouteDiscoveryService_DeltaRoutesServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type routeDiscoveryServiceDeltaRoutesServer struct {
	grpc.ServerStream
}

func (x *routeDiscoveryServiceDeltaRoutesServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeDiscoveryServiceDeltaRoutesServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteDiscoveryService_FetchRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteDiscoveryServiceServer).FetchRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.RouteDiscoveryService/FetchRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteDiscoveryServiceServer).FetchRoutes(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RouteDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.RouteDiscoveryService",
	HandlerType: (*RouteDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRoutes",
			Handler:    _RouteDiscoveryService_FetchRoutes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRoutes",
			Handler:       _RouteDiscoveryService_StreamRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaRoutes",
			Handler:       _RouteDiscoveryService_DeltaRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/rds.proto",
}

// VirtualHostDiscoveryServiceClient is the client API for VirtualHostDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualHostDiscoveryServiceClient interface {
	DeltaVirtualHosts(ctx context.Context, opts ...grpc.CallOption) (VirtualHostDiscoveryService_DeltaVirtualHostsClient, error)
}

type virtualHostDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewVirtualHostDiscoveryServiceClient(cc *grpc.ClientConn) VirtualHostDiscoveryServiceClient {
	return &virtualHostDiscoveryServiceClient{cc}
}

func (c *virtualHostDiscoveryServiceClient) DeltaVirtualHosts(ctx context.Context, opts ...grpc.CallOption) (VirtualHostDiscoveryService_DeltaVirtualHostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_VirtualHostDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.VirtualHostDiscoveryService/DeltaVirtualHosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &virtualHostDiscoveryServiceDeltaVirtualHostsClient{stream}
	return x, nil
}

type VirtualHostDiscoveryService_DeltaVirtualHostsClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type virtualHostDiscoveryServiceDeltaVirtualHostsClient struct {
	grpc.ClientStream
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VirtualHostDiscoveryServiceServer is the server API for VirtualHostDiscoveryService service.
type VirtualHostDiscoveryServiceServer interface {
	DeltaVirtualHosts(VirtualHostDiscoveryService_DeltaVirtualHostsServer) error
}

func RegisterVirtualHostDiscoveryServiceServer(s *grpc.Server, srv VirtualHostDiscoveryServiceServer) {
	s.RegisterService(&_VirtualHostDiscoveryService_serviceDesc, srv)
}

func _VirtualHostDiscoveryService_DeltaVirtualHosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VirtualHostDiscoveryServiceServer).DeltaVirtualHosts(&virtualHostDiscoveryServiceDeltaVirtualHostsServer{stream})
}

type VirtualHostDiscoveryService_DeltaVirtualHostsServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type virtualHostDiscoveryServiceDeltaVirtualHostsServer struct {
	grpc.ServerStream
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _VirtualHostDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.VirtualHostDiscoveryService",
	HandlerType: (*VirtualHostDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaVirtualHosts",
			Handler:       _VirtualHostDiscoveryService_DeltaVirtualHosts_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/rds.proto",
}
