package Implement

import "bytes"

type ChatGPT interface {
	Default(api_key string)
	Start()
	GetReceiveData() string
	SetData(BufferData *bytes.Buffer)
}
