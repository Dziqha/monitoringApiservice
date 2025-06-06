package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testtodo/utils"
	"time"
)

type Service struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Auth struct  {
	Username string `json:"username"`
	Password string `json:"password"`
}

func StartMonitor() {
	file, err := ioutil.ReadFile("config/url.json")
	if err != nil {
		utils.Logger.Println("Gagal membaca file konfigurasi", err)
		return 
	}

	var services []Service
	json.Unmarshal(file, &services)
	
	ticker := time.NewTicker(10 * time.Second)
	for {
		<- ticker.C
		for _, service := range services {
			go checkService(service)
		}
	}
}


func checkService(service Service) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	
	resp, err := client.Get(service.Url)
	if err != nil || resp.StatusCode != 200 {
		utils.Logger.Printf("%s is down (%s)\n" ,service.Name, service.Url)
		return
	}
	utils.Logger.Printf("%s is up (%s)\n" ,service.Name, service.Url)
}
