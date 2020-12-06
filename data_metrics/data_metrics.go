package data_metrics

import (
	"ble_rasbpi/logger"
	"ble_rasbpi/notifications"
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"
)

var once sync.Once

type timeAndData struct {
	data      float32
	timestamp time.Time
}

type dataLogger struct {
	data_map map[string][]timeAndData
}

func check_anomalies(c []timeAndData, sensor string) {
	if len(c) < 5 {
		ed7_logger.Info("Cannot check anomalies as length is " + string(len(c)))
		return
	}
	avg := float32(0.0)
	old_list := ""
	for i := 0; i < 4; i++ {
		avg += c[i].data
		old_list += " __ " + fmt.Sprintf("%f", c[i].data) + " __ "
	}
	avg = avg / 4
	diff := math.Abs(float64((c[4].data-avg)/avg)) * 100.0
	if diff > 20.0 {
		pushover_notification.NotifyPushover("Diff is "+
			fmt.Sprintf("%f", diff), sensor)
		ed7_logger.Warn("Extreme value for " + sensor)
		ed7_logger.Warn("List values: " + old_list)
		ed7_logger.Warn("Current value is: " + fmt.Sprintf("%f", c[4].data))
	} else {
		ed7_logger.Info("Sensor: " + sensor + ". No Anomaly as diff is " + fmt.Sprintf("%f", diff))
	}

}
func (dl *dataLogger) CheckTime() {
	for sensor, elements := range dl.data_map {
		size := len(elements)
		if size == 0 {
			continue
		}
		t1 := time.Now()
		diff := t1.Sub(elements[size-1].timestamp)
		hours := diff.Hours()
		if hours > 5.0 {
			mex := "Sensor " + sensor + " has last update value at " + elements[size-1].timestamp.Format("Mon Jan 2 15:04:05 2006")
			ed7_logger.Warn(mex)
			pushover_notification.NotifyPushover(mex, "Timeout sensor")
		}

	}
}
func (dl *dataLogger) AddValue(value_string string, new_value_s string) {
	ed7_logger.Info("New value in dataLogger. Value is: " + new_value_s + " sensor is " + value_string)
	new_value, _ := strconv.ParseFloat(new_value_s, 32)
	max := 5
	c := dl.data_map[value_string]
	value := timeAndData{
		data:      float32(new_value),
		timestamp: time.Now(),
	}
	if len(c) == max {
		for i := 0; i < max-1; i++ {
			c[i] = c[i+1]
		}
		c[max-1] = value
		check_anomalies(c, value_string)

	} else {
		c = append(c, value)
	}
	dl.data_map[value_string] = c

}

var dl *dataLogger

func GetDataLogger() *dataLogger {
	once.Do(
		func() {
			data_map_init := map[string][]timeAndData{

				"out_humidity":    make([]timeAndData, 0, 5),
				"soil_humidity":   make([]timeAndData, 0, 5),
				"out_temperature": make([]timeAndData, 0, 5),
			}
			dl = &dataLogger{
				data_map: data_map_init,
			}

		})

	return dl
}
