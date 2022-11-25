package constants

// DataDeleted 逻辑删除状态：已删除
func DataDeleted() []uint8 {
	return []uint8{1}
}

// DataUndeleted 逻辑删除状态：正常
func DataUndeleted() []uint8 {
	return []uint8{0}
}

// DataLocked 冻结状态：已冻结
func DataLocked() []uint8 {
	return []uint8{1}
}

// DataUnlocked 冻结状态：正常
func DataUnlocked() []uint8 {
	return []uint8{0}
}
