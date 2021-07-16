package util

import (
	"Group4/models"
	"errors"
)

// format the BRT aka Basic Return Type
// fill the BRT flexibly

func FormatBRT(a ...interface{}) models.BRT {
	if len(a) > 4 {
		return models.BRT{
			ErrMsg: errors.New("arguments counts error"),
		}
	}
	var bRT = models.BRT{}
	for _, val := range a {
		switch val.(type) {
		case int64:
			bRT.Code = val.(int64)
		case int:
			bRT.Code = int64(val.(int))
		case string:
			bRT.Msg = val.(string)
		case bool:
			bRT.B = val.(bool)
		case error:
			bRT.ErrMsg = val.(error)
		}
	}
	return bRT
}
