package db

import "strconv"

type DelStatus int8

const (
	Undeleted DelStatus = 0
	Deleted   DelStatus = 1
)

func (delStatus DelStatus) String() string {
	return strconv.Itoa(int(delStatus))
}
