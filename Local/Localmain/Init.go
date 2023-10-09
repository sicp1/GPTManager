package Localmain

import (
	"GPTManager/GlobalData"
	"GPTManager/GlobalVar"
	"GPTManager/Tool"
)

func Init() {
	GlobalData.ApiKeyQueue_Local.Data = []interface{}{}
	Tool.GET(GlobalVar.Ip_Port_main + "/LocalIpPort?port=" + Tool.IntToStr(GlobalVar.Port))
	go BindOnApiKey()

}
