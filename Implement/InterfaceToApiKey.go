package Implement

func InterfaceToApiKey(inter interface{}) ApiKey {
	apikey, _ := inter.(ApiKey)
	return apikey
}
