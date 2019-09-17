package day07

import (
	"testing"
	"net/http"
	"io/ioutil"
)

func Test_getHttp(t *testing.T) {
	response, err := http.Get("http://www.sogog.com")
	//http.PostForm("http://url", url.Values{"key": {"Value"}, "id": {"123"}})
	//http.Post("http://url", "image/jpeg", &buf)
	if nil != err {
		t.Error("get baidu.com info fail.")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	t.Log(string(body))
}

func Test_getHttp2(t *testing.T) {
	request, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if nil != err {
		t.Error("get google request fail.")
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if nil != err {
		t.Error("get google response fail.")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	t.Log(string(body))
}