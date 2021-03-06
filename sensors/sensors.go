package sensor

import (
	h_utility "ble_rasbpi/http_utility"
	"ble_rasbpi/logger"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/paypal/gatt"
)

var humiditySoil = []byte("-1")
var tempOutside = []byte("-1")
var humidityOut = []byte("-1")

func NewTempHumidityService() *gatt.Service {
	service_ble := os.Getenv("service_ble")
	if len(service_ble) > 0 {
		ed7_logger.Info("Service BLE is defined as " + service_ble)
		ed7_logger.Info("Am gonna use that one")
	} else {
		ed7_logger.Info("Service BLE is not defined using the default ")
		service_ble = "09fc95c0-c111-11e3-9904-0002a5d5c51b"
	}
	s := gatt.NewService(gatt.MustParseUUID(service_ble))

	c := s.AddCharacteristic(gatt.MustParseUUID("11fac9e0-c111-11e3-9246-0002a5d5c51b"))

	// Humidity Soil
	c.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			rsp.Write(humiditySoil)
			log.Println("Sensor has read Humidty Soil value")
		})

	c.HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			ed7_logger.Info("Got Humidity Soil value: " + string(data))
			humiditySoil = data
			h_utility.Send_http_post("soil_humidity", string(data))
			return gatt.StatusSuccess
		})

	// Humidity Out
	c = s.AddCharacteristic(gatt.MustParseUUID("11fac9e0-c111-11e3-9246-0002a5d5c51c"))
	c.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			rsp.Write(humidityOut)
			ed7_logger.Info("Sensor has read Humidity out")
		})

	c.HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			ed7_logger.Info("Got Humidity out value: " + string(data))
			humidityOut = data
			h_utility.Send_http_post("out_humidity", string(data))
			return gatt.StatusSuccess
		})

	// Temperature Outside
	c = s.AddCharacteristic(gatt.MustParseUUID("11fac9e0-c111-11e3-9246-0002a5d5c51d"))
	c.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			rsp.Write(tempOutside)
			ed7_logger.Info("Sensor has read Temperature out")
		})

	c.HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			ed7_logger.Info("Got Temperature out value: " + string(data))
			tempOutside = data
			valueToSend := string(data)
			f, err := strconv.ParseFloat(string(data), 64)
			if err == nil {
				if f < 0 {
					valueToSend = fmt.Sprintf("%.2f", (f + 3275.0))
					ed7_logger.Info("Modifying original, now I have got ", valueToSend)
				}
			}
			h_utility.Send_http_post("out_temperature", valueToSend)
			return gatt.StatusSuccess
		})

	return s
}
