package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadFile("./xx.html")
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	_, _ = w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	bodyObj, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(bodyObj))
	_, _ = w.Write([]byte("ok"))
}

// net/http
func main() {
	http.HandleFunc("/test", f1)
	http.HandleFunc("/xxx", f2)
	_ = http.ListenAndServe("127.0.0.1:9090", nil)
}
