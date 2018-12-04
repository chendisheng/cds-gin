package config

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"log"
	"os"
	"path/filepath"
)

var RunMode string
var AppConfig config.Configer

func init(){
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var filename = "app.conf"
	appConfigPath := filepath.Join(workPath, "conf", filename)
	config, err := config.NewConfig("ini", appConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	AppConfig = config
	RunMode = config.String("runmode")
	fmt.Println("RunMode")

}

//获取string 类型配置
func String(key string)string{
	return AppConfig.String(RunMode+"::"+key)
}
