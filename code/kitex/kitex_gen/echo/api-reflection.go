// Code generated by thriftgo (0.3.6). DO NOT EDIT.

package echo

import (
	"reflect"

	"github.com/cloudwego/thriftgo/thrift_reflection"
)

// IDL Name: api
// IDL Path: api.thrift

var file_api_thrift_go_types = []interface{}{
	(*EchoRequest)(nil),                     // Struct 0: echo.EchoRequest
	(*EchoResponse)(nil),                    // Struct 1: echo.EchoResponse
	(*EchoMultiBoolResponse)(nil),           // Struct 2: echo.EchoMultiBoolResponse
	(*EchoMultiByteResponse)(nil),           // Struct 3: echo.EchoMultiByteResponse
	(*EchoMultiInt16Response)(nil),          // Struct 4: echo.EchoMultiInt16Response
	(*EchoMultiInt32Response)(nil),          // Struct 5: echo.EchoMultiInt32Response
	(*EchoMultiInt64Response)(nil),          // Struct 6: echo.EchoMultiInt64Response
	(*EchoMultiFloatResponse)(nil),          // Struct 7: echo.EchoMultiFloatResponse
	(*EchoMultiDoubleResponse)(nil),         // Struct 8: echo.EchoMultiDoubleResponse
	(*EchoMultiStringResponse)(nil),         // Struct 9: echo.EchoMultiStringResponse
	(*EchoOptionalStructRequest)(nil),       // Struct 10: echo.EchoOptionalStructRequest
	(*EchoOptionalStructResponse)(nil),      // Struct 11: echo.EchoOptionalStructResponse
	(*EchoOptionalMultiBoolRequest)(nil),    // Struct 12: echo.EchoOptionalMultiBoolRequest
	(*EchoOptionalMultiInt32Request)(nil),   // Struct 13: echo.EchoOptionalMultiInt32Request
	(*EchoOptionalMultiStringRequest)(nil),  // Struct 14: echo.EchoOptionalMultiStringRequest
	(*EchoOptionalMultiBoolResponse)(nil),   // Struct 15: echo.EchoOptionalMultiBoolResponse
	(*EchoOptionalMultiInt32Response)(nil),  // Struct 16: echo.EchoOptionalMultiInt32Response
	(*EchoOptionalMultiStringResponse)(nil), // Struct 17: echo.EchoOptionalMultiStringResponse
	(*EchoGenericRequest)(nil),              // Struct 18: echo.EchoGenericRequest
	(*EchoGenericResponse)(nil),             // Struct 19: echo.EchoGenericResponse
	(*EchoGenericEmbedded)(nil),             // Struct 20: echo.EchoGenericEmbedded
	(*EchoCustomizedException)(nil),         // Exception 0: echo.EchoCustomizedException
}
var file_api_thrift *thrift_reflection.FileDescriptor
var file_idl_api_rawDesc = []byte{
	0x1f, 0x8b, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xec, 0x5d, 0x4b, 0x77, 0xdb, 0xb8,
	0x15, 0xc6, 0xd8, 0x72, 0x2c, 0x5b, 0x7e, 0xc4, 0x9e, 0xc4, 0x69, 0xa6, 0x9d, 0x8e, 0x7a, 0xba,
	0xe8, 0x46, 0x95, 0x6a, 0x59, 0x55, 0x1f, 0xbb, 0x26, 0xce, 0xb4, 0x99, 0x53, 0x77, 0x7a, 0x66,
	0xa6, 0xab, 0x9e, 0x2e, 0x20, 0x9, 0x91, 0x38, 0x23, 0x91, 0xa, 0x9, 0xb9, 0x76, 0x7f, 0x7d,
	0xf, 0xc0, 0xb7, 0x24, 0xe2, 0x45, 0x80, 0xa4, 0x15, 0x6e, 0xa2, 0x24, 0x86, 0x71, 0xef, 0xfd,
	0xee, 0x47, 0xf0, 0x12, 0xf7, 0x3, 0xd5, 0x2, 0x9f, 0x1, 0x0, 0x8e, 0xe1, 0xd2, 0xea, 0xe2,
	0x99, 0x6b, 0x7d, 0xc0, 0xa7, 0x60, 0xaf, 0xd5, 0x2, 0x0, 0xec, 0x1, 0x0, 0x1a, 0x3f, 0xc2,
	0x7b, 0x8, 0x0, 0x68, 0x91, 0xcf, 0xe0, 0xe7, 0x64, 0x34, 0x7a, 0xc0, 0xc8, 0xf6, 0x2c, 0xc7,
	0xf6, 0x0, 0x0, 0x17, 0xf1, 0xbf, 0xa2, 0x29, 0xf6, 0xe9, 0x14, 0x64, 0xe6, 0xbd, 0xa9, 0x43,
	0x26, 0x42, 0xe3, 0x99, 0x73, 0xe, 0x1a, 0x27, 0xe4, 0x7f, 0x37, 0x4c, 0xb6, 0xa8, 0xb1, 0xd6,
	0xf, 0xc8, 0xc3, 0xdf, 0x23, 0xf7, 0xde, 0x1a, 0xa3, 0x73, 0xb0, 0x4f, 0x86, 0x2e, 0xb2, 0x86,
	0xbe, 0x1b, 0xcf, 0x9c, 0xef, 0x10, 0x7e, 0xf3, 0x88, 0xd1, 0x9, 0xd8, 0xcf, 0x18, 0xd5, 0x18,
	0x3d, 0x62, 0x4, 0x2, 0xab, 0xe0, 0x14, 0x1c, 0xb4, 0xce, 0xc9, 0x5f, 0x5a, 0xe0, 0x19, 0x0,
	0xe0, 0x45, 0xaf, 0xd7, 0xfe, 0xaf, 0x85, 0x67, 0xce, 0xa, 0xb7, 0x2d, 0xbb, 0xbd, 0x84, 0x2e,
	0x5c, 0x78, 0xe7, 0xe0, 0x90, 0x8e, 0xdd, 0x3, 0x4d, 0x32, 0x8e, 0x6d, 0xdb, 0x71, 0xe6, 0x2c,
	0xdb, 0x8e, 0x33, 0xcf, 0xb2, 0xd, 0x44, 0xec, 0x9c, 0x4, 0x76, 0xde, 0xdb, 0xf8, 0x7a, 0x98,
	0x6d, 0x68, 0xdf, 0xba, 0x1e, 0xea, 0xb2, 0x73, 0xd3, 0x67, 0xd9, 0xb9, 0xe9, 0xeb, 0xb2, 0x33,
	0x1c, 0xb0, 0xec, 0xc, 0x7, 0x5a, 0xec, 0x7c, 0x3d, 0x77, 0x20, 0xce, 0xb6, 0xf3, 0x6c, 0xe2,
	0xac, 0x46, 0xf3, 0x4c, 0x7a, 0x8, 0x99, 0x3a, 0xd, 0x4c, 0xdd, 0xd2, 0xa9, 0x8a, 0xb1, 0xf5,
	0x3d, 0x76, 0x2d, 0x7b, 0xca, 0xb0, 0xe5, 0xd1, 0x1, 0xb9, 0x6c, 0x1d, 0x12, 0x5b, 0xef, 0x6d,
	0x6, 0x7a, 0x49, 0x36, 0x64, 0x5d, 0xcf, 0xfb, 0x2e, 0xfa, 0xc8, 0x99, 0xa1, 0x5, 0x1a, 0xc4,
	0xdc, 0x2d, 0xfa, 0x0, 0x57, 0x73, 0xdc, 0x4, 0x7, 0x64, 0xba, 0x53, 0x70, 0x18, 0x7a, 0xdc,
	0x4, 0xaa, 0x21, 0x34, 0x49, 0x8, 0x12, 0x97, 0xa8, 0x5a, 0x10, 0xfe, 0x14, 0x4a, 0x51, 0x9c,
	0xf6, 0x7a, 0xed, 0x11, 0xf4, 0x50, 0x1b, 0x3f, 0x2e, 0x91, 0xd0, 0xc2, 0xe3, 0x87, 0x24, 0xbe,
	0xe2, 0xa9, 0x86, 0x44, 0xa6, 0x30, 0x97, 0x98, 0xa3, 0x80, 0x5b, 0xa2, 0x6b, 0x9a, 0x2a, 0xbb,
	0xae, 0x87, 0x5, 0x4, 0x21, 0xba, 0x60, 0x56, 0xf2, 0x12, 0x9, 0x83, 0x10, 0x5d, 0x8d, 0x55,
	0x83, 0x18, 0xe, 0x4c, 0x7, 0x21, 0xb7, 0xd4, 0xab, 0xc5, 0x11, 0x4e, 0x22, 0x17, 0xa, 0x99,
	0xea, 0xf9, 0xc, 0x79, 0x9e, 0x5, 0xed, 0x2e, 0x74, 0xa7, 0xde, 0xf, 0x8f, 0x4b, 0x14, 0x96,
	0x46, 0x7, 0xd4, 0x6f, 0x89, 0x60, 0x8f, 0x49, 0xb0, 0x92, 0x37, 0x9b, 0x22, 0xa3, 0x95, 0x49,
	0x1c, 0x8d, 0x45, 0xf2, 0x66, 0xa6, 0x18, 0x4b, 0x30, 0x89, 0xe1, 0x58, 0xde, 0x58, 0x36, 0x74,
	0x1f, 0x19, 0x6e, 0x8c, 0xe8, 0x80, 0x9c, 0xb1, 0x4, 0x93, 0x28, 0xc5, 0xf2, 0xbc, 0xd7, 0x6b,
	0x7b, 0x4b, 0x34, 0xb6, 0xe0, 0x5c, 0xfc, 0xae, 0xd3, 0x20, 0xb1, 0x65, 0x3b, 0x14, 0x54, 0x5b,
	0xde, 0xd2, 0xb1, 0xbd, 0x9c, 0x9c, 0xb, 0xa, 0xeb, 0x8f, 0x2b, 0xe4, 0x61, 0xb5, 0x0, 0x8f,
	0x49, 0x80, 0xd8, 0x5d, 0x8d, 0xb1, 0x50, 0x68, 0x27, 0x61, 0x8d, 0xf0, 0x77, 0xcb, 0x63, 0x2c,
	0x1f, 0x8d, 0x39, 0xfd, 0x71, 0x83, 0x59, 0x3, 0xe4, 0xbc, 0xe9, 0xa, 0x99, 0x50, 0xc2, 0xe4,
	0xa2, 0xd7, 0x6b, 0x8f, 0x1d, 0x1b, 0x43, 0xcb, 0x46, 0x6e, 0x9b, 0xd8, 0x11, 0x87, 0xe6, 0x11,
	0xa3, 0x9c, 0xd0, 0x90, 0x5a, 0xc2, 0x30, 0x34, 0xd4, 0x84, 0xb9, 0x6b, 0xfb, 0x34, 0xaa, 0x57,
	0x72, 0x41, 0x41, 0xb, 0x12, 0xa3, 0x48, 0xf8, 0x16, 0xcc, 0x3, 0x71, 0xd3, 0xcf, 0x9, 0xc4,
	0x4d, 0xdf, 0x34, 0x10, 0xc4, 0x82, 0x79, 0x20, 0x86, 0x83, 0x9c, 0x40, 0xc, 0x7, 0xa6, 0x81,
	0x20, 0x16, 0xc, 0x3, 0x41, 0x6b, 0x98, 0x5c, 0x40, 0x84, 0xb5, 0x86, 0x51, 0x2c, 0x22, 0x23,
	0xe6, 0xe0, 0x38, 0x8b, 0xab, 0xb3, 0x1a, 0x8f, 0x10, 0xf, 0xbf, 0xc2, 0xcb, 0x87, 0x47, 0x50,
	0xbf, 0x99, 0xc5, 0x23, 0x34, 0x62, 0x18, 0xf, 0xbf, 0x4a, 0xcc, 0x87, 0x47, 0x50, 0x3, 0x9a,
	0xc5, 0x23, 0x34, 0x62, 0xe, 0x8f, 0xe7, 0x61, 0xf9, 0xd5, 0x27, 0x7f, 0xdc, 0xc1, 0x25, 0xe3,
	0x21, 0x72, 0xc1, 0xfa, 0xa9, 0x5f, 0x1e, 0x99, 0xad, 0xd0, 0xb4, 0x78, 0xa0, 0x5c, 0xb5, 0xc7,
	0x5, 0xdc, 0x2, 0x2e, 0x25, 0xb1, 0x7d, 0xc4, 0xc8, 0x28, 0xb6, 0xf9, 0x4b, 0x3c, 0x2d, 0x1e,
	0x98, 0xe3, 0xe9, 0x45, 0x84, 0x25, 0x2d, 0x3, 0xd, 0x82, 0xa9, 0xa1, 0x48, 0xd4, 0xe1, 0x40,
	0x41, 0x50, 0xde, 0xf4, 0x8d, 0x42, 0x99, 0xbb, 0xcc, 0xd4, 0xe1, 0x40, 0x41, 0x50, 0xe, 0x7,
	0x46, 0xa1, 0xcc, 0x5d, 0xa8, 0xea, 0x70, 0xa0, 0x8, 0x28, 0x69, 0x31, 0x6b, 0x10, 0x4a, 0x3d,
	0xa5, 0x9d, 0x26, 0x1f, 0xcc, 0x1, 0x7a, 0x19, 0x1, 0xea, 0x97, 0xc3, 0x35, 0xa2, 0xfa, 0x10,
	0xf5, 0xb, 0x6a, 0x93, 0x88, 0x6a, 0x29, 0xb7, 0x35, 0xf9, 0x50, 0x4, 0xa2, 0x7e, 0x49, 0x6e,
	0x12, 0x51, 0x2d, 0x5, 0xbb, 0x26, 0x1f, 0xcc, 0x21, 0xfa, 0x22, 0x55, 0xcf, 0x93, 0x47, 0x1c,
	0x93, 0x75, 0xa7, 0xd0, 0x8e, 0x69, 0xa9, 0x75, 0xa9, 0x98, 0x87, 0x85, 0xe4, 0x23, 0xd8, 0xc8,
	0x2d, 0x33, 0x1f, 0xb4, 0x48, 0xaf, 0x74, 0x3e, 0x7c, 0xf, 0xcd, 0xe5, 0xe3, 0x65, 0xfa, 0x39,
	0xa2, 0xdc, 0x84, 0xf8, 0x95, 0x7e, 0x85, 0xf3, 0x11, 0x38, 0x58, 0x50, 0x3a, 0xfc, 0x4d, 0xed,
	0x32, 0xd3, 0x41, 0x9e, 0x16, 0x2a, 0x9d, 0xe, 0xea, 0x60, 0x41, 0xe9, 0xf0, 0xb7, 0xd6, 0xcb,
	0x4c, 0x7, 0x79, 0xe2, 0xa8, 0x74, 0x3a, 0xa8, 0x83, 0x45, 0xa4, 0x23, 0xda, 0xe0, 0x2f, 0x2f,
	0x1d, 0x51, 0x81, 0x5d, 0xe1, 0x8c, 0xc4, 0x3e, 0x9a, 0x4b, 0xca, 0xd5, 0xda, 0x73, 0x55, 0x9d,
	0x95, 0x6a, 0x65, 0x25, 0x6e, 0x76, 0x94, 0x98, 0x95, 0xf0, 0xd1, 0xa9, 0xca, 0x59, 0x89, 0x7c,
	0x2c, 0x22, 0x2b, 0x71, 0xcb, 0xa5, 0xc4, 0xac, 0x84, 0x8f, 0x5f, 0x55, 0xce, 0x4a, 0xe4, 0xa3,
	0xe1, 0x96, 0xf1, 0xdd, 0x6a, 0x8e, 0x2d, 0xb6, 0x36, 0xf7, 0x65, 0x6a, 0xdc, 0xba, 0xbe, 0x28,
	0xeb, 0xb7, 0xe, 0x47, 0xd0, 0x43, 0xdf, 0xe9, 0x12, 0xec, 0x66, 0x19, 0x21, 0x28, 0xb2, 0x8d,
	0x28, 0xab, 0x79, 0xf6, 0x44, 0x9d, 0x78, 0xb6, 0x80, 0x4b, 0xa6, 0xf, 0xc6, 0xfa, 0x55, 0xfb,
	0xba, 0x79, 0xc0, 0x14, 0x34, 0xbf, 0x4c, 0x8d, 0xd3, 0xc9, 0x3, 0x71, 0x95, 0xb3, 0x39, 0x1e,
	0x64, 0x36, 0xae, 0x8a, 0xe3, 0x1, 0x71, 0x41, 0xc9, 0x45, 0x3d, 0x3c, 0x38, 0x8b, 0xf2, 0xcb,
	0x91, 0x84, 0x5f, 0xa5, 0x7, 0xea, 0x63, 0x82, 0x84, 0x52, 0xdc, 0x14, 0x11, 0xb2, 0xbb, 0x6e,
	0x45, 0xf1, 0x80, 0x7a, 0xa0, 0xe2, 0xa0, 0x1, 0x16, 0xb0, 0x34, 0xf5, 0x57, 0xe9, 0x81, 0x3a,
	0x59, 0x20, 0x2c, 0xb5, 0x37, 0xc7, 0x82, 0xac, 0x86, 0x61, 0x71, 0x2c, 0xb8, 0xe9, 0x2b, 0x75,
	0x34, 0xd, 0xb0, 0x80, 0x75, 0x28, 0xe1, 0x2a, 0x3d, 0x50, 0x27, 0xb, 0x84, 0xcf, 0x2a, 0x98,
	0x63, 0x41, 0x56, 0xaf, 0xb3, 0x38, 0x16, 0xc, 0x7, 0x4a, 0xcd, 0x58, 0xdd, 0x2c, 0xe0, 0x9c,
	0xea, 0xb8, 0x4a, 0xf, 0xd4, 0xc7, 0x2, 0xb9, 0xb3, 0xf, 0xa6, 0x88, 0xc0, 0x6c, 0x2a, 0x16,
	0xc5, 0x85, 0xd0, 0x9, 0x45, 0x37, 0xb3, 0x19, 0xf1, 0x19, 0xeb, 0x5c, 0xcc, 0x21, 0x4d, 0x68,
	0xa7, 0x23, 0xc1, 0x9a, 0xf3, 0x88, 0xc, 0xbc, 0xe3, 0x31, 0xaf, 0xd6, 0x46, 0xd6, 0xbc, 0x79,
	0x4a, 0xbc, 0x91, 0x59, 0x49, 0x62, 0x4e, 0xf0, 0x8e, 0x19, 0xbd, 0x5a, 0x1b, 0xa9, 0x91, 0x13,
	0x52, 0x67, 0x8f, 0x8c, 0x71, 0x82, 0xd5, 0x4e, 0x2f, 0x8c, 0x13, 0x81, 0x13, 0x8a, 0x6e, 0xea,
	0xe1, 0x84, 0x7f, 0xae, 0x5, 0x7a, 0xa8, 0x92, 0x47, 0x83, 0xd9, 0xab, 0x22, 0x99, 0x10, 0x41,
	0xdb, 0x8f, 0xf6, 0x75, 0xaf, 0xd7, 0x5e, 0x20, 0x3c, 0x73, 0x26, 0x6d, 0x32, 0xb0, 0xd, 0x6d,
	0xdb, 0xc1, 0x10, 0x5b, 0x8e, 0x2d, 0x7, 0x43, 0x15, 0x8f, 0x13, 0x33, 0x61, 0xa0, 0x13, 0xca,
	0xee, 0x34, 0x90, 0x50, 0xab, 0x78, 0xe6, 0x98, 0x19, 0xe9, 0x81, 0x37, 0x73, 0x5c, 0x99, 0xe3,
	0xa1, 0xc9, 0x50, 0x2b, 0x77, 0x32, 0x99, 0x19, 0xea, 0xbe, 0x65, 0xab, 0x6, 0x5a, 0xb9, 0xd3,
	0xcb, 0x6c, 0xf6, 0xce, 0x1d, 0x7b, 0xaa, 0x12, 0xe9, 0x53, 0x3d, 0xe2, 0xfc, 0x41, 0xf2, 0x88,
	0xf3, 0x59, 0x18, 0x70, 0x85, 0x8f, 0x39, 0x33, 0x23, 0xe, 0xa6, 0x94, 0x3e, 0x7, 0x11, 0xdc,
	0x93, 0x76, 0xf2, 0x28, 0x2a, 0x13, 0xb0, 0xa3, 0xe0, 0xbe, 0xf6, 0xef, 0xff, 0x28, 0x61, 0xb6,
	0xab, 0x67, 0x54, 0xd9, 0x24, 0x23, 0x33, 0x4a, 0x1, 0x76, 0x91, 0xba, 0xd, 0xee, 0xe0, 0x51,
	0x56, 0x76, 0xed, 0x44, 0x6f, 0xa5, 0xca, 0x80, 0xed, 0xe4, 0x91, 0x57, 0xf6, 0xba, 0x6d, 0xd9,
	0xea, 0x70, 0xed, 0xe4, 0xc1, 0x58, 0xf6, 0xf5, 0x48, 0x6e, 0xeb, 0x6a, 0x78, 0xed, 0xf6, 0xf9,
	0x59, 0xf6, 0x55, 0x49, 0x8b, 0x3, 0x29, 0xd8, 0x2e, 0xd3, 0xe5, 0xc1, 0xa7, 0x89, 0x5b, 0xd3,
	0x9f, 0x53, 0xa, 0xb8, 0xb4, 0x38, 0x9b, 0x20, 0x58, 0x1f, 0xb8, 0x94, 0x2, 0xfd, 0x82, 0xbe,
	0x2, 0x71, 0x85, 0xad, 0x79, 0xf7, 0x6f, 0xd0, 0x9b, 0xdd, 0xc1, 0xa5, 0xb2, 0x14, 0xdb, 0x38,
	0xfa, 0x4f, 0xf2, 0x48, 0xa6, 0x39, 0xf4, 0xd7, 0x84, 0xd7, 0x86, 0xe1, 0x7f, 0x92, 0x87, 0x38,
	0xb, 0x2, 0xff, 0xa6, 0x6f, 0x1c, 0xfc, 0x27, 0x78, 0xec, 0xb3, 0x20, 0xf0, 0x87, 0x3, 0xe3,
	0xe0, 0x3f, 0xc1, 0x83, 0xa2, 0x45, 0x80, 0x4f, 0x6b, 0x3c, 0xc3, 0xe0, 0x3f, 0xe1, 0x83, 0x90,
	0xe6, 0x52, 0xb0, 0x2e, 0x98, 0xae, 0x73, 0x50, 0x7c, 0xe, 0x2e, 0x62, 0x9, 0x1f, 0xb7, 0xef,
	0x52, 0xcb, 0x3e, 0x9f, 0x9e, 0xec, 0x93, 0xc9, 0x9c, 0x5f, 0x7, 0x9b, 0x7b, 0x9d, 0x68, 0x93,
	0xaf, 0xa3, 0x91, 0x4b, 0xb5, 0x74, 0x74, 0xc7, 0xa4, 0xa3, 0x4c, 0x2e, 0x7d, 0x49, 0x66, 0xec,
	0xf8, 0x3b, 0x9f, 0xb9, 0x58, 0x74, 0x99, 0x62, 0x51, 0x2d, 0x3c, 0xdd, 0x2d, 0xe1, 0x29, 0x93,
	0x43, 0x5f, 0xd1, 0x8d, 0xe0, 0x4e, 0xb0, 0x1d, 0xac, 0x97, 0x45, 0xb5, 0x70, 0x75, 0x87, 0x84,
	0xab, 0x4c, 0x16, 0xfd, 0xdc, 0xb2, 0x71, 0x87, 0xee, 0x90, 0xeb, 0x65, 0x50, 0x2d, 0x7a, 0xdd,
	0x21, 0xd1, 0x2b, 0xfb, 0x5e, 0x36, 0x77, 0xec, 0x69, 0xc7, 0xef, 0x1a, 0x68, 0xe4, 0x50, 0x2d,
	0x99, 0xdd, 0x5d, 0xe9, 0x23, 0xfb, 0xbe, 0x46, 0x5b, 0x29, 0x9d, 0xa0, 0xa1, 0x92, 0x8b, 0x51,
	0x9f, 0xa7, 0x18, 0x55, 0xcb, 0x69, 0x3f, 0x59, 0x4e, 0xfd, 0xca, 0x9f, 0xb3, 0x13, 0x76, 0x9b,
	0x72, 0xb1, 0x8a, 0xbe, 0xdc, 0xfd, 0x8e, 0xa, 0x16, 0xff, 0xc2, 0xd7, 0x88, 0x96, 0xa2, 0xbc,
	0x3c, 0xfb, 0x6, 0xde, 0x43, 0xdf, 0xc5, 0x7f, 0xc0, 0x45, 0x4, 0xc3, 0x71, 0xec, 0xf9, 0x86,
	0xf4, 0xd2, 0x86, 0xb, 0x24, 0x29, 0xbd, 0x4c, 0x0, 0xf1, 0xc6, 0x30, 0x10, 0x8a, 0x2a, 0x3d,
	0x51, 0x1c, 0x64, 0xf3, 0xfe, 0xb6, 0x7a, 0xe1, 0xee, 0xf1, 0x44, 0x89, 0x6, 0xe1, 0xb8, 0x15,
	0x86, 0x63, 0x2f, 0x8b, 0xe8, 0x2e, 0xfa, 0x78, 0x6d, 0x74, 0xd3, 0x8b, 0x58, 0xe0, 0x9, 0x49,
	0xc5, 0x16, 0x46, 0x63, 0x4, 0xa3, 0x4a, 0xb0, 0x6f, 0x97, 0xe4, 0xf2, 0x83, 0xf3, 0x4a, 0xaa,
	0xba, 0x3, 0x7d, 0xfa, 0x49, 0xaf, 0xd7, 0x26, 0x2b, 0x68, 0xdb, 0x5e, 0xcd, 0xe7, 0xc2, 0x9b,
	0x5d, 0x61, 0x64, 0x55, 0x94, 0xf4, 0x4a, 0xbf, 0x6e, 0x2f, 0xc, 0x66, 0x47, 0xbe, 0xf6, 0xe5,
	0xc5, 0x3a, 0xf5, 0x76, 0x52, 0xbc, 0x29, 0xfd, 0xd6, 0xa0, 0x14, 0x65, 0x77, 0x50, 0x38, 0x27,
	0xfd, 0x16, 0x92, 0x34, 0xed, 0xeb, 0x77, 0xe1, 0x3, 0x0, 0x7e, 0xb6, 0x7e, 0xe9, 0xd4, 0xef,
	0x80, 0xcf, 0x8d, 0xe9, 0xeb, 0xd, 0x4c, 0xeb, 0x97, 0x6c, 0xe7, 0x84, 0xf4, 0x8b, 0xd, 0x48,
	0xeb, 0xd7, 0xef, 0x9a, 0xa9, 0x7, 0x56, 0x63, 0xc6, 0xa2, 0xf8, 0xc5, 0xe6, 0x60, 0x3d, 0x5f,
	0x39, 0xf6, 0x7a, 0xdb, 0xc4, 0x39, 0xbe, 0x80, 0x4c, 0x28, 0xf4, 0x5f, 0x24, 0x8d, 0x26, 0x9a,
	0xdd, 0xd4, 0xae, 0xd1, 0xea, 0x95, 0x69, 0xd9, 0x60, 0xc4, 0x5f, 0x6e, 0xd8, 0xd, 0xba, 0x20,
	0x9c, 0x90, 0xf3, 0x57, 0xb5, 0x6c, 0xcb, 0x6, 0x43, 0xfe, 0xe5, 0x86, 0xe1, 0xf0, 0x68, 0x31,
	0x27, 0x66, 0x2d, 0xb5, 0x2f, 0xc7, 0x7a, 0xa1, 0xa9, 0x4e, 0x2a, 0x39, 0x64, 0x12, 0xb5, 0x4d,
	0x1, 0x52, 0xf8, 0x43, 0x9b, 0x62, 0xae, 0x53, 0x4d, 0x3e, 0x99, 0x34, 0x6d, 0xed, 0xe, 0x56,
	0xf2, 0x81, 0xee, 0xab, 0x4c, 0x8a, 0xf1, 0xa2, 0xe6, 0xfd, 0x66, 0xd5, 0x1f, 0xfb, 0xe8, 0x99,
	0xcc, 0x77, 0xf, 0x63, 0x44, 0x43, 0xa8, 0xf2, 0x76, 0x3, 0xa, 0x9d, 0x14, 0x9, 0x8b, 0xee,
	0xe4, 0xbf, 0x5d, 0x79, 0xd8, 0x59, 0x58, 0xff, 0x43, 0x93, 0xca, 0x7, 0x28, 0xbe, 0xdf, 0xf6,
	0x57, 0x64, 0x23, 0xd7, 0x1a, 0x67, 0x7b, 0xf1, 0x79, 0x62, 0x94, 0x1e, 0x16, 0x5e, 0xa6, 0x66,
	0xcc, 0xb1, 0xe8, 0x9e, 0x87, 0xbb, 0x46, 0x53, 0x7f, 0x32, 0xe1, 0xd3, 0xfd, 0xdf, 0xc0, 0x7b,
	0x78, 0xb, 0x59, 0x2, 0xa9, 0x23, 0xba, 0x9f, 0x4f, 0xc6, 0xe4, 0x8b, 0x35, 0x31, 0x8f, 0x52,
	0x88, 0xaf, 0xc2, 0x10, 0xc7, 0x73, 0xe8, 0x79, 0x6d, 0xf4, 0x80, 0x91, 0xed, 0x9, 0x92, 0xf6,
	0x79, 0x32, 0xd4, 0x5c, 0x4f, 0xd6, 0x89, 0x28, 0x8c, 0x3e, 0x5c, 0x27, 0xed, 0x18, 0x2e, 0xaf,
	0x9, 0x2e, 0x6f, 0xac, 0xe9, 0x2d, 0x1a, 0x5b, 0xb, 0xc8, 0x52, 0x5d, 0x46, 0x90, 0x7b, 0xdd,
	0x78, 0x7c, 0x3e, 0x14, 0x32, 0xe6, 0x2c, 0x26, 0xe2, 0xf7, 0x36, 0x46, 0x53, 0xe4, 0xa, 0x7b,
	0x17, 0x8c, 0xd7, 0x1a, 0x71, 0x38, 0xa7, 0x86, 0x88, 0x4f, 0x41, 0x23, 0xda, 0x3b, 0x3f, 0x25,
	0x21, 0xbe, 0x25, 0x17, 0x4a, 0x72, 0xeb, 0xfc, 0x37, 0x8e, 0x3b, 0xed, 0xc2, 0x25, 0x1c, 0xcf,
	0x50, 0x77, 0xb2, 0x1a, 0x8d, 0x9c, 0x2e, 0x46, 0x1e, 0xf6, 0xba, 0xc4, 0xc7, 0x7f, 0x79, 0xc8,
	0xfd, 0xa7, 0xeb, 0xdc, 0x5b, 0x13, 0xe4, 0xb6, 0xa8, 0x79, 0x32, 0xfd, 0x1, 0x99, 0xfe, 0x25,
	0x6b, 0xd5, 0xc, 0x96, 0xad, 0x73, 0xb0, 0xcf, 0xc2, 0xe4, 0xc0, 0x12, 0xd8, 0x98, 0xf6, 0x31,
	0x68, 0x92, 0x19, 0x2d, 0x17, 0x4d, 0xb2, 0x41, 0xe0, 0xc5, 0x79, 0xf5, 0x93, 0x85, 0xd1, 0x43,
	0x17, 0x8d, 0x67, 0x4e, 0x37, 0xe1, 0x63, 0x18, 0x16, 0x6b, 0x3d, 0xc, 0xd7, 0xf5, 0x4a, 0xc5,
	0xf3, 0x6a, 0x23, 0x1e, 0xdf, 0x49, 0x4e, 0x40, 0xdb, 0x65, 0xd2, 0x41, 0x64, 0x59, 0x9e, 0x37,
	0xfd, 0x8e, 0xbb, 0xc7, 0xdd, 0x17, 0x10, 0x8c, 0x2e, 0xcb, 0x8c, 0xdf, 0x71, 0x67, 0x9a, 0x91,
	0xd8, 0xa5, 0x5e, 0x73, 0x43, 0xb8, 0xe3, 0x7e, 0x48, 0x3b, 0xee, 0x2c, 0x2f, 0x34, 0xee, 0xe6,
	0xad, 0x39, 0xb9, 0xad, 0xdf, 0xce, 0x63, 0xc2, 0xef, 0x18, 0x57, 0xf0, 0xd6, 0x6c, 0x8b, 0x53,
	0x24, 0xa1, 0x7e, 0xd6, 0x41, 0x91, 0x58, 0xfd, 0x5c, 0x26, 0x45, 0x12, 0xe2, 0xe2, 0xf2, 0x28,
	0x22, 0xa1, 0x80, 0x2e, 0x8c, 0x22, 0x89, 0x6c, 0x73, 0x28, 0x92, 0xa1, 0x6c, 0xce, 0xcf, 0x91,
	0x84, 0xb2, 0xb9, 0x3c, 0x8a, 0x24, 0xa5, 0xc3, 0x65, 0x31, 0x44, 0x42, 0xdd, 0xac, 0x81, 0x1f,
	0xd7, 0x22, 0xfc, 0x48, 0xa5, 0x5a, 0x86, 0x20, 0xf1, 0xb6, 0x84, 0xe, 0x82, 0x88, 0xdf, 0x44,
	0xd, 0x12, 0x24, 0xee, 0x1d, 0x94, 0x47, 0x10, 0x61, 0xe1, 0x72, 0x81, 0x4, 0x89, 0x53, 0x2d,
	0x43, 0x90, 0x58, 0x93, 0xac, 0x83, 0x20, 0x91, 0x26, 0xb9, 0x4c, 0x82, 0xc4, 0xa2, 0xdf, 0xf2,
	0x8, 0x22, 0xac, 0x4b, 0x2e, 0x90, 0x20, 0x71, 0xaa, 0x85, 0x9, 0x92, 0x12, 0x1c, 0xe7, 0x27,
	0x48, 0x5a, 0x1c, 0x5a, 0x1e, 0x47, 0xd6, 0x74, 0x97, 0x65, 0xd1, 0x44, 0x5a, 0x20, 0x5a, 0x14,
	0x59, 0x52, 0x69, 0xe7, 0x90, 0x25, 0x4b, 0x4a, 0x5c, 0xb3, 0x65, 0x7, 0xd8, 0xd2, 0x17, 0x61,
	0x4b, 0x3a, 0xef, 0xc2, 0x74, 0x49, 0xb7, 0xf, 0x34, 0xd0, 0x25, 0xd5, 0x42, 0x28, 0x91, 0x2e,
	0xe9, 0x36, 0x7c, 0x69, 0x74, 0x91, 0x7d, 0x71, 0x6f, 0x51, 0x74, 0x49, 0xe7, 0x9d, 0x43, 0x97,
	0xec, 0x56, 0x3f, 0x67, 0x47, 0x48, 0xbc, 0x81, 0xd1, 0xc, 0xa7, 0x57, 0xdf, 0x10, 0x1a, 0x70,
	0xe2, 0xde, 0x1a, 0x0, 0x27, 0x72, 0x86, 0x7a, 0x82, 0x13, 0x7a, 0xc3, 0x15, 0xdf, 0x2e, 0xca,
	0x1f, 0xfb, 0xef, 0x25, 0x63, 0x17, 0x4a, 0x3b, 0x53, 0xf2, 0x10, 0x84, 0x9f, 0x45, 0x6b, 0xb2,
	0x54, 0xd0, 0xf6, 0x8d, 0x26, 0x8, 0xb2, 0xae, 0xcf, 0x25, 0x1c, 0xff, 0xa4, 0x6a, 0x45, 0x7c,
	0x15, 0xd0, 0xfa, 0xfe, 0x82, 0x35, 0x2f, 0x36, 0x2f, 0xf1, 0xac, 0x95, 0xa2, 0xc0, 0x17, 0x18,
	0xac, 0xf9, 0xd8, 0x50, 0xa0, 0xe4, 0x1f, 0x4, 0x29, 0xb9, 0x4e, 0x2c, 0xe, 0x29, 0xd9, 0xaa,
	0x94, 0xfc, 0xac, 0x4c, 0x3e, 0x5f, 0x1b, 0x23, 0x65, 0xb6, 0x91, 0xe2, 0x38, 0x99, 0x7a, 0x40,
	0xae, 0x26, 0x25, 0x19, 0x2e, 0xaa, 0x30, 0xf2, 0x8f, 0x32, 0x8c, 0x4c, 0xb2, 0x8a, 0x43, 0x49,
	0x8e, 0x62, 0x88, 0x53, 0x54, 0xc9, 0x7e, 0x11, 0x82, 0x32, 0x2b, 0x35, 0x7f, 0x13, 0x82, 0x2a,
	0x71, 0xd, 0x73, 0x86, 0xed, 0xa5, 0x4a, 0x3d, 0xf5, 0x27, 0x19, 0xda, 0xa4, 0x32, 0x2f, 0xbd,
	0x94, 0x6d, 0x69, 0x49, 0x65, 0xf6, 0xc2, 0x83, 0xa5, 0x4c, 0x5b, 0x91, 0x91, 0xb5, 0x64, 0xfa,
	0x8b, 0x99, 0xa2, 0x19, 0x61, 0x56, 0xe8, 0x6e, 0x7d, 0x29, 0x2e, 0x68, 0xc5, 0xb6, 0xbe, 0xa,
	0x5e, 0xd2, 0x24, 0x5a, 0x60, 0x1c, 0x75, 0x9d, 0x6, 0x6e, 0x6a, 0xb8, 0xcf, 0xa, 0x50, 0x53,
	0xc3, 0x8d, 0x56, 0xf3, 0x6e, 0x79, 0x45, 0x89, 0xa9, 0xf9, 0x56, 0x2b, 0xb5, 0x66, 0xca, 0x6c,
	0x9b, 0xf3, 0x4, 0x90, 0xda, 0x77, 0x30, 0x54, 0xb9, 0xa9, 0x7d, 0x7, 0x43, 0xb5, 0x4e, 0x34,
	0x4c, 0x1c, 0xfd, 0xf7, 0xdb, 0x3f, 0xab, 0xdc, 0x6f, 0x85, 0xc8, 0xb3, 0x45, 0x59, 0x18, 0xf0,
	0x25, 0xeb, 0x98, 0x70, 0xd3, 0x45, 0x1f, 0xbf, 0xb6, 0xd0, 0x7c, 0x62, 0xb4, 0x21, 0x17, 0x30,
	0x41, 0x91, 0x28, 0x2d, 0xaa, 0x8a, 0xfb, 0x76, 0xf4, 0x23, 0x1a, 0x63, 0xe1, 0xfd, 0x2e, 0x7e,
	0x1a, 0x7e, 0xcb, 0x49, 0x43, 0x1a, 0x47, 0xe, 0xf2, 0xdb, 0x54, 0xa2, 0x1c, 0xe8, 0x8f, 0x5c,
	0xe4, 0x2d, 0x3f, 0x51, 0xec, 0xbb, 0xa2, 0xd8, 0xb, 0xd1, 0x3e, 0x9, 0xfe, 0xbb, 0xc5, 0x8,
	0x4d, 0x26, 0x68, 0xc2, 0x1, 0xff, 0xc2, 0xb2, 0x31, 0x72, 0xa3, 0x73, 0x97, 0x32, 0x8f, 0x27,
	0x8a, 0x79, 0xb8, 0xc, 0x2d, 0x6, 0xe2, 0x3f, 0xe1, 0xe7, 0xe7, 0xc2, 0xd0, 0xe, 0xa1, 0x8b,
	0x85, 0x80, 0xcf, 0x58, 0x1b, 0x80, 0x59, 0x9a, 0x70, 0x1e, 0xed, 0x11, 0x5f, 0x3c, 0x7e, 0x46,
	0x59, 0x17, 0xcd, 0xc8, 0x41, 0x9e, 0x6, 0xee, 0xff, 0x7a, 0xd7, 0x46, 0x1e, 0x26, 0x11, 0x4,
	0xdf, 0xa8, 0x84, 0xdd, 0x15, 0xe2, 0xe4, 0xe5, 0x62, 0x1c, 0x85, 0x70, 0x87, 0x3c, 0xf, 0x4e,
	0x59, 0xdf, 0x2b, 0xc4, 0x62, 0x82, 0x4a, 0x66, 0x78, 0x5b, 0xd9, 0x5b, 0xe0, 0x8d, 0xb3, 0xe3,
	0xab, 0x40, 0xcf, 0x41, 0x33, 0xf8, 0x3c, 0xa, 0x3e, 0x8f, 0xe9, 0x27, 0xf8, 0x7f, 0x0, 0x0,
	0x0, 0xff, 0xff, 0xee, 0x19, 0xc6, 0x25, 0xbd, 0xaf, 0x0, 0x0,
}

