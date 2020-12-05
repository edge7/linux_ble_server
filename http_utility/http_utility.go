package http_utility

import (
	"ble_rasbpi/data_metrics"
	"ble_rasbpi/notifications"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Send_http_post(type_sensor string, value string) {

	api_url_map := map[string]string{
		"out_humidity":    "http://localhost:8080/api/data/put_out_humidity",
		"soil_humidity":   "http://localhost:8080/api/data/put_soil_humidity",
		"out_temperature": "http://localhost:8080/api/data/put_out_temperature",
	}
	url := api_url_map[type_sensor]
	dl := data_metrics.GetDataLogger()
	dl.AddValue(type_sensor, value)
	payload := map[string]interface{}{"value": value}
	jsonStr, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Error when trying to connect to Server HTTP")
		pushover_notification.NotifyPushover("Unable to push HTTP from GO", "Irrigation")
		return

	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

}
