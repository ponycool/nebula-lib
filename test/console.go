package main

import (
	"flag"
	"fmt"
	"github.com/ponycool/nebula-lib/command"
	"os"
	"time"
)

// initFs 初始化命令行参数
func initFs() {

	// 初始化命令参数
	fs := flag.NewFlagSet("task", flag.ExitOnError)
	var (
		DBDriver   = fs.String("DBDriver", "mariadb", "DataBase Driver, default mariadb")
		MQLifetime = fs.Duration("MQLifetime", 0*time.Second, "lifetime of process before shutdown (0s=infinite)")
	)
	fs.Usage = command.Usage(fs, os.Args[0]+" [flags]")
	_ = fs.Parse(os.Args[1:])

	fmt.Println(fmt.Sprintf("%v", *DBDriver))
	fmt.Println(MQLifetime)
}
