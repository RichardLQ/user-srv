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
		Username: req.Account,
		Password: req.Password,
	}
	info,err:=m.Encryption()
	if err != nil {
		rsp.Code = refer.Token_err
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
		rsp.Code = refer.Find_User_Err
		rsp.Msg = err.Error()
		return &rsp,nil
	}
	user := &User{
		Id: list.Id,
		Openid: list.Openid,
		Username: list.Username,
		Address: list.Address,
		Nickname: list.Nickname,
		Desc: list.Desc,
	}
	rsp.User = user
	return &rsp,nil
}

//DelUserInfo 删除用户信息
func (this *userServiceServer) DelUserInfo(ctx context.Context, req *DelUserInfoReq) (*DelUserInfoRsp, error) {
	rsp := DelUserInfoRsp{
		Code: 200,
		Msg:  "请求失败！",
		Data: "",
	}
	user := &user.Users{
		Id: req.UserId,
	}
	err := user.DelUserById()
	if err != nil {
		rsp.Msg = err.Error()
		rsp.Code = refer.Del_User_Err
		return &rsp,err
	}
	rsp.Msg = "请求成功！"
	return &rsp,nil
}

//UpdateUserInfo 更新用户信息
func (this *userServiceServer) UpdateUserInfo(ctx context.Context, req *UpdateUserInfoReq) (*UpdateUserInfoRsp, error) {
	rsp := UpdateUserInfoRsp{
		Code: 200,
		Msg:  "请求失败！",
	}
	user := &user.Users{
		Id: req.Id,
		Openid: req.Openid,
		Gender: req.Gender,
		Username: req.Username,
		Password: req.Password,
		Address: req.Address,
		Nickname: req.Nickname,
		Desc: req.Desc,
	}
	err := user.UpdateUser()
	if err != nil {
		rsp.Msg = err.Error()
		rsp.Code = refer.Update_User_Err
		return &rsp,err
	}
	rsp.Msg = "更新成功！"
	return &rsp,nil
}
