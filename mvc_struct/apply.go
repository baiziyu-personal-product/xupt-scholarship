package mvc_struct

type moralFormItem struct {
	Level string   `json:"level"`
	Name  string   `json:"name"`
	Info  string   `json:"info"`
	Time  string   `json:"time"`
	Files []string `json:"files"`
}

type moralFormValue struct {
	List []moralFormItem `json:"list"`
}

type practiceFormResultItem struct {
	Level    string   `json:"level"`
	Time     string   `json:"time"`
	Name     string   `json:"name"`
	Order    int      `json:"order"`
	Partners int      `json:"partners"`
	Files    []string `json:"files"`
}

type socialCadreItem struct {
	Level      []string `json:"level"`
	Department string   `json:"department"`
}

type socialActivityItem struct {
	Level string   `json:"level"`
	Name  string   `json:"name"`
	Time  []string `json:"time"`
	Files []string `json:"files"`
}

type practiceSocialFormValue struct {
	Cadre    []socialCadreItem    `json:"cadre"`
	Activity []socialActivityItem `json:"activity"`
}

type practiceCompetitionFormValue struct {
	Level []string `json:"level"`
	Name  string   `json:"name"`
	Time  string   `json:"time"`
	Files []string `json:"files"`
}

type practiceFormValue struct {
	Result      []practiceFormResultItem       `json:"result"`
	Social      practiceSocialFormValue        `json:"social"`
	Competition []practiceCompetitionFormValue `json:"competition"`
}

type academicAwardFormItem struct {
	Level []string `json:"level"`
	Name  string   `json:"name"`
	Time  string   `json:"time"`
	Files []string `json:"files"`
}

type academicScientificFormItem struct {
	Level                 string   `json:"level"`
	Name                  string   `json:"name"`
	Time                  []string `json:"time"`
	FundsActuallyReceived int      `json:"funds_actually_received"`
	FundsDue              int      `json:"funds_due"`
	Distribute            int      `json:"distribute"`
	Files                 []string `json:"files"`
}

type academicDissertationFormItem struct {
	Level    string   `json:"level"`
	Name     string   `json:"name"`
	IdNumber string   `json:"id_number"`
	Time     string   `json:"time"`
	Files    []string `json:"files"`
}

type academicPublishFormItem struct {
	Level            string   `json:"level"`
	Name             string   `json:"name"`
	Time             string   `json:"time"`
	PublishHouseName string   `json:"publish_house_name"`
	FontsCount       int      `json:"fonts_count"`
	Files            []string `json:"files"`
}

type academicFormValue struct {
	Award        []academicAwardFormItem        `json:"award"`
	Scientific   []academicScientificFormItem   `json:"scientific"`
	Dissertation []academicDissertationFormItem `json:"dissertation"`
	Publish      []academicPublishFormItem      `json:"publish"`
}

type ApplicationValue struct {
	Moral    moralFormValue    `json:"moral"`
	Practice practiceFormValue `json:"practice"`
	Academic academicFormValue `json:"academic"`
}

type BaseApply struct {
	Form      interface{} `json:"form"`
	StudentId string      `json:"student_id"`
	Type      string      `json:"type"`
}
