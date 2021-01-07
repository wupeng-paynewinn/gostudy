package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	_ "strings"
)

func main() {

	urlRequest := "http://127.0.0.1:9090/xxx"
	urlObj, err := url.Parse(urlRequest)
	if err != nil {
		fmt.Println(err)
		return
	}
	method := "POST"

	data := url.Values{}
	data.Set("name", "张家界")
	data.Set("age", "185")
	queryStr := data.Encode()
	fmt.Println(queryStr)

	urlObj.RawQuery = queryStr

	//payload := strings.NewReader("name=值紧机&age=185")
	//
	//client := &http.Client {
	//}
	//req, err := http.NewRequest(method, urlRequest, payload)

	req, err := http.NewRequest(method, urlObj.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "*/*")

	//res, err := client.Do(req)
	res, err := http.DefaultClient.Do(req)

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
