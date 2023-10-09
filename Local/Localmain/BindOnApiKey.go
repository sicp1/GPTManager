package Localmain

import (
	"GPTManager/GlobalData"
	"GPTManager/Implement"
	"GPTManager/Local/Api"
	"time"
)

func BindOnApiKey() {
	ticker_StartApiKey := time.NewTicker(5 * time.Second) //开始apikey

	//"ch_user":   make(chan string, 1),
	//	"ch_system": make(chan string, 1),
	//	"status":    0, //1代表正在连接，0代表准备中，-1代表销毁
	for {
		select {
		case <-ticker_StartApiKey.C:
			for index, item := range GlobalData.LocalChannelWithStatus.ChannelWithStatus {
				//1代表正在连接，0代表准备中，-1代表销毁
				if item["status"] == 0 {
					apikey, flag := GlobalData.ApiKeyQueue_Local.Pop()
					if flag == -1 {
						break
					}
					if flag == 0 {
						println("最后一个")
						go Api.ChatGPTMain(index, Implement.InterfaceToApiKey(apikey))
					}
					if flag == 1 {
						go Api.ChatGPTMain(index, Implement.InterfaceToApiKey(apikey))
					}
					println("状态从准备到正在连接")
					item["status"] = 1
				} else if item["status"] == 1 {
					println("正在连接")
				} else if item["status"] == -1 {
					println("已销毁")
					//销毁

					//销毁
				}

			}

		}
	}
}
