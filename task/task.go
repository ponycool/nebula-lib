package task

import (
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sync"
)

var taskQueue []TFunc

type Task struct {
	// Closed chan bool
	Wg     sync.WaitGroup
	Logger *zap.Logger
}

type TFunc struct {
	Func func(task *Task)
}

// AddFunc 新增运行方法
func (t *Task) AddFunc(f TFunc) {
	taskQueue = append(taskQueue, f)
}

// Run 运行所有任务
func (t *Task) Run() {
	count := len(taskQueue)
	t.Wg.Add(count)
	t.Logger.Info("[task] start task...")

	// 运行任务
	for i := range taskQueue {
		go taskQueue[i].Func(t)
	}

	//for {
	//	select {
	//	case <-t.Closed:
	//		t.Logger.Info("[task] close...")
	//		return
	//	}
	//}
}

// Stop 优雅退出所有任务
func (t *Task) Stop() {
	t.Logger.Info("[task] got close sig...")
	// 等待所有协程退出
	t.Wg.Wait()
	t.Logger.Info("[task] all goroutine done...")
}

// RunListener 启动系统监听退出信号
func (t *Task) RunListener() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	// 监听`Ctrl+C`消息，停止所有任务
	select {
	case sig := <-c:
		t.Logger.Info("[task] got signal. Aborting...",
			zap.Any("sig", sig),
		)
		t.Stop()
	}
}