func init() {
	if file_api_thrift != nil {
		return
	}
	type x struct{}
	builder := &thrift_reflection.FileDescriptorBuilder{
		Bytes:         file_idl_api_rawDesc,
		GoTypes:       file_api_thrift_go_types,
		GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
	}
	file_api_thrift = thrift_reflection.BuildFileDescriptor(builder)
}

func GetFileDescriptorForApi() *thrift_reflection.FileDescriptor {
	return file_api_thrift
}
func (p *EchoRequest) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoRequest")
}

func (p *EchoRequest) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoRequest"
	return ret
}
func (p *EchoResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoResponse")
}

func (p *EchoResponse) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoResponse"
	return ret
}
func (p *EchoMultiBoolResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoMultiBoolResponse")
}

func (p *EchoMultiBoolResponse) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoMultiBoolResponse"
	return ret
}
func (p *EchoMultiByteResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoMultiByteResponse")
}

func (p *EchoMultiByteResponse) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoMultiByteResponse"
	return ret
}
func (p *EchoMultiInt16Response) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoMultiInt16Response")
}

func (p *EchoMultiInt16Response) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoMultiInt16Response"
	return ret
}
func (p *EchoMultiInt32Response) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoMultiInt32Response")
}

func (p *EchoMultiInt32Response) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoMultiInt32Response"
	return ret
}
func (p *EchoMultiInt64Response) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoMultiInt64Response")
}

