package datetime

import "time"

type Time struct {
}

var timeTemplates = []string{
	"2006-01-02T00:00:00+08:00",
	"2006-01-02 15:04:05",
	"2006/01/02 15:04:05",
	"2006-01-02",
	"2006/01/02",
	"15:04:05",
}

// TimeStringToTime 时间格式字符串转换
func (t Time) TimeStringToTime(timeStr string) time.Time {
	for i := range timeTemplates {
		t, err := time.ParseInLocation(timeTemplates[i], timeStr, time.Local)
		if nil == err && !t.IsZero() {
			return t
		}
	}
	return time.Time{}
}

// GetNowSec 获取Unix时间戳，精确到秒
func GetNowSec() int64 {
	return time.Now().Unix()
}

// GetNowMs 获取Unix时间戳，精确到毫秒
func GetNowMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// GetNowUs 获取Unix时间戳，精确到微妙
func GetNowUs() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

// GetNowNs 获取Unix时间戳，精确到纳秒
func GetNowNs() int64 {
	return time.Now().UnixNano()
}
