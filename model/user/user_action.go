package user

import (
	"fmt"
	"github.com/RichardLQ/user-srv/client"
	"github.com/RichardLQ/user-srv/refer"
)


//FindForName 根据username查询
func (u *Users) FindForName() (*[]Users ,error) {
	list := &[]Users{}
	err:=client.Global.Mini.Table(refer.Table_Users).Where("username = ?", u.Username).
	Find(list).Error
	if err !=nil {
		return list,err
	}
	return list,nil
}

//FindByIdOrOpenid 根据user_id或者openid查询
func (u *Users) FindByIdOrOpenid() (*Users ,error) {
	if u.Id == 0 && u.Openid == ""{
		return &Users{},fmt.Errorf("查询为空")
	}
	list := &Users{}
	sql := client.Global.Mini.Table(refer.Table_Users)
	if u.Id != 0 {
		sql = sql.Where("id = ?", u.Id)
	}
	if u.Openid != "" {
		sql = sql.Where("openid = ?", u.Openid)
	}
	err:=sql.Find(list).Error
	if err !=nil {
		return list,err
	}
	return list,nil
}

//DelUserById 根据id删除用户
func (u *Users) DelUserById() error {
	err:=client.Global.Mini.Table(refer.Table_Users).Delete(u).Error
	if err !=nil {
		return err
	}
	return nil
}



