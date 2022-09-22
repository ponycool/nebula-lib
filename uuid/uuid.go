package uuid

import uid "github.com/satori/go.uuid"

// GetUuid 获取UUID
func GetUuid() (uuid string) {
	return uid.NewV5(uid.NewV4(), "uuid").String()
}

// Check 检查输入的字符串是否为UUID格式
func Check(uuid string) bool {
	_, err := uid.FromString(uuid)
	if err != nil {
		return false
	}
	return true
}
