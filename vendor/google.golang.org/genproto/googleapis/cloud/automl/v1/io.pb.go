// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1/io.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Input configuration for ImportData Action.
//
// The format of input depends on dataset_metadata the Dataset into which
// the import is happening has. As input source the
// [gcs_source][google.cloud.automl.v1beta1.InputConfig.gcs_source]
// is expected, unless specified otherwise. Additionally any input .CSV file
// by itself must be 100MB or smaller, unless specified otherwise.
// If an "example" file (that is, image, video etc.) with identical content
// (even if it had different GCS_FILE_PATH) is mentioned multiple times, then
// its label, bounding boxes etc. are appended. The same file should be always
// provided with the same ML_USE and GCS_FILE_PATH, if it is not, then
// these values are nondeterministically selected from the given ones.
//
//  Errors:
//  If any of the provided CSV files can't be parsed or if more than certain
//  percent of CSV rows cannot be processed then the operation fails and
//  nothing is imported. Regardless of overall success or failure the per-row
//  failures, up to a certain count cap, is listed in
//  Operation.metadata.partial_failures.
//
type InputConfig struct {
	// The source of the input.
	//
	// Types that are valid to be assigned to Source:
	//	*InputConfig_GcsSource
	Source isInputConfig_Source `protobuf_oneof:"source"`
	// Additional domain-specific parameters describing the semantic of the
	// imported data, any string must be up to 25000
	// characters long.
	Params               map[string]string `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *InputConfig) Reset()         { *m = InputConfig{} }
func (m *InputConfig) String() string { return proto.CompactTextString(m) }
func (*InputConfig) ProtoMessage()    {}
func (*InputConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_929966d18309cd53, []int{0}
}

func (m *InputConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputConfig.Unmarshal(m, b)
}
func (m *InputConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputConfig.Marshal(b, m, deterministic)
}
func (m *InputConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputConfig.Merge(m, src)
}
func (m *InputConfig) XXX_Size() int {
	return xxx_messageInfo_InputConfig.Size(m)
}
func (m *InputConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_InputConfig.DiscardUnknown(m)
}

var xxx_messageInfo_InputConfig proto.InternalMessageInfo

type isInputConfig_Source interface {
	isInputConfig_Source()
}

type InputConfig_GcsSource struct {
	GcsSource *GcsSource `protobuf:"bytes,1,opt,name=gcs_source,json=gcsSource,proto3,oneof"`
}

func (*InputConfig_GcsSource) isInputConfig_Source() {}

func (m *InputConfig) GetSource() isInputConfig_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *InputConfig) GetGcsSource() *GcsSource {
	if x, ok := m.GetSource().(*InputConfig_GcsSource); ok {
		return x.GcsSource
	}
	return nil
}

func (m *InputConfig) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*InputConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*InputConfig_GcsSource)(nil),
	}
}

// *  For Translation:
//         CSV file `translation.csv`, with each line in format:
//         ML_USE,GCS_FILE_PATH
//         GCS_FILE_PATH leads to a .TSV file which describes examples that have
//         given ML_USE, using the following row format per line:
//         TEXT_SNIPPET (in source language) \t TEXT_SNIPPET (in target
//         language)
//
// `export_data_<automl-dataset-display-name>_<timestamp-of-export-call>`
//           where <automl-dataset-display-name> will be made
//           BigQuery-dataset-name compatible (e.g. most special characters will
//           become underscores), and timestamp will be in
//           YYYY_MM_DDThh_mm_ss_sssZ "based on ISO-8601" format. In that
//           dataset a new table called `primary_table` will be created, and
//           filled with precisely the same data as this obtained on import.
type OutputConfig struct {
	// Required. The destination of the output.
	//
	// Types that are valid to be assigned to Destination:
	//	*OutputConfig_GcsDestination
	Destination          isOutputConfig_Destination `protobuf_oneof:"destination"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *OutputConfig) Reset()         { *m = OutputConfig{} }
func (m *OutputConfig) String() string { return proto.CompactTextString(m) }
func (*OutputConfig) ProtoMessage()    {}
func (*OutputConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_929966d18309cd53, []int{1}
}

func (m *OutputConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputConfig.Unmarshal(m, b)
}
func (m *OutputConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputConfig.Marshal(b, m, deterministic)
}
func (m *OutputConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputConfig.Merge(m, src)
}
func (m *OutputConfig) XXX_Size() int {
	return xxx_messageInfo_OutputConfig.Size(m)
}
func (m *OutputConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OutputConfig proto.InternalMessageInfo

type isOutputConfig_Destination interface {
	isOutputConfig_Destination()
}

type OutputConfig_GcsDestination struct {
	GcsDestination *GcsDestination `protobuf:"bytes,1,opt,name=gcs_destination,json=gcsDestination,proto3,oneof"`
}

func (*OutputConfig_GcsDestination) isOutputConfig_Destination() {}

func (m *OutputConfig) GetDestination() isOutputConfig_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *OutputConfig) GetGcsDestination() *GcsDestination {
	if x, ok := m.GetDestination().(*OutputConfig_GcsDestination); ok {
		return x.GcsDestination
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OutputConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OutputConfig_GcsDestination)(nil),
	}
}

