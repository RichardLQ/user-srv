package stub

import "context"

var CollectService = &collectServiceServer{}

type collectServiceServer struct {
	CollectServiceServer
}

//GetCollectList 删除用户信息
func (this *userServiceServer) GetCollectList(ctx context.Context, req *GetCollectListReq) (*GetCollectListRsp, error) {
	rsp := GetCollectListRsp{
		Code: 200,
		Msg: "请求成功",
	}
	return &rsp,nil
}
