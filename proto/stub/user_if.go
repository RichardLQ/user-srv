package stub

import (
	"context"
	"github.com/RichardLQ/user-srv/auth"
	"github.com/RichardLQ/user-srv/model/user"
	"github.com/RichardLQ/user-srv/refer"
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
		Data: &TokenInfo{},
	}
	if req.Account == "" || req.Password == "" {
		rsp.Msg = "账号错误"
		rsp.Code = refer.Login_Miss
		return &rsp,nil
	}
	m := &auth.MyClaims{
		UserName: req.Account,
		Password: req.Password,
	}
	info,err:=m.Encryption()
	if err != nil {
		rsp.Code = 201
		rsp.Msg = err.Error()
		return &rsp,nil
	}
	rsp.Data.Token = info["token"].(string)
	rsp.Data.Openid = info["openid"].(string)
	rsp.Data.UserId = info["user_id"].(int32)
	return &rsp,nil
}

//GetUserInfo 获取用户信息
func (this *userServiceServer) GetUserInfo(ctx context.Context, req *GetUserInfoReq) (*GetUserInfoRsp,error) {
	rsp := GetUserInfoRsp{
		Code: 200,
	}
	u := user.Users{
		Id: req.UserId,
		Openid: req.Openid,
	}
	list ,err :=u.FindByIdOrOpenid()
	if err!= nil {
		return &rsp,nil
	}
	user := &User{
		Id: list.Id,
		Openid: list.Openid,
		Username: list.UserName,
		Address: list.Address,
		Nickname: list.NickName,
		Desc: list.Desc,
	}
	rsp.User = user
	return &rsp,nil
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
