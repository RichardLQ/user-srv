package stub

import (
	"context"
)

//GetCollectList 获取收藏信息
func (this *userServiceServer) GetCollectList(ctx context.Context, req *GetCollectListReq) (*GetCollectListRsp, error) {
	rsp := GetCollectListRsp{
		Code: 200,
	}
	return &rsp,nil
}
