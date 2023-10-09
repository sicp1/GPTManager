package Implement

type ChatGPT_SendData struct {
	Model       string                  `json:"model"`
	Messages    []ChatGPT_SendData_item `json:"messages"`
	Temperature float64                 `json:"temperature"`
}
type ChatGPT_SendData_item struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
