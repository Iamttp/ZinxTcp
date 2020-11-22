package untils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type GlobalObject struct {
	Name        string
	Ip          string
	Port        int
	IpVersion   string
	MaxReadSize int
}

var GlobalObj GlobalObject

func init() {
	path, _ := os.Getwd()
	log.Println("Ready Read Json, Now Work Path : ", path)

	context, err := ioutil.ReadFile("myZinx/conf/zinx.json")
	if err != nil {
		log.Println("Cannot Read conf/zinx.json", err)
	}

	// default value
	GlobalObj.MaxReadSize = 512
	GlobalObj.IpVersion = "tcp4"
	GlobalObj.Ip = "0.0.0.0"
	GlobalObj.Port = 8999
	GlobalObj.Name = "myZinx"

	err = json.Unmarshal(context, &GlobalObj)
	if err != nil {
		log.Println("conf.zinx.json Error", err)
	}
}
