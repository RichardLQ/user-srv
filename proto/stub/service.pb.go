// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.9
// source: service.proto

package stub

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x73, 0x74, 0x75, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf9, 0x02, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x15, 0x2e, 0x73,
	0x74, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x10, 0x22, 0x0b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a,
	0x01, 0x2a, 0x12, 0x5a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x14, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x22, 0x1f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x5a,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e,
	0x73, 0x74, 0x75, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x2a, 0x17, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x5c, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x73,
	0x74, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x75, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x73, 0x74,
	0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*GetUserTokenReq)(nil),   // 0: stub.GetUserTokenReq
	(*GetUserInfoReq)(nil),    // 1: stub.GetUserInfoReq
	(*DelUserInfoReq)(nil),    // 2: stub.DelUserInfoReq
	(*GetCollectListReq)(nil), // 3: stub.GetCollectListReq
	(*GetUserTokenRsp)(nil),   // 4: stub.GetUserTokenRsp
	(*GetUserInfoRsp)(nil),    // 5: stub.GetUserInfoRsp
	(*DelUserInfoRsp)(nil),    // 6: stub.DelUserInfoRsp
	(*GetCollectListRsp)(nil), // 7: stub.GetCollectListRsp
}
var file_service_proto_depIdxs = []int32{
	0, // 0: stub.UserService.GetUserToken:input_type -> stub.GetUserTokenReq
	1, // 1: stub.UserService.GetUserInfo:input_type -> stub.GetUserInfoReq
	2, // 2: stub.UserService.DelUserInfo:input_type -> stub.DelUserInfoReq
	3, // 3: stub.UserService.GetCollectList:input_type -> stub.GetCollectListReq
	4, // 4: stub.UserService.GetUserToken:output_type -> stub.GetUserTokenRsp
	5, // 5: stub.UserService.GetUserInfo:output_type -> stub.GetUserInfoRsp
	6, // 6: stub.UserService.DelUserInfo:output_type -> stub.DelUserInfoRsp
	7, // 7: stub.UserService.GetCollectList:output_type -> stub.GetCollectListRsp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_user_struct_proto_init()
	file_collect_struct_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
