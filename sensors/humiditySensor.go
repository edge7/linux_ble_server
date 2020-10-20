package sensor

import (
	"log"

	"github.com/paypal/gatt"
)

var humidity = []byte("-1")

func NewHumidityService() *gatt.Service {
	s := gatt.NewService(gatt.MustParseUUID("09fc95c0-c111-11e3-9904-0002a5d5c51b"))

	c := s.AddCharacteristic(gatt.MustParseUUID("11fac9e0-c111-11e3-9246-0002a5d5c51b"))
	c.HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			rsp.Write(humidity)
			log.Println("Sensor has read charateristics")
		})

	c.HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			log.Println("Got Humidity value: ", string(data))
			humidity = data
			return gatt.StatusSuccess
		})

	return s
}
