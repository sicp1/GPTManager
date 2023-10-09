package Implement

import (
	"GPTManager/GlobalVar"
	"bytes"
	"encoding/json"
)

//	type ChatGPT_SendData_Inter interface {
//		Default()
//		AddMessage(message string)
//		AddSystem(system string)
//		Clean()
//	 ToBufferData()  *bytes.Buffer
//	}
//
//	type ChatGPT_SendData struct {
//		Model       string                  `json:"model"`
//		Messages    []ChatGPT_SendData_item `json:"messages"`
//		Temperature float64                 `json:"temperature"`
//	}
//
//	type ChatGPT_SendData_item struct {
//		Role    string `json:"role"`
//		Content string `json:"content"`
//	}
func (chatgpt_senddata *ChatGPT_SendData) ToBufferData() *bytes.Buffer {
	jsonStr, _ := json.Marshal(chatgpt_senddata)
	var jsonstr = []byte(jsonStr) //转换二进制
	buffer := bytes.NewBuffer(jsonstr)
	return buffer
}
func (chatgpt_senddata *ChatGPT_SendData) Default() {
	chatgpt_senddata.Model = GlobalVar.ChatGPT_Model
	chatgpt_senddata.Temperature = GlobalVar.ChatGPT_Temperature
	chatgpt_senddata.Messages = []ChatGPT_SendData_item{}
	println("chatgpt数据创建成功")
}
func (chatgpt_senddata *ChatGPT_SendData) AddMessage(message string) {
	chatgpt_senddata_item := ChatGPT_SendData_item{
		Role:    "user",
		Content: message,
	}
	chatgpt_senddata.Messages = append(chatgpt_senddata.Messages, chatgpt_senddata_item)
	println("user添加成功")
}
func (chatgpt_senddata *ChatGPT_SendData) AddSystem(system string) {
	chatgpt_senddata_item := ChatGPT_SendData_item{
		Role:    "system",
		Content: system,
	}
	chatgpt_senddata.Messages = append(chatgpt_senddata.Messages, chatgpt_senddata_item)
	println("system添加成功")
}
func (chatgpt_senddata *ChatGPT_SendData) Clean() {
	chatgpt_senddata.Messages = []ChatGPT_SendData_item{}
	println("清理成功")
}
