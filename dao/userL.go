package dao

import (
	"Group4/models"
	"Group4/util"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type User models.User

func IsExisted(UserName, Password string) bool {
	u := &User{
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

func NotExisted(UserName string) models.BRT {
	str := fmt.Sprintf("%s doesn`t exist", UserName)
	return models.BRT{
		ErrMsg: errors.New(str),
	}
}

//type BRT models.BRT

func (u *User) Insert() models.BRT {
	return NoIf(mysqlDB.Model(u).Create(u), "create user error")
}

func (u *User) Update() models.BRT {
	return NoIf(mysqlDB.Model(u).Where("user_name = ?", u.UserName).Updates(u), "update user error")
}

func (u *User) Delete() models.BRT {
	return NoIf(mysqlDB.Model(u).Where("user_name = ?", u.UserName).Delete(u), "delete user error")
}

func (u *User) Select() models.BRT {
	//user1 := user.(*User)
	return NoIf(mysqlDB.Model(u).Where("user_name = ?", u.UserName).First(u), "select user error")
}

func NoIf(tx *gorm.DB, em string) models.BRT {
	if tx.RowsAffected != 0 {
		return util.FormatBRT(tx.RowsAffected)
	}
	return util.FormatBRT(errors.New(em))
}
