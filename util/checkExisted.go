package util

import (
	"Group4/dao"
	"Group4/service"
	"errors"
	"fmt"
)

func IsExisted(UserName, Password string) bool {
	u := &dao.User{
		UserName: UserName,
		//Password: Password,
	}
	//u1 := dao.User{}
	brt := u.Select()
	if brt.ErrMsg != nil {
		return false
	}
	if u.Password == Password {
		return true
	}
	return false
}

func NotExisted(UserName string) service.BRT {
	str := fmt.Sprintf("%s doesn`t exist", UserName)
	return service.BRT{
		ErrMsg: errors.New(str),
	}
}
