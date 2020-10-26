package sensor

import (
	"log"

	"github.com/paypal/gatt"
)

var humiditySoil = []byte("-1")
var tempOutside = []byte("-1")
var humidityOut = []byte("-1")

func NewTempHumidityService() *gatt.Service {
	s := gatt.NewService(gatt.MustParseUUID("09fc95c0-c111-11e3-9904-0002a5d5c51b"))

	c := s.AddCharacteristic(gatt.MustParseUUID("11fac9e0-c111-11e3-9246-0002a5d5c51b"))

	// Humidity Soil
	c.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			rsp.Write(humiditySoil)
			log.Println("Sensor has read Humidty Soil value")
		})

	c.HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			log.Println("Got Humidity Soil value: ", string(data))
			humiditySoil = data
			return gatt.StatusSuccess
		})

	// Humidity Out
	c := s.AddCharacteristic(gatt.MustParseUUID("11fac9e0-c111-11e3-9246-0002a5d5c51c"))
	c.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			rsp.Write(humidityOut)
			log.Println("Sensor has read Humidity out")
		})

	c.HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			log.Println("Got Humidity out value: ", string(data))
			humidityOut = data
			return gatt.StatusSuccess
		})

	// Temperature Outside
	c := s.AddCharacteristic(gatt.MustParseUUID("11fac9e0-c111-11e3-9246-0002a5d5c51d"))
	c.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			rsp.Write(tempOutside)
			log.Println("Sensor has read Temperature out")
		})

	c.HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			log.Println("Got Temperature out value: ", string(data))
			tempOutside = data
			return gatt.StatusSuccess
		})

	return s
}
