package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const IP_GEOLOCATION = "http://ip-api.com/json"

type Geolocation struct {
	Status      string `json:"status"`
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
	RegionName  string `json:"regionName"`
	City        string `json:"city"`
	Zip         string `json:"zip"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	Timezone    string `json:"timezone"`
	Isp         string `json:"isp"`
	Org         string `json:"org"`
	As          string `json:"as"`
	Query       string `json:"query"`
}

func LogLocation(ip string) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", IP_GEOLOCATION, ip))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	log.Printf("Connecting from %s", ip)
	log.Printf("which is : %s", fmt.Sprint(body))
}
