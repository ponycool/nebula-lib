package main

import (
	"fmt"
	"github.com/ponycool/nebula-lib/job"
	"github.com/ponycool/nebula-lib/log"
	"github.com/ponycool/nebula-lib/task"
	"testing"
)

func testCronTask(t *task.Task) {
	count := 1
	j := job.Job{
		Callback: func() (err error) {
			if count > 5 {
				defer t.Wg.Done()
				t.Closed <- true
			}
			fmt.Println(fmt.Sprintf("启动一个定时Task任务，这是第%d次执行", count))
			count++
			return nil
		},
		Logger: log.Get(),
	}

	go j.Start("@every 2s", j)
}

func testTask(t *task.Task) {
	defer t.Wg.Done()
	fmt.Println("启动一个Task任务")
}

func TestTask(t *testing.T) {
	t.Helper()

	fmt.Println("======== Task 运行测试 ========")

	tk := task.Task{
		Logger: log.Get(),
	}

	t1 := task.TFunc{
		Func: testCronTask,
	}
	t2 := task.TFunc{
		Func: testTask,
	}

	tk.AddFunc(t1)
	tk.AddFunc(t2)
	tk.Run()
	tk.RunListener()
}
