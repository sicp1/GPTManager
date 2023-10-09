package Api

import (
	"GPTManager/GlobalData"
	"GPTManager/Implement"
	"GPTManager/Local/Atom"
	"GPTManager/Tool"
)

// "ch_return": make(chan string,1),
// "ch_user":   make(chan string, 1),
// "ch_system": make(chan string, 1),
// "status":    0, //1代表正在连接，0代表准备中，-1代表销毁
func ChatGPTMain(channel_index int, ApiKey Implement.ApiKey) {
	for {
		ch_return := Tool.InterfaceToChannel_string(GlobalData.LocalChannelWithStatus.ChannelWithStatus[channel_index]["ch_return"])
		ch_user := Tool.InterfaceToChannel_string(GlobalData.LocalChannelWithStatus.ChannelWithStatus[channel_index]["ch_user"])
		ch_system := Tool.InterfaceToChannel_string(GlobalData.LocalChannelWithStatus.ChannelWithStatus[channel_index]["ch_system"])
		user := <-ch_user
		system := <-ch_system
		receive_data := Atom.ChatGPT_FullProcess(ApiKey.ApiKeyNum, system, user)
		ch_return <- receive_data
		println(receive_data)

	}
}
