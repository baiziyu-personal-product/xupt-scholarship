package mvc_struct

type ProcessStepValue struct {
	Date     []string `json:"date"`
	Desc     string   `json:"desc"`
	Mentions string   `json:"mentions"`
}

type ProcessStepData struct {
	DeploymentMobilizationPhase                                ProcessStepValue `json:"deployment_mobilization_phase"`
	IndividualApplicationStage                                 ProcessStepValue `json:"individual_application_stage"`
	PoliciesForAllGrades                                       ProcessStepValue `json:"policies_for_all_grades"`
	FirstSelfAssessment                                        ProcessStepValue `json:"first_self_assessment"`
	FirstClassAnnouncement                                     ProcessStepValue `json:"first_class_announcement"`
	SecondPersonalSelfAssessment                               ProcessStepValue `json:"second_personal_self_assessment"`
	SecondClassAnnouncement                                    ProcessStepValue `json:"second_class_announcement"`
	GradeAnnouncement                                          ProcessStepValue `json:"grade_announcement"`
	ExaminationAndReviewOfTheDisciplineOffice                  ProcessStepValue `json:"examination_and_review_of_the_discipline_office"`
	VerificationAndDeliberationByTheScholarshipEvaluationGroup ProcessStepValue `json:"verification_and_deliberation_by_the_scholarship_evaluation_group"`
	SchoolReviewStage                                          ProcessStepValue `json:"school_review_stage"`
	Finish                                                     ProcessStepValue `json:"finish"`
}

type ProcessReqData struct {
	Form   ProcessStepData  `json:"form"`
	Upload []UploadFileItem `json:"upload"`
}
