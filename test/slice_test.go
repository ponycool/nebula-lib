package main

import (
	"fmt"
	"github.com/ponycool/nebula-lib/utils"
	"testing"
)

func TestUniqueSlice(t *testing.T) {
	s := []int{1, 1, 2, 3}
	uniqueSlice, err := utils.UniqueSlice(s)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("删除重复值后的新切片为：")
	fmt.Println(uniqueSlice)
}
