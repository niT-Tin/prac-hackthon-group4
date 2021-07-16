package interfaces

import (
	"Group4/dao"
	"Group4/models"
)

type User models.User

// CRUD the basic interface of the CRUD
type CRUD interface {
	Insert() dao.BRT
	Update() dao.BRT
	Delete() dao.BRT
	Select() dao.BRT
}

//type CRUDText interface {
//	InsertText(text *models.Text)(int64, error)
//	UpdateText(text *models.Text)(int64, error)
//	DeleteText(text *models.Text)(int64, error)
//	SelectT(TextId int)([]models.Text, error)
//}
