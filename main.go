package main

import (
	"github.com/electricbubble/gadb"
	"log"
	//"os"
	//"strings"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	adbClient, err := gadb.NewClient()
	checkErr(err, "fail to connect adb server")
	devices, err := adbClient.DeviceList()
	checkErr(err)
	fmt.Println(devices)
}

func checkErr(err error, msg ...string) {
	if err == nil {
		return
	}

	var output string
	if len(msg) != 0 {
		output = msg[0] + " "
	}
	output += err.Error()
	log.Fatalln(output)
}