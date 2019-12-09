package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	_ "net/http/pprof"
	"runtime/debug"
	"strings"
)

func main() {
	//分析器启动
	//go pprof()
	id := uuid.NewV4()
	fmt.Println(strings.Replace(id.String(), "-", "", -1))
}

func pprof() {
	// 关闭GC
	debug.SetGCPercent(-1)

}
