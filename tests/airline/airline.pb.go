// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.3
// source: airline.proto

package test

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Offer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Options        []*Option `protobuf:"bytes,2,rep,name=options,proto3" json:"options,omitempty"`
	Price          float64   `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Currency       string    `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	AvailableSeats int32     `protobuf:"varint,5,opt,name=available_seats,json=availableSeats,proto3" json:"available_seats,omitempty"`
}

func (x *Offer) Reset() {
	*x = Offer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_airline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Offer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Offer) ProtoMessage() {}

func (x *Offer) ProtoReflect() protoreflect.Message {
	mi := &file_airline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Offer.ProtoReflect.Descriptor instead.
func (*Offer) Descriptor() ([]byte, []int) {
	return file_airline_proto_rawDescGZIP(), []int{0}
}

func (x *Offer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Offer) GetOptions() []*Option {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *Offer) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Offer) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Offer) GetAvailableSeats() int32 {
	if x != nil {
		return x.AvailableSeats
	}
	return 0
}

type Option struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duration string   `protobuf:"bytes,1,opt,name=duration,proto3" json:"duration,omitempty"`
	Routes   []*Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *Option) Reset() {
	*x = Option{}
	if protoimpl.UnsafeEnabled {
		mi := &file_airline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Option) ProtoMessage() {}

func (x *Option) ProtoReflect() protoreflect.Message {
	mi := &file_airline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Option.ProtoReflect.Descriptor instead.
func (*Option) Descriptor() ([]byte, []int) {
	return file_airline_proto_rawDescGZIP(), []int{1}
}

func (x *Option) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *Option) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Departure     *Airport `protobuf:"bytes,1,opt,name=departure,proto3" json:"departure,omitempty"`
	Arrival       *Airport `protobuf:"bytes,2,opt,name=arrival,proto3" json:"arrival,omitempty"`
	CarrierCode   string   `protobuf:"bytes,3,opt,name=carrier_code,json=carrierCode,proto3" json:"carrier_code,omitempty"`
	Carrier       string   `protobuf:"bytes,4,opt,name=carrier,proto3" json:"carrier,omitempty"`
	OperatedBy    string   `protobuf:"bytes,5,opt,name=operated_by,json=operatedBy,proto3" json:"operated_by,omitempty"`
	Duration      string   `protobuf:"bytes,6,opt,name=duration,proto3" json:"duration,omitempty"`
	NumberOfStops int32    `protobuf:"varint,7,opt,name=number_of_stops,json=numberOfStops,proto3" json:"number_of_stops,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_airline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_airline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_airline_proto_rawDescGZIP(), []int{2}
}

func (x *Route) GetDeparture() *Airport {
	if x != nil {
		return x.Departure
	}
	return nil
}

func (x *Route) GetArrival() *Airport {
	if x != nil {
		return x.Arrival
	}
	return nil
}

func (x *Route) GetCarrierCode() string {
	if x != nil {
		return x.CarrierCode
	}
	return ""
}

func (x *Route) GetCarrier() string {
	if x != nil {
		return x.Carrier
	}
	return ""
}

func (x *Route) GetOperatedBy() string {
	if x != nil {
		return x.OperatedBy
	}
	return ""
}

func (x *Route) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *Route) GetNumberOfStops() int32 {
	if x != nil {
		return x.NumberOfStops
	}
	return 0
}

type Airport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IataCode string `protobuf:"bytes,1,opt,name=iata_code,json=iataCode,proto3" json:"iata_code,omitempty"`
	Terminal string `protobuf:"bytes,2,opt,name=terminal,proto3" json:"terminal,omitempty"`
	Datetime string `protobuf:"bytes,3,opt,name=datetime,proto3" json:"datetime,omitempty"`
}

func (x *Airport) Reset() {
	*x = Airport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_airline_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Airport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Airport) ProtoMessage() {}

func (x *Airport) ProtoReflect() protoreflect.Message {
	mi := &file_airline_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Airport.ProtoReflect.Descriptor instead.
func (*Airport) Descriptor() ([]byte, []int) {
	return file_airline_proto_rawDescGZIP(), []int{3}
}

func (x *Airport) GetIataCode() string {
	if x != nil {
		return x.IataCode
	}
	return ""
}

func (x *Airport) GetTerminal() string {
	if x != nil {
		return x.Terminal
	}
	return ""
}

func (x *Airport) GetDatetime() string {
	if x != nil {
		return x.Datetime
	}
	return ""
}

type ArrayOfOffers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offers []*Offer `protobuf:"bytes,1,rep,name=offers,proto3" json:"offers,omitempty"`
}

func (x *ArrayOfOffers) Reset() {
	*x = ArrayOfOffers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_airline_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrayOfOffers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayOfOffers) ProtoMessage() {}

func (x *ArrayOfOffers) ProtoReflect() protoreflect.Message {
	mi := &file_airline_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayOfOffers.ProtoReflect.Descriptor instead.
func (*ArrayOfOffers) Descriptor() ([]byte, []int) {
	return file_airline_proto_rawDescGZIP(), []int{4}
}

func (x *ArrayOfOffers) GetOffers() []*Offer {
	if x != nil {
		return x.Offers
	}
	return nil
}

type SearchRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//	*SearchRes_Data
	Result isSearchRes_Result `protobuf_oneof:"result"`
}

func (x *SearchRes) Reset() {
	*x = SearchRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_airline_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRes) ProtoMessage() {}

func (x *SearchRes) ProtoReflect() protoreflect.Message {
	mi := &file_airline_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRes.ProtoReflect.Descriptor instead.
func (*SearchRes) Descriptor() ([]byte, []int) {
	return file_airline_proto_rawDescGZIP(), []int{5}
}

func (m *SearchRes) GetResult() isSearchRes_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *SearchRes) GetData() *ArrayOfOffers {
	if x, ok := x.GetResult().(*SearchRes_Data); ok {
		return x.Data
	}
	return nil
}

type isSearchRes_Result interface {
	isSearchRes_Result()
}

type SearchRes_Data struct {
	Data *ArrayOfOffers `protobuf:"bytes,1,opt,name=data,proto3,oneof"`
}

func (*SearchRes_Data) isSearchRes_Result() {}

type SearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude    float64 `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude   float64 `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Adults      int32   `protobuf:"varint,4,opt,name=adults,proto3" json:"adults,omitempty"`
	Children    *int32  `protobuf:"varint,5,opt,name=children,proto3,oneof" json:"children,omitempty"`
	TravelClass string  `protobuf:"bytes,6,opt,name=travel_class,json=travelClass,proto3" json:"travel_class,omitempty"`
	OneWay      bool    `protobuf:"varint,7,opt,name=one_way,json=oneWay,proto3" json:"one_way,omitempty"`
}

func (x *SearchReq) Reset() {
	*x = SearchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_airline_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReq) ProtoMessage() {}

func (x *SearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_airline_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReq.ProtoReflect.Descriptor instead.
func (*SearchReq) Descriptor() ([]byte, []int) {
	return file_airline_proto_rawDescGZIP(), []int{6}
}

func (x *SearchReq) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *SearchReq) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *SearchReq) GetAdults() int32 {
	if x != nil {
		return x.Adults
	}
	return 0
}

func (x *SearchReq) GetChildren() int32 {
	if x != nil && x.Children != nil {
		return *x.Children
	}
	return 0
}

func (x *SearchReq) GetTravelClass() string {
	if x != nil {
		return x.TravelClass
	}
	return ""
}

func (x *SearchReq) GetOneWay() bool {
	if x != nil {
		return x.OneWay
	}
	return false
}

var file_airline_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51234,
		Name:          "amadeus.airline.select",
		Tag:           "bytes,51234,opt,name=select",
		Filename:      "airline.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51235,
		Name:          "amadeus.airline.query",
		Tag:           "bytes,51235,opt,name=query",
		Filename:      "airline.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string select = 51234;
	E_Select = &file_airline_proto_extTypes[0]
	// optional string query = 51235;
	E_Query = &file_airline_proto_extTypes[1]
)

var File_airline_proto protoreflect.FileDescriptor

var file_airline_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x61, 0x6d, 0x61, 0x64, 0x65, 0x75, 0x73, 0x2e, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xfe, 0x01, 0x0a, 0x05, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0x92, 0x82, 0x19, 0x02, 0x69, 0x64,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x6d, 0x61, 0x64, 0x65, 0x75, 0x73, 0x2e,
	0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0f,
	0x92, 0x82, 0x19, 0x0b, 0x69, 0x74, 0x69, 0x6e, 0x65, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x42, 0x0f, 0x92, 0x82, 0x19, 0x0b, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x2e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x2e, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x12, 0x92, 0x82, 0x19, 0x0e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x42, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x61,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x19, 0x92, 0x82, 0x19, 0x15, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x42, 0x6f, 0x6f, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65,
	0x61, 0x74, 0x73, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65,
	0x61, 0x74, 0x73, 0x22, 0x70, 0x0a, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0c, 0x92, 0x82, 0x19, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x6d, 0x61, 0x64, 0x65, 0x75,
	0x73, 0x2e, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x42,
	0x0c, 0x92, 0x82, 0x19, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x06, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0xca, 0x03, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12,
	0x45, 0x0a, 0x09, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x6d, 0x61, 0x64, 0x65, 0x75, 0x73, 0x2e, 0x61, 0x69, 0x72,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x0d, 0x92, 0x82,
	0x19, 0x09, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x6d, 0x61, 0x64, 0x65, 0x75,
	0x73, 0x2e, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72,
	0x74, 0x42, 0x0b, 0x92, 0x82, 0x19, 0x07, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x52, 0x07,
	0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x72, 0x69,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x92,
	0x82, 0x19, 0x0b, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0b,
	0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x64, 0x0a, 0x07, 0x63,
	0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x4a, 0x9a, 0x82,
	0x19, 0x46, 0x57, 0x48, 0x45, 0x52, 0x45, 0x28, 0x46, 0x52, 0x4f, 0x4d, 0x28, 0x60, 0x24, 0x52,
	0x4f, 0x4f, 0x54, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73,
	0x2e, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x73, 0x60, 0x29, 0x2c, 0x20, 0x60, 0x24, 0x4b,
	0x45, 0x59, 0x60, 0x20, 0x2c, 0x20, 0x46, 0x52, 0x4f, 0x4d, 0x28, 0x63, 0x61, 0x72, 0x72, 0x69,
	0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x29, 0x29, 0x52, 0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65,
	0x72, 0x12, 0x3a, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0x92, 0x82, 0x19, 0x15, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x28, 0x0a,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0c, 0x92, 0x82, 0x19, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x11, 0x92, 0x82, 0x19, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x53, 0x74,
	0x6f, 0x70, 0x73, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x53, 0x74, 0x6f,
	0x70, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x07, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x29,
	0x0a, 0x09, 0x69, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0c, 0x92, 0x82, 0x19, 0x08, 0x69, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x08, 0x69, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0x92, 0x82, 0x19,
	0x08, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x08, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0x92, 0x82, 0x19, 0x02, 0x61, 0x74, 0x52, 0x08, 0x64,
	0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x0d, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x4f, 0x66, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x12, 0x38, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x6d, 0x61, 0x64, 0x65,
	0x75, 0x73, 0x2e, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72,
	0x42, 0x08, 0x92, 0x82, 0x19, 0x04, 0x64, 0x61, 0x74, 0x61, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x65,
	0x72, 0x73, 0x22, 0x4b, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x12,
	0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x61, 0x6d, 0x61, 0x64, 0x65, 0x75, 0x73, 0x2e, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x41, 0x72, 0x72, 0x61, 0x79, 0x4f, 0x66, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0xc7, 0x01, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x64, 0x75, 0x6c, 0x74, 0x73, 0x12,
	0x1f, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x6e, 0x65, 0x5f, 0x77, 0x61, 0x79, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6f, 0x6e, 0x65, 0x57, 0x61, 0x79, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x3a, 0x3a, 0x0a, 0x06, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xa2, 0x90, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x88, 0x01, 0x01, 0x3a, 0x38, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa3, 0x90,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x88, 0x01, 0x01, 0x42,
	0x19, 0x5a, 0x17, 0x61, 0x75, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x6d, 0x61, 0x64, 0x65,
	0x75, 0x73, 0x2f, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_airline_proto_rawDescOnce sync.Once
	file_airline_proto_rawDescData = file_airline_proto_rawDesc
)

func file_airline_proto_rawDescGZIP() []byte {
	file_airline_proto_rawDescOnce.Do(func() {
		file_airline_proto_rawDescData = protoimpl.X.CompressGZIP(file_airline_proto_rawDescData)
	})
	return file_airline_proto_rawDescData
}

var file_airline_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_airline_proto_goTypes = []interface{}{
	(*Offer)(nil),                     // 0: amadeus.airline.Offer
	(*Option)(nil),                    // 1: amadeus.airline.Option
	(*Route)(nil),                     // 2: amadeus.airline.Route
	(*Airport)(nil),                   // 3: amadeus.airline.Airport
	(*ArrayOfOffers)(nil),             // 4: amadeus.airline.ArrayOfOffers
	(*SearchRes)(nil),                 // 5: amadeus.airline.SearchRes
	(*SearchReq)(nil),                 // 6: amadeus.airline.SearchReq
	(*descriptorpb.FieldOptions)(nil), // 7: google.protobuf.FieldOptions
}
var file_airline_proto_depIdxs = []int32{
	1, // 0: amadeus.airline.Offer.options:type_name -> amadeus.airline.Option
	2, // 1: amadeus.airline.Option.routes:type_name -> amadeus.airline.Route
	3, // 2: amadeus.airline.Route.departure:type_name -> amadeus.airline.Airport
	3, // 3: amadeus.airline.Route.arrival:type_name -> amadeus.airline.Airport
	0, // 4: amadeus.airline.ArrayOfOffers.offers:type_name -> amadeus.airline.Offer
	4, // 5: amadeus.airline.SearchRes.data:type_name -> amadeus.airline.ArrayOfOffers
	7, // 6: amadeus.airline.select:extendee -> google.protobuf.FieldOptions
	7, // 7: amadeus.airline.query:extendee -> google.protobuf.FieldOptions
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	6, // [6:8] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_airline_proto_init() }
func file_airline_proto_init() {
	if File_airline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_airline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Offer); i {
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
		file_airline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Option); i {
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
		file_airline_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
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
		file_airline_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Airport); i {
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
		file_airline_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrayOfOffers); i {
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
		file_airline_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRes); i {
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
		file_airline_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReq); i {
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
	file_airline_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*SearchRes_Data)(nil),
	}
	file_airline_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_airline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_airline_proto_goTypes,
		DependencyIndexes: file_airline_proto_depIdxs,
		MessageInfos:      file_airline_proto_msgTypes,
		ExtensionInfos:    file_airline_proto_extTypes,
	}.Build()
	File_airline_proto = out.File
	file_airline_proto_rawDesc = nil
	file_airline_proto_goTypes = nil
	file_airline_proto_depIdxs = nil
}
