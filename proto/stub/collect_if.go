package stub

import (
	"context"
	"github.com/RichardLQ/user-srv/model/collect"
	"github.com/RichardLQ/user-srv/refer"
)

//GetCollectList 获取收藏信息
func (this *userServiceServer) GetCollectList(ctx context.Context, req *GetCollectListReq) (*GetCollectListRsp, error) {
	rsp := GetCollectListRsp{
		Code: 200,
	}
	c :=collect.Collect{
		Id: req.UserId,
		Openid: req.Openid,
		Collects: req.Collects,
	}
	list,err := c.FindCollectList()
	if err != nil {
		rsp.Msg = err.Error()
		rsp.Code = refer.Find_Collect_Err
		return &rsp,err
	}
	var collectList []*CollectList
	for _, col := range *list {
		temp := &CollectList{
			Id: col.Id,
			Openid: col.Openid,
			Address: col.Address,
			Collects: col.Collects,
		}
		collectList = append(collectList,temp)
	}
	rsp.List = collectList
	return &rsp,nil
}


//UpdateCollect 更新收藏信息
func (this *userServiceServer) UpdateCollect(ctx context.Context, req *UpdateCollectReq) (*UpdateCollectRsp, error) {
	rsp := UpdateCollectRsp{
		Code: 200,
		Msg: "更新成功！",
	}
	c :=collect.Collect{
		Id: req.Id,
		Collects: req.Collects,
	}
	err := c.UpdateCollect()
	if err != nil {
		rsp.Msg = err.Error()
		rsp.Code = refer.Update_Collect_Err
		return &rsp,err
	}
	return &rsp,nil
}

