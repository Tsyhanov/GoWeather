package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apikey string = "secret token"

type reqFlags struct {
	city string
	days int
}

type respData struct {
	Weather []struct {
		Status string `json:"main"`
		Desc   string `json:"description"`
	} `json:"weather"`
	System struct {
		Id      int    `json:"id"`
		Country string `json:"country"`
	} `json:"sys"`
	City string `json:"name"`
}

func main() {

	c := flag.String("city", "Kharkiv", "Name of city")
	d := flag.Int("d", 0, "Number of forecast days otherwise 0")
	flag.Parse()
	fmt.Println(*c)
	fmt.Println(*d)
	fmt.Println(apikey)

	sendRequest()
	//http.HandleFunc("/", mainHandler)
	//http.ListenAndServe(":8080", nil)

}

func sendRequest() {
	fmt.Println("hello from sendRequest")
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=London,uk&APPID=" + apikey)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var r respData
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println("Something goes wrong")
	}
	fmt.Println(r)
}
