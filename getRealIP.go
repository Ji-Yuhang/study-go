package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getRealIP() {
	url := "https://api.ipify.org?format=json"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	fmt.Printf("%s\n", body)

	type RealIpResponse struct {
		RealIp string `json:"ip"`
	}

	//https://blog.csdn.net/zxy_666/article/details/80173288
	var realIpresponse RealIpResponse
	err = json.Unmarshal(body, &realIpresponse)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(realIpresponse.RealIp)

}
