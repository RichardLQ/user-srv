package auth

import (
	"fmt"
	"github.com/RichardLQ/user-srv/client"
	"github.com/RichardLQ/user-srv/model/user"
	"github.com/dgrijalva/jwt-go"
	"time"
)
//MyClaims 加密信息
type MyClaims struct {
	UserName string `json:"username"`//账号
	Password string `json:"password"` //密码
	Source   string `json:"source"` //来源
	Token string `json:"token"` //token
	jwt.StandardClaims
}

func (m *MyClaims) authVerify() error {
	u := user.Users{
		UserName: m.UserName,
	}
	list ,err := u.FindForName()
	if err != nil {
		return err
	}
	var flag = false
	for _, users := range *list {
		fmt.Println(users.Password)
		if users.Password == m.Password {
			flag = true
		}
	}
	if !flag {
		return fmt.Errorf("暂无此账号")
	}
	return nil
}

//Encryption 加密信息
func (m *MyClaims) Encryption() (string, error) {
	err :=m.authVerify()
	if err!=nil{
		return "", err
	}

	claims := MyClaims{
		UserName: m.UserName,
		Password: m.Password,
		Source: m.Source,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,      // 一分钟之前开始生效
			ExpiresAt: time.Now().Unix() + 60*60*2, // 两个小时后失效
			Issuer:    "Sourcandy.com",             // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token_string, err := token.SignedString([]byte(client.Global.UserConf.JwtKey))
	if err != nil {
		return "", err
	}
	return token_string, nil
}

//Decryption 解密信息
func (m *MyClaims) Decryption() (interface{}, error) {
	parseToken,err := jwt.ParseWithClaims(m.Token, &MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(client.Global.UserConf.JwtKey), nil
		})
	if err != nil {
		return "", fmt.Errorf("Unauthorized access to this resource%+v",err)
	}
	if !parseToken.Valid {
		return "", fmt.Errorf("Token is not valid！%+v",err)
	}
	return parseToken.Claims.(*MyClaims), nil
}
