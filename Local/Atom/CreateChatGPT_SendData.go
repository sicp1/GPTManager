package Atom

import (
	"GPTManager/Implement"
)

//type ChatGPT_SendData struct {
//	Model       string        `json:"model"`
//	Messages    []ChatGPT_SendData_item `json:"messages"`
//	Temperature float64       `json:"temperature"`
//}
//type ChatGPT_SendData_item struct {
//	Role    string `json:"role"`
//	Content string `json:"content"`
//}

func CreateChatGPT_SendData_system_user(system string, user string) *Implement.ChatGPT_SendData {
	chatgpt_senddata := &Implement.ChatGPT_SendData{}
	chatgpt_senddata.Default()
	chatgpt_senddata.AddSystem(system)
	chatgpt_senddata.AddMessage(user)
	return chatgpt_senddata
}

func CreateChatGPT_SendData_user(user string) *Implement.ChatGPT_SendData {
	chatgpt_senddata := &Implement.ChatGPT_SendData{}
	chatgpt_senddata.Default()
	chatgpt_senddata.AddMessage(user)
	return chatgpt_senddata
}
