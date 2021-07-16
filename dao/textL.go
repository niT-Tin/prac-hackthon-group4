package dao

import (
	"Group4/models"
	"errors"
)

type Text models.Text

func (t *Text) Insert() models.BRT {
	return NoIf(mysqlDB.Where("text_id = ?", t.TextId).Create(t), "insert text error")
}

func (t *Text) Update() models.BRT {
	return NoIf(mysqlDB.Model(t).Where("text_id = ?", t.TextId).Update("text_data", t.TextData), "update text error")
}

func (t *Text) Delete() models.BRT {
	return NoIf(mysqlDB.Where("text_id = ?", t.TextId).Delete(t), "delete user error")
}

func SelectMultiText(texts interface{}) (int64, error) {
	text1 := texts.(*[]Text) // Mandatory type conversion
	textId := (*text1)[0].TextId
	if find := mysqlDB.Where("text_id = ?", textId).Find(text1); text1 != nil {
		return find.RowsAffected, nil
	} else {
		return 0, errors.New("find text error")
	}
}

// Select select the TextId in the first number of the text array
// the arguments is the pointer of array or slice of the Text
// process will get the texts of the specified TextId
func (t *Text) Select() models.BRT {
	return NoIf(mysqlDB.Where("text_id = ?", t.TextId).Find(t), "find text error")
}
