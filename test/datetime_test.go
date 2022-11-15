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
