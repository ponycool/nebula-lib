package main

import (
	"fmt"
	"github.com/ponycool/nebula-lib/datetime"
	"testing"
)

func TestTimeStringToTime(t *testing.T) {
	t.Helper()

	timeStr := "2014-10-01T00:00:00+08:00"
	time := datetime.Time{}.TimeStringToTime(timeStr)
	fmt.Println(time)
}

func TestGetNowSec(t *testing.T) {
	timestamp := datetime.GetNowSec()
	fmt.Println(fmt.Sprintf("当前Unix时间戳为：%d，精确到秒", timestamp))
}

func TestGetNowMs(t *testing.T) {
	timestamp := datetime.GetNowMs()
	fmt.Println(fmt.Sprintf("当前Unix时间戳为：%d，精确到毫秒", timestamp))
}

func TestGetNowUs(t *testing.T) {
	timestamp := datetime.GetNowUs()
	fmt.Println(fmt.Sprintf("当前Unix时间戳为：%d，精确到微妙", timestamp))
}

func TestGetNowNs(t *testing.T) {
	timestamp := datetime.GetNowNs()
	fmt.Println(fmt.Sprintf("当前Unix时间戳为：%d，精确到纳秒", timestamp))
}
