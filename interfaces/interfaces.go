package interfaces

import "Group4/models"

type User models.User

//type Text models.Text

// CRUD the basic interface of the CRUD
type CRUD interface {
	Insert(user interface{}) (int64, error)
	Update(user interface{}) (int64, error)
	Delete(user interface{}) (int64, error)
	Select(user interface{}) (int64, error)
}

//type CRUDText interface {
//	InsertText(text *models.Text)(int64, error)
//	UpdateText(text *models.Text)(int64, error)
//	DeleteText(text *models.Text)(int64, error)
//	SelectT(TextId int)([]models.Text, error)
//}
