package main

import (
	"flag"
	"fmt"
	"goserver/init/cmd"
	"os"
)

// 현재경로
var pwd, _ = os.Getwd()

// 설정파일 플래그 지정
var configPathFlag = flag.String("config", fmt.Sprintf("%s/init/config.toml", pwd), "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPathFlag)
}
