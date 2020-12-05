package main

import (
	"ble_rasbpi/data_metrics"
	"ble_rasbpi/logger"
	"ble_rasbpi/sensors"
	"fmt"
	"log"
	"time"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
	"github.com/paypal/gatt/examples/service"
)

func main() {

	ed7_logger.Info("Server BLE has just started")

	d, err := gatt.NewDevice(option.DefaultServerOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s", err)
	}
	// Register optional handlers.
	d.Handle(
		gatt.CentralConnected(func(c gatt.Central) { log.Println("Connect: " + c.ID()) }),
		gatt.CentralDisconnected(func(c gatt.Central) { log.Println("Disconnect: " + c.ID()) }),
	)

	ed7_logger.Info("Device has been opened")

	// A mandatory handler for monitoring device state.
	onStateChanged := func(d gatt.Device, s gatt.State) {
		fmt.Printf("State: %s\n", s)
		switch s {
		case gatt.StatePoweredOn:
			// Setup GAP and GATT services for Linux implementation.
			// OS X doesn't export the access of these services.
			d.AddService(service.NewGapService("Gopher")) // no effect on OS X
			d.AddService(service.NewGattService())        // no effect on OS X

			// A simple count service for demo.
			s1 := sensor.NewTempHumidityService()
			d.AddService(s1)

			// Advertise device name and service's UUIDs.
			d.AdvertiseNameAndServices("Raspi!", []gatt.UUID{s1.UUID()})

		default:
		}
	}

	d.Init(onStateChanged)
	for {
		go check_sensors_stuck()
		time.Sleep(1 * time.Hour)
	}
}
func check_sensors_stuck() {
	ed7_logger.Info("Checking if sensors are stuck")
	dl := data_metrics.GetDataLogger()
	dl.CheckTime()

}