func (p *EchoMultiInt64Response) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoMultiInt64Response"
	return ret
}
func (p *EchoMultiFloatResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoMultiFloatResponse")
}

func (p *EchoMultiFloatResponse) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoMultiFloatResponse"
	return ret
}
func (p *EchoMultiDoubleResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoMultiDoubleResponse")
}

func (p *EchoMultiDoubleResponse) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoMultiDoubleResponse"
	return ret
}
func (p *EchoMultiStringResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoMultiStringResponse")
}

func (p *EchoMultiStringResponse) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoMultiStringResponse"
	return ret
}
func (p *EchoOptionalStructRequest) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoOptionalStructRequest")
}

func (p *EchoOptionalStructRequest) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoOptionalStructRequest"
	return ret
}
func (p *EchoOptionalStructResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoOptionalStructResponse")
}

func (p *EchoOptionalStructResponse) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoOptionalStructResponse"
	return ret
}
func (p *EchoOptionalMultiBoolRequest) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoOptionalMultiBoolRequest")
}

func (p *EchoOptionalMultiBoolRequest) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoOptionalMultiBoolRequest"
	return ret
}
func (p *EchoOptionalMultiInt32Request) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoOptionalMultiInt32Request")
}

func (p *EchoOptionalMultiInt32Request) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoOptionalMultiInt32Request"
	return ret
}
func (p *EchoOptionalMultiStringRequest) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoOptionalMultiStringRequest")
}

