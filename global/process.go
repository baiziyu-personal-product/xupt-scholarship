package global

type ProcessStepMapInfo struct {
	DeploymentMobilizationPhase                                string `json:"deployment_mobilization_phase"`
	IndividualApplicationStage                                 string `json:"individual_application_stage"`
	PoliciesForAllGrades                                       string `json:"policies_for_all_grades"`
	FirstSelfAssessment                                        string `json:"first_self_assessment"`
	FirstClassAnnouncement                                     string `json:"first_class_announcement"`
	SecondPersonalSelfAssessment                               string `json:"second_personal_self_assessment"`
	SecondClassAnnouncement                                    string `json:"second_class_announcement"`
	GradeAnnouncement                                          string `json:"grade_announcement"`
	ExaminationAndReviewOfTheDisciplineOffice                  string `json:"examination_and_review_of_the_discipline_office"`
	VerificationAndDeliberationByTheScholarshipEvaluationGroup string `json:"verification_and_deliberation_by_the_scholarship_evaluation_group"`
	SchoolReviewStage                                          string `json:"school_review_stage"`
	Finish                                                     string `json:"finish"`
}

var ProcessStepMap = ProcessStepMapInfo{
	DeploymentMobilizationPhase:                                "部署动员阶段",
	IndividualApplicationStage:                                 "个人申请阶段",
	PoliciesForAllGrades:                                       "各年级政策宣讲",
	FirstSelfAssessment:                                        "第一次个人自评",
	FirstClassAnnouncement:                                     "第一次班级公示",
	SecondPersonalSelfAssessment:                               "第二次个人自评",
	SecondClassAnnouncement:                                    "第二次班级公示",
	GradeAnnouncement:                                          "年级公示",
	ExaminationAndReviewOfTheDisciplineOffice:                  "学科办审核、复查",
	VerificationAndDeliberationByTheScholarshipEvaluationGroup: "奖学金评定小组核查、审议",
	SchoolReviewStage:                                          "学校评审阶段",
	Finish:                                                     "结束",
}
