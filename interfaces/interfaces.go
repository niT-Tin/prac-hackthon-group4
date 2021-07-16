package interfaces

import (
	"Group4/models"
)

//type User models.User

// CRUD the basic interface of the CRUD
type CRUD interface {
	Insert() models.BRT
	Update() models.BRT
	Delete() models.BRT
	Select() models.BRT
}
