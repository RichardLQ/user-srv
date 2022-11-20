package stub

import (
	"context"
	"github.com/RichardLQ/user-srv/auth"
)

var UserService = &userServiceServer{}

type userServiceServer struct {
	UserServiceServer
}

//GetUserToken 获取token
func (this *userServiceServer) GetUserToken(ctx context.Context, req *GetUserTokenReq) (*GetUserTokenRsp,error) {
	rsp := GetUserTokenRsp{
		Code: 200,
		Msg: "token请求成功",
		Data: "",
	}
	m := &auth.MyClaims{
		UserName: req.Account,
		Password: req.Password,
	}
	token,err:=m.Encryption()
	if err != nil {
		rsp.Code = 201
		rsp.Msg = err.Error()
		return &rsp,nil
	}
	rsp.Data = token
	return &rsp,nil
}

//GetUserInfo 获取用户信息
func (this *userServiceServer) GetUserInfo(ctx context.Context, req *GetUserInfoReq) (rsp *GetUserInfoRsp, err error) {
	rsp = &GetUserInfoRsp{
		Code: 200,
		User: &User{
			Id:       req.UserId,
			Username: "Jerome",
			Desc:     "这是一条数据",
		},
	}
	return
}

//DelUserInfo 删除用户信息
func (this *userServiceServer) DelUserInfo(ctx context.Context, req *DelUserInfoReq) (rsp *DelUserInfoRsp, err error) {
	rsp = &DelUserInfoRsp{
		Code: 200,
		Msg:  "删除成功！",
		Data: "一大堆杂数据",
	}
	return
}
