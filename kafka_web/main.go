package main

import (
	_ "kafka_web/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

