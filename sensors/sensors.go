package sensor

import (
	hutility "ble_server/http_utility"
	ed7logger "ble_server/logger"
	"fmt"
	"os"
	"strconv"

	"github.com/paypal/gatt"
)

var humiditySoil = []byte("-1")
var tempOutside = []byte("-1")
var humidityOut = []byte("-1")
var current = []byte("-1")

func NewTempHumidityService() *gatt.Service {
	serviceBle := os.Getenv("service_ble")
	if len(serviceBle) > 0 {
		ed7logger.Info("Service BLE is defined as " + serviceBle)
		ed7logger.Info("Am gonna use that one")
	} else {
		ed7logger.Info("Service BLE is not defined using the default ")
		serviceBle = "09fc95c0-c111-11e3-9904-0002a5d5c51b"
	}
	s := gatt.NewService(gatt.MustParseUUID(serviceBle))

	c := s.AddCharacteristic(gatt.MustParseUUID("11fac9e0-c111-11e3-9246-0002a5d5c51b"))

	// Humidity Soil
	c.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			rsp.Write(humiditySoil)
			ed7logger.Info("Sensor has read Humidity Soil value")
		})

	c.HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			ed7logger.Info("Got Humidity Soil value: " + string(data))
			humiditySoil = data
			hutility.Send_http_post("soil_humidity", string(data))
			return gatt.StatusSuccess
		})

	// Humidity Out
	c = s.AddCharacteristic(gatt.MustParseUUID("11fac9e0-c111-11e3-9246-0002a5d5c51c"))
	c.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			rsp.Write(humidityOut)
			ed7logger.Info("Sensor has read Humidity out")
		})

	c.HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			ed7logger.Info("Got Humidity out value: " + string(data))
			humidityOut = data
			hutility.Send_http_post("out_humidity", string(data))
			return gatt.StatusSuccess
		})

	// Current sensor
	c = s.AddCharacteristic(gatt.MustParseUUID("11cac9e0-c111-11e3-9246-0002a5d5c51c"))
	c.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			rsp.Write(current)
			ed7logger.Info("Sensor has read current")
		})

	c.HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			ed7logger.Info("Got Current value: " + string(data))
			current = data
			hutility.Send_http_post("current", string(data))
			return gatt.StatusSuccess
		})

	// Temperature Outside
	c = s.AddCharacteristic(gatt.MustParseUUID("11fac9e0-c111-11e3-9246-0002a5d5c51d"))
	c.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			rsp.Write(tempOutside)
			ed7logger.Info("Sensor has read Temperature out")
		})

	c.HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			ed7logger.Info("Got Temperature out value: " + string(data))
			tempOutside = data
			valueToSend := string(data)
			f, err := strconv.ParseFloat(string(data), 64)
			if err == nil {
				if f < 0 {
					valueToSend = fmt.Sprintf("%.2f", (f + 3275.0))
					ed7logger.Info("Modifying original, now I have got ", valueToSend)
				}
			}
			hutility.Send_http_post("out_temperature", valueToSend)
			return gatt.StatusSuccess
		})

	return s
}
