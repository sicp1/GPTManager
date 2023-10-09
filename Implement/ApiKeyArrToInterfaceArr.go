package Implement

func ApiKeyArrToInterfaceArr(ApiKeyArr []ApiKey) []interface{} {
	interface_arr := make([]interface{}, len(ApiKeyArr))
	for i, apikey := range ApiKeyArr {
		interface_arr[i] = apikey
	}
	return interface_arr

}
