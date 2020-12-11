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
	MaxConnect  int
}

var GlobalObj GlobalObject

func init() {
	path, _ := os.Getwd()
	log.Println("Ready Read Json, Now Work Path : ", path)

	dirs := [...]string{"myZinx", "awesomeProject/myZinx", "../.."}
	var context []byte
	var err error
	for _, dir := range dirs {
		context, err = ioutil.ReadFile(dir + "/conf/zinx.json")
		if err != nil {
			log.Println("Cannot Read ", err)
			continue
		} else {
			log.Println("Read Success! Read From ", dir)
			break
		}
	}

	// default value
	GlobalObj.MaxReadSize = 2048
	GlobalObj.MaxConnect = 512
	GlobalObj.IpVersion = "tcp4"
	GlobalObj.Ip = "0.0.0.0"
	GlobalObj.Port = 8999
	GlobalObj.Name = "myZinx"

	err = json.Unmarshal(context, &GlobalObj)
	if err != nil {
		log.Println("conf.zinx.json Error", err)
	}
}
