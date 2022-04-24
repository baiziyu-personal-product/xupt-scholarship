package service

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"os"
	"sync"
	"time"
	"xupt-scholarship/global"
)

type EmailService struct {
	Email *email.Email
}

// ProcessStepEnd 流程终止提醒
func (es *EmailService) ProcessStepEnd(processStep string, endDate string) {
	e := es.Email
	e.Subject = time.Now().Format("2006") + "年研究生奖学金评定流程结束的通知"
	e.HTML = []byte(`
		<div style="background-color: #014a95;">
      <div style="width: 100%; height: 100px; padding-top: 3px;">
        <img
          style="width: 398px; height: 94px;"
		src="https://s1.ax1x.com/2022/04/14/L1dUSI.png"
        />
      </div>
      <div style="background-color: #fff; padding-top: 20px;">
        <div
          style="
            font-size: 16px;
            margin: 0 20px;
            border-radius: 4px;
            box-shadow: 1px 1px 2px #ccc;
            padding: 20px 40px;
            border: 1px solid rgb(240, 239, 239);
          "
        >
          <p style="text-indent: 2em;">
            本次研究生奖学金：◣` + processStep + `◥流程将于` + endDate + `结束，请尚未完成当前流程的同学以及管理员抓紧完善当前流程工作内容。
          </p>
          <p style="text-indent: 2em;">
            本通知仅用于通知流程变更以及提醒用户注意当前流程轮转。
          </p>
        </div>
        <p
          style="
            text-align: right;
            color: #aaa;
            margin: 0 20px;
            font-size: 14px;
          "
        >
          ——西安邮电大学大学研究生院
        </p>
      </div>
    </div>
	`)
}

// ProcessStepStart 流程开启通知
func (es *EmailService) ProcessStepStart(processStep string, startDate, endDate string) {
	e := es.Email
	e.Subject = time.Now().Format("2006") + "年研究生奖学金评定流程轮转的通知"
	e.HTML = []byte(`
		<div style="background-color: #014a95;">
      <div style="width: 100%; height: 100px; padding-top: 3px;">
        <img
          style="width: 398px; height: 94px;"
		src="https://s1.ax1x.com/2022/04/14/L1dUSI.png"
        />
      </div>
      <div style="background-color: #fff; padding-top: 20px;">
        <div
          style="
            font-size: 16px;
            margin: 0 20px;
            border-radius: 4px;
            box-shadow: 1px 1px 2px #ccc;
            padding: 20px 40px;
            border: 1px solid rgb(240, 239, 239);
          "
        >

          <p style="text-indent: 2em;">
            研究生奖学金：◣` + processStep + `◥流程，将于` + startDate + `开始，` + endDate + `结束` +
		` 
          </p>
          <p style="text-indent: 2em;">
            本通知仅用于通知流程变更以及提醒用户注意当前流程轮转。
          </p>
        </div>
        <p
          style="
            text-align: right;
            color: #aaa;
            margin: 0 20px;
            font-size: 14px;
          "
        >
          ——西安邮电大学大学研究生院
        </p>
      </div>
    </div>
	`)
}

