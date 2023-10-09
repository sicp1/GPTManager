package Tool

import "fmt"

func InterfaceToStr(inter interface{}) string {
	str := fmt.Sprintf("%v", inter)
	return str
}
