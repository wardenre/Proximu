package main

import (
	"Proximu/proxy"
	"Proximu/utils"
)

func main() {
	utils.InitLogger()
	proxy.Start(":19132", "127.0.0.1:19133")
}
