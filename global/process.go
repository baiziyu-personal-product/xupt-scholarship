package global

var ProcessStepList = []string{
	"部署动员阶段",
	"个人申请阶段",
	"各年级政策宣讲",
	"第一次个人自评",
	"第一次班级公示",
	"第二次个人自评",
	"第二次班级公示",
	"年级公示",
	"学科办审核、复查",
	"奖学金评定小组核查、审议",
	"学校评审阶段",
	"结束",
}

// ProcessStep 状态
const (
	ProcessStart = "start"
	ProcessEnd   = "end"
	ProcessInit  = "init"
)

// Process 流程轮转距离正常开始的时间
const (
	ProcessInitDurationHours  = 24
	ProcessStartDurationHours = 8
)
