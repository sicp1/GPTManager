package main

import (
	"GPTManager/GlobalVar"
	"GPTManager/Local/Localmain"
	"GPTManager/Main/Mainmain"
	"GPTManager/Reader"
)

func main() {
	Reader.Reader()
	if GlobalVar.Type == "slave" {
		Localmain.LocalServer()

	}
	if GlobalVar.Type == "main" {
		Mainmain.MainServer()
	}
}
