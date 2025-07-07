package main

import (
	"fmt"
	"github.com/ponycool/nebula-lib/job"
	"github.com/ponycool/nebula-lib/log"
	"testing"
)

var times = 1

func testJob() error {
	fmt.Println("启动了一个任务")
	return nil
}

func testCronJob() error {
	fmt.Println(fmt.Sprintf("启动了一个定时任务，这是第%d次调用", times))
	times++
	return nil
}

func TestJobRun(t *testing.T) {
	t.Helper()

	// 暂不执行测试
	t.Skip()

	logInit()

	fmt.Println("========== 任务运行测试 ============")

	j := job.Job{
		Callback: testJob,
		Logger:   log.Get(),
	}

	j.Run()
}

func TestJobStart(t *testing.T) {
	t.Helper()

	// 暂不执行测试
	t.Skip()

	logInit()

	fmt.Println("========== 定时任务运行测试 ===========")

	j := job.Job{
		Callback: testCronJob,
		Logger:   log.Get(),
	}

	go j.Start("@every 1s", j)
}
