package user

//Users 用户表结构
type Users struct {
	Id         int32  `json:"id"`
	Openid     string `json:"openid"`
	Gender     int32  `json:"gender"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Address    string `json:"address"`
	NickName   string `json:"nickname"`
	Desc       string `json:"desc"`
	LogoutTime string `json:"logoutime"`
	UpdateTime string `json:"updatetime"`
	CreateTime string `json:"createtime"`
}
