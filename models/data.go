package models

import "gorm.io/gorm"

type G4Data struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	DataU     []User `json:"datau,omitempty"`
	DataT     []Text `json:"datat,omitempty"`
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
