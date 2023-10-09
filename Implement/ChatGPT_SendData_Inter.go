package Implement

import "bytes"

type ChatGPT_SendData_Inter interface {
	Default()
	AddMessage(message string)
	AddSystem(system string)
	Clean()
	ToBufferData() *bytes.Buffer
}
