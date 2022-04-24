package model

import (
	"encoding/json"
	"xupt-scholarship/db"
	"xupt-scholarship/mvc_struct"
)

type AnnouncementModel struct {
}

func (a *AnnouncementModel) GetAnnouncementDataByProcedureId(procedureId int) BaseModelFmtData {
	var applicationList []db.Application
	var announcementList []mvc_struct.AnnouncementData
	result := db.Mysql.Where("procedure_id = ?", procedureId).Find(&applicationList)
	if result.Error == nil {
		for _, v := range applicationList {
			var usermodel UserModel
			var scoreInfo mvc_struct.ApplyScoreInfo
			json.Unmarshal(v.ScoreInfo, &scoreInfo)
			stu := usermodel.GetStudent(v.UserId).Data.(mvc_struct.Student)
			announcementList = append(announcementList, mvc_struct.AnnouncementData{
				Id:           v.ID,
				Moral:        scoreInfo.Moral,
				Academic:     scoreInfo.Academic,
				Practice:     scoreInfo.Practice,
				Score:        v.Score,
				Name:         stu.Name,
				StudentId:    stu.UserId,
				College:      stu.Info.College,
				Professional: stu.Info.Professional,
				Grade:        stu.Info.Grade,
				Class:        stu.Info.Class,
				CourseCredit: stu.Course,
				ShipType:     v.ShipType,
			})
		}
	}
	return HandleDBData(result, announcementList)
}
