package Tool

func InterfaceToChannel_string(inter interface{}) chan string {
	channel, _ := inter.(chan string)
	return channel
}
