package collect

//Collects 收藏表结构
type Collect struct {
	Id         int32  `json:"id"`
	Openid     string `json:"openid"`
	Types      int32  `json:"types"`
	Address    string `json:"address"`
	Collects   int32 `json:"collects"`
	Updatetime string `json:"updatetime"`
	Createtime string `json:"createtime"`
}

//CollectList 查询列表
type CollectList struct {
	Id         int32  `json:"id"`
	Openid     string `json:"openid"`
	Address    string `json:"address"`
	Collects   int32 `json:"collects"`
}