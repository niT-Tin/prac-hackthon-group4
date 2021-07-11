package dao

import (
	"Group4/models"
	"errors"
)

type User models.User

func (User) Insert(user interface{}) (int64, error) {
	user1 := user.(*User)
	if create := mysqlDB.Model(user1).Create(user1); create.RowsAffected != 0 {
		return create.RowsAffected, nil
	}
	return 0, errors.New("create user error")
}

func (User) Update(user interface{}) (int64, error) {
	user1 := user.(*User)
	if updates := mysqlDB.Model(user1).Where("user_name = ?", user1.UserName).Updates(user1); updates.RowsAffected != 0 {
		return updates.RowsAffected, nil
	}
	return 0, errors.New("update user error")
}

func (User) Delete(user interface{}) (int64, error) {
	user1 := user.(*User)
	if deleted := mysqlDB.Where("user_name = ?", user1.UserName).Delete(user); deleted.RowsAffected != 0 {
		return deleted.RowsAffected, nil
	}
	return 0, errors.New("delete user error")
}

func (User) Select(user interface{}) (int64, error) {
	user1 := user.(*User)
	if selected := mysqlDB.Where("user_name = ?", user1.UserName).First(user); selected.RowsAffected != 0 {
		return selected.RowsAffected, nil
	}
	return 0, errors.New("select user error")
}
