package client

import (
	"github.com/jinzhu/gorm"
)

var Global = struct {
	UserConf CostVar `confs:"read=rainbow;format=json"`
	Work     gorm.DB      `confs:"name=Work;read=rainbow"`
}{}

//CostVar 定义变量
type CostVar struct {
	JwtKey string
}
