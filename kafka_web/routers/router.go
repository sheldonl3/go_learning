package routers

import (
	"kafka_web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/kafka_web", &controllers.KafkaController{})
}