package testing

import (
	"Group4/dao"
	"fmt"
	"testing"
)

func TestUser_Insert(t *testing.T) {

	dao.G4AutoMigrate()
	//fmt.Printf("%T\n", migrate)

	users := []dao.User{
		{UserName: "Json", Password: "123321", AvatarUrl: "https://2343453", ActiveDays: 12, TextCount: 23},
		{UserName: "Barry", Password: "123abc", AvatarUrl: "https://43vserwe", ActiveDays: 19, TextCount: 26},
		{UserName: "Constantins", Password: "qwerty", AvatarUrl: "https://127.0.0.1", ActiveDays: 2700, TextCount: 99},
	}

	//var s = dao.User{}
	for _, user := range users {
		insert := user.Insert()
		if insert.Code != 1 {
			t.Errorf("User: %+v", user)
		}
	}
}

func TestUser_Update(t *testing.T) {
	dao.G4AutoMigrate()
	//fmt.Printf("%T\n", migrate)

	users := []dao.User{
		{UserName: "Json", Password: "123321", AvatarUrl: "https://2343453", ActiveDays: 12, TextCount: 28},
		{UserName: "Barry", Password: "123abc", AvatarUrl: "https://43vserwe", ActiveDays: 19, TextCount: 31},
		{UserName: "Constantins", Password: "qwerty", AvatarUrl: "https://127.0.0.1", ActiveDays: 2700, TextCount: 104},
	}

	//var s = dao.User{}
	for _, user := range users {
		actual := user.Update()
		if actual.Code != 1 {
			t.Errorf("User: %+v", user)
		}
	}
}

func TestUser_Delete(t *testing.T) {
	dao.G4AutoMigrate()
	//fmt.Printf("%T\n", migrate)

	users := []dao.User{
		{UserName: "Json", Password: "123321", AvatarUrl: "https://2343453", ActiveDays: 12, TextCount: 28},
		{UserName: "Barry", Password: "123abc", AvatarUrl: "https://43vserwe", ActiveDays: 19, TextCount: 31},
		{UserName: "Constantins", Password: "qwerty", AvatarUrl: "https://127.0.0.1", ActiveDays: 2700, TextCount: 104},
	}

	//var s = dao.User{}
	for _, user := range users {
		actual := user.Delete()
		if actual.Code != 1 {
			t.Errorf("User: %+v", user)
		}
	}
}

func TestUser_Select(t *testing.T) {
	dao.G4AutoMigrate()
	//fmt.Printf("%T\n", migrate)

	users := []dao.User{
		{UserName: "Json", Password: "123321", AvatarUrl: "https://2343453", ActiveDays: 12, TextCount: 28},
		{UserName: "Barry", Password: "123abc", AvatarUrl: "https://43vserwe", ActiveDays: 19, TextCount: 31},
		{UserName: "Constantins", Password: "qwerty", AvatarUrl: "https://127.0.0.1", ActiveDays: 2700, TextCount: 104},
	}

	//var s = dao.User{}
	for _, user := range users {
		actual := user.Select()
		if actual.Code != 1 {
			t.Errorf("User: %+v", user)
		}
		fmt.Println("User: ", user)
	}
	var ss = dao.User{UserName: "Barry"}
	fmt.Printf("%+v\n", ss)
	ss.Select()
	fmt.Printf("%+v\n", ss)
}
