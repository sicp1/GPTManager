package Tool

func StrArrToInterfaceArr(str_arr []string) []interface{} {
	interface_arr := make([]interface{}, len(str_arr))
	for i, str := range str_arr {
		interface_arr[i] = str
	}
	return interface_arr
}
