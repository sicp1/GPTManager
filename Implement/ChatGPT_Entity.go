package Implement

import (
	"bytes"
	"sync"
)

type ChatGPTItem struct {
	GPT_Api_URL   string
	ContentType   string
	Authorization string
	ReceiveData   string
	BufferData    *bytes.Buffer
	MUTEX         sync.Mutex
}
