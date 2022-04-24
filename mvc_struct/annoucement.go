package mvc_struct

type AnnouncementData struct {
	Id           int     `json:"id"`
	Moral        float32 `json:"moral"`
	Academic     float32 `json:"academic"`
	Practice     float32 `json:"practice"`
	Score        float32 `json:"score"`
	Name         string  `json:"name"`
	StudentId    string  `json:"student_id"`
	College      string  `json:"college"`
	Professional string  `json:"professional"`
	Grade        string  `json:"grade"`
	Class        int     `json:"class"`
	CourseCredit float32 `json:"course_credit"`
	ShipType     string  `json:"ship_type"`
}
