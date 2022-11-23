package user

//Users 用户表结构
type Users struct {
	Id         int32  `json:"id"`
	Openid     string `json:"openid"`
	Gender     int32  `json:"gender"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Address    string `json:"address"`
	Nickname   string `json:"nickname"`
	Desc       string `json:"desc"`
	Logoutime string `json:"logoutime"`
	Updatetime string `json:"updatetime"`
	Createtime string `json:"createtime"`
}


