package Reader

import (
	"GPTManager/GlobalVar"
	"gopkg.in/ini.v1"
	"os"
)

// [default]
// type=slave
// ChatGPT_API_URL = https://api.openai.com/v1/chat/completions
// ChatGPT_ContentType = application/json;charset=UTF-8
// ChatGPT_Model = gpt-3.5-turbo
// ChatGPT_Temperature = 0.7
// [slave]
// Port = 8080
// Ip_Port_main =
// [Localmain]
// Main_Port = 8080
func Reader() {
	file, e := ini.Load("Manager.ini")
	if e != nil {
		println(e.Error())
		os.Exit(1)
	}
	GlobalVar.Type = file.Section("default").Key("type").String()
	GlobalVar.ChatGPT_API_URL = file.Section("default").Key("ChatGPT_API_URL").String()
	GlobalVar.ChatGPT_ContentType = file.Section("default").Key("ChatGPT_ContentType").String()
	GlobalVar.ChatGPT_Model = file.Section("default").Key("ChatGPT_Model").String()
	GlobalVar.ChatGPT_Temperature, _ = file.Section("default").Key("ChatGPT_Temperature").Float64()
	if GlobalVar.Type == "slave" {
		GlobalVar.Port, _ = file.Section("slave").Key("Port").Int()
		GlobalVar.Ip_Port_main = file.Section("slave").Key("Ip_Port_main").String()
		//Port = 8080
		//Ip_Port_main =
		println("slave_Init")
	}
	if GlobalVar.Type == "Localmain" {
		//Main_Port = 8080
		println("main_Init")
		GlobalVar.Main_Port, _ = file.Section("Localmain").Key("Main_Port").Int()

	}

}