// ProcessInit  设置发起奖学金评定流程的邮件模板
func (es *EmailService) ProcessInit(startDate string) {
	e := es.Email
	e.Subject = "开启" + time.Now().Format("2006") + "年研究生奖学金评定流程的通知"
	e.HTML = []byte(`
		<div style="background-color: #014a95;">
		  <div style="width:100%;height: 100px;padding-top: 3px">
			  <img style="width: 398px;height: 94px;" src="https://s1.ax1x.com/2022/04/14/L1dUSI.png" />
		  </div>
		  <div style="background-color: #fff;padding-top: 20px;">
			  <h2 style="text-align:center;">《西安邮电大学研究生奖学金流程开始的通知》</h1>
				<p style="text-align: center;font-size: 13px;color: #aaa;">` + startDate + `</p>
            <div style="font-size: 16px;margin:20px;border-radius: 4px;box-shadow: 1px 1px 2px #ccc;padding: 20px 40px;border: 1px solid rgb(240, 239, 239);">
              <p style="text-indent: 2em;">本次奖学金评定流程将于` + startDate + `开始，您已经被添加为本次评定流程的管理人员，负责参与奖学金的评定。</p>
              <p style="text-indent: 2em;">流程包含12个阶段，分别是部署动员阶段，个人申请阶段，各年级政策宣讲，第一次个人自评，第一次班级公示，第二次个人自评，第二次班级公示，年级公示，学科办审核、复查，奖学金评定小组核查、审议，学校评审阶段，结束。 流程顺序执行如下图，当前支持同时执行学科办审核、复查和年级公示两个部分同时进行进行。</p>
              <p style="text-indent: 2em;">流程在开始前会陆续向下一流程的参与人员（管理人员以及研究生）发送邮件通知，并且告知对应的操作方式。 流程在进行至个人申请阶段后开放个人申请奖学金通道（申请奖学金→发起申请），支持研究生按照个人实际情况进行填写和完善对应的申请信息。在其他非自评阶段，不会开启申请奖学金通道。</p>
              <p style="text-indent: 2em;">奖学金审核、复查包括个人自评时在评定流程中流程管理进行审批和处理，对于存疑的信息进行异常处理，通知对应的研究生处理。</p>
              <p style="text-indent: 2em;">在确认创建奖学金评定流程后，流程距离进入第一阶段（部署动员阶段）的24小时之前，可以进行流程的修改和撤销。在流程创建后进行对应的修改和撤销操作，都会被认为当前的操作是一种风险操作。</p>
              <p style="text-indent: 2em;">在本系统的所有流程管理以及对应的通知，都是采用邮件进行通知和反馈，希望对应的参与人员在流程开始后能及时核对个人奖学金信息。 对于存疑的个人信息可进入用户中心进行修改和反馈。</p>
              <p style="text-indent: 2em;">更多详情可登录<a href="">西安邮电大学研究生奖学金官网进行</a>查看。</p>
            </div>
            <p style="text-align: right;margin: 0 20px;font-size: 14px;">——西安邮电大学大学研究生院</p>
      </div>
    </div>
	`)
}

type ProcessBaseEmailOptions struct {
	Step      string   `json:"step"`
	EndDate   string   `json:"end_date"`
	StartDate string   `json:"start_date"`
	Status    string   `json:"status"`
	Receiver  []string `json:"receiver"`
}

// handleProcessEmailByStatus 通过状态处理发送的Process邮件
func handleProcessEmailByStatus(e EmailService, options ProcessBaseEmailOptions) {
	status, step := options.Status, options.Step
	if status == global.ProcessInit {
		e.ProcessInit(options.StartDate)
	} else if status == global.ProcessStart {
		e.ProcessStepStart(step, options.StartDate, options.EndDate)
	} else {
		e.ProcessStepEnd(step, options.StartDate)
	}
}

// sendProcessStatusToReceivers 发送Process 信息邮件
func sendProcessStatusToReceivers(ch chan *email.Email, options ProcessBaseEmailOptions) {
	for _, receiver := range options.Receiver {
		e := EmailService{Email: email.NewEmail()}
		e.Email.From = "西安邮电大学研究生院<" + global.EmailAddress + ">"
		e.Email.To = []string{receiver}
		handleProcessEmailByStatus(e, options)
		ch <- e.Email
	}
}

func SendEmail(options ProcessBaseEmailOptions) {
	ch := make(chan *email.Email, 10)
	p, err := email.NewPool(
		global.EmailHost+":"+global.EmailPort,
		4,
		smtp.PlainAuth("", global.EmailAddress, global.EmailPassword, global.EmailHost),
	)

	if err != nil {
		log.Fatal("failed to create pool:", err)
	}
	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func() {
			defer wg.Done()
			for e := range ch {
				err := p.Send(e, 10*time.Second)
				if err != nil {
					fmt.Fprintf(os.Stderr, "email:%v sent error:%v\n", e, err)
				}
			}
		}()
	}
	sendProcessStatusToReceivers(ch, options)
	close(ch)
	wg.Wait()
}
