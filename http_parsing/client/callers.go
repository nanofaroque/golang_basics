package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"../models"
)

func main() {
	var res models.Response
	url := "https://api.adviceslip.com/advice"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &res)
	fmt.Println("Advice: ", res.Slip.Advice)
	fmt.Println("Advice ID: ", res.Slip.Slip_ID)

}