// The Google Cloud Storage location for the input content.
type GcsSource struct {
	// Required. Google Cloud Storage URIs to input files, up to 2000 characters
	// long. Accepted forms:
	// * Full object path, e.g. gs://bucket/directory/object.csv
	InputUris            []string `protobuf:"bytes,1,rep,name=input_uris,json=inputUris,proto3" json:"input_uris,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GcsSource) Reset()         { *m = GcsSource{} }
func (m *GcsSource) String() string { return proto.CompactTextString(m) }
func (*GcsSource) ProtoMessage()    {}
func (*GcsSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_929966d18309cd53, []int{2}
}

func (m *GcsSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GcsSource.Unmarshal(m, b)
}
func (m *GcsSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GcsSource.Marshal(b, m, deterministic)
}
func (m *GcsSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GcsSource.Merge(m, src)
}
func (m *GcsSource) XXX_Size() int {
	return xxx_messageInfo_GcsSource.Size(m)
}
func (m *GcsSource) XXX_DiscardUnknown() {
	xxx_messageInfo_GcsSource.DiscardUnknown(m)
}

var xxx_messageInfo_GcsSource proto.InternalMessageInfo

func (m *GcsSource) GetInputUris() []string {
	if m != nil {
		return m.InputUris
	}
	return nil
}

// The Google Cloud Storage location where the output is to be written to.
type GcsDestination struct {
	// Required. Google Cloud Storage URI to output directory, up to 2000
	// characters long.
	// Accepted forms:
	// * Prefix path: gs://bucket/directory
	// The requesting user must have write permission to the bucket.
	// The directory is created if it doesn't exist.
	OutputUriPrefix      string   `protobuf:"bytes,1,opt,name=output_uri_prefix,json=outputUriPrefix,proto3" json:"output_uri_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GcsDestination) Reset()         { *m = GcsDestination{} }
func (m *GcsDestination) String() string { return proto.CompactTextString(m) }
func (*GcsDestination) ProtoMessage()    {}
func (*GcsDestination) Descriptor() ([]byte, []int) {
	return fileDescriptor_929966d18309cd53, []int{3}
}

func (m *GcsDestination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GcsDestination.Unmarshal(m, b)
}
func (m *GcsDestination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GcsDestination.Marshal(b, m, deterministic)
}
func (m *GcsDestination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GcsDestination.Merge(m, src)
}
func (m *GcsDestination) XXX_Size() int {
	return xxx_messageInfo_GcsDestination.Size(m)
}
func (m *GcsDestination) XXX_DiscardUnknown() {
	xxx_messageInfo_GcsDestination.DiscardUnknown(m)
}

var xxx_messageInfo_GcsDestination proto.InternalMessageInfo

func (m *GcsDestination) GetOutputUriPrefix() string {
	if m != nil {
		return m.OutputUriPrefix
	}
	return ""
}

func init() {
	proto.RegisterType((*InputConfig)(nil), "google.cloud.automl.v1.InputConfig")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.automl.v1.InputConfig.ParamsEntry")
	proto.RegisterType((*OutputConfig)(nil), "google.cloud.automl.v1.OutputConfig")
	proto.RegisterType((*GcsSource)(nil), "google.cloud.automl.v1.GcsSource")
	proto.RegisterType((*GcsDestination)(nil), "google.cloud.automl.v1.GcsDestination")
}

func init() { proto.RegisterFile("google/cloud/automl/v1/io.proto", fileDescriptor_929966d18309cd53) }

var fileDescriptor_929966d18309cd53 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x18, 0xc5, 0x97, 0x54, 0x54, 0xe4, 0x0b, 0x6c, 0x10, 0xa1, 0xa9, 0x54, 0x20, 0x4a, 0x2f, 0x50,
	0xb5, 0x0b, 0x5b, 0x19, 0x37, 0x10, 0x76, 0x43, 0x06, 0xea, 0x90, 0x40, 0x94, 0xa0, 0xf5, 0x02,
	0x55, 0xaa, 0x4c, 0x96, 0x59, 0x16, 0xa9, 0xbf, 0xc8, 0x7f, 0x2a, 0xf6, 0x4a, 0x7b, 0x14, 0x5e,
	0x82, 0x7b, 0x9e, 0x02, 0xc5, 0xce, 0x58, 0x87, 0x28, 0x77, 0xf6, 0xf9, 0x9d, 0xf3, 0xd9, 0xc7,
	0x09, 0x3c, 0xe1, 0x88, 0xbc, 0xae, 0x68, 0x59, 0xa3, 0x3d, 0xa3, 0xcc, 0x1a, 0x5c, 0xd5, 0x74,
	0x9d, 0x52, 0x81, 0xa4, 0x51, 0x68, 0x30, 0xd9, 0xf7, 0x06, 0xe2, 0x0c, 0xc4, 0x1b, 0xc8, 0x3a,
	0x1d, 0x3e, 0xea, 0x82, 0xac, 0x11, 0x94, 0x49, 0x89, 0x86, 0x19, 0x81, 0x52, 0xfb, 0xd4, 0xf8,
	0x67, 0x00, 0xf1, 0x3b, 0xd9, 0x58, 0x73, 0x8c, 0xf2, 0x5c, 0xf0, 0x24, 0x07, 0xe0, 0xa5, 0x5e,
	0x6a, 0xb4, 0xaa, 0xac, 0x06, 0xc1, 0x28, 0x98, 0xc4, 0x87, 0x4f, 0xc9, 0xbf, 0x47, 0x93, 0x69,
	0xa9, 0x3f, 0x3b, 0xe3, 0xc9, 0x4e, 0x11, 0xf1, 0xab, 0x4d, 0x32, 0x85, 0x7e, 0xc3, 0x14, 0x5b,
	0xe9, 0x41, 0x38, 0xea, 0x4d, 0xe2, 0x43, 0xba, 0x2d, 0xbf, 0x71, 0x30, 0x99, 0xb9, 0xc4, 0x5b,
	0x69, 0xd4, 0x45, 0xd1, 0xc5, 0x87, 0x2f, 0x21, 0xde, 0x90, 0x93, 0x7b, 0xd0, 0xfb, 0x56, 0x5d,
	0xb8, 0x4b, 0x45, 0x45, 0xbb, 0x4c, 0x1e, 0xc0, 0xad, 0x35, 0xab, 0x6d, 0x35, 0x08, 0x9d, 0xe6,
	0x37, 0x59, 0xf8, 0x22, 0xc8, 0x6f, 0x43, 0xdf, 0x77, 0x18, 0x37, 0x70, 0xe7, 0xa3, 0x35, 0xd7,
	0x0d, 0x3f, 0xc1, 0x5e, 0xdb, 0xf0, 0xac, 0xd2, 0x46, 0x48, 0xf7, 0x16, 0x5d, 0xcd, 0x67, 0xff,
	0xa9, 0xf9, 0xe6, 0xda, 0x7d, 0xb2, 0x53, 0xec, 0xf2, 0x1b, 0x4a, 0x7e, 0x17, 0xe2, 0x8d, 0x71,
	0xe3, 0x03, 0x88, 0xfe, 0xbc, 0x4c, 0xf2, 0x18, 0x40, 0xb4, 0x35, 0x97, 0x56, 0x09, 0x3d, 0x08,
	0x46, 0xbd, 0x49, 0x54, 0x44, 0x4e, 0x39, 0x55, 0x42, 0x8f, 0x8f, 0x60, 0xf7, 0xe6, 0xf8, 0xe4,
	0x00, 0xee, 0xa3, 0xbb, 0x6f, 0x9b, 0x58, 0x36, 0xaa, 0x3a, 0x17, 0xdf, 0xbb, 0xce, 0x7b, 0x1e,
	0x9c, 0x2a, 0x31, 0x73, 0x72, 0x7e, 0x19, 0xc0, 0xb0, 0xc4, 0xd5, 0x96, 0x8b, 0xcf, 0x82, 0x2f,
	0x47, 0x1d, 0xe1, 0x58, 0x33, 0xc9, 0x09, 0x2a, 0x4e, 0x79, 0x25, 0xdd, 0xa7, 0xa7, 0x1e, 0xb1,
	0x46, 0xe8, 0xbf, 0x7f, 0xaa, 0x57, 0x7e, 0x75, 0x19, 0xee, 0x4f, 0x7d, 0xfc, 0xd8, 0x0d, 0x7e,
	0x6d, 0x0d, 0x7e, 0x78, 0x4f, 0xe6, 0xe9, 0x8f, 0x2b, 0xb0, 0x70, 0x60, 0xe1, 0xc1, 0x62, 0x9e,
	0xfe, 0x0a, 0x1f, 0x7a, 0x90, 0x65, 0x8e, 0x64, 0x99, 0x47, 0x59, 0x36, 0x4f, 0xbf, 0xf6, 0xdd,
	0xb1, 0xcf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x09, 0x41, 0xdb, 0xca, 0x02, 0x00, 0x00,
}