func (p *EchoOptionalMultiStringRequest) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoOptionalMultiStringRequest"
	return ret
}
func (p *EchoOptionalMultiBoolResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoOptionalMultiBoolResponse")
}

func (p *EchoOptionalMultiBoolResponse) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoOptionalMultiBoolResponse"
	return ret
}
func (p *EchoOptionalMultiInt32Response) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoOptionalMultiInt32Response")
}

func (p *EchoOptionalMultiInt32Response) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoOptionalMultiInt32Response"
	return ret
}
func (p *EchoOptionalMultiStringResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoOptionalMultiStringResponse")
}

func (p *EchoOptionalMultiStringResponse) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoOptionalMultiStringResponse"
	return ret
}
func (p *EchoGenericRequest) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoGenericRequest")
}

func (p *EchoGenericRequest) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoGenericRequest"
	return ret
}
func (p *EchoGenericResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoGenericResponse")
}

func (p *EchoGenericResponse) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoGenericResponse"
	return ret
}
func (p *EchoGenericEmbedded) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetStructDescriptor("EchoGenericEmbedded")
}

func (p *EchoGenericEmbedded) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoGenericEmbedded"
	return ret
}
func (p *EchoCustomizedException) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api_thrift.GetExceptionDescriptor("EchoCustomizedException")
}

func (p *EchoCustomizedException) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api_thrift.Filepath
	ret.Name = "EchoCustomizedException"
	return ret
}
