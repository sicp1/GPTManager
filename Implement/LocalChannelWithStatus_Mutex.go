package Implement

import "sync"

type LocalChannelWithStatus_Mutex struct {
	ChannelWithStatus []map[string]interface{}
	Mutex             sync.Mutex
}
