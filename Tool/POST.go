package Tool

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

func POST_FormData(formdata map[string]string, Url string) string {
	formValues := url.Values{}
	for k, v := range formdata {
		println(k)
		println(v)
		formValues.Set(k, v)
	}

	formDataStr := formValues.Encode()
	formDataBytes := []byte(formDataStr)
	formBytesReader := bytes.NewReader(formDataBytes)

	//生成post请求
	client := &http.Client{}
	req, _ := http.NewRequest("POST", Url, formBytesReader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	respBytes, _ := ioutil.ReadAll(resp.Body)
	return string(respBytes)
}
