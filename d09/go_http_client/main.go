package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func getQueryKey() {
	defer wg.Done()
	url := "https://speech.10jqka.com.cn/info/api/v2/auth/get/queryAESKey"
	method := "POST"

	payload := strings.NewReader("userId=-1&account=18557531475")
	//client := &http.Client{}
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "*/*")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func main() {
	start := time.Now()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go getQueryKey()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
