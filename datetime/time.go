package datetime

import "time"

type Time struct {
}

var timeTemplates = []string{
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
