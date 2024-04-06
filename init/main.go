package main

import (
	"flag"
	"fmt"
	"goserver/init/cmd"
	"net/http"
	"os"
)

// 현재경로
var pwd, _ = os.Getwd()

// 설정파일 플래그 지정
var configPathFlag = flag.String("config", fmt.Sprintf("%s/init/config.toml", pwd), "config file not found")

func main() {
	//http.HandleFunc("/", helloWorld)
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	fmt.Println("에러 발생")
	//	panic(err)
	//}

	flag.Parse()
	cmd.NewCmd(*configPathFlag)
}

func helloWorld(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("hello world")
}
