package GlobalData

import (
	"GPTManager/Implement"
	"GPTManager/Struct"
)

var LocalServerList = []string{}
var LocalServerStatusList = map[string][]string{}
var ApiKeyQueue_Main = &Struct.Queue{}
var ApiKeyQueue_Local = &Struct.Queue{}
var ApiKey_WithStatus = []Implement.ApiKey{}
var LocalChannelWithStatus Implement.LocalChannelWithStatus_Mutex

var System string
var Prompt string
