package service

import (
	"GoBlog/internal/dao"
	"GoBlog/internal/middleware/jwt"
	"GoBlog/internal/model"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"log"
	"strconv"
)

// Register 用户注册后直接登录
func Register(UserName string, Password string) (*model.LoginRes, error) {
	//用户名查重
	isDuplicate, _ := dao.GetUserByName(UserName)

	log.Println("is duplicate?", isDuplicate)

	if isDuplicate != nil {
		//数据库中已有该用户名
		return nil, errors.New("duplicate username")
		//一般不能将错误详细信息返回，本项目简化了状态码
	} else {
		//无重复,在数据库中添加用户信息
		user := model.User{
			UserName: UserName,
			Password: Password,
			Uid:      int(idgen.NextId()),
		}
		err := dao.CreateUser(&user)
		if err != nil {
			return nil, errors.New("create user failed")
		}
		token, err := jwt.Award(&user.Uid)
		if err != nil {
			return nil, errors.New("create token failed")
		}

		var registerRes = &model.LoginRes{
			Token:     token,
			Uid:       strconv.Itoa(user.Uid),
			UserName:  user.UserName,
			AvatarUrl: user.AvatarUrl,
		}
		return registerRes, nil
	}

}

// Login 用户登录，携带token
func Login(UserName string, Password string) (*model.LoginRes, error) {
	user := dao.GetUser(UserName, Password)
	if user == nil {
		return nil, errors.New("账号或密码错误！")
	}
	token, err := jwt.Award(&user.Uid)
	if err != nil {
		return nil, errors.New("create token failed")
	}

	var loginRes = &model.LoginRes{
		Token:     token,
		Uid:       strconv.Itoa(user.Uid),
		UserName:  user.UserName,
		AvatarUrl: user.AvatarUrl,
	}
	return loginRes, nil
}
