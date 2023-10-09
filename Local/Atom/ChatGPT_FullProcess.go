package Atom

import "GPTManager/Implement"

func ChatGPT_FullProcess(apikey string, system string, user string) string {
	chatgpt := &Implement.ChatGPTItem{}
	chatgpt.Default(apikey)
	chatgpt_senddata := CreateChatGPT_SendData_system_user(system, user)
	bufferdata := chatgpt_senddata.ToBufferData()
	chatgpt.SetData(bufferdata)
	chatgpt.Start()
	return chatgpt.GetReceiveData()
}
