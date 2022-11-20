package user

import (
	"github.com/RichardLQ/user-srv/client"
	"github.com/RichardLQ/user-srv/refer"
)


//FindForName 根据username查询
func (u *Users) FindForName() (*[]Users ,error) {
	list := &[]Users{}
	err:=client.Global.Work.Table(refer.Table_Users).Where("username = ?", u.UserName).
	Find(list).Error
	if err !=nil {
		return list,err
	}
	return list,nil
}