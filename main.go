package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type cmdParams struct {
	logFilePath string
	routimeNum  int
}

type digData struct {
	time string
	url  string
}

type urlData struct {
	data digData
	uid  string
}

type storageBlock struct {
	counterType string
}

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	// 获取参数, 分别是 log的路径， 最大的并发度。 最后l是 写到本地的路径。 最后记得 调用 flag.Parse()
	logFilepath := flag.String("logFilePath", "D:/Develop/GO_Develop/src/github.com/zxccl0518/GoRoutinePractice/log/dig.log", "logFilePath")
	routineNum := flag.Int("routineNum", 5, "consumer number by goroutine")
	l := flag.String("l", "D:/Develop/GO_Develop/src/github.com/zxccl0518/GoRoutinePractice/log/dig.log", "this prgrame runtime log target file path")
	flag.Parse()

	fmt.Println("---> l = ", *l)

	params := cmdParams{logFilePath: *logFilepath, routimeNum: *routineNum}

	// 打日志
	logFd, err := os.OpenFile(*l, os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		log.Out = logFd
		defer logFd.Close()
	}
	log.Infoln("Exec start.")
	log.Infof("params : logFilePath = %s, routineNum = %d", params.logFilePath, params.routimeNum)
	// 初始化一些channel
	var logChannel = make(chan string, 3*(params.routimeNum))
	var pvChannel = make(chan urlData, params.routimeNum)
	var uvChannel = make(chan urlData, params.routimeNum)
	var storageChannel = make(chan storageBlock, params.routimeNum)

	// 日至消费者。
	go readFileLineByLine(params, logChannel)

	// 创建一组日志处理
	for i := 0; i < params.routimeNum; i++ {
		go logConsumer(logChannel, pvChannel, uvChannel)
	}

	// 创建能见pv uv 统计器
	go pvCounter(pvChannel, storageChannel)
	go uvCounter(uvChannel, storageChannel)

	// 创建存储器。
	go dataStorage(storageChannel)

	time.Sleep(time.Second * 10)
}

func readFileLineByLine(params cmdParams, logChannel chan string) {

}

func logConsumer(logChannel chan string, pvChannel chan urlData, uvChannel chan urlData) {

}

func pvCounter(pvChannel chan urlData, storageChannel chan storageBlock) {

}

func uvCounter(uvChannel chan urlData, storageChannel chan storageBlock) {

}

func dataStorage(storageChannel chan storageBlock) {

}
