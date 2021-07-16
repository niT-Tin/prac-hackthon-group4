package models

import "gorm.io/gorm"

// BRT stands for `Basic Return Type`

type Rf func()

type BRT struct {
	Code   int64
	Msg    string
	ErrMsg error
	B      bool
	Token  string
}

type G4Data struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	Data      DataS  `json:"data"`
}

type DataS struct {
	DataU []User `json:"datau,omitempty"`
	DataT []Text `json:"datat,omitempty"`
}

type ModifyData struct {
	UserName  string `json:"user_name"`
	Password  string `json:"change_password"`
	AvatarUrl string `json:"avatar_url"`
}

type User struct {
	gorm.Model
	ID         int    `gorm:"primary_key" json:"id"`
	UserName   string `json:"user_name"`
	Password   string
	AvatarUrl  string `json:"avatar_url"`
	ActiveDays int    `json:"active_days"`
	TextCount  int    `json:"text_count"`
}

type Text struct {
	gorm.Model
	TextId   int    `json:"text_id"`
	TextData string `json:"text"`
}
