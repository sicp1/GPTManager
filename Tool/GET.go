package Tool

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GET(url string) map[string]string {
	var result map[string]string
	response, _ := http.Get(url)

	reader := response.Body

	defer reader.Close()
	b, _ := ioutil.ReadAll(reader)
	json.Unmarshal([]byte(string(b)), &result)
	return result
}
