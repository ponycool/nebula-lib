package job

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"os"
	"os/signal"
)

type Job struct {
	Callback func() (err error)
	JobID    cron.EntryID
	Done     chan bool
	Logger   *zap.Logger
}

func (j Job) Run() {
	if j.Callback != nil {
		err := j.Callback()
		if err != nil {
			j.Logger.Error("job error", zap.Any("error", err))
		}
	}
}

// Start 开始定时计划任务 -spec 传入cron时间设置 -job 执行的任务
func (j Job) Start(spec string, job Job) {
	logger := &DefaultLog{
		logger: j.Logger,
	}

	c := cron.New(cron.WithChain(cron.SkipIfStillRunning(logger)))
	id, err := c.AddJob(spec, &job)
	if err != nil {
		return
	}

	j.JobID = id

	// 启动执行任务
	c.Start()
	// 退出时关闭计划任务
	defer c.Stop()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	select {
	case <-ch:
		return
	case <-j.Done:
		c.Remove(j.JobID)
	}
}
