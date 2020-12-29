package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpGet(count int) {
	for i := 0; i < count; i++ {
		resp, err := http.Get("http://www.baidu.com")
		if err != nil {
			// handle error
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
			return
		}
		fmt.Println("========================================================", i)
		fmt.Println(string(body))
		resp.Body.Close()
	}

}

func main() {
	for i := 0; i < 5; i++ {
		go httpGet(1)
	}
}
