package model

import (
	"encoding/json"
	"xupt-scholarship/db"
	"xupt-scholarship/mvc_struct"
)

type ApplyModel struct {
}

func (a *ApplyModel) CreateApplyForm(data mvc_struct.BaseApply) BaseModelFmtData {
	jsonForm, _ := json.Marshal(data.Form)
	apply := db.Application{
		Form:    jsonForm,
		History: "{}",
		Creator: data.StudentId,
		Status:  data.Type,
		Step:    "",
	}

	result := db.Mysql.Create(&apply)
	return HandleDBData(result, apply.ID)
}
