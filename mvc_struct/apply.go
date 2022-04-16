package mvc_struct

type moralFormItem struct {
	Level string           `json:"level"`
	Name  string           `json:"name"`
	Info  string           `json:"info"`
	Time  string           `json:"time"`
	Files []UploadFileItem `json:"files"`
	Score float32          `json:"score" default:"0"`
}

type moralFormValue struct {
	List  []moralFormItem `json:"list"`
	Score float32         `json:"score" default:"0"`
}

type practiceFormResultItem struct {
	Level    string           `json:"level"`
	Time     string           `json:"time"`
	Name     string           `json:"name"`
	Order    int              `json:"order"`
	Partners int              `json:"partners"`
	Files    []UploadFileItem `json:"files"`
	Score    float32          `json:"score" default:"0"`
}

type socialCadreItem struct {
	Level      []string `json:"level"`
	Department string   `json:"department"`
	Score      float32  `json:"score" default:"0"`
}

type socialActivityItem struct {
	Level string           `json:"level"`
	Name  string           `json:"name"`
	Time  []string         `json:"time"`
	Files []UploadFileItem `json:"files"`
	Score float32          `json:"score" default:"0"`
}

type practiceSocialFormValue struct {
	Cadre    []socialCadreItem    `json:"cadre"`
	Activity []socialActivityItem `json:"activity"`
	Score    float32              `json:"score" default:"0"`
}

type practiceCompetitionFormValue struct {
	Level    []string         `json:"level"`
	Name     string           `json:"name"`
	Time     string           `json:"time"`
	Files    []UploadFileItem `json:"files"`
	Score    float32          `json:"score" default:"0"`
	Order    int              `json:"order" default:"1"`
	Partners int              `json:"partners" default:"1"`
}

type practiceFormValue struct {
	Result      []practiceFormResultItem       `json:"result"`
	Social      practiceSocialFormValue        `json:"social"`
	Competition []practiceCompetitionFormValue `json:"competition"`
	Score       float32                        `json:"score" default:"0"`
}

type academicAwardFormItem struct {
	Level []string         `json:"level"`
	Name  string           `json:"name"`
	Time  string           `json:"time"`
	Files []UploadFileItem `json:"files"`
	Score float32          `json:"score" default:"0"`
}

type academicScientificFormItem struct {
	Level                 string           `json:"level"`
	Name                  string           `json:"name"`
	Time                  []string         `json:"time"`
	FundsActuallyReceived int              `json:"funds_actually_received"`
	FundsDue              int              `json:"funds_due"`
	Distribute            int              `json:"distribute"`
	Files                 []UploadFileItem `json:"files"`
	Score                 float32          `json:"score" default:"0"`
	Order                 int              `json:"order" default:"1"`
	Partners              int              `json:"partners" default:"1"`
}

type academicDissertationFormItem struct {
	Level    string           `json:"level"`
	Name     string           `json:"name"`
	IdNumber string           `json:"id_number"`
	Time     string           `json:"time"`
	Files    []UploadFileItem `json:"files"`
	Score    float32          `json:"score" default:"0"`
}

type academicPublishFormItem struct {
	Level            string           `json:"level"`
	Name             string           `json:"name"`
	Time             string           `json:"time"`
	PublishHouseName string           `json:"publish_house_name"`
	FontsCount       int              `json:"fonts_count"`
	Files            []UploadFileItem `json:"files"`
	Score            float32          `json:"score" default:"0"`
}

type academicFormValue struct {
	Award        []academicAwardFormItem        `json:"award"`
	Scientific   []academicScientificFormItem   `json:"scientific"`
	Dissertation []academicDissertationFormItem `json:"dissertation"`
	Publish      []academicPublishFormItem      `json:"publish"`
	Score        float32                        `json:"score" default:"0"`
	ScoreInfo    ApplyScoreInfo                 `json:"score_info"`
}

type ApplicationValue struct {
	Moral    moralFormValue    `json:"moral"`
	Practice practiceFormValue `json:"practice"`
	Academic academicFormValue `json:"academic"`
}

type ApplicationRequest struct {
	ApplicationValue
	ScoreInfo ApplyScoreInfo `json:"score_info"`
}

type CreateApplyByBaseInfo struct {
	Form      ApplicationValue `json:"form"`
	StudentId string           `json:"student_id"`
	Type      string           `json:"type"`
	ScoreInfo ApplyScoreInfo   `json:"score_info"`
}

type SearchApplyByBaseInfo struct {
	Id        int    `json:"id"`
	StudentId string `json:"student_id"`
}

type UpdateApplyBaseInfo struct {
	Id        int              `json:"id"`
	Form      ApplicationValue `json:"form"`
	StudentId string           `json:"student_id"`
	Type      string           `json:"type"`
	ScoreInfo ApplyScoreInfo   `json:"score_info"`
}

type ApplyListParams struct {
	PageCount int    `json:"page_count" default:"10"`
	PageIndex int    `json:"page_index" default:"1"`
	IsCheck   bool   `json:"is_check" default:"false"`
	LastDate  string `json:"last_date" default:""`
}

type ApplyListFilterParams struct {
	UserId      string `json:"user_id"`
	PageCount   int    `json:"page_count"`
	PageIndex   int    `json:"page_index"`
	IsCheck     string `json:"is_check"`
	ProcedureId int    `json:"procedure_id"`
}

type ApplyScoreInfo struct {
	Moral    float32 `json:"moral"`
	Practice float32 `json:"practice"`
	Academic float32 `json:"academic"`
	Sum      float32 `json:"sum"`
}

type ApplyHistoryStep struct {
	UserId   string `json:"user_id"`
	EditAt   string `json:"edit_at"`
	Comment  string `json:"comment"`
	Identity string `json:"identity"`
}
type ApplyHistoryData struct {
	Step    ApplyHistoryStep   `json:"step"`
	History []ApplyHistoryStep `json:"history"`
}
