// Copyright 2018 morgine.com. All rights reserved.

package log_test

import (
	"github.com/morgine/log"
	"os"
)

func ExampleSetOutput() {

	log.Info.SetFlags(0)

	// 自定义写入接口
	log.SetOutput(log.FlagInfo|log.FlagDanger, os.Stdout)

	// 打印一行 info 信息
	log.Info.Println("hello")

	// Output:
	// [info] hello
}
