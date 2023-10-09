package Implement

import (
	"GPTManager/GlobalVar"
	"GPTManager/Tool"
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
)

//type ChatGPT interface {
//	Default(auth string)
//	Start()
//	GetReceiveData() string
//	SetData(BufferData *bytes.Buffer)
//}

//	type ChatGPTItem struct {
//		GPT_Api_URL   string
//		ContentType   string
//		Authorization string
//		ReceiveData   string
//		BufferData    *bytes.Buffer
//	}
func (chatgpt_item *ChatGPTItem) SetData(BufferData *bytes.Buffer) {
	chatgpt_item.BufferData = BufferData
}
func (chatgpt_item *ChatGPTItem) GetReceiveData() string {
	return chatgpt_item.ReceiveData
}
func (chatgpt_item *ChatGPTItem) Start() {
	ch := make(chan string, 1)
	go func() {
		chatgpt_item.MUTEX.Lock()
		request, _ := http.NewRequest("POST", chatgpt_item.GPT_Api_URL, chatgpt_item.BufferData)
		request.Header.Set("Content-Type", chatgpt_item.ContentType) //添加请求头
		request.Header.Set("Authorization", chatgpt_item.Authorization)
		client := http.Client{}                                   //创建客户端
		resp, _ := client.Do(request.WithContext(context.TODO())) //发送请求
		respBytes, _ := ioutil.ReadAll(resp.Body)
		ch <- string(respBytes)
		chatgpt_item.MUTEX.Unlock()
	}()
	ret_str := <-ch
	chatgpt_item.ReceiveData = ret_str
	println("请求完成")
}
func (chatgpt_item *ChatGPTItem) Default(api_key string) {
	chatgpt_item.Authorization = Tool.CreateAuth(api_key)
	chatgpt_item.ContentType = GlobalVar.ChatGPT_ContentType
	chatgpt_item.GPT_Api_URL = GlobalVar.ChatGPT_API_URL
	println("创建Chatgpt item成功")

}
