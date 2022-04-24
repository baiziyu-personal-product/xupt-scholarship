package mvc_struct

type ProcessStepValue struct {
	Step     string   `json:"step"`
	Desc     string   `json:"desc"`
	Date     []string `json:"date"`
	Mentions []string `json:"mentions"`
}

type ProcessStepData []ProcessStepValue

type uploadFilesField struct {
	files []UploadFileItem `json:"files"`
}

type ProcessFormData struct {
	Form   ProcessStepData  `json:"form"`
	Upload uploadFilesField `json:"upload"`
}

type ProcessHistoryItem struct {
	StartAt string `json:"start_at"`
	Step    string `json:"step"`
}

type ProcessStepSchedule struct {
	Name       string   `json:"name"`
	Step       string   `json:"step"`
	Duration   float64  `json:"duration"`
	NotifyList []string `json:"notify_list"`
	Date       []string `json:"date"`
	Status     string   `json:"status" default:"start"`
}
