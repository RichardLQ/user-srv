package collect

import (
	"fmt"
	"github.com/RichardLQ/user-srv/client"
	"github.com/RichardLQ/user-srv/refer"
)

//FindCollectList 查询收藏列表
func (c *Collect) FindCollectList() (*[]CollectList ,error) {
	list := &[]CollectList{}
	if c.Id == 0 && c.Openid == ""{
		return &[]CollectList{},fmt.Errorf("查询地址不可以为空")
	}
	sql := client.Global.Mini.Table(refer.Table_Collects)
	if c.Id != 0 {
		sql = sql.Where("id = ?", c.Id)
	}
	if c.Collects != 0 {
		sql = sql.Where("collects = ?", c.Collects)
	}
	if c.Openid != "" {
		sql = sql.Where("openid = ?", c.Openid)
	}
	err:=sql.Find(list).Error
	if err !=nil {
		return list,err
	}
	return list,nil
}

//UpdateCollect 更新点赞
func (c *Collect) UpdateCollect() error {
	if c.Id == 0 {
		return fmt.Errorf("查询地址不可以为空")
	}
	err:= client.Global.Mini.Table(refer.Table_Collects).
		Where("id = ?",c.Id).Update("collects",c.Collects).Error
	if err != nil {
		return err
	}
	return nil
}