package http_utility

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Send_http_post(type_sensor string, value string) {

	api_url_map := map[string]string{
		"out_humidity":    "http://localhost:8080/api/data/put_out_humidity",
		"soil_humidity":   "http://localhost:8080/api/data/put_soil_humidity",
		"out_temperature": "http://localhost:8080/api/data/put_out_temperature",
	}
	url := api_url_map[type_sensor]
	payload := map[string]interface{}{"value": value}
	jsonStr, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
